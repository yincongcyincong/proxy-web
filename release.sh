#!/bin/bash
rm -rf ./zip
mkdir ./zip
set CGO_ENABLED=0
#linux
GOOS=linux GOARCH=386 go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-linux-386.tar.gz" proxy-web && cd ..
GOOS=linux GOARCH=amd64 go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-linux-amd64.tar.gz" proxy-web && cd ..
#darwin
GOOS=darwin GOARCH=386 go build go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-darwin-386.tar.gz" proxy-web && cd ..  
GOOS=darwin GOARCH=amd64 go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-darwin-amd64.tar.gz" proxy-web && cd ..
GOOS=darwin GOARCH=arm go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-darwin-arm.tar.gz" proxy-web && cd ..
GOOS=darwin GOARCH=arm64 go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-darwin-arm64.tar.gz"proxy-web && cd ..
#windows
cd proxy/proxy-web
rm -rf proxy-web
cd ..
cd ..
GOOS=windows GOARCH=386 go build -ldflags "-H windowsgui" && mv proxy-web.exe proxy/proxy-web/proxy-web.exe && cd proxy && tar zcfv "../zip/proxy-web-windows-386.tar.gz" proxy-web && cd ..
GOOS=windows GOARCH=amd64 go build -ldflags "-H windowsgui" && mv proxy-web.exe proxy/proxy-web/proxy-web.exe && cd proxy && tar zcfv "../zip/proxy-web-windows-amd64.tar.gz" proxy-web && cd ..

rm -rf proxy-web proxy-web.exe
