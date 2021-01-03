package services

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"mentatfoundation/stock-journal/server/logger"
	"mentatfoundation/stock-journal/server/models"
)

type AuthService interface {
	SignUp(newUser models.NewUser) error
}

type authService struct {
	logger logger.Logger
	db     *dynamodb.DynamoDB
}

func NewAuthService(logger logger.Logger, db *dynamodb.DynamoDB) AuthService {
	return &authService{
		logger: logger,
		db:     db,
	}
}

func (a *authService) Test() {
	fmt.Println("hello")
}

func (a *authService) SignUp(newUser models.NewUser) error {
	newUserId := uuid.New()
	a.logger.Info("CreateUser", "creating new user with id: "+newUserId.String())
	item := struct {
		Id       string `json:"id"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}{
		Id:       newUserId.String(),
		Name:     newUser.Username,
		Password: newUser.Password,
	}

	av, err := dynamodbattribute.MarshalMap(item)

	a.logger.Info("CreateUser", "unmarshalling user "+newUserId.String())

	if err != nil {
		fmt.Println("Got error marshalling new movie item:")
		return err
	}

	tableName := "stock-journal-users-dev"

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = a.db.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		return err

	}

	return nil
}
