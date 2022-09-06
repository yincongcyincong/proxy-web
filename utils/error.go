package utils

import (
	"encoding/json"
	"net/http"
)

type Result struct {
	Msg string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func ReturnJson(msg string, data interface{}, v http.ResponseWriter) {
	r := Result{Msg: msg, Data: data}
	json.NewEncoder(v).Encode(r)
}
