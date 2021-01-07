package services

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"mentatfoundation/stock-journal/server/errors"
	"mentatfoundation/stock-journal/server/logger"
	"mentatfoundation/stock-journal/server/models"
)

type AuthService interface {
	SignUp(newUser models.NewUser) error
	Login(newUser models.NewUser) (string, *errors.Error)
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

func (a *authService) SignUp(newUser models.NewUser) error {
	newUserId := uuid.New()

	hashedPassword, err := HashPassword(newUser.Password)

	if err != nil {
		a.logger.Info("CreateUser", "Error hashing password"+err.Error())
	}

	item := struct {
		Id       string `json:"id"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}{
		Id:       newUserId.String(),
		Name:     newUser.Username,
		Password: hashedPassword,
	}

	av, err := dynamodbattribute.MarshalMap(item)

	if err != nil {
		a.logger.Info("CreateUser", "Error unmarshalling user "+err.Error())
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

func (a *authService) Login(user models.NewUser) (string, *errors.Error) {
	result, err := a.db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("stock-journal-users-dev"),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(user.Username),
			},
		},
	})

	if err != nil {
		fmt.Println(err.Error())
		return "", &errors.Error{Code: 500}
	}

	if result.Item == nil {
		return "", nil
	}

	userFound := &models.User{}

	if err := dynamodbattribute.UnmarshalMap(result.Item, &userFound); err != nil {
		return "", errors.Create("LoginSvcUnmarshall", "Unable to login", 500)
		//return "", &appError.Error{Code: 500, Message: "Error logging in"}
	}

	if !PasswordsMatch(userFound.Password, user.Password) {
		return "", &errors.Error{Code: 400, Message: "Login information invalid", Operation: "PasswordsMatch"}
	}

	return "", &errors.Error{
		Code:      500,
		Message:   "hello",
		Operation: "LoginSvcCatchAll"}
}
