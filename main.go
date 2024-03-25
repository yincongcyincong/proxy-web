//go:generate goversioninfo -icon=school.ico
package main

import (
	"github.com/yincongcyincong/proxy-web/server"
	_ "net/http/pprof"
)


func main() {
	server.StartServer()
	//clean()
}


