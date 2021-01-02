package authentication

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"mentatfoundation/stock-journal/server/logger"
	"net/http"
	"net/http/httptest"
	"testing"
)

var localLogger logger.Logger
var e *echo.Echo
var req *http.Request
var rec *httptest.ResponseRecorder
var c echo.Context

func init() {
	localLogger = logger.New("dev")
}

func TestLogin(t *testing.T) {

	// Setup
	setupTest()

	// configure handler
	h := New(localLogger)

	// Assertions
	if assert.NoError(t, h.Login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "login", rec.Body.String())
	}
}

func setupTest() {
	e = echo.New()
	req = httptest.NewRequest(http.MethodGet, "/", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

}
