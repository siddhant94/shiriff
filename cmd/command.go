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
type UserRole string
//NORMAL - Of R,W,D - can have R or W or both but not W.
const NORMAL UserRole = "normal"
//SUPERUSER - Of R,W,D - will have all three.
const SUPERUSER UserRole = "superuser"

// UserDetails - Details for one user.
type UserDetails struct {
	UserName string `json:"username"`
	Email	 string `json:"email"`
	Password string
	Access   string	`json:"access"`
	RequestPending string `json:"accessRequestsPending,omitempty"`
	Role 	UserRole `json:"user_role"`
}

func StartApp() {
	command.Start()
}

// SetCommands - Sets the Commands available and links handlers.
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

	command = getGrantSuperUserRoleCommand()
	command.AddCommandWithArgs(grantSuperuserRoleToUser)

	command = getListUsersCommand()
	command.AddCommand(listUsersInfo)

	command = getAccessToUsersCommand()
	command.AddCommand(grantAccessToPendingUserRequests)
}

func getRegisterUserCommand() command.Command {
	command := command.Command {
		Name: "register",
		Description: "Add client as a user",
		Category: "Authentication",
		UsageText: "register {UserName} {Email} {Password}",
	}
	return command
}

func getLoginUserCommand() command.Command {
	command := command.Command {
		Name: "login",
		Description: "Log in as an existing user",
		Category: "Authentication",
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

func getGrantSuperUserRoleCommand() command.Command {
	command := command.Command {
		Name: "makeSuperuser",
		Description: "Would grant superuser role to the user specified. Accepts secret as input.",
		Category: "Access Control",
		UsageText: "makeSuperuser {secret} : Accepts app secret and grants all access to specified user.",
	}
	return command
}

func getListUsersCommand() command.Command {
	command := command.Command {
		Name: "listUsers",
		Description: "List the users for Shiriff.",
		Category: "Resource",
		UsageText: "listUsers: List the users registered in Shiriff and displays their information.",
	}
	return command
}

func getAccessToUsersCommand() command.Command {
	command := command.Command {
		Name: "grantAccess",
		Description: "Grant access to pending access-requests for users.",
		Category: "Resource",
		UsageText: "grantAccess: Grant access to users for pending access requests.",
	}
	return command
}