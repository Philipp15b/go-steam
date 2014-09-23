# Steam for Go

It is based on [SteamKit2](https://github.com/SteamRE/SteamKit), a .NET library.

[go-steam](https://github.com/Philipp15b/go-steam), a go library.

## Compile on Windows 
    mkdir -p /e/git/myapp
    mkdir -p /e/git/myapp/src
    mkdir -p /e/git/myapp/bin
    mkdir -p /e/git/myapp/protobuf_gen_tool
    download protoc-2.5.0-win32.zip  on http://code.google.com/p/protobuf/downloads/ to /e/git/myapp/protobuf_gen_tool
    unzip protoc.exe from protoc-2.5.0-win32.zip to /e/git/myapp/protobuf_gen_tool
    export PATH=$PATH:/e/git/myapp/protobuf_gen_tool
    export GOPATH=/e/git/myapp
    cd /e/git/myapp/src
    mkdir -p code.google.com/p
    cd code.google.com/p
    git clone https://github.com/smithfox/goprotobuf.git
    cd /e/git/myapp/src
    mkdir -p github.com/Philipp15b
    cd github.com/Philipp15b
    git clone https://github.com/smithfox/go-steam.git
    cd /e/git/myapp/src/github.com/Philipp15b/go-steam
    git submodule init && git submodule update
    cd /e/git/myapp/protobuf_gen_tool
    go build ../src/code.google.com/p/goprotobuf/protoc-gen-go/main.go -o protoc-gen-go.exe
    download [mono](http://www.go-mono.com/mono-downloads/download.html)
    install mono and add it's bin path to $PATH
    cd /e/git/myapp/src/github.com/Philipp15b/go-steam/generator/GoSteamLanguageGenerator
    run xbuild will create bin/Debug/GoSteamLanguageGenerator.exe  and SteamLanguageParser.exe
    cd /e/git/myapp/src/github.com/Philipp15b/go-steam/generator
    go run generator.go clean proto steamlang
    will generator protocbuf go file under /e/git/myapp/src/github.com/Philipp15b/internal
    
## Usage

You can view the documentation with the [`godoc`](http://golang.org/cmd/godoc) tool or
[online on godoc.org](http://godoc.org/github.com/Philipp15b/go-steam).

You should also take a look at the following sub-packages:

  * [`trade`](http://godoc.org/github.com/Philipp15b/go-steam/trade) for trading
  * [`tradeoffer`](http://godoc.org/github.com/Philipp15b/go-steam/tradeoffer) for trade offers
  * [`economy/inventory`](http://godoc.org/github.com/Philipp15b/go-steam/economy/inventory) for inventories
  * [`tf2`](http://godoc.org/github.com/Philipp15b/go-steam/tf2) for Team Fortress 2 related things


Steam for Go is licensed under the New BSD License. More information can be found in LICENSE.txt.
