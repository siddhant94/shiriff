package cmd

import (
	"fmt"
	"os"
	"strings"
	"encoding/json"
	"io/ioutil"
	"shiriff/cmd/internal/command"
)


const DBPATH = "/home/sid/Desktop/Workspace/go/src/shiriff"
const USERSLISTFILE = "/shiriffDB/users.json"
const LoggedInUsersFile = "/shiriffDB/logged-in-users.txt"

type UserDetails struct {
	UserName string `json:"username"`
	Email	 string `json:"email"`
	Password string `json:"password"`
	Access   string
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
		Description: "Log in existing user",
		Category: "Auth",
		UsageText: "login {Email} {Password}",
	}
	return command
}

func registerUser(args ...string) {
	if len(args) != 3 {
		fmt.Println("Please provide all the details - username, email and password")
		return
	}

	userName := args[0]
	email := args[1]
	password := args[2]

	oneUser := UserDetails{
		UserName: userName,
		Email: email,
		Password: password,
		Access: "READ", // By default READ Access given
	}

	filepath := DBPATH + USERSLISTFILE
	// Get existing users and append new user to the list.
	usersList := getUnmarshalledUsersList()
	// Add new user to existing list.
	usersList = append(usersList, oneUser)

	res, err := json.MarshalIndent(usersList, "", " ")
	if err != nil {
		fmt.Println("Error marshalling user data in register, ", err)
		return
	}
	err = ioutil.WriteFile(filepath, res, 0644)
	if err != nil {
		fmt.Println("Error writing user data in register, ", err)
		return
	}
	fmt.Println("Write Successful. Successfully registered "+ userName)
	return
}

func loginUser(args ...string) {
	if len(args) != 2 {
		fmt.Println("Please provide all the details - email and password")
		return
	}
	email := args[0]
	password:= args[1]
	// Check if user logged in already
	loggedInUsersFilePath := DBPATH + LoggedInUsersFile
	res := checkIfFileContains(loggedInUsersFilePath, email)
	if res == true {
		fmt.Println("You are already logged in with "+email)
		return
	}
	// Check users list for credentials
	usersList := getUnmarshalledUsersList()
	for i := 0; i < len(usersList); i++ {
		if usersList[i].Email == email {
			if usersList[i].Password == password {
				//Write to logged in users store.
				writeToLoggedInFileStore(usersList[i].Email)
				fmt.Println("Yay! you are now logged in!")
				return
			} else {
				fmt.Println("You shall not pass (INCORRECT PASSWORD)")
			}
		} else {
			fmt.Println("Oops, you need to register first")
		}
	}

}

func getUnmarshalledUsersList() (usersList []UserDetails) {
	filepath := DBPATH + USERSLISTFILE
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("Unable to open file, ", err)
		return
	}
	// Unmarshall existing JSON of users.
	err = json.Unmarshal([]byte(file), &usersList)
	if err != nil {
		fmt.Println("Unable to unmarhsall list of users, ", err)
		return
	}
	return
}

func writeToLoggedInFileStore(email string) {
	filename := DBPATH + LoggedInUsersFile
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		// panic(err)
		fmt.Println("Cannot open file, ", err)
	}
	
	defer f.Close()
	
	if _, err = f.WriteString(email+"\n"); err != nil {
		// panic(err)
		fmt.Println("Cannot write to  file, ", err)
	}
}

func checkIfFileContains(filepath string, str string) bool {
	b, err := ioutil.ReadFile(filepath)
    if err != nil {
		fmt.Println("Cannot read file", err)
		return false
    }
    s := string(b)
    // check whether s contains substring str
	if strings.Contains(s, str) == false {
		return false
	}
	return true
}