package handler

import (
	"io"
	"io/ioutil"
	"net/http"
)

func UploadHandler(w http.ResponseWriter, r*http.Request)  {
	if r.Method == "GET" {
		data,err := ioutil.ReadFile("./static/index.html")
		if err != nil {
			io.WriteString(w, "internal error")
			return
		}
		io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		// 接收文件流及存储到本地目录
	}
}