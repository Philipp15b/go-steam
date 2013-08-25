/*
This program generates the protobuf and SteamLanguage files from the SteamKit data.
*/
package main

import (
	"bytes"
	"go/parser"
	"go/token"
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
	protos, _ := filepath.Glob("../internal/*.pb.go")
	for _, proto := range protos {
		err := os.Remove(proto)
		if err != nil {
			panic(err)
		}
	}

	os.Remove("../internal/steam_language_enums.go")
	os.Remove("../internal/steam_language_internal.go")
}

func buildSteamLanguage(debug bool) {
	print("# Building Steam Language")
	exePath := "./GoSteamLanguageGenerator/bin/Debug/GoSteamLanguageGenerator.exe"
	d := ""
	if debug {
		d = "debug"
	}
	if runtime.GOOS != "windows" {
		execute("mono", exePath, "./SteamKit", "../internal", d)
	} else {
		execute(exePath, "./SteamKit", "../internal", d)
	}
	execute("gofmt", "-w", "../internal/steam_language_enums.go", "../internal/steam_language_internal.go")
}

var pkgRegex = regexp.MustCompile("(package \\w+)")

func buildProto() {
	print("# Building Proto")

	print("## Compiling")
	protobufs := "SteamKit/Resources/Protobufs"
	protos, err := filepath.Glob(protobufs + "/steamclient/*.proto")
	if err != nil {
		panic(err)
	}

	filteredProtos := make([]string, 0, 0)
	for _, proto := range protos {
		if strings.Contains(proto, "unified") {
			continue
		}
		filteredProtos = append(filteredProtos, proto)
	}
	protos = filteredProtos

	for _, proto := range protos {
		execute("protoc", "--go_out=../internal", "-I="+protobufs+"/steamclient", "-I="+protobufs, proto)
	}

	goprotos, _ := filepath.Glob("../internal/*.pb.go")

	print("## Fixing imports")
	// goprotobuf is really bad at dependencies, so we must fix them manually...
	// It tries to load each dependency of a file as a seperate package (but in a very, very wrong way).
	// Because we want each file in the same package, we'll remove all references to
	// imported protobuf packages.
	for _, proto := range goprotos {
		file, err := ioutil.ReadFile(proto)
		if err != nil {
			panic(err)
		}

		// find the protobuf imports by parsing the file
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, proto, file, parser.ImportsOnly)
		if err != nil {
			panic("Error parsing " + proto + ": " + err.Error())
		}

		importsToRemove := make([]string, 0, 0)
		for _, i := range f.Imports {
			// google's protobuf stays in a seperate package
			if i.Name.Name == "google_protobuf" {
				continue
			}
			// only search for .pb imports
			if matched, _ := filepath.Match("*.pb", strings.Replace(i.Path.Value, "\"", "", -1)); matched {
				importsToRemove = append(importsToRemove, i.Name.Name)
			}
		}

		for _, itr := range importsToRemove {
			// remove the package name from all types
			file = bytes.Replace(file, []byte(itr+"."), []byte{}, -1)
			// and remove the import itself
			file = bytes.Replace(file, []byte("import "+itr+" \"pb\"\n"), []byte{}, -1)
		}

		// make the package local
		file = pkgRegex.ReplaceAll(file, []byte("package internal"))

		// fix the google dependency
		// we just reuse the one from protoc-gen-go
		file = bytes.Replace(file, []byte("google/protobuf/descriptor.pb"), []byte("code.google.com/p/goprotobuf/protoc-gen-go/descriptor"), -1)

		err = ioutil.WriteFile(proto, file, 0666)
		if err != nil {
			panic(err)
		}
	}

	print("## Formatting")
	for _, proto := range goprotos {
		execute("gofmt", "-w", proto)
	}
}

func print(text string) { os.Stdout.WriteString(text + "\n") }

func printerr(text string) { os.Stderr.WriteString(text + "\n") }

func execute(command string, args ...string) {
	if printCommands {
		print(command + " " + strings.Join(args, " "))
	}
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		printerr(err.Error())
		os.Exit(1)
	}
}
