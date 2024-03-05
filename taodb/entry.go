package taodb

// 这个文件实现了一个基础存储单元Entry要存储的内容

import (
	"math"
	"time"
)

type Meta struct {
	operationType uint8 // 操作类型
	version       uint32 // 事务代号，只有commit的事务才会生效
	keySize       uint32 // key长度
	valueSize       uint32 // value长度
	expireDate       int64 // timestamp, 为0则永不过期
}


const (
	PUT = iota
	UPDATE = iota
	DELETE = iota
	COMMIT = iota
	ROLLBACK = iota

)

func (m *Meta) KeySize() uint32 { return m.keySize }

func (m *Meta) ValueSize() uint32 { return m.valueSize }

type Entry struct {
	key   string
	value string
	meta  *Meta
}

func (e *Entry) defaultMeta(pos uint32, version uint32, operationType uint8) {
	e.meta = &Meta{
		keySize:       uint32(len(e.key)),
		valueSize:       uint32(len(e.value)),
		expireDate:       math.MaxInt64,
		version:       version,
		operationType: operationType,
	}
}

func newEntry(key, value string, pos, version uint32,operationType uint8) *Entry {
	entry := &Entry{
		key:   key,
		value: value,
		meta: nil,
	}
	entry.defaultMeta(pos, version, operationType)
	return entry
}

func (e *Entry) setTTL(ttl uint32)  {
	e.meta.expireDate = time.Now().Unix() + int64(ttl)
}