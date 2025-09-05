package utils

import (
	"testing"

	"github.com/fossoreslp/uuid"
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

// TestGenerateIds test generating new ids.
func (suite *IdentifierTestSuite) TestGenerateIds() {

	id1 := NewId()
	suite.Len(id1, 36)
	_, err := uuid.Parse(id1)
	suite.Nil(err)

	id2 := NewV7Id()
	suite.Len(id2, 36)
	_, err = uuid.Parse(id2)
	suite.Nil(err)
}

// Test asserting ids in UUID version 4.
func (suite *IdentifierTestSuite) TestValidateIds() {

	suite.True(IsId(uuid.NewV4().String()))
	suite.True(IsId(uuid.NewV7().String()))
	suite.False(IsId("xxx"))
}
