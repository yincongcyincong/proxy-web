package server

import (
	"net/http"
	"io"
	"proxy-web/utils"
	"html/template"
)

func login(v http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(dir + "/view/login.html")
	if err != nil {
		io.WriteString(v, err.Error())
		return
	}
	t.Execute(v, nil)
}

func doLogin(v http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username, password, err := utils.NewConfig().GetUsernameAndPassword()
	if err != nil {
		v.WriteHeader(http.StatusInternalServerError)
		utils.ReturnJson(err.Error(), "", v)
		return
	}
	sess, _ := globalSessions.SessionStart(v, r)
	newSessionId := sess.SessionID()
	if lock && (sessionId != newSessionId && sessionId != "") {
		v.WriteHeader(http.StatusInternalServerError)
		utils.ReturnJson("已有人登陆", "", v)
		return
	}
	if (r.Form.Get("username") == username) && (r.Form.Get("password") == password) {
		lock = true
		sessionId = sess.SessionID()
		defer sess.SessionRelease(v)
		utils.ReturnJson("success", "", v)
		return
	}

	v.WriteHeader(http.StatusInternalServerError)
	utils.ReturnJson("登陆失败", "", v)
}

func logout(v http.ResponseWriter, r *http.Request){
	r.ParseForm()
	logoutType := r.Form.Get("type")
	if logoutType == "1" {
		lock = false
		sessionId = ""
	} else {
		lock = false
	}

	utils.ReturnJson("success", "", v)
}
