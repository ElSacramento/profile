package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/profile/dto"
	"github.com/profile/tests/testutils"
)

func TestUpdateUser(t *testing.T) {
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

		update := dto.UpdateUser{
			FirstName: "kate",
			LastName:  "rogushkova",
			Nickname:  "nick",
			Password:  "pwd",
			Email:     "kate@mail.ru",
			Country:   "russia",
		}
		updated := testutils.UpdateUserWithResponse(t, created.ID, update)

		expected := dto.User{
			ID:        1,
			FirstName: "kate",
			LastName:  "rogushkova",
			Nickname:  "nick",
			Password:  "pwd",
			Email:     "kate@mail.ru",
			Country:   "Russia",
		}
		require.Equal(t, expected, updated)

		got := testutils.GetUserWithResponse(t, created.ID)
		require.Equal(t, expected, got)
	})
}

func TestUpdateUser_WrongParams(t *testing.T) {
	testutils.RunWithServer(t, func() {
		t.Run("empty email and password", func(t *testing.T) {
			request := dto.CreateUser{
				Password: "pwd",
				Email:    "test@mail.ru",
				Country:  "sweden",
			}
			created := testutils.CreateUserWithResponse(t, request)

			update := dto.UpdateUser{
				FirstName: "kate",
				LastName:  "rogushkova",
				Country:   "russia",
			}
			response := testutils.UpdateUserWithErrorResponse(t, created.ID, update, http.StatusUnprocessableEntity)
			require.Contains(t, response.Message, "empty")
		})
		t.Run("wrong country", func(t *testing.T) {
			request := dto.CreateUser{
				Password: "pwd",
				Email:    "test@mail.ru",
				Country:  "sweden",
			}
			created := testutils.CreateUserWithResponse(t, request)

			update := dto.UpdateUser{
				Password: "pwd",
				Email:    "test@mail.ru",
				Country:  "blabla",
			}
			response := testutils.UpdateUserWithErrorResponse(t, created.ID, update, http.StatusUnprocessableEntity)
			require.Contains(t, response.Message, "wrong country")
		})
		t.Run("not found", func(t *testing.T) {
			update := dto.UpdateUser{
				Password: "pwd",
				Email:    "test@mail.ru",
			}
			// no response
			testutils.UpdateUserWithErrorResponse(t, 100500, update, http.StatusNotFound)
		})
	})
}
