package server

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
	"github.com/yincongcyincong/proxy-web/utils"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/yincongcyincong/proxy-web/lib/goproxy/sdk/android-ios"
)

func add(v http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	command := r.Form.Get("command")
	autoStart := r.Form.Get("auto")
	keyFile := r.Form.Get("key")
	crtFile := r.Form.Get("crt")
	log := r.Form.Get("log")

	serviceId, err := utils.SaveParams(name, command, autoStart, keyFile, crtFile, log)
	if err != nil {
		v.WriteHeader(http.StatusInternalServerError)
		utils.ReturnJson(err.Error(), "", v)
		return
	}

	data := make(map[string]interface{})
	data["id"] = serviceId
	data["command"] = command
	data["auto_start"] = autoStart
	data["name"] = name
	data["log"] = log
	data["status"] = "未开启"
	utils.ReturnJson("success", data, v)
}

func show(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles(dir + "/view/index.html")
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		autoStart := utils.NewConfig().GetAutoStart()
		isProxy := utils.NewConfig().GetProxySetting()
		proxySetting, _ := utils.GetProxy()
		var ip, port string
		if _, ok := proxySetting["ip"]; ok {
			ip = proxySetting["ip"]
		}
		if _, ok := proxySetting["port"]; ok {
			port = proxySetting["port"]
		}

		proxyVersion := proxy.Version()
		data := map[string]interface{}{"auto_start": autoStart, "proxy_version": proxyVersion, "version": version, "proxy": isProxy, "ip": ip, "port": port}

		t.Execute(w, data)
	}
}

func getData(v http.ResponseWriter, r *http.Request) {
	var data interface{}
	var err error
	r.ParseForm()
	id := r.Form.Get("id")

	if id == "0" {
		data, err = utils.GetAllParams()
	} else {
		data, err = utils.GetParamsById(id)
	}
	if err != nil {
		v.WriteHeader(http.StatusInternalServerError)
		utils.ReturnJson(err.Error(), "", v)
		return
	}
	utils.ReturnJson("success", data, v)
}

func link(v http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		var command string
		var err error
		id := r.Form.Get("id")
		command, err = getCommand(id)
		if err != nil {
			v.WriteHeader(http.StatusInternalServerError)
			utils.ReturnJson(err.Error(), "", v)
			return
		}
		fmt.Println(command)
		errStr := proxy.Start(id, command)
		if errStr != "" {
			v.WriteHeader(http.StatusInternalServerError)
			utils.ReturnJson(errStr, "", v)
			return
		}
		utils.ChangeParameterDataById(id, "已开启")
		utils.ReturnJson("success", "", v)
	}
}

func getCommand(id string) (command string, err error) {
	parameter, err := utils.GetParamsById(id)
	if err != nil {
		return "", err
	}

	command += parameter["command"].(string)
	command = strings.Replace(command, "\n", " ", -1)
	command = strings.Replace(command, "\r", " ", -1)
	command = strings.Replace(command, "  ", " ", -1)

	if parameter["key_file"].(string) != "" {
		command += " -K " + parameter["key_file"].(string)
	}
	if parameter["crt_file"].(string) != "" {
		command += " -C " + parameter["crt_file"].(string)
	}
	if parameter["log"] == "是" {
		command += " --log ./log/" + parameter["id"].(string) + ".log"
	}
	s, err := os.Stat("./log/")
	if err != nil || !s.IsDir() {
		os.Mkdir("./log/", os.ModePerm)
	}
	return command, nil
}

func close(v http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.Form.Get("id")
	if id == "undefined" {
		v.WriteHeader(http.StatusInternalServerError)
		utils.ReturnJson("id not found", "", v)
		return
	}
	err := utils.ChangeParameterDataById(id, "未开启")
	if err != nil {
		v.WriteHeader(http.StatusInternalServerError)
		utils.ReturnJson(err.Error(), "", v)
		return
	}
	proxy.Stop(id)
	utils.ReturnJson("success", "", v)
	return
}

