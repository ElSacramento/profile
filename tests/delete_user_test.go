package tests

import (
	"net/http"
	"testing"

	"github.com/profile/dto"
	"github.com/profile/tests/testutils"
)

func TestDeleteUser(t *testing.T) {
	testutils.RunWithServer(t, func() {
		request := dto.CreateUser{
			Password: "pwd",
			Email:    "test@mail.ru",
		}
		created := testutils.CreateUserWithResponse(t, request)
		testutils.DeleteUser(t, created.ID)
	})
}

func TestDeleteUser_WrongParams(t *testing.T) {
	testutils.RunWithServer(t, func() {
		t.Run("not found", func(t *testing.T) {
			// no response for 404
			testutils.DeleteUserWithErrorResponse(t, 100500, http.StatusNotFound)
		})
	})
}
