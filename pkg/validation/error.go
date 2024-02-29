package validation

import "errors"

type ValidationError struct {
	error
	messageBag *MessageBag
}

func (e ValidationError) Errors() *MessageBag {
	return e.messageBag
}

func (e ValidationError) Error() string {
	return e.error.Error()
}

func NewValidationError(err error, messageBag *MessageBag) ValidationError {
	return ValidationError{
		error: errors.New("Invalid input"),
		messageBag: messageBag,
	}
}
