package roles

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

const someRight1 Right = 1
const someRight2 Right = 2
const someRight3 Right = 3
const someRight4 Right = 4

const Role1 Role = 1
const Role2 Role = 2
const Role3 Role = 3

func TestRoleTestSuite(t *testing.T) {
	suite.Run(t, new(RoleTestSuite))
}

type RoleTestSuite struct {
	suite.Suite
	registry *Registry
}

func (suite *RoleTestSuite) BeforeTest(name, test string) {
	suite.registry = NewRegistry()

	suite.registry.RegisterRole(Role1, someRight1, someRight2)
	suite.registry.RegisterRole(Role2)
	suite.registry.RegisterRole(Role3, someRight3)
}

func (suite *RoleTestSuite) TestEmptyRole() {
	user := User{}

	suite.False(user.Check(someRight1))
	suite.True(user.Check())
	suite.Panics(func() {
		user.Guard(someRight1)
	})
}

func (suite *RoleTestSuite) TestRoleCheck() {
	user1 := suite.registry.NewUser(1, Role1)
	user2 := suite.registry.NewUser(2, Role2)
	user3 := suite.registry.NewUser(3, Role3)
	user4 := suite.registry.NewUser(3, Role1, Role3)

	suite.True(user1.Check(someRight1))
	suite.True(user1.Check(someRight2))
	suite.False(user1.Check(someRight3))
	suite.False(user1.Check(someRight4))

	suite.False(user2.Check(someRight1))
	suite.False(user2.Check(someRight2))
	suite.False(user2.Check(someRight3))
	suite.False(user2.Check(someRight4))

	suite.False(user3.Check(someRight1))
	suite.False(user3.Check(someRight2))
	suite.True(user3.Check(someRight3))
	suite.False(user3.Check(someRight4))

	suite.True(user4.Check(someRight1))
	suite.True(user4.Check(someRight2))
	suite.True(user4.Check(someRight3))
	suite.False(user4.Check(someRight4))
}

func (suite *RoleTestSuite) TestRoleGuard() {
	user1 := suite.registry.NewUser(1, Role1)
	user2 := suite.registry.NewUser(2, Role2)
	user3 := suite.registry.NewUser(3, Role3)
	user4 := suite.registry.NewUser(3, Role1, Role3)

	user1.Guard(someRight1)
	user1.Guard(someRight2)
	suite.Panics(func() { user1.Guard(someRight3) })
	suite.Panics(func() { user1.Guard(someRight4) })

	suite.Panics(func() { user2.Guard(someRight1) })
	suite.Panics(func() { user2.Guard(someRight2) })
	suite.Panics(func() { user2.Guard(someRight3) })
	suite.Panics(func() { user2.Guard(someRight4) })

	suite.Panics(func() { user3.Guard(someRight1) })
	suite.Panics(func() { user3.Guard(someRight2) })
	user3.Guard(someRight3)
	suite.Panics(func() { user3.Guard(someRight4) })

	user4.Guard(someRight1)
	user4.Guard(someRight2)
	user4.Guard(someRight3)
	suite.Panics(func() { user4.Guard(someRight4) })
}
