package taodb

/*
	事务
*/
type Transaction struct {
	id uint32 // 事务id
	db *DB
	index TempIndex // 

}

func NewTransaction() *Transaction {
	TransactionMutex.Lock()
	defer TransactionMutex.Unlock()
	transaction := &Transaction{
		id: TransactionId,
	}
	TransactionId++
	return transaction
}

// 以下未完成

func (transaction *Transaction) Put(key, value string) error {
	return nil
}

func (transaction *Transaction) Get(key string) (string, error) {
    return "", nil
}

func (transaction *Transaction) Delete(key string) error {
	return nil
}

func (transaction *Transaction) Exist(key string) (bool, error) {
    return false, nil
}

func (transaction *Transaction) Commit() error {
	return nil
}

func (transaction *Transaction) Rollback() error {
	return nil
}