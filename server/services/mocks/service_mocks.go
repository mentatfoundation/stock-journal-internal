package auth_service_mocks

import (
	"mentatfoundation/stock-journal/server/errors"
	"mentatfoundation/stock-journal/server/models"
)

type AuthServiceMock struct{}

var SignUpMock func(newUser models.NewUser) error
var LoginMock func(user models.NewUser) (string, *errors.Error)

func (as AuthServiceMock) SignUp(newUser models.NewUser) error {
	return SignUpMock(newUser)
}

func (as AuthServiceMock) Login(user models.NewUser) (string, *errors.Error) {
	return LoginMock(user)
}
