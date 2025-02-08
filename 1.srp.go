package main

import "fmt"

/*
Single Responsibility Principle (SRP):
A struct should have only one reason to change.
*/

// `User` only holds data.
type User struct {
	FirstName string
	LastName  string
}

// Separate struct for formatting user data.
type UserFormatter struct{}

func (f UserFormatter) GetFullName(u User) string {
	return u.FirstName + " " + u.LastName
}

// Separate struct for database operations.
type UserRepository struct{}

func (r *UserRepository) Save(u User) error {
	fmt.Println("Saving user:", u.FirstName, u.LastName)
	return nil
}

// func main() {
// 	user := User{"John", "Doe"}
// 	formatter := UserFormatter{}
// 	repo := UserRepository{}

// 	fmt.Println("Full Name:", formatter.GetFullName(user))
// 	repo.Save(user)
// }
