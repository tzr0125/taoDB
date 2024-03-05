package taodb

import (
	"fmt"
	"os"
	"sync"
)

type DB struct {
	dataFiles    map[uint32]*Storage // 已经存储好但没有merge的文件
	hintFile     map[uint32]*Storage // 已经merge好的文件
	activateFile *Storage            //正在活跃的文件
	mu           sync.RWMutex
	index        Index   // 索引，key与文件的对应
	options      Options //
}

func Open(options Options) (*DB, error) {
	db := &DB{
		dataFiles:    nil,
		hintFile:     nil,
		activateFile: nil,
		mu:           sync.RWMutex{},
		index:        DefaultIndex,
		options:      options,
	}

	storage, err := db.newFile()
	if err != nil {
		return nil, err
	}
	db.activateFile = storage
	return db, nil
}

// 取新的空文件，新的空文件一定是active的
func (db *DB) newFile() (*Storage, error) {
	storage, err := NewEmtpyStorage(db.options.DirPath, uint32(len(db.dataFiles)+len(db.hintFile)))
	return storage, err
}

// 写入
func (db *DB) Put(key []byte, value []byte) error {
	// 首先创建新的entry

	// 加入文件中，并记录index

	return nil
}

func (db *DB) Get(key []byte) ([]byte, error) {
	// 读取index
	pos, ok := db.index.keyhash[string(key)]
	if !ok {
		return []byte(""), nil
	}

	// 寻找文件
	// 先从merge里面找
	sto, ok := db.hintFile[pos.fileId]
	if !ok {
		// 再从dataFiles中找
		sto, ok = db.dataFiles[pos.fileId]
		if !ok {
			// 最后从active中找
			if db.activateFile.fileId == pos.fileId {
				sto = db.activateFile
			} else {
				return nil, fmt.Errorf("没有找到表")
			}

		}
	}

	// !:这里直接将file整个读到内存中，绝对不能这样
	file, err := os.ReadFile(sto.file)
	if err!= nil {
        return nil, fmt.Errorf("文件打开失败")
    }
	return file[pos.vpos: pos.vpos + pos.vsz], nil
}
