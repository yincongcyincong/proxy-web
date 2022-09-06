//go:generate goversioninfo -icon=school.ico
package main

import (
	_ "net/http/pprof"
	"proxy-web/server"
)

//var daemon = flag.Bool("d", true, "default run daemon")

func init() {
	//if !flag.Parsed() {
	//	flag.Parse()
	//}
	//if *daemon {
	//	args := make([]string, 1)
	//	args[0] = "-d=false"
	//	cmd := exec.Command(os.Args[0], args...)
	//	cmd.Start()
	//	fmt.Println("[PID]", cmd.Process.Pid)
	//	os.Exit(0)
	//}
}

func main() {
	server.StartServer()
	//clean()
}

//func clean() {
//	signalChan := make(chan os.Signal, 1)
//	cleanupDone := make(chan bool)
//	signal.Notify(signalChan,
//		os.Interrupt,
//		syscall.SIGHUP,
//		syscall.SIGINT,
//		syscall.SIGTERM,
//		syscall.SIGQUIT)
//	go func() {
//		fmt.Println("close")
//	}()
//	<-cleanupDone
//}
