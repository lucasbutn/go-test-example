package balance

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

const contentTypeJSON = "application/json"

func TestGetAllMovements(t *testing.T) {

	testServer := getServer(200, `[{"userId":"anyUserId","time":1596699342,"description":"DEPOSIT","value":1000.0},{"userId":"anyUserId","time":1596789342,"description":"PURCHASE","value":-100.0}]`)
	defer testServer.Close()

	restClient, _ := NewRestClient(testServer.URL)

	movements, err := restClient.GetAllMovements("anyUserId")

	require.Nil(t, err)

	assert.True(t, len(movements) == 2)
	assert.Equal(t, &Movement{"anyUserId", 1596699342, "DEPOSIT", 1000.0}, movements[0])
	assert.Equal(t, &Movement{"anyUserId", 1596789342, "PURCHASE", -100.0}, movements[1])
}

func TestGetAllMovementsBadRequest(t *testing.T) {

	testServer := getServer(400, "")
	defer testServer.Close()

	restClient,_ := NewRestClient(testServer.URL)

	_, err := restClient.GetAllMovements("anyUserId")

	assert.Error(t, err)
	assert.Equal(t, err.Error(), "response status was not OK")
}

// generate a test server so we can capture and inspect the request
func getServer(statusCode int, body string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(statusCode)
			res.Write([]byte(body))
			res.Header().Add("Content-Type", contentTypeJSON)
	}))
}


