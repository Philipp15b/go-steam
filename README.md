# Steam for Go

This library allows you to interact with Steam as if it was an actual Steam client.
It's a port of [SteamKit2](https://github.com/SteamRE/SteamKit) to Go.

## Installation

    go get github.com/Philipp15b/go-steam

## Usage

You can view the documentation with the `godoc` tool or
[online on godoc.org](http://godoc.org/github.com/Philipp15b/go-steam).

## Updating go-steam to a new SteamKit version

To update go-steam to a new version of SteamKit, do the following:

	go get code.google.com/p/goprotobuf/protoc-gen-go
    git submodule init && git submodule update
    cd generator
    go run generator.go clean proto steamlang

Replace `steamlang` with `steamlang:nodebug` if you want to exclude the `String() string` methods
for the generated SteamLanguage enums. They make debugging much easier, but do increase binary size a bit.

Make sure that `protoc-gen-go`/`$GOPATH/bin` is in your `$PATH`.

To compile the Steam Langugae files, you also need the [.NET Framework](https://www.microsoft.com/net/downloads)
on Windows or [mono](http://www.go-mono.com/mono-downloads/download.html) on other operating systems.

Apply the protocol changes where necessary.

## License

Steam for Go is licensed under the New BSD License. More information can be found in LICENSE.txt.