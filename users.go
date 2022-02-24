package types

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	passwordSalt = 10
)

type User struct {
	StampedEntity
	IDEntity
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (u *User) SetPassword(password, confirmPassword string) error {
	if len(password) == 0 || len(confirmPassword) == 0 || password != confirmPassword {
		return fmt.Errorf("password and confirm password must exist and be the same")
	}

	passwordByte, err := bcrypt.GenerateFromPassword([]byte(password), passwordSalt)
	if err != nil {
		return err
	}

	u.Password = string(passwordByte)

	return nil
}

func (u *User) IsPasswordValid(password string) bool {
	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) != nil {
		return false
	}

	return true
}

type UpdateUser struct {
	IDEntity
	FirstName       *string `json:"firstName"`
	LastName        *string `json:"lastName"`
	Password        *string `json:"password"`
	PasswordConfirm *string `json:"passwordConfirm"`
}
