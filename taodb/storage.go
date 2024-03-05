package taodb

import (
	"os"
	"strconv"
	"sync"
)

type Storage struct {
	// 锁控制读，如果已经merge但依然有事务在读，则暂时不删除，否则可以删除。
	mu sync.RWMutex
	// 是否为活跃，只有活跃才会被开启
	isactive bool
	fileId   uint32
	file     string
}

// 根据ID获取storage
func NewEmtpyStorage(filePath string, fileId uint32) (*Storage, error) {
	file, err := os.Create(filePath + "/" + strconv.FormatInt(int64(fileId), 10) + ".tao")
	if err != nil {
		return nil, err
	}
	file.Close()
	return &Storage{
		file:     filePath + "/" + strconv.FormatInt(int64(fileId), 10) + ".tao",
		isactive: true, // 新的文件一定是active的
		fileId:   fileId,
	}, nil
}
