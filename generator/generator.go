/*
This program generates the protobuf and SteamLanguage files from the SteamKit data.
*/
package main

import (
	"bytes"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
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
		buildSteamLanguage(!strings.Contains(args, "steamlang:nodebug"))
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
	protos, _ := filepath.Glob("../internal/**/*.pb.go")
	for _, proto := range protos {
		err := os.Remove(proto)
		if err != nil {
			panic(err)
		}
	}

	os.Remove("../internal/steamlang/enums.go")
	os.Remove("../internal/steamlang/messages.go")
}

func buildSteamLanguage(debug bool) {
	print("# Building Steam Language")
	exePath := "./GoSteamLanguageGenerator/bin/Debug/GoSteamLanguageGenerator.exe"
	d := ""
	if debug {
		d = "debug"
	}
	if runtime.GOOS != "windows" {
		execute("mono", exePath, "./SteamKit", "../internal/steamlang", d)
	} else {
		execute(exePath, "./SteamKit", "../internal/steamlang", d)
	}
	execute("gofmt", "-w", "../internal/steamlang/enums.go", "../internal/steamlang/messages.go")
}

func buildProto() {
	print("# Building Protobufs")

	base := "SteamKit/Resources/Protobufs"
	outDir := "../internal/protobuf"
	os.MkdirAll(outDir, os.ModePerm)
	for proto, out := range steamclient_protoFiles {
		full := filepath.Join(outDir, out)
		compileProto(base, "steamclient", proto, full)
		fixProto(full)
	}

	for proto, out := range dota2_protoFiles {
		full := filepath.Join(outDir, out)
		compileProto(base, "dota", proto, full)
		fixProto(full)
	}
}

// Maps the proto files to their target files.
// See `SteamKit/Resources/Protobufs/dota/generate-base.bat` for reference.
var steamclient_protoFiles = map[string]string{
	"steammessages_base.proto":   "base.pb.go",
	"encrypted_app_ticket.proto": "app_ticket.pb.go",

	"steammessages_clientserver.proto":   "client_server.pb.go",
	"steammessages_clientserver_2.proto": "client_server_2.pb.go",

	"content_manifest.proto": "content_manifest.pb.go",

	"iclient_objects.proto": "iclient_objects.pb.go",

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

// Maps the proto files to their target files.
// See `SteamKit/Resources/Protobufs/steamclient/generate-base.bat` for reference.
var dota2_protoFiles = map[string]string{
	"steammessages.proto":          "dota/steammessages.pb.go",
	"gcsystemmsgs.proto":           "dota/gcsystemmsgs.pb.go",
	"base_gcmessages.proto":        "dota/basegcmessages.pb.go",
	"gcsdk_gcmessages.proto":       "dota/gcsdkgcmessages.pb.go",
	"econ_gcmessages.proto":        "dota/econgcmessages.pb.go",
	"matchmaker_common.proto":      "dota/matchmakercommon.pb.go",
	"network_connection.proto":     "dota/networkconnection.pb.go",
	"dota_gcmessages_common.proto": "dota/gcmessagescommon.pb.go",
	"dota_gcmessages_client.proto": "dota/gcmessagesclient.pb.go",
	"dota_gcmessages_server.proto": "dota/gcmessagesserver.pb.go",
}

func compileProto(base, subbase1, proto, target string) {
	outDir, _ := filepath.Split(target)
	err := os.MkdirAll(outDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	execute("protoc", "--go_out="+outDir, "-I="+base+"/"+subbase1, "-I="+base, filepath.Join(base, subbase1, proto))
	out := strings.Replace(filepath.Join(outDir, proto), ".proto", ".pb.go", 1)
	err = forceRename(out, target)
	if err != nil {
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
var pkgCommentRegex = regexp.MustCompile(`(?s)(\/\*.*?\*\/\n)package`)

func fixProto(path string) {
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

	importsToRemove := make([]string, 0)
	for _, i := range f.Imports {
		// We remove all imports that include ".pb". This assumes unified and protobuf packages don't share anything.
		if i.Name.Name != "google_protobuf" && strings.Contains(i.Path.Value, ".pb") {
			importsToRemove = append(importsToRemove, i.Name.Name)
		}
	}

	for _, itr := range importsToRemove {
		// remove the package name from all types
		file = bytes.Replace(file, []byte(itr+"."), []byte{}, -1)
		// and remove the import itself
		file = bytes.Replace(file, []byte("import "+itr+" \"pb\"\n"), []byte{}, -1)
	}

	// remove the package comment because it just includes a list of all messages and
	// creates collisions between the others.
	file = cutAllSubmatch(pkgCommentRegex, file, 1)

	// fix the package name
	file = pkgRegex.ReplaceAll(file, []byte("package "+inferPackageName(path)))

	// fix the google dependency;
	// we just reuse the one from protoc-gen-go
	file = bytes.Replace(file, []byte("google/protobuf/descriptor.pb"), []byte("github.com/smithfox/goprotobuf/protoc-gen-go/descriptor"), -1)

	err = ioutil.WriteFile(path, file, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func inferPackageName(path string) string {
	pieces := strings.Split(path, string(filepath.Separator))
	return pieces[len(pieces)-2]
}

func cutAllSubmatch(r *regexp.Regexp, b []byte, n int) []byte {
	i := r.FindSubmatchIndex(b)
	return bytesCut(b, i[2*n], i[2*n+1])
}

// Removes the given section from the byte array
func bytesCut(b []byte, from, to int) []byte {
	buf := new(bytes.Buffer)
	buf.Write(b[:from])
	buf.Write(b[to:])
	return buf.Bytes()
}

func print(text string) { os.Stdout.WriteString(text + "\n") }

func printerr(text string) { os.Stderr.WriteString(text + "\n") }

// This writer appends a "> " after every newline so that the outpout appears quoted.
type QuotedWriter struct {
	w       io.Writer
	started bool
}

func NewQuotedWriter(w io.Writer) *QuotedWriter {
	return &QuotedWriter{w, false}
}

func (w *QuotedWriter) Write(p []byte) (n int, err error) {
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

func execute(command string, args ...string) {
	if printCommands {
		print(command + " " + strings.Join(args, " "))
	}
	cmd := exec.Command(command, args...)
	cmd.Stdout = NewQuotedWriter(os.Stdout)
	cmd.Stderr = NewQuotedWriter(os.Stderr)
	err := cmd.Run()
	if err != nil {
		printerr(err.Error())
		os.Exit(1)
	}
}
