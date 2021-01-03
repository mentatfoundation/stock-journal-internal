package authentication

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"mentatfoundation/stock-journal/server/config"
	"mentatfoundation/stock-journal/server/logger"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var e *echo.Echo
var req *http.Request
var rec *httptest.ResponseRecorder
var c echo.Context
var testConfig config.ConfigurationSettings

func init() {

	testConfig = config.ConfigurationSettings{
		Env: "test",
	}
}

func TestTheTest(t *testing.T) {

	// Setup
	setupTest("get", "")

	// configure handler
	l := logger.New(testConfig)
	h := NewAuthHandler(l)

	// Assertions
	if assert.NoError(t, h.Test(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "login", rec.Body.String())
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
