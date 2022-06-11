package datastore

import "fmt"

var _ TransactionStacker = &TransactionStackerImpl{}

type TransactionStacker interface {
	Peek() *transaction
	PushTransaction()
	PopTransaction() error
	Commit() error
}

// Maintains a list of active/suspended transactions.
type TransactionStackerImpl struct {
	top  *transaction
	size int
}

type transaction struct {
	// Each transaction has its own local store.
	store map[string]string
	next  *transaction
}

// Push a new transaction.
func (ts *TransactionStackerImpl) PushTransaction() {
	// Push a new transaction, this is the current active transaction.
	temp := transaction{store: make(map[string]string)}
	temp.next = ts.top
	ts.top = &temp
	ts.size++
}

// Pop.
func (ts *TransactionStackerImpl) PopTransaction() error {
	if ts.top == nil {
		return fmt.Errorf("error: no active transaction.")
	} else {
		node := &transaction{}
		ts.top = ts.top.next
		node.next = nil
		ts.size--
	}
	return nil
}

// Peek.
func (ts *TransactionStackerImpl) Peek() *transaction {
	return ts.top
}

// Commit.
func (ts *TransactionStackerImpl) Commit() error {
	// Fetch active transaction from stack.
	activeTransaction := ts.Peek()
	if activeTransaction != nil {
		for key, val := range activeTransaction.store {
			// INSERT: db store.
			if activeTransaction.next != nil {
				// Update the parent transaction.
				activeTransaction.next.store[key] = val
			}
		}
		ts.PushTransaction()
		return nil
	} else {
		return fmt.Errorf("error: no active transaction")
	}
}

// RollbackTransaction clears all keys SET within a transaction.
func (ts *TransactionStackerImpl) RollbackTransaction() {
	if ts.top == nil {
		fmt.Printf("error: no active transaction/\n")
	} else {
		for key := range ts.top.store {
			delete(ts.top.store, key)
		}
		ts.top = ts.top.next
	}
}

func NewTransaction() TransactionStacker {
	return newTransaction()
}

func newTransaction() *TransactionStackerImpl {
	return &TransactionStackerImpl{}
}
