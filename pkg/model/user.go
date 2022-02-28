package model

import "fmt"

type User struct {
	Id       int
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) CheckName() error {
	if len(u.Name) < 6 {
		return fmt.Errorf("Invalid name")
	}
	return nil
}

func (u *User) CheckPassword() error {
	if len(u.Password) < 6 || len(u.Password) > 20 {
		return fmt.Errorf("Invalid password")
	}
	return nil
}
