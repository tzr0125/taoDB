package taodb

// 这个文件实现了一个基础存储单元Entry要存储的内容

import (
	"errors"
	"time"
)

// 原信息，要和所有的东西一起存
type Meta struct {
	pos       uint32 // 位置 //位置有什么用？
	timestamp int64  // timestamp
	ksz       uint32 // key长度
	vsz       uint32 // value长度
	ttl       uint32 // timestamp
}

func (m *Meta) Pos() uint32 { return m.pos }

func (m *Meta) Timestamp() int64 { return m.timestamp }

func (m *Meta) KeySize() uint32 { return m.ksz }

func (m *Meta) ValueSize() uint32 { return m.vsz }

func (m *Meta) TTL() uint32 { return m.ttl }

type Entry struct {
	key   []byte
	value []byte
	meta  *Meta
}

func (e *Entry) Key() []byte {
	return e.key
}

func (e *Entry) Value() []byte {
	return e.value
}

func (e *Entry) Meta() *Meta {
	return e.meta
}

func (e *Entry) CalculateMeta() {
	e.meta = &Meta{
		pos:       uint32(len(e.key)),
		timestamp: time.Now().Unix(),
		ksz:       uint32(len(e.key)),
		vsz:       uint32(len(e.value)),
	}
	// e.meta.pos = uint32(e.meta.ksz + e.meta.vsz + e.meta.csc + e.meta.timestamp)
}

type IMetaBuilder interface {
	newMeta(pos uint32, ksz uint32, vsz uint32)
	setTTL(time.Duration)
	getMeta() (*Meta, error)
}

type MetaBuilder struct {
	meta *Meta
}

func (m *MetaBuilder) newMeta(pos uint32, ksz uint32, vsz uint32) {
	m.meta = &Meta{
		pos:       pos,
		timestamp: time.Now().Unix(),
		ksz:       ksz,
		vsz:       vsz,
	}
}

func (m *MetaBuilder) getMeta() (*Meta, error) {
	if m.meta == nil {
		return nil, errors.New("meta havent been initialized")
	}
	return m.meta, nil
}

func GetEntry(key, value []byte) *Entry {
	entry := &Entry{
		key:   key,
		value: value,
	}
	// entry.CalculateMeta()
	return entry
}
