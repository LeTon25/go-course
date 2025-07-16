package users

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	// Embedding User struct to inherit its fields and methods
	User User
}

func NewAdmin(email, password, firstName, lastName, birthDate string) (*Admin, error) {
	if email == "" || password == "" || firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("Email, password, first name, last name, and birth date cannot be empty")
	}
	user, err := New(firstName, lastName, birthDate)
	if err != nil {
		return nil, err
	}
	return &Admin{
		email:    email,
		password: password,
		User:     *user,
	}, nil
}

func (user User) OutputUserDetails() {
	fmt.Println("Your full name is:", user.firstName+" "+user.lastName)
	fmt.Println("Your birth date is:", user.birthDate)
}

func (user *User) ClearUserName() {
	user.firstName = ""
	user.lastName = ""
}
func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("First name, last name, and birth date cannot be empty")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
