package authentication

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"mentatfoundation/stock-journal/server/config"
	"mentatfoundation/stock-journal/server/handlers/authentication/mocks"
	"mentatfoundation/stock-journal/server/logger"
	"mentatfoundation/stock-journal/server/models"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var e *echo.Echo
var req *http.Request
var rec *httptest.ResponseRecorder
var c echo.Context
var testConfig config.ConfigurationSettings
var testLogger logger.Logger
var mockAuthService mocks.AuthServiceMock
var authHandler *Handler

func TestMain(m *testing.M) {
	testConfig = config.ConfigurationSettings{
		Env: "test",
	}
	testLogger = logger.New(testConfig)
	mockAuthService = mocks.AuthServiceMock{}
	authHandler = NewAuthHandler(testLogger, mockAuthService)
	os.Exit(m.Run())
}

func TestTheTest(t *testing.T) {

	// Setup
	body := `{"username":"brian", "password":"password"}`
	setupTest("post", body)

	mocks.SignUpMock = func(newUser models.NewUser) error {
		return nil
	}

	//Assertions
	if assert.NoError(t, authHandler.SignUp(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestShouldBreak(t *testing.T) {

	// Setup
	body := `{"username":"brian"}`
	setupTest("post", body)

	//Assertions
	if assert.NoError(t, authHandler.SignUp(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
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
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec = httptest.NewRecorder()

	}

	c = e.NewContext(req, rec)
}
