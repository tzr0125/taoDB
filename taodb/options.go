package taodb

import (
	"os"
)
type Options struct {
	// 数据库路径
	DirPath string

}


var DefaultOptions = Options{
	DirPath: tempDBDir(),
}

func tempDBDir() string {
	dir, _ := os.MkdirTemp("", "taodb-temp")
	return dir
}