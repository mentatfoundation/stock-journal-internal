package server

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	authHandler "mentatfoundation/stock-journal/server/handlers/authentication"
	profileHandler "mentatfoundation/stock-journal/server/handlers/profile"
	globalLogger "mentatfoundation/stock-journal/server/logger"
	"mentatfoundation/stock-journal/server/services"
)

func (a App) ConfigureRoutes() {

	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess)

	// setup logger
	logger := globalLogger.New(a.Config)

	// setup services
	as := services.NewAuthService(logger, svc)

	// configure handler & dependencies
	ah := authHandler.New(logger, as)
	ph := profileHandler.New(as, logger)

	// api group
	api := a.Server.Group("/api")

	api.GET("/test", ah.Test)
	api.GET("/shit", ph.Login)
	api.GET("/auth/signup", ah.SignUp)
}
