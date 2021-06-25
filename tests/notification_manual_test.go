package tests

import (
	"testing"

	"github.com/profile/dto"
	"github.com/profile/tests/testutils"
)

// Expect notifications with create, update, get, delete.
func TestNotifications(t *testing.T) {
	testutils.RunWithGRPCServer(t, func() {
		request := dto.CreateUser{
			FirstName: "fname",
			LastName:  "lname",
			Nickname:  "nick",
			Password:  "pwd",
			Email:     "test@mail.ru",
			Country:   "sweden",
		}
		created := testutils.CreateUserWithResponse(t, request)

		update := dto.UpdateUser{
			FirstName: "kate",
			LastName:  "rogushkova",
			Nickname:  "nick",
			Password:  "pwd",
			Email:     "kate@mail.ru",
			Country:   "russia",
		}
		testutils.UpdateUserWithResponse(t, created.ID, update)

		testutils.GetUserWithResponse(t, created.ID)

		testutils.DeleteUser(t, created.ID)
	})
}
