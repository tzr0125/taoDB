package taodb


import (
	"sync"
)
const (
	ERROR_STRORAGE_NOT_ACTIVE = "storage is not active"
	ERROR_META_NOT_SET = "meta is not set"
	
)


var (
	TransactionId = uint32(0)
	TransactionMutex = &sync.Mutex{}
	
)
