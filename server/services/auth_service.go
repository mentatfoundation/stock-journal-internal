package authentication

import (
	"fmt"
)

type AuthService struct{}

func New() *AuthService {
	return &AuthService{}
}

func (a *AuthService) Test() {
	fmt.Println("hello")
}
