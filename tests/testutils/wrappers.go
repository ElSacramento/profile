package testutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/profile/dto"
)

func getURL(method string) string {
	return fmt.Sprintf("http://%s%s", ServerAddr, method)
}

func postRequest(t *testing.T, url string, reqBody interface{}, status int) []byte {
	body, err := json.Marshal(reqBody)
	require.NoError(t, err)

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	require.NoError(t, err)

	request.Header.Add("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)
	require.NoError(t, err)
	require.Equal(t, status, response.StatusCode)

	bodyBytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	require.NoError(t, err)

	return bodyBytes
}

func deleteRequest(t *testing.T, url string, status int) []byte {
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	require.NoError(t, err)

	response, err := http.DefaultClient.Do(request)
	require.NoError(t, err)
	require.Equal(t, status, response.StatusCode)

	bodyBytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	require.NoError(t, err)

	return bodyBytes
}

func getRequest(t *testing.T, url string, status int) []byte {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	response, err := http.DefaultClient.Do(request)
	require.NoError(t, err)
	require.Equal(t, status, response.StatusCode)

	bodyBytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	require.NoError(t, err)

	return bodyBytes
}

func CreateUserWithResponse(t *testing.T, createUser dto.CreateUser) dto.User {
	bodyBytes := postRequest(t, getURL("/users"), createUser, http.StatusOK)

	var user dto.User
	err := json.Unmarshal(bodyBytes, &user)
	require.NoError(t, err)

	return user
}

func CreateUserWithErrorResponse(t *testing.T, createUser dto.CreateUser, status int) dto.ErrorResponse {
	bodyBytes := postRequest(t, getURL("/users"), createUser, status)

	var errResponse dto.ErrorResponse
	err := json.Unmarshal(bodyBytes, &errResponse)
	require.NoError(t, err)

	return errResponse
}

func DeleteUser(t *testing.T, id uint64) {
	url := getURL("/users/") + strconv.Itoa(int(id))
	deleteRequest(t, url, http.StatusNoContent)
}

func DeleteUserWithErrorResponse(t *testing.T, id uint64, status int) dto.ErrorResponse {
	url := getURL("/users/") + strconv.Itoa(int(id))
	bodyBytes := deleteRequest(t, url, status)

	var errResponse dto.ErrorResponse
	err := json.Unmarshal(bodyBytes, &errResponse)
	require.NoError(t, err)

	return errResponse
}

func UpdateUserWithResponse(t *testing.T, id uint64, update dto.UpdateUser) dto.User {
	url := getURL("/users/") + strconv.Itoa(int(id))
	bodyBytes := postRequest(t, url, update, http.StatusOK)

	var user dto.User
	err := json.Unmarshal(bodyBytes, &user)
	require.NoError(t, err)

	return user
}

func UpdateUserWithErrorResponse(t *testing.T, id uint64, update dto.UpdateUser, status int) dto.ErrorResponse {
	url := getURL("/users/") + strconv.Itoa(int(id))
	bodyBytes := postRequest(t, url, update, status)

	var errResponse dto.ErrorResponse
	err := json.Unmarshal(bodyBytes, &errResponse)
	require.NoError(t, err)

	return errResponse
}

func GetUserWithResponse(t *testing.T, id uint64) dto.User {
	url := getURL("/users/") + strconv.Itoa(int(id))
	bodyBytes := getRequest(t, url, http.StatusOK)

	var user dto.User
	err := json.Unmarshal(bodyBytes, &user)
	require.NoError(t, err)

	return user
}

func GetUserWithErrorResponse(t *testing.T, id uint64, status int) dto.ErrorResponse {
	url := getURL("/users/") + strconv.Itoa(int(id))
	bodyBytes := deleteRequest(t, url, status)

	var errResponse dto.ErrorResponse
	err := json.Unmarshal(bodyBytes, &errResponse)
	require.NoError(t, err)

	return errResponse
}

func UsersWithResponse(t *testing.T, filter dto.Filter) []dto.User {
	url := getURL("/users")
	if filter.Country != "" {
		urlCreation, err := url2.Parse(url)
		require.NoError(t, err)

		values := urlCreation.Query()
		values.Add("country", filter.Country)
		urlCreation.RawQuery = values.Encode()

		url = urlCreation.String()
	}
	bodyBytes := getRequest(t, url, http.StatusOK)

	var users []dto.User
	err := json.Unmarshal(bodyBytes, &users)
	require.NoError(t, err)

	return users
}

func UsersWithErrorResponse(t *testing.T, filter dto.Filter, status int) dto.ErrorResponse {
	url := getURL("/users")
	if filter.Country != "" {
		urlCreation, err := url2.Parse(url)
		require.NoError(t, err)

		values := urlCreation.Query()
		values.Add("country", filter.Country)
		urlCreation.RawQuery = values.Encode()

		url = urlCreation.String()
	}
	bodyBytes := getRequest(t, url, status)

	var errResponse dto.ErrorResponse
	err := json.Unmarshal(bodyBytes, &errResponse)
	require.NoError(t, err)

	return errResponse
}
