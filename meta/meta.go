package meta

type FileMeta struct {
	FileShal string
	FileNmae string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

// 新增更新文件元信息
func UploadFileMeta(fmeta FileMeta) {
	fileMetas[fmeta.FileShal] = fmeta
}

// 获取文件元信息
func GetFileMeta(fileShal string) FileMeta {
	return fileMetas[fileShal]
}