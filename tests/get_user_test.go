package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/profile/dto"
	"github.com/profile/tests/testutils"
)

func TestGetUser(t *testing.T) {
	testutils.RunWithServer(t, func() {
		request := dto.CreateUser{
			FirstName: "fname",
			LastName:  "lname",
			Nickname:  "nick",
			Password:  "pwd",
			Email:     "test@mail.ru",
			Country:   "sweden",
		}
		created := testutils.CreateUserWithResponse(t, request)

		expected := dto.User{
			ID:        1,
			FirstName: "fname",
			LastName:  "lname",
			Nickname:  "nick",
			Password:  "pwd",
			Email:     "test@mail.ru",
			Country:   "Sweden",
		}
		got := testutils.GetUserWithResponse(t, created.ID)
		require.Equal(t, expected, got)
	})
}

func TestGetUser_WrongParams(t *testing.T) {
	testutils.RunWithServer(t, func() {
		t.Run("not found", func(t *testing.T) {
			// no response for 404
			testutils.GetUserWithErrorResponse(t, 100500, http.StatusNotFound)
		})
	})
}
