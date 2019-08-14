# Generator for steamlang and protobuf

We generate Go code from SteamKit protocol descriptors, namely `steamlang` files and protocol buffer files.

## Dependencies
1. Get SteamKit submodule: `git submodule update --init --recursive`.
2. Install [`protoc`](https://developers.google.com/protocol-buffers/docs/downloads), the protocol buffer compiler. At the moment, we use Protocol Buffers 2.6.1.
3. Install `proco-gen-go`: `go get github.com/golang/protobuf/protoc-gen-go/`. At the moment, we use
`4c88cc3`.
4. Build the `GoSteamLanguageGenerator` C# project with either Visual Studio or mono.

## Execute generator

Execute `go run generator.go clean proto steamlang` to clean build files, then build protocol buffer files and then build steamlang files.
