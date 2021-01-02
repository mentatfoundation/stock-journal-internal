package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"mentatfoundation/stock-journal/server/config"
	authHandler "mentatfoundation/stock-journal/server/handlers/authentication"
	profileHandler "mentatfoundation/stock-journal/server/handlers/profile"
	globalLogger "mentatfoundation/stock-journal/server/logger"
	"mentatfoundation/stock-journal/server/services"
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
	a.Server.Use(middleware.Recover())
	if a.Config.IsDev() {
		a.Server.Use(middleware.Logger())
		a.Server.Use(middleware.CORS())
	}

	a.Server.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "client/build",
		HTML5: true,
	}))
}

func (a App) ConfigureRoutes() {

	// setup logger
	logger := globalLogger.New(a.Config)

	// setup services
	as := services.NewAuthService()
	//ss := services.NewStocksService()

	// configure handler & dependencies
	ah := authHandler.New(logger)
	ph := profileHandler.New(as, logger)

	// api group
	api := a.Server.Group("/api")

	api.GET("/test", ah.Test)
	api.GET("/shit", ph.Login)
}

func (a App) Run() {
	a.Configure()
	a.Server.Logger.Fatal(a.Server.Start(":" + a.Config.Port))
}
