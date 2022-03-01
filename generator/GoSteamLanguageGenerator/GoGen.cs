using System;
using System.Collections.Generic;
using System.Text;
using System.Linq;
using SteamLanguageParser;
using System.Text.RegularExpressions;

namespace GoSteamLanguageGenerator
{
    public class GoGen
    {
        /// <summary>
        /// Maps the SteamLanguage types to Go types.
        /// </summary>
        private static Dictionary<String, String> typeMap = new Dictionary<String, String> {
            {"byte", "uint8"},
            {"short", "int16"},
            {"ushort", "uint16"},
            {"int", "int32"},
            {"uint", "uint32"},
            {"long", "int64"},
            {"ulong", "uint64"},
            {"char", "uint16"}
        };

        /// <summary>
        /// Maps the enum names to their underlying Go types. In most cases, this is int32.
        /// </summary>
        private Dictionary<string, string> enumTypes = new Dictionary<string, string>();

        public void EmitEnums(Node root, StringBuilder sb)
        {
            sb.AppendLine("// Generated code");
            sb.AppendLine("// DO NOT EDIT");
            sb.AppendLine();
            sb.AppendLine("package steamlang");
            sb.AppendLine();
            sb.AppendLine("import (");
            sb.AppendLine("    \"strings\"");
            sb.AppendLine("    \"sort\"");
            sb.AppendLine("    \"fmt\"");
            sb.AppendLine(")");
            sb.AppendLine();
            foreach (Node n in root.childNodes)
            {
                EmitEnumNode(n as EnumNode, sb);
            }
        }

        public void EmitClasses(Node root, StringBuilder sb)
        {
            sb.AppendLine("// Generated code");
            sb.AppendLine("// DO NOT EDIT");
            sb.AppendLine();
            sb.AppendLine("package steamlang");
            sb.AppendLine();
            sb.AppendLine("import (");
            sb.AppendLine("    \"io\"");
            sb.AppendLine("    \"encoding/binary\"");
            sb.AppendLine("    \"google.golang.org/protobuf/proto\"");
            sb.AppendLine("    \"github.com/Philipp15b/go-steam/v3/steamid\"");
            sb.AppendLine("    \"github.com/Philipp15b/go-steam/v3/rwu\"");
            sb.AppendLine("    \"github.com/Philipp15b/go-steam/v3/protocol/protobuf\"");
            sb.AppendLine(")");
            sb.AppendLine();

            foreach (Node n in root.childNodes)
            {
                EmitClassNode(n as ClassNode, sb);
            }
        }

        private static string EmitSymbol(Symbol sym)
        {
            string s = null;
            if (sym is WeakSymbol)
            {
                s = (sym as WeakSymbol).Identifier;
            }
            else if (sym is StrongSymbol)
            {
                StrongSymbol ssym = sym as StrongSymbol;
                if (ssym.Prop == null)
                {
                    s = ssym.Class.Name;
                }
                else
                {
                    s = ssym.Class.Name + "_" + ssym.Prop.Name;
                }
            }
            return s.Replace("ulong.MaxValue", "^uint64(0)").Replace("SteamKit2.Internal.", "").Replace("SteamKit2.GC.Internal.", "");
        }

        private static string EmitType(Symbol sym)
        {
            string type = EmitSymbol(sym);
            return typeMap.ContainsKey(type) ? typeMap[type] : type;
        }

        private static string GetUpperName(string name)
        {
            return name.Substring(0, 1).ToUpper() + name.Remove(0, 1);
        }

        private void EmitEnumNode(EnumNode enode, StringBuilder sb)
        {
            string type = enode.Type != null ? EmitType(enode.Type) : "int32";
            enumTypes.Add(enode.Name, type);

            sb.AppendLine("type " + enode.Name + " " + type);
            sb.AppendLine();

            sb.AppendLine("const (");
            bool first = true;
            foreach (PropNode prop in enode.childNodes)
            {
                if (!prop.Emit) continue;

                string val = String.Join(" | ", prop.Default.Select(item =>
                {
                    var name = EmitSymbol(item);
                    // if this is an element of this enum, make sure to prefix it with its name
                    return (enode.childNodes.Exists(node => node.Name == name) ? enode.Name + "_" : "") + name;
                }));

                sb.Append("    " + enode.Name + "_" + prop.Name + " " + enode.Name + " = " + val);

                if (prop.Obsolete != null)
                {
                    if (prop.Obsolete.Length > 0)
                        sb.Append(" // Deprecated: " + prop.Obsolete);
                    else
                        sb.Append(" // Deprecated");
                }

                sb.AppendLine();

                if (first)
                    first = false;
            }
            sb.AppendLine(")");
            sb.AppendLine();

            EmitEnumStringer(enode, sb);
        }

