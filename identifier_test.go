package utils

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/suite"
)

type IdentifierTestSuite struct {
	suite.Suite
}

func TestIdentifierTestSuite(t *testing.T) {
	suite.Run(t, new(IdentifierTestSuite))
}

func (suite *IdentifierTestSuite) TestGenrateIds() {

	id1 := NewId()
	suite.Len(id1, 36)
	_, err := uuid.FromString(id1)
	suite.Nil(err)
}

func (suite *IdentifierTestSuite) TestValidaeIds() {

	suite.True(IsId(uuid.NewV4().String()))
	suite.False(IsId("xxx"))
}
