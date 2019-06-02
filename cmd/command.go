package cmd

import (
	"shiriff/cmd/internal/command"
)


const DBPATH = "/home/sid/Desktop/Workspace/go/src/shiriff"
const USERSLISTFILE = "/shiriffDB/users.json"
const LoggedInUsersFile = "/shiriffDB/logged-in-users.txt"
const RESOURCEPATH = DBPATH + "/shiriffDB/resource.txt"
var AccessLevelAbbrToTextMap = map[string]string {
	"R" : "READ",
	"W" : "WRITE",
	"D" : "DELETE",
}

type UserDetails struct {
	UserName string `json:"username"`
	Email	 string `json:"email"`
	Password string `json:"password"`
	Access   string
	RequestPending string `json:"accessRequestsPending, omitempty"`
}

type UserList struct {
	List []UserDetails `json:"users_list"`
}

func StartApp() {
	command.Start()
}

func SetCommands() {
	command:= command.Command{}

	command = getRegisterUserCommand()
	command.AddCommandWithArgs(registerUser)

	command = getLoginUserCommand()
	command.AddCommandWithArgs(loginUser)

	command = getRequestAccessCommand()
	command.AddCommandWithArgs(requestAccess)

	command = getCheckAccessLevelsCommand()
	command.AddCommandWithArgs(checkUserAccessLevels)
}

func getRegisterUserCommand() command.Command {
	command := command.Command {
		Name: "register",
		Description: "Add client as a user",
		Category: "Auth",
		UsageText: "register {UserName} {Email} {Password}",
	}
	return command
}

func getLoginUserCommand() command.Command {
	command := command.Command {
		Name: "login",
		Description: "Log in as an existing user",
		Category: "Auth",
		UsageText: "login {Email} {Password}",
	}
	return command
}

func getRequestAccessCommand() command.Command {
	command := command.Command {
		Name: "requestAccess",
		Description: "Request for Access for a registered user",
		Category: "Access Control",
		UsageText: "requestAccess {Email} {AccessAbbreviations} : Abbreviations are `R`- READ,`W`- WRITE ,`D`-DELETE. Ex - requestAccess abc@gmail.com WD when requesting for Write and Delete",
	}
	return command
}