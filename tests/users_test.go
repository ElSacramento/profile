package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/profile/dto"
	"github.com/profile/tests/testutils"
)

func TestUsers_Empty(t *testing.T) {
	testutils.RunWithServer(t, func() {
		response := testutils.UsersWithResponse(t, dto.Filter{})
		require.Equal(t, make([]dto.User, 0), response)
	})
}

func TestUsers(t *testing.T) {
	testutils.RunWithServer(t, func() {
		ukUser1 := "test@mail.ru"
		ukUser2 := "kate@mail.ru"
		usUser := "tropical@mail.ru"
		denmarkUser := "actual@mail.ru"

		// prepare - create users with different countries
		{
			request := dto.CreateUser{
				Password: "pwd",
				Email:    ukUser1,
				Country:  "United Kingdom",
			}
			testutils.CreateUserWithResponse(t, request)

			request = dto.CreateUser{
				Password: "pwd",
				Email:    ukUser2,
				Country:  "united kingdom",
			}
			testutils.CreateUserWithResponse(t, request)

			request = dto.CreateUser{
				Password: "pwd",
				Email:    usUser,
				Country:  "united states",
			}
			testutils.CreateUserWithResponse(t, request)

			request = dto.CreateUser{
				Password: "pwd",
				Email:    denmarkUser,
				Country:  "denmark",
			}
			testutils.CreateUserWithResponse(t, request)
		}

		// no filter
		{
			expected := []dto.User{
				{
					ID:       1,
					Password: "pwd",
					Email:    ukUser1,
					Country:  "United Kingdom",
				},
				{
					ID:       2,
					Password: "pwd",
					Email:    ukUser2,
					Country:  "United Kingdom",
				},
				{
					ID:       3,
					Password: "pwd",
					Email:    usUser,
					Country:  "United States",
				},
				{
					ID:       4,
					Password: "pwd",
					Email:    denmarkUser,
					Country:  "Denmark",
				},
			}
			response := testutils.UsersWithResponse(t, dto.Filter{})
			require.Equal(t, expected, response)
		}

		// with uk filter
		{
			expected := []dto.User{
				{
					ID:       1,
					Password: "pwd",
					Email:    ukUser1,
					Country:  "United Kingdom",
				},
				{
					ID:       2,
					Password: "pwd",
					Email:    ukUser2,
					Country:  "United Kingdom",
				},
			}
			response := testutils.UsersWithResponse(t, dto.Filter{Country: "united kingdom"})
			require.Equal(t, expected, response)
		}

		// with sweden filter
		{
			response := testutils.UsersWithResponse(t, dto.Filter{Country: "sweden"})
			require.Equal(t, make([]dto.User, 0), response)
		}

		// with wrong country name filter
		{
			response := testutils.UsersWithErrorResponse(t, dto.Filter{Country: "blabla"}, http.StatusUnprocessableEntity)
			require.Contains(t, response.Message, "wrong country")
		}
	})
}
