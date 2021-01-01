package authentication

import (
	"net/http"

	authservice "mentatfoundation/stock-journal/server/services"

	"github.com/labstack/echo"
)

type AuthHandler struct {
	AuthService *authservice.AuthService
}

func (a *AuthHandler) Login(c echo.Context) error {
	return c.String(http.StatusOK, "login")
}
