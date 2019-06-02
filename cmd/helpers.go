package cmd

import (
	"fmt"
	"os"
	"strings"
	"encoding/json"
	"io/ioutil"
	"shiriff/cmd/internal/command"
)

func validateArgumentsLength(args []string, expLength int, errMsg string) bool {
	if len(args) != expLength {
		fmt.Println(errMsg)
		return false
	}
	return true
}

func getCheckAccessLevelsCommand() command.Command {
	command := command.Command {
		Name: "checkAccessLevel",
		Description: "Check Access Levels for a registered user.",
		Category: "Access Control",
		UsageText: "checkAccessLevel {Email}",
	}
	return command
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

func updateUsersList(usersList []UserDetails) {
	filepath := DBPATH + USERSLISTFILE
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

func getAccessLevelsForAUser(email string) string {
	usersList := getUnmarshalledUsersList()
	for i := 0; i < len(usersList); i++ {
		if email == usersList[i].Email {
			return usersList[i].Access
		}
	}
	return ""
}

func viewResource(filepath string) {
	b, err := ioutil.ReadFile(filepath)
    if err != nil {
        fmt.Println(err)
    }

    str := string(b)

    fmt.Println(str)
}