package services

import (
	"fmt"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (a *AuthService) Test() {
	fmt.Println("hello")
}
