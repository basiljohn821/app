package models

import (
	"errors"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	Name     string `gorm:"size:255;not null;unique" json:"name"`
	Password string `gorm:"size:100;not null;" json:"password"`
}

func HashAdmin(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPasswordAdmin(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (a *Admin) Prepare() {

	a.Name = html.EscapeString(strings.TrimSpace(a.Name))
	a.Password = html.EscapeString(strings.TrimSpace(a.Password))
}

func (a *Admin) Validate(action string) error {
	switch strings.ToLower(action) {

	case "Adminlogin":
		if a.Password == "" {
			return errors.New("Required Password")
		}
		if a.Name == "" {
			return errors.New("Required name")
		}
		return nil

	default:

		if a.Password == "" {
			return errors.New("Required Password")
		}
		if a.Name == "" {
			return errors.New("Required name")
		}
		return nil
	}
}
