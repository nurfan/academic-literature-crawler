package errors

import (
	"errors"
)

// Error stsruct
type Error struct{}

// NotFoundError is for something that not tound
func (e *Error) NotFoundError(model string) error {
	return errors.New("The " + model + " resounce not found")
}

// UpdateError indicates that the update is failed and need to be fixed
func (e *Error) UpdateError() error {
	return errors.New("Update failed")
}

// CreateError indicates that the create is failed and need to be fixed
func (e *Error) CreateError() error {
	return errors.New("Create failed")
}

// NotEnoughBalance indicates that the create is failed and need to be fixed
func (e *Error) NotEnoughBalance() error {
	return errors.New("Not Enough Balance")
}

// ProcessingError indicates that the create is custome failed message
func (e *Error) ProcessingError(message string) error {
	return errors.New(message)
}

// DuplicateTransaction indicates that transcation already exist
func (e *Error) DuplicateTransaction() error {
	return errors.New("Duplicate Transaction")
}

// TransactionDeclined indicates that transaction is declined by user
func (e *Error) TransactionDeclined() error {
	return errors.New("Transaction Declined")
}

// TimeoutRequest indicates the request exceeds the time limit
func (e *Error) TimeoutRequest() error {
	return errors.New("Timeout Request")
}

// AlreadySuccess indicates that transaction is already success
func (e *Error) AlreadySuccess() error {
	return errors.New("Transaction already success")
}
