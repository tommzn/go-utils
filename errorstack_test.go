package utils

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ErrorsStackTestSuite struct {
	suite.Suite
}

func TestErrorsStackTestSuite(t *testing.T) {
	suite.Run(t, new(ErrorsStackTestSuite))
}

func (suite *ErrorsStackTestSuite) TestErrorHandling() {

	errorStack := NewErrorStack()
	suite.Equal("", errorStack.Error())
	suite.Nil(errorStack.AsError())
	suite.False(errorStack.HasErrors())
	suite.Equal(0, errorStack.Len())

	err1 := errors.New("Error 1 occured!")
	err2 := errors.New("Error 2 occured!")
	err3 := errors.New("Error 3 occured!")
	errorStack.Append(err1)
	errorStack.Append(err2)
	errorStack.Append(nil)
	errorStack.Append(err3)
	suite.Equal(3, errorStack.Len())
	suite.True(errorStack.HasErrors())

	errorsAsString := strings.Join([]string{err1.Error(), err2.Error(), err3.Error()}, "\n")
	suite.Equal(errorsAsString, errorStack.Error())
	err := errorStack.AsError()
	suite.NotNil(err)
	suite.Equal(errorsAsString, err.Error())
}
