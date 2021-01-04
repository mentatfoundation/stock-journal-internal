package models

import "errors"

type User struct {
}

type NewUser struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required, email"`
}

func (u NewUser) IsValid() error {
	if len(u.Password) > 0 && len(u.Username) > 0 {
		return nil
	}

	return errors.New("invalid")
}
