package util

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"hash"
	"io"
	"os"
	"path/filepath"
)

type ShalStream struct {
	_shal hash.Hash
}

func (obj*ShalStream) Upload(data[]byte) {
	if obj._shal==nil{
		obj._shal=sha1.New()
	}
	obj._shal.Write(data)
}

func (obj*ShalStream) Sum()string {
	return hex.EncodeToString(obj._shal.Sum([]byte("")))
}

func Shal(data[]byte)string{
	_shal:=sha1.New()
	_shal.Write(data)
	return hex.EncodeToString(_shal.Sum([]byte("")))
}

func FileShal(file*os.File)string {
	_shal:=sha1.New()
	io.Copy(_shal,file)
	return hex.EncodeToString(_shal.Sum(nil))
}

func MD5(data[]byte)string {
	_md5:=md5.New()
	_md5.Write(data)
	return hex.EncodeToString(_md5.Sum([]byte("")))
}

func FileMD5(file*os.File)string {
	_md5:=md5.New()
	io.Copy(_md5,file)
	return hex.EncodeToString(_md5.Sum(nil))
}

func PathExists(path string) (bool,error) {
	_,err:=os.Stat(path)
	if err==nil {
		return true,nil
	}
	if os.IsNotExist(err) {
		return false,nil
	}
	return false,err
}

func GetFileSize(filname string) int64 {
	var result int64
	// filepath.Walk(filname,func(path string, info fs.FileInfo, err error) error {
	// 	result=fs.Size()
	// 	return nil
	// })
	filepath.Walk(filname,func(path string, fos.FileInfo, err error) error {
		result=f.Size()
		return nil
	})
	return result
}