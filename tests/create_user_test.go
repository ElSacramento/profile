package tests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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
		created := CreateUserWithResponse(t, request)
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
			response := CreateUserWithErrorResponse(t, request, http.StatusUnprocessableEntity)
			require.Contains(t, response.Message, "empty")
		})
		t.Run("wrong country", func(t *testing.T) {
			request := dto.CreateUser{
				Email:    "test@mail.ru",
				Password: "pwd",
				Country:  "blabla",
			}
			response := CreateUserWithErrorResponse(t, request, http.StatusUnprocessableEntity)
			require.Contains(t, response.Message, "wrong country")
		})
	})
}

func CreateUserWithResponse(t *testing.T, createUser dto.CreateUser) dto.User {
	body, err := json.Marshal(createUser)
	require.NoError(t, err)

	request, err := http.NewRequest(http.MethodPost, "http://localhost:8080/users", bytes.NewReader(body))
	require.NoError(t, err)

	request.Header.Add("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode)

	bodyBytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var user dto.User
	err = json.Unmarshal(bodyBytes, &user)
	require.NoError(t, err)

	return user
}

func CreateUserWithErrorResponse(t *testing.T, createUser dto.CreateUser, status int) dto.ErrorResponse {
	body, err := json.Marshal(createUser)
	require.NoError(t, err)

	request, err := http.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
	require.NoError(t, err)

	response, err := http.DefaultClient.Do(request)
	require.NoError(t, err)
	require.Equal(t, status, response.StatusCode)

	bodyBytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var errResponse dto.ErrorResponse
	err = json.Unmarshal(bodyBytes, &errResponse)
	require.NoError(t, err)

	return errResponse
}
