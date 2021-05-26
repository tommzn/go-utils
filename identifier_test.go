package utils

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/suite"
)

// IdentifierTestSuite runs all test for id relaated functions.
type IdentifierTestSuite struct {
	suite.Suite
}

// TestIdentifierTestSuite executes all tests from IdentifierTestSuite.
func TestIdentifierTestSuite(t *testing.T) {
	suite.Run(t, new(IdentifierTestSuite))
}

// TestGenrateIds test generating new ids.
func (suite *IdentifierTestSuite) TestGenrateIds() {

	id1 := NewId()
	suite.Len(id1, 36)
	_, err := uuid.FromString(id1)
	suite.Nil(err)
}

// Test asserting ids in UUID version 4.
func (suite *IdentifierTestSuite) TestValidaeIds() {

	suite.True(IsId(uuid.NewV4().String()))
	suite.False(IsId("xxx"))
}
