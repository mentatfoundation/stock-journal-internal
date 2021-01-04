package server

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"mentatfoundation/stock-journal/server/handlers/authentication"
	"mentatfoundation/stock-journal/server/handlers/profile"
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

	// dynamo
	svc := dynamodb.New(sess)

	// setup logger
	logger := globalLogger.New(a.Config)

	// setup services
	authSvc := services.NewAuthService(logger, svc)

	// configure handler
	ah := authentication.NewAuthHandler(logger, authSvc)
	ph := profile.NewProfileHandler(logger, authSvc)

	// api group
	api := a.Server.Group("/api")

	api.GET("/shit", ph.Login)
	api.POST("/auth/signup", ah.SignUp)
}
