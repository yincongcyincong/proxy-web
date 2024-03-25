#!/bin/bash
rm -rf ./zip
mkdir ./zip
set CGO_ENABLED=0
#linux
GOOS=linux GOARCH=386 go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-linux-386.tar.gz" proxy-web && cd ..
GOOS=linux GOARCH=amd64 go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-linux-amd64.tar.gz" proxy-web && cd ..
GOOS=linux GOARCH=arm GOARM=7 go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-linux-arm.tar.gz" proxy-web && cd ..
GOOS=linux GOARCH=arm64 GOARM=7 go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-linux-arm64.tar.gz" proxy-web && cd ..
GOOS=linux GOARCH=mips go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-linux-mips.tar.gz" proxy-web && cd ..
GOOS=linux GOARCH=mips64 go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-linux-mips64.tar.gz" proxy-web && cd ..
GOOS=linux GOARCH=mips64le go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-linux-mips64le.tar.gz" proxy-web && cd .. 
GOOS=linux GOARCH=mipsle go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-linux-mipsle.tar.gz" proxy-web && cd ..
GOOS=linux GOARCH=ppc64 go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-linux-ppc64.tar.gz" proxy-web && cd ..
GOOS=linux GOARCH=ppc64le go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-linux-ppc64le.tar.gz" proxy-web && cd ..
GOOS=linux GOARCH=s390x go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-linux-s390x.tar.gz" proxy-web && cd ..
#android
GOOS=android GOARCH=386 go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-android-386.tar.gz" proxy-web && cd ..
GOOS=android GOARCH=amd64 go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-android-amd64.tar.gz" proxy-web && cd .. 
GOOS=android GOARCH=arm go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-android-arm.tar.gz" proxy-web && cd ..
GOOS=android GOARCH=arm64 go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-android-arm64.tar.gz" proxy-web && cd ..
#darwin
GOOS=darwin GOARCH=386 go build go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-darwin-386.tar.gz" proxy-web && cd ..  
GOOS=darwin GOARCH=amd64 go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-darwin-amd64.tar.gz" proxy-web && cd ..
GOOS=darwin GOARCH=arm go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-darwin-arm.tar.gz" proxy-web && cd ..
GOOS=darwin GOARCH=arm64 go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-darwin-arm64.tar.gz"proxy-web && cd ..
#dragonfly
GOOS=dragonfly GOARCH=amd64 go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-dragonfly-amd64.tar.gz" proxy-web && cd ..  
#freebsd
GOOS=freebsd GOARCH=386 go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-freebsd-386.tar.gz" proxy-web && cd ..
GOOS=freebsd GOARCH=amd64 go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-freebsd-amd64.tar.gz" proxy-web && cd ..
GOOS=freebsd GOARCH=arm go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-freebsd-arm.tar.gz" proxy-web && cd .. 
#nacl
GOOS=nacl GOARCH=386 go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-nacl-386.tar.gz" proxy-web && cd ..
GOOS=nacl GOARCH=amd64p32 go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-nacl-amd64p32.tar.gz" proxy-web && cd ..
GOOS=nacl GOARCH=arm go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-nacl-arm.tar.gz" proxy-web && cd ..
#netbsd
GOOS=netbsd GOARCH=386 go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-netbsd-386.tar.gz" proxy-web && cd .. 
GOOS=netbsd GOARCH=amd64 go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-netbsd-amd64.tar.gz" proxy-web && cd .. 
GOOS=netbsd GOARCH=arm go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-netbsd-arm.tar.gz" proxy-web && cd .. 
#openbsd
GOOS=openbsd GOARCH=386 go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-openbsd-386.tar.gz" proxy-web && cd ..
GOOS=openbsd GOARCH=amd64 go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-openbsd-amd64.tar.gz" proxy-web && cd ..  
GOOS=openbsd GOARCH=arm go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-openbsd-arm.tar.gz" proxy-web && cd ..
#plan9
GOOS=plan9 GOARCH=386 go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-plan9-386.tar.gz" proxy-web && cd .. 
GOOS=plan9 GOARCH=amd64 go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-plan9-amd64.tar.gz" proxy-web && cd ..
GOOS=plan9 GOARCH=arm go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-plan9-arm.tar.gz" proxy-web && cd .. 
#solaris
GOOS=solaris GOARCH=amd64 go build && mv proxy-web proxy/proxy-web/proxy-web && cd proxy && tar zcfv "../zip/proxy-web-solaris-amd64.tar.gz" proxy-web && cd ..  
#windows
cd proxy/proxy-web
rm -rf proxy-web
cd ..
cd ..
GOOS=windows GOARCH=386 go build -ldflags "-H windowsgui" && mv proxy-web.exe proxy/proxy-web/proxy-web.exe && cd proxy && tar zcfv "../zip/proxy-web-windows-386.tar.gz" proxy-web && cd ..
GOOS=windows GOARCH=amd64 go build -ldflags "-H windowsgui" && mv proxy-web.exe proxy/proxy-web/proxy-web.exe && cd proxy && tar zcfv "../zip/proxy-web-windows-amd64.tar.gz" proxy-web && cd ..

rm -rf proxy-web proxy-web.exe
