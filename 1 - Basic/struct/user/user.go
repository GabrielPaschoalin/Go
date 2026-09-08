package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string // Acessível
	lastName  string // Não acessível
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			FirstName: "ADMIN",
			lastName:  "ADMIN",
			birthDate: "---",
			createdAt: time.Now(),
		},
	}
}

func (u User) OutputUserDetails() {
	fmt.Println(u.FirstName, u.lastName, u.birthDate)
}

func (u *User) ClearUserName() {
	u.FirstName = ""
	u.lastName = ""
}

func New(firstName, lastName, birthDate string) (*User, error) {

	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("First name, last name and birthdate are required")
	}

	return &User{
		FirstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
