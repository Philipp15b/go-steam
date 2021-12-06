/*
This program generates the protobuf and SteamLanguage files from the SteamKit data.
*/
package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var printCommands = false

func main() {
	args := strings.Join(os.Args[1:], " ")

	found := false
	if strings.Contains(args, "clean") {
		clean()
		found = true
	}
	if strings.Contains(args, "steamlang") {
		buildSteamLanguage()
		found = true
	}
	if strings.Contains(args, "proto") {
		buildProto()
		found = true
	}

	if !found {
		os.Stderr.WriteString("Invalid target!\nAvailable targets: clean, proto, steamlang\n")
		os.Exit(1)
	}
}

func clean() {
	print("# Cleaning")
	cleanGlob("../protocol/**/*.pb.go")
	cleanGlob("../tf2/protocol/**/*.pb.go")
	cleanGlob("../dota/protocol/**/*.pb.go")
	cleanGlob("../csgo/protocol/**/*.pb.go")

	os.Remove("../protocol/steamlang/enums.go")
	os.Remove("../protocol/steamlang/messages.go")
}

func cleanGlob(pattern string) {
	protos, _ := filepath.Glob(pattern)
	for _, proto := range protos {
		err := os.Remove(proto)
		if err != nil {
			panic(err)
		}
	}
}

func buildSteamLanguage() {
	print("# Building Steam Language")
	execute("dotnet", []string{"run", "-c", "release", "-p", "./GoSteamLanguageGenerator", "./SteamKit", "../protocol/steamlang"})
	execute("gofmt", []string{"-w", "../protocol/steamlang/enums.go", "../protocol/steamlang/messages.go"})
}

func buildProto() {
	print("# Building Protobufs")

	buildProtoMap("steam", clientProtoFiles, "../protocol/protobuf")
	buildProtoMap("tf2", tf2ProtoFiles, "../tf2/protocol/protobuf")
	buildProtoMap("dota2", dotaProtoFiles, "../dota/protocol/protobuf")
	buildProtoMap("csgo", csgoProtoFiles, "../csgo/protocol/protobuf")
}

func buildProtoMap(srcSubdir string, files map[string]string, outDir string) {
	_ = os.MkdirAll(outDir, os.ModePerm)

	var opt []string

	opt = append(opt, "--go_opt=Mgoogle/protobuf/descriptor.proto=google/protobuf/descriptor.proto")

	for proto := range files {
		opt = append(opt, "--go_opt=Msteammessages.proto=Protobufs/"+srcSubdir+"/steammessages.proto")
		opt = append(opt, "--go_opt=M"+proto+"=Protobufs/"+srcSubdir+"/"+proto)
	}

	if srcSubdir == "dota2" {
		opt = append(opt, "--go_opt=Mecon_shared_enums.proto=Protobufs/"+srcSubdir+"/econ_shared_enums.proto")
	}

	for proto, out := range files {
		full := filepath.Join(outDir, out)
		print("# Building: " + full)
		compileProto("Protobufs", srcSubdir, proto, full, opt)
		fixProto(outDir, full)
	}
}

// Maps the proto files to their target files.
// See `SteamKit/Resources/Protobufs/steamclient/generate-base.bat` for reference.
var clientProtoFiles = map[string]string{
	"steammessages_base.proto":   "base.pb.go",
	"encrypted_app_ticket.proto": "app_ticket.pb.go",

	"steammessages_clientserver.proto":         "client_server.pb.go",
	"steammessages_clientserver_2.proto":       "client_server_2.pb.go",
	"steammessages_clientserver_friends.proto": "client_server_friends.pb.go",
	"steammessages_clientserver_login.proto":   "client_server_login.pb.go",
	"steammessages_sitelicenseclient.proto":    "client_site_license.pb.go",

	"content_manifest.proto": "content_manifest.pb.go",

	"steammessages_unified_base.steamclient.proto":      "unified/base.pb.go",
	"steammessages_cloud.steamclient.proto":             "unified/cloud.pb.go",
	"steammessages_credentials.steamclient.proto":       "unified/credentials.pb.go",
	"steammessages_deviceauth.steamclient.proto":        "unified/deviceauth.pb.go",
	"steammessages_gamenotifications.steamclient.proto": "unified/gamenotifications.pb.go",
	"steammessages_offline.steamclient.proto":           "unified/offline.pb.go",
	"steammessages_parental.steamclient.proto":          "unified/parental.pb.go",
	"steammessages_partnerapps.steamclient.proto":       "unified/partnerapps.pb.go",
	"steammessages_player.steamclient.proto":            "unified/player.pb.go",
	"steammessages_publishedfile.steamclient.proto":     "unified/publishedfile.pb.go",
}

var tf2ProtoFiles = map[string]string{
	"base_gcmessages.proto":  "base.pb.go",
	"econ_gcmessages.proto":  "econ.pb.go",
	"gcsdk_gcmessages.proto": "gcsdk.pb.go",
	"tf_gcmessages.proto":    "tf.pb.go",
	"gcsystemmsgs.proto":     "system.pb.go",
}

var dotaProtoFiles = map[string]string{
	"base_gcmessages.proto":   "base.pb.go",
	"econ_shared_enums.proto": "econ_shared_enum.pb.go",
	"econ_gcmessages.proto":   "econ.pb.go",
	"gcsdk_gcmessages.proto":  "gcsdk.pb.go",
	"gcsystemmsgs.proto":      "system.pb.go",
	"steammessages.proto":     "steam.pb.go",
}

