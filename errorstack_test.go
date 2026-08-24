package utils

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewErrorStackIsEmpty(t *testing.T) {
	t.Parallel()

	stack := NewErrorStack()
	require.NotNil(t, stack)
	assert.Equal(t, "", stack.Error())
	assert.Nil(t, stack.AsError())
	assert.False(t, stack.HasErrors())
	assert.Equal(t, 0, stack.Len())
}

func TestErrorStackAppendIgnoresNil(t *testing.T) {
	t.Parallel()

	stack := NewErrorStack()
	stack.Append(nil)
	stack.Append(nil)
	assert.False(t, stack.HasErrors())
	assert.Equal(t, 0, stack.Len())
	assert.Nil(t, stack.AsError())
	assert.Equal(t, "", stack.Error())
}

func TestErrorStackAggregatesErrors(t *testing.T) {
	t.Parallel()

	err1 := errors.New("Error 1 occured!")
	err2 := errors.New("Error 2 occured!")
	err3 := errors.New("Error 3 occured!")

	stack := NewErrorStack()
	stack.Append(err1)
	stack.Append(err2)
	stack.Append(nil)
	stack.Append(err3)

	require.True(t, stack.HasErrors())
	assert.Equal(t, 3, stack.Len())

	joined := strings.Join([]string{err1.Error(), err2.Error(), err3.Error()}, "\n")
	assert.Equal(t, joined, stack.Error())

	err := stack.AsError()
	require.NotNil(t, err)
	assert.Equal(t, joined, err.Error())
}

// TestErrorStackAsErrorSupportsErrorsIs verifies that AsError returns a joined
// error that supports errors.Is / errors.As for each contained error, so
// callers can pattern match on sentinel errors.
func TestErrorStackAsErrorSupportsErrorsIs(t *testing.T) {
	t.Parallel()

	sentinel := errors.New("sentinel")
	other := errors.New("other")

	stack := NewErrorStack()
	stack.Append(other)
	stack.Append(sentinel)

	err := stack.AsError()
	require.NotNil(t, err)
	assert.True(t, errors.Is(err, sentinel), "AsError should support errors.Is for contained errors")
	assert.True(t, errors.Is(err, other))
}

// TestErrorStackImplementsErrorInterface makes sure ErrorStack satisfies the
// error interface at compile time and at runtime.
func TestErrorStackImplementsErrorInterface(t *testing.T) {
	t.Parallel()

	var _ error = (*ErrorStack)(nil)

	stack := NewErrorStack()
	stack.Append(errors.New("boom"))
	var err error = stack
	assert.Equal(t, "boom", err.Error())
}