        private void EmitEnumStringer(EnumNode enode, StringBuilder sb)
        {
            sb.AppendLine("var " + enode.Name + "_name = map[" + enode.Name + "]string{");

            // duplicate elements don't really make sense here, so we filter them
            var uniqueNodes = new List<PropNode>();
            {
                var allValues = new List<long>();
                foreach (var node in enode.childNodes)
                {
                    if (!(node is PropNode))
                        continue;
                    var prop = node as PropNode;
                    long val = EvalEnum(enode, prop);
                    if (allValues.Contains(val))
                        continue;
                    allValues.Add(val);
                    uniqueNodes.Add(prop);
                }
            }

            foreach (PropNode prop in uniqueNodes)
            {
                sb.Append("    " + EvalEnum(enode, prop) + ": ");
                sb.AppendLine("\"" + enode.Name + "_" + prop.Name + "\",");
            }

            sb.AppendLine("}");
            sb.AppendLine();

            sb.AppendLine("func (e " + enode.Name + ") String() string {");
            sb.AppendLine("    if s, ok := " + enode.Name + "_name[e]; ok {");
            sb.AppendLine("         return s");
            sb.AppendLine("    }");
            sb.AppendLine("    var flags []string");
            sb.AppendLine("    for k, v := range " + enode.Name + "_name {");
            sb.AppendLine("        	if e&k != 0 {");
            sb.AppendLine("        		flags = append(flags, v)");
            sb.AppendLine("	        }");
            sb.AppendLine("    }");
            sb.AppendLine("    if len(flags) == 0 {");
            sb.AppendLine("        return fmt.Sprintf(\"%d\", e)");
            sb.AppendLine("    }");
            sb.AppendLine("    sort.Strings(flags)");
            sb.AppendLine("    return strings.Join(flags, \" | \")");
            sb.AppendLine("}");
            sb.AppendLine();
        }

        private long EvalEnum(EnumNode enode, PropNode prop)
        {
            long number = 0;
            foreach (Symbol sym in prop.Default)
            {
                long n = 0;
                var val = (sym as WeakSymbol).Identifier;

                if (new Regex("^[-0-9]+$|^0x[-0-9a-fA-F]+$").IsMatch(val))
                {
                    int bas = 10;
                    if (val.StartsWith("0x", StringComparison.Ordinal))
                    {
                        bas = 16;
                        val = val.Substring(2);
                    }
                    n = Convert.ToInt64(val, bas);
                }
                else
                {
                    var otherNode = enode.childNodes.Single(o => o is PropNode && (o as PropNode).Name == val) as PropNode;
                    n = EvalEnum(enode, otherNode);
                }
                number = number | n;
            }
            return number;
        }

        private void EmitClassNode(ClassNode cnode, StringBuilder sb)
        {
            if (!cnode.Emit) return;

            EmitClassConstants(cnode, sb);
            EmitClassDef(cnode, sb);
            EmitClassConstructor(cnode, sb);
            EmitClassEMsg(cnode, sb);
            EmitClassSerializer(cnode, sb);
            EmitClassDeserializer(cnode, sb);
        }

        private void EmitClassConstants(ClassNode cnode, StringBuilder sb)
        {
            Func<PropNode, string> emitElement = node => cnode.Name + "_" + node.Name + " " + EmitType(node.Type) + " = " + EmitType(node.Default.FirstOrDefault());

            var statics = cnode.childNodes.Where(node => node is PropNode && (node as PropNode).Flags == "const");

            if (statics.Count() == 1)
            {
                sb.AppendLine("const " + emitElement(cnode.childNodes[0] as PropNode));
                sb.AppendLine();
                return;
            }

            bool first = true;
            foreach (PropNode node in statics)
            {
                if (first)
                {
                    sb.AppendLine("const (");
                    first = false;
                }
                sb.AppendLine("    " + emitElement(node));
            }
            if (!first)
            {
                sb.AppendLine(")");
                sb.AppendLine();
            }
        }

