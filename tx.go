package nero

// Tx is a transaction type
type Tx interface {
	Commit() error
	Rollback() error
}
