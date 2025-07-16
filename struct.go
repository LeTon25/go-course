package main

import (
	"fmt"

	"example.com/structs-sample/users"
)

func main() {
	/*
		firstName := getUserData("Enter your first name: ")
		lastName := getUserData("Enter your last name: ")
		birthDate := getUserData("Enter your birth date(MM/DD/YYYY):")

		var appUser *users.User
		appUser, err := users.New(firstName, lastName, birthDate)

		if err != nil {
			fmt.Println("Error creating user:", err)
			return
		}

		appUser.OutputUserDetails()
	*/

	var appAdmin *users.Admin
	appAdmin, err := users.NewAdmin("admin@gmail.com", "admin123", "John", "Doe", "01/01/1990")
	if err != nil {
		fmt.Println("Error creating admin:", err)
		return
	}
	appAdmin.User.OutputUserDetails()
	appAdmin.User.ClearUserName()
	appAdmin.User.OutputUserDetails()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)

	return value
}
