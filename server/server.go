package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"mentatfoundation/stock-journal/server/config"
	authHandler "mentatfoundation/stock-journal/server/handlers/authentication"
	profileHandler "mentatfoundation/stock-journal/server/handlers/profile"
	globalLogger "mentatfoundation/stock-journal/server/logger"
	authService "mentatfoundation/stock-journal/server/services"
	"net/http"
)

type App struct {
	Server *echo.Echo
	Config config.ConfigurationSettings
}

func New(c config.ConfigurationSettings) App {
	e := echo.New()
	return App{
		Server: e,
		Config: c,
	}
}

func (a App) Configure() {
	a.ConfigureMiddleware()
	a.ConfigureRoutes()
}

func (a App) ConfigureMiddleware() {

	if a.Config.IsDev() {
		a.Server.Use(middleware.CORS())
	}

	a.Server.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "client/build",
		HTML5: true,
	}))
}

func (a App) ConfigureRoutes() {

	// setup logger
	logger := globalLogger.New(a.Config.Env)

	// setup services
	as := authService.New()

	// configure handler & dependencies
	ah := authHandler.New(logger)
	ph := profileHandler.New(as, logger)

	// api group
	api := a.Server.Group("/api")

	api.GET("/auth/login", ah.Login)
	api.GET("/home", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello")
	})

	a.Server.GET("/test", ph.Login)
}

func (a App) Run() {
	a.Configure()
	a.Server.Logger.Fatal(a.Server.Start(":" + a.Config.Port))
}
