package repository

type TransactionInterface interface {
	Get() any
	Commit() error
	Rollback() error
}

type DatabaseInterface interface {
	Close()
	Begin() (TransactionInterface, []error)
	Commit(tx TransactionInterface) []error
	Rollback(tx TransactionInterface) []error
	ApiFaculty() ApiFacultyInterface
}