        private void EmitClassDef(ClassNode cnode, StringBuilder sb)
        {
            sb.AppendLine("type " + cnode.Name + " struct {");
            foreach (PropNode node in cnode.childNodes)
            {
                if (node.Flags == "const")
                {
                    continue;
                }
                else if (node.Flags == "boolmarshal" && EmitSymbol(node.Type) == "byte")
                {
                    sb.AppendLine("    " + GetUpperName(node.Name) + " bool");
                }
                else if (node.Flags == "steamidmarshal" && EmitSymbol(node.Type) == "ulong")
                {
                    sb.AppendLine("    " + GetUpperName(node.Name) + " steamid.SteamId");
                }
                else if (node.Flags == "proto")
                {
                    sb.AppendLine("    " + GetUpperName(node.Name) + " *protobuf." + EmitType(node.Type));
                }
                else
                {
                    int arraySize;
                    string typestr = EmitType(node.Type);
                    if (!String.IsNullOrEmpty(node.FlagsOpt) && Int32.TryParse(node.FlagsOpt, out arraySize))
                    {
                        // TODO: use arrays instead of slices?
                        typestr = "[]" + typestr;
                    }
                    sb.AppendLine("    " + GetUpperName(node.Name) + " " + typestr);
                }
            }
            sb.AppendLine("}");
            sb.AppendLine();
        }

        private void EmitClassConstructor(ClassNode cnode, StringBuilder sb)
        {
            sb.AppendLine("func New" + cnode.Name + "() *" + cnode.Name + " {");
            sb.AppendLine("    return &" + cnode.Name + "{");
            foreach (PropNode node in cnode.childNodes)
            {
                string ctor = null;
                Symbol firstDefault = node.Default.FirstOrDefault();
                if (node.Flags == "const")
                {
                    continue;
                }
                else if (node.Flags == "proto")
                {
                    ctor = "new(protobuf." + GetUpperName(EmitType(node.Type)) + ")";
                }
                else if (firstDefault == null)
                {
                    // only arrays/slices need explicit initialization
                    if (String.IsNullOrEmpty(node.FlagsOpt))
                        continue;
                    ctor = "make([]" + EmitType(node.Type) + ", " + node.FlagsOpt + ", " + node.FlagsOpt + ")";
                }
                else
                {
                    ctor = EmitSymbol(firstDefault);
                }
                sb.AppendLine("        " + GetUpperName(node.Name) + ": " + ctor + ",");
            }
            sb.AppendLine("    }");
            sb.AppendLine("}");
            sb.AppendLine();
        }

        private void EmitClassEMsg(ClassNode cnode, StringBuilder sb)
        {
            if (cnode.Ident == null)
                return;
            string ident = EmitSymbol(cnode.Ident);
            if (!ident.Contains("EMsg_"))
                return;
            sb.AppendLine("func (d *" + cnode.Name + ") GetEMsg() EMsg {");
            sb.AppendLine("    return " + ident);
            sb.AppendLine("}");
            sb.AppendLine();
        }

        private void EmitClassSerializer(ClassNode cnode, StringBuilder sb)
        {
            var nodeToBuf = new Dictionary<PropNode, string>();

            int tempNum = 0;
            sb.AppendLine("func (d *" + cnode.Name + ") Serialize(w io.Writer) error {");
            sb.AppendLine("    var err error");
            foreach (PropNode node in cnode.childNodes)
            {
                if (node.Flags == "proto")
                {
                    sb.AppendLine("    buf" + tempNum + ", err := proto.Marshal(d." + GetUpperName(node.Name) + ")");
                    sb.AppendLine("    if err != nil { return err }");

                    if (node.FlagsOpt != null)
                    {
                        sb.AppendLine("    d." + GetUpperName(node.FlagsOpt) + " = int32(len(buf" + tempNum + "))");
                    }

                    nodeToBuf[node] = "buf" + tempNum;
                    tempNum++;
                }
            }
            foreach (PropNode node in cnode.childNodes)
            {
                if (node.Flags == "const")
                {
                    continue;
                }
                else if (node.Flags == "boolmarshal")
                {
                    sb.AppendLine("    err = rwu.WriteBool(w, d." + GetUpperName(node.Name) + ")");
                }
                else if (node.Flags == "steamidmarshal")
                {
                    sb.AppendLine("    err = binary.Write(w, binary.LittleEndian, d." + GetUpperName(node.Name) + ")");
                }
                else if (node.Flags == "protomask")
                {
                    sb.AppendLine("    err = binary.Write(w, binary.LittleEndian, EMsg(uint32(d." + GetUpperName(node.Name) + ") | ProtoMask))");
                }
                else if (node.Flags == "protomaskgc")
                {
                    sb.AppendLine("    err = binary.Write(w, binary.LittleEndian, EMsg(uint32(d." + GetUpperName(node.Name) + ") | ProtoMask))");
                }
                else if (node.Flags == "proto")
                {
                    sb.AppendLine("    _, err = w.Write(" + nodeToBuf[node] + ")");
                }
                else
                {
                    sb.AppendLine("    err = binary.Write(w, binary.LittleEndian, d." + GetUpperName(node.Name) + ")");
                }
                if (node != cnode.childNodes[cnode.childNodes.Count - 1])
                    sb.AppendLine("    if err != nil { return err }");
            }
            sb.AppendLine("    return err");
            sb.AppendLine("}");
            sb.AppendLine();
        }

