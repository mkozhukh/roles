package roles

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestHelpersSuite(t *testing.T) {
	suite.Run(t, new(HelpersTestSuite))
}

type HelpersTestSuite struct {
	suite.Suite
}

func (suite *HelpersTestSuite) TestParseRight() {
	const someRight1 Right = 1
	const someRight2 Right = 2
	const someRight3 Right = 3
	const someRight4 Right = 4

	set1 := ParseRights("a,1,3")
	suite.Equal([]Right{someRight1, someRight3}, set1)

	set2 := ParseRights("admin")
	suite.Equal([]Right{}, set2)

	set3 := ParseRights("")
	suite.Equal([]Right{}, set3)
}

func (suite *HelpersTestSuite) TestSerializeRight() {
	const someRight1 Right = 1
	const someRight2 Right = 2
	const someRight3 Right = 3
	const someRight4 Right = 4

	set1 := SerializeRights(someRight1, someRight3)
	suite.Equal("1,3", set1)

	set2 := SerializeRights()
	suite.Equal("", set2)
}
