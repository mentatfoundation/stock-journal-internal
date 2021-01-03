package server

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	authHandler "mentatfoundation/stock-journal/server/handlers/authentication"
	profileHandler "mentatfoundation/stock-journal/server/handlers/profile"
	globalLogger "mentatfoundation/stock-journal/server/logger"
	"mentatfoundation/stock-journal/server/services"
	"os"
)

func (a App) ConfigureRoutes() {

	sess, err := session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String("us-east-1")},
		Profile: "stock-journal",
	})

	if err != nil {
		fmt.Println("Error configuring AWS Session: " + err.Error())
		os.Exit(1)
	}

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
