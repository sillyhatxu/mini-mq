package dto

import "fmt"

type AddUserDTO struct {
	LoginName string `json:"login_name"`
	Password  string `json:"password"`
}

func (dto AddUserDTO) Validate() error {
	if dto.LoginName == "" {
		return fmt.Errorf("LoginName cannot be empty")
	}
	if dto.Password == "" {
		return fmt.Errorf("Password cannot be empty")
	}
	return nil
}
