package taodb

import (
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
		index:        make(Index),
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

func (db *DB) Close() error {
    return nil
}

// 写入
func (db *DB) Put(key string, value string) error {
	trans := NewTransaction()
	err := trans.Put(key, value)
	if err != nil{
		trans.Rollback()
		return err
	}
	err = trans.Commit()
	return err
}


func (db *DB) Get(key string) (string, error) {
	trans := NewTransaction()
	val, err := trans.Get(key)
	if err!= nil{
        return "", err
    }
	return val, nil
}


func (db *DB) Delete(key string) error {
	trans := NewTransaction()
	return trans.Delete(key)
}


func (db *DB) Exist(key string) (bool, error) {
	trans := NewTransaction()
	return trans.Exist(key)
}