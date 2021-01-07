package models

import "fmt"

type User struct {
	Id       string
	Username string
	Password string
}

type NewUser struct {
	Username string
	Password string
}

type ValidationError struct {
	Field   string
	Message string
}

type UserInvalidResponse struct {
	Message    string             `json:"message,omitempty"`
	Error      string             `json:"error,omitempty"`
	Validation []*ValidationError `json:"validation,omitempty"`
}

func (u NewUser) IsValid() (bool, *UserInvalidResponse) {
	var valError []*ValidationError

	if len(u.Password) > 0 && len(u.Username) > 0 {
		return true, nil
	}

	if u.Password == "" {
		valError = append(valError, &ValidationError{Message: "Password is required", Field: "password"})
	}

	fmt.Println(valError)
	return false, &UserInvalidResponse{Message: "Invalid", Validation: valError}
}
