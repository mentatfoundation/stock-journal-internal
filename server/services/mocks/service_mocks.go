package auth_service_mocks

import "mentatfoundation/stock-journal/server/models"

type AuthServiceMock struct{}

var SignUpMock func(newUser models.NewUser) error

func (as AuthServiceMock) SignUp(newUser models.NewUser) error {
	return SignUpMock(newUser)
}
