# Steam for Go

This library implements Steam's protocol to allow automation of different actions on Steam without running an actual Steam client. It is based on [SteamKit2](https://github.com/SteamRE/SteamKit), a .NET library.

Some of the currently implemented features:

  * Trading
  * Friend and group management
  * Chatting with friends
  * Persona states (online, offline, looking to trade, etc.)
  * SteamGuard
  * Trade Offer notifications
  * Team Fortress 2: Crafting, moving, naming and deleting items

## Installation

    go get github.com/Philipp15b/go-steam

## Usage

You can view the documentation with the [`godoc`](http://golang.org/cmd/godoc) tool or
[online on godoc.org](http://godoc.org/github.com/Philipp15b/go-steam).

There's also more information on trading in the documentation of the `trade` sub-package.

## Updating go-steam to a new SteamKit version

To update go-steam to a new version of SteamKit, do the following:

	go get code.google.com/p/goprotobuf/protoc-gen-go
    git submodule init && git submodule update
    cd generator
    go run generator.go clean proto steamlang

Replace `steamlang` with `steamlang:nodebug` if you want to exclude the `String() string` methods
for the generated SteamLanguage enums. They make debugging much easier, but do increase binary size a bit.

Make sure that `protoc-gen-go`/`$GOPATH/bin` is in your `$PATH`.

To compile the Steam Language files, you also need the [.NET Framework](https://www.microsoft.com/net/downloads)
on Windows or [mono](http://www.go-mono.com/mono-downloads/download.html) on other operating systems.

Apply the protocol changes where necessary.

## License

Steam for Go is licensed under the New BSD License. More information can be found in LICENSE.txt.