func uploade(v http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		file, head, err := r.FormFile("file")
		fileSuffix := path.Ext(head.Filename)
		if err != nil {
			v.WriteHeader(http.StatusInternalServerError)
			utils.ReturnJson(err.Error(), "", v)
			return
		}
		defer file.Close()
		t := time.Now().Unix()
		fw, err := os.Create(dir + "/upload/" + strconv.FormatInt(t, 10) + fileSuffix)
		defer fw.Close()
		if err != nil {
			v.WriteHeader(http.StatusInternalServerError)
			utils.ReturnJson(err.Error(), "", v)
			return
		}
		_, err = io.Copy(fw, file)
		if err != nil {
			v.WriteHeader(http.StatusInternalServerError)
			utils.ReturnJson(err.Error(), "", v)
			return
		}
		name := fw.Name()
		utils.ReturnJson("", name, v)
		return
	}
}

func update(v http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.Form.Get("id")
	name := r.Form.Get("name")
	command := r.Form.Get("command")
	autoStart := r.Form.Get("auto")
	keyFile := r.Form.Get("key")
	crtFile := r.Form.Get("crt")
	log := r.Form.Get("log")

	err := utils.UpdateParams(id, name, command, autoStart, keyFile, crtFile, log)
	if err != nil {
		v.WriteHeader(http.StatusInternalServerError)
		utils.ReturnJson(err.Error(), "", v)
		return
	}
	utils.ReturnJson("success", "", v)
}

func deleteParameter(v http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.Form.Get("id")
	err := utils.DeleteParam(id)
	if err != nil {
		v.WriteHeader(http.StatusInternalServerError)
		utils.ReturnJson(err.Error(), "", v)
	}
	utils.ReturnJson("success", "", v)
}

