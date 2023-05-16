package main

import (
	"filestore-server/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload",handler.UploadHandler)
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		fmt.Printf("Failed toStart server, err:%s",err.Error())
	}
}