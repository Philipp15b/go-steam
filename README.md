# Steam for Go

This library allows you to interact with Steam as if it was an actual Steam client.
It's a port of [SteamKit2](https://github.com/SteamRE/SteamKit) to Go.

## Installation

    go get github.com/Philipp15b/go-steam

## Usage

You can view the documentation with the `godoc` tool or
[online on godoc.org](http://godoc.org/github.com/Philipp15b/go-steam).

When updating, always check the [`CHANGELOG.md`](CHANGELOG.md) first.

## Updating go-steam for a new SteamKit version

To update go-steam for a new version of SteamKit, do the following:

	go get code.google.com/p/goprotobuf/protoc-gen-go
    git submodule init && git submodule update
    cd generator
    go run generator.go clean proto steamlang

Replace `steamlang` with `steamlang:nodebug` if you want to exclude the `String() string` methods
for the generated SteamLanguage enums.

Make sure that `protoc-gen-go`/`$GOPATH/bin` is in your `$PATH`.
On Windows you also have to have [.NET Framework](https://www.microsoft.com/net/downloads) installed,
on other operating systems [mono](http://www.go-mono.com/mono-downloads/download.html)
is used (must be in `$PATH` too).

Apply the protocol changes where necessary.

## License

Steam for Go is licensed under the New BSD License. More information can be found in LICENSE.txt.