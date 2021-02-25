package nero

// Tx is an interface that wraps the Commit and Rollback method
type Tx interface {
	Commit() error
	Rollback() error
}
