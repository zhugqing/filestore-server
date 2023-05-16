package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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
		file,head,err := r.FormFile("file")
		if err != nil {
			fmt.Printf("failed to get data, err:%s",err.Error())
			return
		}
		defer file.Close()
		fmt.Println(head.Filename)

		newFile,err:=os.Create("./static/"+head.Filename)
		if err != nil {
			fmt.Printf("Failed to create file,err:%s",err.Error())
			return
		}
		defer newFile.Close()

		_,err=io.Copy(newFile,file)
		if err != nil {
			fmt.Printf("Failed to save data into file,err:%s",err.Error())
			return
		}
		http.Redirect(w,r,"/file/upload/suc",http.StatusFound)
	}
}

func UploadSucHandler(w http.ResponseWriter,r*http.Request) {
	io.WriteString(w,"Upload finished!")
}