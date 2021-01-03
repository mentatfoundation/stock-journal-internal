package models

import "fmt"

type User struct {
}

type NewUser struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required, email"`
}

func (u NewUser) IsValid() {
	fmt.Println(u.Username)
	fmt.Println(u.Password)
}
