#get steamkit submodule
    git submodule update --init --recursive

#windows
    build project in visual studio
    go run generator.go

#linux
    install monodevelop 4.x
    install mono-xbuild
    xbuild GoSteamLanguageGenerator.csproj /p:OutputPath=bin/Debug
    go run generator.go clean proto steamlang
