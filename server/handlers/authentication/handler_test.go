package authentication

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"mentatfoundation/stock-journal/server/logger"
	"mentatfoundation/stock-journal/server/models"
	"mentatfoundation/stock-journal/server/services"
	authServiceMocks "mentatfoundation/stock-journal/server/services/mocks"
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
var mockLogger logger.Logger
var mockAuthService services.AuthService
var authHandler *Handler

func TestMain(m *testing.M) {
	mockLogger = logger.Mock{}
	mockAuthService = authServiceMocks.AuthServiceMock{}
	authHandler = NewAuthHandler(mockLogger, mockAuthService)
	os.Exit(m.Run())
}

func TestTheTest(t *testing.T) {

	// Setup
	body := `{"username":"brian", "password":"password"}`
	setupTest("post", body)

	authServiceMocks.SignUpMock = func(newUser models.NewUser) error {
		return nil
	}

	// Assert
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

func TestUserModelValidation(t *testing.T) {

	// Setup
	body := `{"name":"brian"}`
	setupTest("post", body)

	//Assertions
	if assert.NoError(t, authHandler.SignUp(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, "user invalid.", rec.Body.String())
	}
}

func TestUserModelBindShouldFail(t *testing.T) {

	// Setup
	body := `{"name":"brian`
	setupTest("post", body)

	//Assertions
	if assert.NoError(t, authHandler.SignUp(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, "Unable to process request.", rec.Body.String())
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
