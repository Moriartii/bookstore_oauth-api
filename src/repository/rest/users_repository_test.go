package rest

import (
	//"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeOutFromApi(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https://localhost:8081/users/login",
		ReqBody:      `{"email":"test@test.com","password":"1234_test_5678"}`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})

	repository := usersRepository{}

	user, err := repository.LoginUser("email@gmail.com", "1234_test_5678")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid restclient response when trying to login user", err.Message)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	// rest.FlushMockups()
	// rest.AddMockups(&rest.Mock{
	// 	HTTPMethod:   http.MethodPost,
	// 	URL:          "https://localhost:8081/users/login",
	// 	ReqBody:      `{"email":"test@test.com","password":"1234_test_5678"}`,
	// 	RespHTTPCode: http.StatusNotFound,
	// 	RespBody:     `{"message":"invalid login credentials","status":404,"error":"not_found"}`,
	// })

	// repository := usersRepository{}

	// user, err := repository.LoginUser("test@test.com", "1234_test_5678")

	// assert.Nil(t, user)
	// assert.NotNil(t, err)
	// assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	// assert.EqualValues(t, "invalid error interface when tryng to login user", err.Message)
}

func TestLoginInvalidLoginCredentials(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "https://localhost:8081/users/login",
		ReqBody:      `{"email":"test@test.com","password":"1234_test_5678"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id":44,"first_name":"Nikolay","last_name":"Sharapov","email":"nik0@nik0.com"}`,
	})

	repository := usersRepository{}

	user, err := repository.LoginUser("email@gmail.com", "1234_test_5678")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshal users response", err.Message)
}

func TestLoginUserNoError(t *testing.T) {

}
