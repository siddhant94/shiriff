package cmd

import (
	"fmt"
	// "shiriff/internal/command"
)

func SayHello() string {
	fmt.Println("Yo")
	return "Hi from Command pkg"
}

// func registerUserCommand() {
// 	command.Command {
// 		Name: "Register a User",
// 		Description: "Add client as a user"
// 		Category: "Third Party"
// 	}
// }