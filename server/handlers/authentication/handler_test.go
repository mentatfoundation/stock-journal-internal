package authentication

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"mentatfoundation/stock-journal/server/logger"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var testLogger logger.Logger
var e *echo.Echo
var req *http.Request
var rec *httptest.ResponseRecorder
var c echo.Context

type log struct{}

func (tl log) Info(operator string, message string) {
}

func init() {
	testLogger = log{}
}

func TestLogin(t *testing.T) {

	// Setup
	setupTest("get", "")

	// configure handler
	h := New(testLogger)

	// Assertions
	if assert.NoError(t, h.Login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "login", rec.Body.String())
	}
}

func TestLogout(t *testing.T) {

	// Setup
	setupTest("get", "")

	// configure handler
	h := New(testLogger)

	// Assertions
	if assert.NoError(t, h.Logout(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "logout", rec.Body.String())
	}
}

func TestCustomer(t *testing.T) {

	// Setup
	setupTest("get", "")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// configure handler
	h := New(testLogger)

	// Assertions
	if assert.NoError(t, h.User(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "1", rec.Body.String())
	}
}

func setupTest(method string, body string) {
	e = echo.New()
	switch method {
	case "get":
		req = httptest.NewRequest(http.MethodGet, "/", nil)
		rec = httptest.NewRecorder()

	case "post":
		req = httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body))
		rec = httptest.NewRecorder()

	}

	c = e.NewContext(req, rec)
}
