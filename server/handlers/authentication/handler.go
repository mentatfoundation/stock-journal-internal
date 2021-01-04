package authentication

import (
	"fmt"
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

func (h *Handler) Login(c echo.Context) error {
	return c.String(http.StatusOK, "login")
}

func (h *Handler) Logout(c echo.Context) error {
	return c.String(http.StatusOK, "logout")
}

func (h *Handler) User(c echo.Context) error {
	return c.String(http.StatusOK, c.Param("id"))
}

func (h *Handler) SignUp(c echo.Context) error {

	// process data
	newUser := new(models.NewUser)
	if err := c.Bind(newUser); err != nil {
		return c.String(http.StatusBadRequest, "Unable to process request.")
	}

	if err := newUser.IsValid(); err != nil {
		fmt.Println("here")
		return c.String(http.StatusBadRequest, "user invalid.")
	}

	// call service
	err := h.authService.SignUp(*newUser)

	if err != nil {
		return c.String(http.StatusBadRequest, "0")
	}

	return c.NoContent(http.StatusCreated)
}
