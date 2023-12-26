#!/usr/bin/env fish
GOOS=windows GOARCH=386 \
CGO_ENABLED=1 CXX=i686-w64-mingw32-g++ CC=i686-w64-mingw32-gcc \
go build ../../flxDelister.go


GOOS=windows GOARCH=386 \
CGO_ENABLED=1 CXX=i686-w64-mingw32-g++ CC=i686-w64-mingw32-gcc \
go build ../../flxUpdater.go


GOOS=windows GOARCH=386 \
CGO_ENABLED=1 CXX=i686-w64-mingw32-g++ CC=i686-w64-mingw32-gcc \
go build ../../flxStageCounter.go