        private void EmitClassDeserializer(ClassNode cnode, StringBuilder sb)
        {
            int tempNum = 0;
            sb.AppendLine("func (d *" + cnode.Name + ") Deserialize(r io.Reader) error {");
            sb.AppendLine("    var err error");
            foreach (PropNode node in cnode.childNodes)
            {
                if (node.Flags == "const")
                {
                    continue;
                }
                else if (node.Flags == "boolmarshal")
                {
                    sb.AppendLine("    d." + GetUpperName(node.Name) + ", err = rwu.ReadBool(r)");
                }
                else if (node.Flags == "steamidmarshal")
                {
                    sb.AppendLine("    t" + tempNum + ", err := rwu.ReadUint64(r)");
                    sb.AppendLine("    if err != nil { return err }");
                    sb.AppendLine("    d." + GetUpperName(node.Name) + " = steamid.SteamId(t" + tempNum + ")");
                    tempNum++;
                    continue;
                }
                else if (node.Flags == "protomask")
                {
                    string type = EmitType(node.Type);
                    if (enumTypes.ContainsKey(type))
                        type = enumTypes[type];
                    sb.AppendLine("    t" + tempNum + ", err := rwu.Read" + GetUpperName(type) + "(r)");
                    sb.AppendLine("    if err != nil { return err }");
                    sb.AppendLine("    d." + GetUpperName(node.Name) + " = EMsg(uint32(t" + tempNum + ") & EMsgMask)");
                    tempNum++;
                    continue;
                }
                else if (node.Flags == "protomaskgc")
                {
                    sb.AppendLine("    t" + tempNum + ", err := rwu.Read" + GetUpperName(EmitType(node.Type)) + "(r)");
                    sb.AppendLine("    if err != nil { return err }");
                    sb.AppendLine("    d." + GetUpperName(node.Name) + " = uint32(t" + tempNum + ") & EMsgMask");
                    tempNum++;
                    continue;
                }
                else if (node.Flags == "proto")
                {
                    sb.AppendLine("    buf" + tempNum + " := make([]byte, d." + GetUpperName(node.FlagsOpt) + ", d." + GetUpperName(node.FlagsOpt) + ")");
                    sb.AppendLine("    _, err = io.ReadFull(r, buf" + tempNum + ")");
                    sb.AppendLine("    if err != nil { return err }");
                    sb.AppendLine("    err = proto.Unmarshal(buf" + tempNum + ", d." + GetUpperName(node.Name) + ")");
                    tempNum++;
                }
                else
                {
                    string type = EmitType(node.Type);
                    if (!String.IsNullOrEmpty(node.FlagsOpt))
                    {
                        sb.AppendLine("    err = binary.Read(r, binary.LittleEndian, d." + GetUpperName(node.Name) + ")");
                    }
                    else if (!enumTypes.ContainsKey(type))
                    {
                        sb.AppendLine("    d." + GetUpperName(node.Name) + ", err = rwu.Read" + GetUpperName(type) + "(r)");
                    }
                    else
                    {
                        sb.AppendLine("    t" + tempNum + ", err := rwu.Read" + GetUpperName(enumTypes[type]) + "(r)");
                        if (node != cnode.childNodes[cnode.childNodes.Count - 1])
                            sb.AppendLine("    if err != nil { return err }");
                        sb.AppendLine("    d." + GetUpperName(node.Name) + " = " + type + "(t" + tempNum + ")");
                        tempNum++;
                        continue;
                    }
                }
                if (node != cnode.childNodes[cnode.childNodes.Count - 1])
                    sb.AppendLine("    if err != nil { return err }");
            }
            sb.AppendLine("    return err");
            sb.AppendLine("}");
            sb.AppendLine();
        }
    }
}
