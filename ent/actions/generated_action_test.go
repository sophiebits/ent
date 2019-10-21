package actions_test

import (
	"testing"

	"github.com/lolopinto/ent/ent"
	"github.com/lolopinto/ent/ent/actions"
	"github.com/lolopinto/ent/ent/viewer"
	"github.com/lolopinto/ent/ent/viewertesting"

	"github.com/lolopinto/ent/ent/test_schema/models"
	"github.com/lolopinto/ent/ent/test_schema/models/configs"
	"github.com/lolopinto/ent/ent/test_schema/models/user/action"
	"github.com/lolopinto/ent/internal/testingutils"
	testhelpers "github.com/lolopinto/ent/internal/testschemautils"
	"github.com/lolopinto/ent/internal/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type generatedActionSuite struct {
	testingutils.Suite
}

func (suite *generatedActionSuite) SetupSuite() {
	suite.Tables = []string{
		"users",
		"event_invited_edges",
		"events",
		"user_family_members_edges",
		"user_friends_edges",
	}
	suite.Suite.SetupSuite()
}

func (suite *generatedActionSuite) createUser() *models.User {
	v := viewer.LoggedOutViewer()

	email := util.GenerateRandEmail()

	user, err := action.CreateUser(v).
		SetEmailAddress(email).
		SetFirstName("Ola").
		SetLastName("Okelola").
		Save()

	assert.Nil(suite.T(), err)
	testhelpers.VerifyUserObj(suite.T(), user, email)
	return user
}
func (suite *generatedActionSuite) TestCreation() {
	suite.createUser()
}

func (suite *generatedActionSuite) TestEditing() {
	user := suite.createUser()

	v := viewertesting.LoggedinViewerContext{ViewerID: user.ID}

	editedUser, err := action.EditUser(v, user).
		SetFirstName("Ola2").
		Save()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), editedUser.EmailAddress, user.EmailAddress)
	assert.Equal(suite.T(), editedUser.FirstName, "Ola2")
	assert.Equal(suite.T(), editedUser.LastName, user.LastName)
}

func (suite *generatedActionSuite) TestDeleting() {
	user := suite.createUser()

	v := viewertesting.LoggedinViewerContext{ViewerID: user.ID}

	err := action.DeleteUser(v, user).
		Save()

	assert.Nil(suite.T(), err)

	user2, err := models.LoadUser(v, user.ID)
	assert.NotNil(suite.T(), err)
	assert.Zero(suite.T(), *user2)
}

func (suite *generatedActionSuite) TestAddEdgeAction() {
	user := suite.createUser()
	user2 := suite.createUser()

	v := viewertesting.LoggedinViewerContext{ViewerID: user.ID}

	updatedUser, err := action.AddFriends(v, user).
		AddUser(user2).
		Save()

	assert.Nil(suite.T(), err)
	testhelpers.VerifyUserObj(suite.T(), updatedUser, user.EmailAddress)

	testhelpers.VerifyFriendsEdge(suite.T(), user, user2)
}

func (suite *generatedActionSuite) addFamilyEdge(v viewer.ViewerContext, user, user2 *models.User) {
	b := actions.EntMutationBuilder{
		Viewer:         v,
		EntConfig:      &configs.UserConfig{},
		Operation:      ent.EditOperation,
		ExistingEntity: user,
	}
	// manually adding this until we fix the API and generating this correctly
	b.AddOutboundEdge(models.UserToFamilyMembersEdge, user2.ID, user2.GetType())

	var updatedUser models.User
	c, err := b.GetChangeset(&updatedUser)
	assert.Nil(suite.T(), err)
	err = ent.SaveChangeset(c)
	assert.Nil(suite.T(), err)

	testhelpers.VerifyFamilyEdge(suite.T(), user, user2)
}

func (suite *generatedActionSuite) TestRemoveEdgeAction() {
	user := suite.createUser()
	user2 := suite.createUser()

	v := viewertesting.LoggedinViewerContext{ViewerID: user.ID}

	suite.addFamilyEdge(v, user, user2)

	// hmm this API :(
	// remove
	updatedUser, err := action.RemoveFamilyMembers(v, user).
		AddUser(user2).
		Save()

	assert.Nil(suite.T(), err)
	testhelpers.VerifyUserObj(suite.T(), updatedUser, user.EmailAddress)
	testhelpers.VerifyNoFamilyEdge(suite.T(), user, user2)
}

func TestGeneratedAction(t *testing.T) {
	suite.Run(t, new(generatedActionSuite))
}
