// Package utils provides several helper functions for Go projects.
package utils

import (
	"errors"
	"strings"
)

// ErrorStack handles a collection of errors. Useful e.g. in a loop where you
// want to collect every error that occured instead of stopping at the first
// one.
type ErrorStack struct {
	errorList []error
}

// NewErrorStack returns a new, empty error stack.
func NewErrorStack() *ErrorStack {
	return &ErrorStack{}
}

// Append adds the passed error to the internal list if it is not nil.
func (stack *ErrorStack) Append(err error) {
	if err != nil {
		stack.errorList = append(stack.errorList, err)
	}
}

// Error implements the error interface. Returns all existing errors joined by
// newlines, or an empty string if the stack is empty.
func (stack *ErrorStack) Error() string {
	if len(stack.errorList) == 0 {
		return ""
	}
	parts := make([]string, len(stack.errorList))
	for i, err := range stack.errorList {
		parts[i] = err.Error()
	}
	return strings.Join(parts, "\n")
}

// HasErrors returns true if there's at least one error in the stack.
func (stack *ErrorStack) HasErrors() bool {
	return len(stack.errorList) > 0
}

// Len returns the number of errors in the stack.
func (stack *ErrorStack) Len() int {
	return len(stack.errorList)
}

// AsError returns nil if the stack is empty. Otherwise it returns an error
// built with errors.Join so callers can inspect the individual errors with
// errors.Is and errors.As.
func (stack *ErrorStack) AsError() error {
	if len(stack.errorList) == 0 {
		return nil
	}
	return errors.Join(stack.errorList...)
}
