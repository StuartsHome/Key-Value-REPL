package errors

import (
	"errors"
	"fmt"
)

var ErrKeyNotFound = errors.New("not found")

var ErrIncorrectNumberArguments = errors.New("incorrect number of arguments")

var ErrNoActiveTransaction = errors.New("no active transaction")

func ErrorKeyNotFound(v interface{}) error {
	return fmt.Errorf("error: key %v %w\n", v, ErrKeyNotFound)
}

func ErrorIncorrectNumberArguments() error {
	return fmt.Errorf("error: %w\n", ErrIncorrectNumberArguments)
}

func ErrorNoActiveTransaction() error {
	return fmt.Errorf("error: %w\n", ErrNoActiveTransaction)
}
