package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/profile/dto"
	"github.com/profile/tests/testutils"
)

func TestCreateUser(t *testing.T) {
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
		require.Equal(t, expected, created)
	})
}

func TestCreateUser_WrongParams(t *testing.T) {
	testutils.RunWithServer(t, func() {
		t.Run("empty email and password", func(t *testing.T) {
			request := dto.CreateUser{
				FirstName: "fname",
				LastName:  "lname",
				Nickname:  "nick",
				Country:   "sweden",
			}
			response := testutils.CreateUserWithErrorResponse(t, request, http.StatusUnprocessableEntity)
			require.Contains(t, response.Message, "empty")
		})
		t.Run("wrong country", func(t *testing.T) {
			request := dto.CreateUser{
				Email:    "test@mail.ru",
				Password: "pwd",
				Country:  "blabla",
			}
			response := testutils.CreateUserWithErrorResponse(t, request, http.StatusUnprocessableEntity)
			require.Contains(t, response.Message, "wrong country")
		})
	})
}
