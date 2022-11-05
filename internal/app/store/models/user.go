package models

import (
	"fmt"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {

		b, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
		if err != nil {
			return err
		}

		u.Password = string(b)

		fmt.Println(string(b))
	}

	return nil
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.NotIn(""), validation.Length(6, 100)),
	)
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) != nil
}
