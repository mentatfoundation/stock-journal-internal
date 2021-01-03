package authentication

import (
	"github.com/labstack/echo/v4"
	"mentatfoundation/stock-journal/server/logger"
	"mentatfoundation/stock-journal/server/models"
	"mentatfoundation/stock-journal/server/services"
	"net/http"
)

type Handler struct {
	authService services.AuthService
	logger      logger.Logger
}

func NewAuthHandler(logger logger.Logger, authService services.AuthService) *Handler {
	return &Handler{
		logger:      logger,
		authService: authService,
	}
}

func (a *Handler) Login(c echo.Context) error {
	return c.String(http.StatusOK, "login")
}

func (a *Handler) Logout(c echo.Context) error {
	return c.String(http.StatusOK, "logout")
}

func (a *Handler) User(c echo.Context) error {
	return c.String(http.StatusOK, c.Param("id"))
}

func (a *Handler) Test(c echo.Context) error {
	return c.String(http.StatusOK, "login")
}

func (a *Handler) SignUp(c echo.Context) error {

	// process data
	newUser := new(models.NewUser)
	if err := c.Bind(newUser); err != nil {
		return c.String(http.StatusBadRequest, "")
	}

	newUser.IsValid()

	// call service
	err := a.authService.SignUp(*newUser)
	if err != nil {
		return c.String(http.StatusBadRequest, "0")
	}

	return c.NoContent(http.StatusCreated)
}
