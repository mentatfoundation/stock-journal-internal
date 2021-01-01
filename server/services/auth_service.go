package auth_service

import (
	"fmt"
)

type AuthService struct{}

func (a *AuthService) Test() {
	fmt.Println("hello")
}
