package utils

import (
	"errors"
	"strings"
)

// ErrorStack handles a collection of errors. Usefull e.g. in loop where you want to
// get a list of all occured errors.
type ErrorStack struct {
	errorList []error
}

// NewErrorStack returns a new, empty error stack.
func NewErrorStack() *ErrorStack {
	return &ErrorStack{errorList: []error{}}
}

// Append will add passed error to internal list if it's no nil.
func (stack *ErrorStack) Append(err error) {
	if err != nil {
		stack.errorList = append(stack.errorList, err)
	}
}

// Error used to fulfill error interface. Returns all existing errors composed by new line
// or an empty string.
func (stack *ErrorStack) Error() string {
	errorStrings := []string{}
	for _, err := range stack.errorList {
		errorStrings = append(errorStrings, err.Error())
	}
	return strings.Join(errorStrings, "\n")
}

// AsError returns nil of there's no error in internal stack or a new erros
// which contains all existing errors, composed by new line.
func (stack *ErrorStack) AsError() error {
	if len(stack.errorList) == 0 {
		return nil
	}
	return errors.New(stack.Error())
}
