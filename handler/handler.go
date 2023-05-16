package handler

import (
	"filestore-server/meta"
	"filestore-server/util"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
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

		FileMeta:=meta.FileMeta{
			FileNmae: head.Filename,
			Location: "./static/"+head.Filename,
			UploadAt: time.Now().Format("2006-01-01 13:16:08"),
		}

		newFile,err:=os.Create(FileMeta.Location)
		if err != nil {
			fmt.Printf("Failed to create file,err:%s",err.Error())
			return
		}
		defer newFile.Close()

		FileMeta.FileSize,err=io.Copy(newFile,file)
		if err != nil {
			fmt.Printf("Failed to save data into file,err:%s",err.Error())
			return
		}

		newFile.Seek(0,0)
		FileMeta.FileShal = util.FileShal(newFile)
		meta.UploadFileMeta(FileMeta)

		http.Redirect(w,r,"/file/upload/suc",http.StatusFound)
	}
}

func UploadSucHandler(w http.ResponseWriter,r*http.Request) {
	io.WriteString(w,"Upload finished!")
}