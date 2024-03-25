package server

import (
	"fmt"
	proxy "github.com/yincongcyincong/proxy-web/lib/goproxy/sdk/android-ios"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego/session"
	"github.com/yincongcyincong/proxy-web/utils"
)

var globalSessions *session.Manager
var version = "v2.0"
var lock = false
var sessionId string
var dir string

func basicAuth(handler func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess, _ := globalSessions.SessionStart(w, r)
		newSessionId := sess.SessionID()
		if sessionId != newSessionId {
			login(w, r)
			return
		}
		handler(w, r)
	})
}

func StartServer() {
	// 文件路径
	dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	dir = strings.Replace(dir, "\\", "/", -1)

	// 启动一个websocket，判断是否有人登陆
	//go StartWebscoket()
	SetProxy()
	AutoStart()
	InitShowLog()
	initSession()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(dir+"/static"))))
	http.Handle("/", basicAuth(show))
	http.HandleFunc("/add", add)
	http.HandleFunc("/update", update)
	http.HandleFunc("/close", close)
	http.HandleFunc("/link", link)
	http.HandleFunc("/getData", getData)
	http.HandleFunc("/uploade", uploade)
	http.HandleFunc("/delete", deleteParameter)
	http.HandleFunc("/saveSetting", saveSetting)
	http.HandleFunc("/login", login)
	http.HandleFunc("/doLogin", doLogin)
	http.HandleFunc("/logout", logout)
	//http.Handle("/keygen", basicAuth(keygen))
	port, err := utils.NewConfig().GetServerPort()
	if err != nil {
		log.Fatal("get port failure: ", err)
	}
	fmt.Println("proxy-web: 127.0.0.1" + port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("listen port failure", err)
	}
}

func AutoStart() {
	datas, err := utils.InitParams()
	if err != nil {
		return
	}
	for _, data := range datas {
		var command string
		command += data["command"].(string)
		command = strings.Replace(command, "\n", "", -1)
		command = strings.Replace(command, "\r", "", -1)
		command = strings.Replace(command, "  ", " ", -1)
		if data["key_file"].(string) != "" {
			command += " -K " + data["key_file"].(string)
		}
		if data["crt_file"].(string) != "" {
			command += " -C " + data["crt_file"].(string)
		}
		if data["log"] == "yes" {
			command += " --log " + dir + "/log/" + data["id"].(string) + ".log"
		}
		s, err := os.Stat(dir + "/log/")
		if err != nil || !s.IsDir() {
			os.Mkdir(dir+"/log/", os.ModePerm)
		}
		go autoRunCommand(data["id"].(string), command)
	}
}

func autoRunCommand(id, command string) {
	fmt.Println(command)
	errStr := proxy.Start(id, command)
	if errStr != "" {
		utils.ChangeParameterDataById(id, "close")
	}
}

func initSession() {
	sessionConfig := &session.ManagerConfig{
		CookieName:      "sessionid",
		EnableSetCookie: true,
		Gclifetime:      360000,
		Maxlifetime:     360000,
		Secure:          false,
		CookieLifeTime:  360000,
		ProviderConfig:  dir + "/tmp",
	}
	globalSessions, _ = session.NewManager("file", sessionConfig)
	go globalSessions.GC()
}

func SetProxy() {
	data, err := utils.GetProxy()
	if err != nil {
		return
	}
	proxy := utils.NewConfig().GetProxySetting()
	if !proxy {
		return
	}
	utils.StartProxy(data["ip"], data["port"])
}

//func StartWebscoket() {
//	http.Handle("/websocket", websocket.Handler(svrConnHandler))
//	log.Fatal(http.ListenAndServe(":8222", nil))
//}
//
//func svrConnHandler(conn *websocket.Conn) {
//	request := make([]byte, 128)
//	defer conn.Close()
//	readLen, err := conn.Read(request)
//	if err != nil {
//		return
//	}
//
//	if string(request[:readLen]) == "close" {
//		lock = false
//	} else {
//		lock = true
//	}
//
//}
