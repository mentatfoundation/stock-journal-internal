package services

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"mentatfoundation/stock-journal/server/logger"
)

type AuthService interface {
	CreateUser(id string) uuid.UUID
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

func (a *authService) CreateUser(id string) uuid.UUID {
	newUserId := uuid.New()
	a.logger.Info("CreateUser", "creating new user with id: "+newUserId.String())
	item := struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}{
		Id:   newUserId.String(),
		Name: "Brian",
	}

	av, err := dynamodbattribute.MarshalMap(item)

	a.logger.Info("CreateUser", "unmarshalling user "+newUserId.String())

	if err != nil {
		fmt.Println("Got error marshalling new movie item:")
		fmt.Println(err.Error())
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

	}

	return newUserId
}
