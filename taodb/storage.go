package taodb

import (
	"os"
	"strconv"
	"sync"
	"fmt"
)

type Storage struct {
	mu sync.RWMutex 
	// 是否为活跃，只有活跃才会被开启
	isactive bool
	fileId   uint32
	filePath     string
	nextPos 	uint32
}

/*
Storage 设计
1. 创建
创建时创建一个在数据库目录下名为fileId.tao的文件，并激活Storage作为新的activeStorage
2. 转为只读
在占用空间达到1GB后isactive = false，同时要求db创建新的activeStorage
3. merge 
遍历文件中数据，询问该数据在Index中位置和是否已过期，如果未过期且位置是当前位置，则合并到hintStorage中，否则丢弃。
*/

// 根据ID获取storage
func NewEmtpyStorage(filePath string, fileId uint32) (*Storage, error) {
	file, err := os.Create(filePath + "/" + strconv.FormatInt(int64(fileId), 10) + ".tao")
	if err != nil {
		return nil, err
	}
	file.Close()
	return &Storage{
		filePath:     filePath + "/" + strconv.FormatInt(int64(fileId), 10) + ".tao",
		isactive: true, // 新的文件一定是active的
		fileId:   fileId,
	}, nil
}

func (sto *Storage) NextPos() uint32 {
	return sto.nextPos
}


func (sto *Storage) WriteEntry() error {
	if !sto.isactive{
		return fmt.Errorf(ERROR_STRORAGE_NOT_ACTIVE)
	}

	// 加锁，准备写入
	sto.mu.Lock()
	defer sto.mu.Unlock()
	// 1. 在文件尾部写入

	// 2. 写在hashtable上
	return nil
	
}


func (sto *Storage) Merge() error {
	return nil
}