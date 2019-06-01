package cmd

import (
	"fmt"
	"os"
	// "path"
	// "runtime"
	"shiriff/cmd/internal/command"
)


const DBPATH = "/home/sid/Desktop/Workspace/go/src/shiriff"
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
	command.AddCommandWithArgs(registerUser)
}

func getRegisterUserCommand() command.Command {
	command := command.Command {
		Name: "register",
		Description: "Add client as a user",
		Category: "Auth",
	}
	return command
}

func registerUser(args ...string) {
	firstName := args[0]
	lastName := args[1]
	email := args[2]
	filepath := DBPATH + "/shiriffDB/users.txt";

	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file : \n", err)
		return
	}
	_, err = fmt.Fprintln(f, firstName+"		"+lastName+"			"+email)
    if err != nil {
        fmt.Println(err)
                f.Close()
        return
	}
}