func saveSetting(v http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	auto := r.Form.Get("auto")
	proxy := r.Form.Get("proxy")
	ip := r.Form.Get("ip")
	port := r.Form.Get("port")

	config := utils.NewConfig()
	isAutoStart := config.GetAutoStart()
	isProxy := config.GetProxySetting()

	// 判断是否开启全局代理
	if proxy == "proxy" {
		if !isProxy {
			err := utils.StartProxy(ip, port)
			if err != nil {
				v.WriteHeader(http.StatusInternalServerError)
				utils.ReturnJson("修改配置失败,请使用root权限操作", err.Error(), v)
				return
			}
		}

	} else {
		if isProxy {
			utils.StopProxy(ip, port)
		}
	}

	switch runtime.GOOS {
	case "windows":

		if auto == "auto" {
			if !isAutoStart {
				command := dir + `/config/autostart.exe enable -k proxy-web -n proxy-web -c`
				commandSlice := strings.Split(command, " ")
				commandSlice = append(commandSlice, dir+`/proxy-web.exe c:`)
				cmd := exec.Command(commandSlice[0], commandSlice[1:]...)
				output, _ := cmd.CombinedOutput()
				outputStr := string(output)
				if !strings.Contains(outputStr, "Done") {
					v.WriteHeader(http.StatusInternalServerError)
					utils.ReturnJson("修改配置失败,请使用root权限操作", outputStr, v)
					return
				}
				is_success := utils.NewConfig().UpdateAutoStart("true")
				if !is_success {
					v.WriteHeader(http.StatusInternalServerError)
					utils.ReturnJson("修改配置失败,请使用root权限操作", "", v)
					return
				}
			}

		} else {
			if isAutoStart {
				command := dir + `/config/autostart.exe disable -k proxy-web`
				commandSlice := strings.Split(command, " ")
				cmd := exec.Command(commandSlice[0], commandSlice[1:]...)
				output, _ := cmd.CombinedOutput()
				outputStr := string(output)
				if !strings.Contains(outputStr, "Done") {
					v.WriteHeader(http.StatusInternalServerError)
					utils.ReturnJson("修改配置失败,请使用root权限操作", outputStr, v)
					return
				}
				is_success := utils.NewConfig().UpdateAutoStart("false")
				if !is_success {
					v.WriteHeader(http.StatusInternalServerError)
					utils.ReturnJson("修改配置失败,请使用root权限操作", "", v)
					return
				}
			}
		}

	case "darwin":
		if auto == "auto" {
			if !isAutoStart {
				command := dir + `/config/autostart enable -k proxy -n proxy -c`
				commandSlice := strings.Split(command, " ")
				commandSlice = append(commandSlice, dir+"/proxy-web")
				cmd := exec.Command(commandSlice[0], commandSlice[1:]...)
				output, err := cmd.CombinedOutput()
				if err != nil {
					v.WriteHeader(http.StatusInternalServerError)
					utils.ReturnJson("修改配置失败,请使用root权限操作", string(output), v)
					return
				}
				is_success := utils.NewConfig().UpdateAutoStart("true")
				if !is_success {
					v.WriteHeader(http.StatusInternalServerError)
					utils.ReturnJson("修改配置失败,请使用root权限操作", "", v)
					return
				}
			}
		} else {
			if isAutoStart {
				command := dir + `/config/autostart disable -k proxy`
				commandSlice := strings.Split(command, " ")
				cmd := exec.Command(commandSlice[0], commandSlice[1:]...)
				output, _ := cmd.CombinedOutput()
				is_success := utils.NewConfig().UpdateAutoStart("false")
				if !is_success {
					v.WriteHeader(http.StatusInternalServerError)
					utils.ReturnJson("修改配置失败,请使用root权限操作", string(output), v)
					return
				}
			}
		}
	case "linux":
		if auto == "auto" {
			if !isAutoStart {
				data := `#!/bin/sh
` + dir + `/proxy-web`
				err := ioutil.WriteFile(dir+"/config/autostart.sh", []byte(data), 0777)
				if err != nil {
					v.WriteHeader(http.StatusInternalServerError)
					utils.ReturnJson("修改配置失败,请使用root权限操作", "", v)
					return
				}
				fd, err := os.OpenFile("/etc/crontab", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
				if err != nil {
					v.WriteHeader(http.StatusInternalServerError)
					utils.ReturnJson("修改配置失败,请使用root权限操作", "", v)
					return
				}
				defer fd.Close()
				fileData, _ := ioutil.ReadAll(fd)
				if !strings.Contains(string(fileData), dir+"/config/autostart.sh") {
					fd.Write([]byte(`@reboot root ` + dir + `/config/autostart.sh
`))
				}

				is_success := utils.NewConfig().UpdateAutoStart("true")
				if !is_success {
					v.WriteHeader(http.StatusInternalServerError)
					utils.ReturnJson("修改配置失败,请使用root权限操作", "", v)
					return
				}
			}
		} else {
			if isAutoStart {
				os.Remove(dir + "/config/autostart.sh")
				is_success := utils.NewConfig().UpdateAutoStart("false")
				if !is_success {
					v.WriteHeader(http.StatusInternalServerError)
					utils.ReturnJson("修改配置失败,请使用root权限操作", "", v)
					return
				}
			}
		}

	}

	// 修改数据
	if proxy == "proxy" {
		if !isProxy {
			is_success := utils.NewConfig().UpdateProxy("true")
			if !is_success {
				v.WriteHeader(http.StatusInternalServerError)
				utils.ReturnJson("修改配置失败,请使用root权限操作", "", v)
				return
			}
			err := utils.UpdateProxy(ip, port)
			if err != nil {
				v.WriteHeader(http.StatusInternalServerError)
				utils.ReturnJson(err.Error(), "", v)
				return
			}
		}
	} else {
		if isProxy {
			is_success := utils.NewConfig().UpdateProxy("false")
			if !is_success {
				v.WriteHeader(http.StatusInternalServerError)
				utils.ReturnJson("修改配置失败,请使用root权限操作", "", v)
				return
			}
		}
	}

	utils.ReturnJson("success", "", v)
	return
}
