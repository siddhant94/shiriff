package cmd

import (
	// "fmt"
	"shiriff/cmd/internal/command"
)

// func SayHello() string {
// 	fmt.Println("Yo")
// 	setCommands()
// 	return "Hi from Command pkg"
// }

func StartApp() {
	command.Start()
}

func SetCommands() {
	command:= command.Command{}

	command = getRegisterUserCommand()
	command.AddCommandWithArgs()
}

func getRegisterUserCommand() command.Command {
	command := command.Command {
		Name: "Register a User",
		Description: "Add client as a user",
		Category: "Third Party",
	}
	return command
}