var csgoProtoFiles = map[string]string{
	"base_gcmessages.proto":        "base.pb.go",
	"cstrike15_gcmessages.proto":   "cstrike15gc.pb.go",
	"cstrike15_usermessages.proto": "cstrike15user.pb.go",
	"econ_gcmessages.proto":        "econ.pb.go",
	"engine_gcmessages.proto":      "enginegc.pb.go",
	"fatdemo.proto":                "fatdemo.pb.go",
	"gcsdk_gcmessages.proto":       "gcsdk.pb.go",
	"gcsystemmsgs.proto":           "system.pb.go",
	"netmessages.proto":            "net.pb.go",
	"network_connection.proto":     "networkconnection.pb.go",
	"steammessages.proto":          "steam.pb.go",
	"uifontfile_format.proto":      "uifontfile.pb.go",
}

func compileProto(srcBase, srcSubdir, proto, target string, opt []string) {
	outDir, _ := filepath.Split(target)
	err := os.MkdirAll(outDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	args := []string{
		"-I=" + srcBase,
		"-I=" + srcBase + "/google",
		"-I=" + srcBase + "/" + srcSubdir,
		"--go_out=" + outDir,
	}
	args = append(args, opt...)
	args = append(args, proto)

	execute("protoc", args)

	dir := outDir + "Protobufs/" + srcSubdir + "/" + proto
	file := filepath.Join(dir, strings.Replace(proto, ".proto", ".pb.go", 1))

	err = forceRename(file, target)
	if err != nil {
		panic(err)
	}

	// clean stuff
	if err := os.RemoveAll(outDir + "/Protobufs"); err != nil {
		panic(err)
	}
}

func forceRename(from, to string) error {
	if from != to {
		os.Remove(to)
	}
	return os.Rename(from, to)
}

var pkgRegex = regexp.MustCompile(`(package \w+)`)
var unusedImportCommentRegex = regexp.MustCompile("// discarding unused import .*\n")
var fileDescriptorVarRegex = regexp.MustCompile(`fileDescriptor\d+`)

func fixProto(outDir, path string) {
	// goprotobuf is really bad at dependencies, so we must fix them manually...
	// It tries to load each dependency of a file as a seperate package (but in a very, very wrong way).
	// Because we want some files in the same package, we'll remove those imports to local files.

	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, file, parser.ImportsOnly)
	if err != nil {
		panic("Error parsing " + path + ": " + err.Error())
	}

	importsToRemove := make([]*ast.ImportSpec, 0)
	for _, i := range f.Imports {
		if strings.Contains(i.Path.Value, "google/protobuf/descriptor.proto") {
			continue
		}
		// We remove all local imports
		if strings.Contains(i.Path.Value, ".proto") {
			importsToRemove = append(importsToRemove, i)
		}
	}

	for _, itr := range importsToRemove {
		// remove the package name from all types
		file = bytes.Replace(file, []byte(itr.Name.Name+"."), []byte{}, -1)
		// and remove the import itself
		file = bytes.Replace(file, []byte(fmt.Sprintf("%v %v", itr.Name.Name, itr.Path.Value)), []byte{}, -1)
	}

	// remove warnings
	file = unusedImportCommentRegex.ReplaceAllLiteral(file, []byte{})

	// fix the package name
	file = pkgRegex.ReplaceAll(file, []byte("package "+inferPackageName(path)))

	// fix the google dependency;
	// we just reuse the one from protoc-gen-go
	file = bytes.Replace(file, []byte("google/protobuf/descriptor.proto"), []byte("google.golang.org/protobuf/types/descriptorpb"), -1)

	// we need to prefix local variables created by protoc-gen-go so that they don't clash with others in the same package
	filename := strings.Split(filepath.Base(path), ".")[0]
	file = fileDescriptorVarRegex.ReplaceAllFunc(file, func(match []byte) []byte {
		return []byte(filename + "_" + string(match))
	})

	err = ioutil.WriteFile(path, file, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func inferPackageName(path string) string {
	pieces := strings.Split(path, string(filepath.Separator))
	return pieces[len(pieces)-2]
}

func print(text string) { os.Stdout.WriteString(text + "\n") }

func printerr(text string) { os.Stderr.WriteString(text + "\n") }

// This writer appends a "> " after every newline so that the outpout appears quoted.
type quotedWriter struct {
	w       io.Writer
	started bool
}

func newQuotedWriter(w io.Writer) *quotedWriter {
	return &quotedWriter{w, false}
}

func (w *quotedWriter) Write(p []byte) (n int, err error) {
	if !w.started {
		_, err = w.w.Write([]byte("> "))
		if err != nil {
			return n, err
		}
		w.started = true
	}

	for i, c := range p {
		if c == '\n' {
			nw, err := w.w.Write(p[n : i+1])
			n += nw
			if err != nil {
				return n, err
			}

			_, err = w.w.Write([]byte("> "))
			if err != nil {
				return n, err
			}
		}
	}
	if n != len(p) {
		nw, err := w.w.Write(p[n:len(p)])
		n += nw
		return n, err
	}
	return
}

func execute(command string, args []string) {
	if printCommands {
		print(command + " " + strings.Join(args, " "))
	}

	cmd := exec.Command(command, args...)
	cmd.Stdout = newQuotedWriter(os.Stdout)
	cmd.Stderr = newQuotedWriter(os.Stderr)

	err := cmd.Run()
	if err != nil {
		printerr(err.Error())
		os.Exit(1)
	}
}
