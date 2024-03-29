package cmd

import (
	"fmt"
	"os"
	"strings"
	"io/ioutil"
	"shiriff/logs"
	"encoding/json"
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
		if logToFile {
			logging.Error("Unable to open file ", err)
		} else {
			fmt.Println("Unable to open file, ", err)
		}
		return
	}
	// Unmarshall existing JSON of users.
	err = json.Unmarshal([]byte(file), &usersList)
	if err != nil {
		if logToFile {
			logging.Error("Unable to unmarhsall list of users, ", err)
		} else {
			fmt.Println("Unable to unmarhsall list of users, ", err)
		}
		return
	}
	return
}

func updateUsersList(usersList []UserDetails) {
	filepath := DBPATH + USERSLISTFILE
	res, err := json.MarshalIndent(usersList, "", " ")
	if err != nil {
		if logToFile {
			logging.Error("Error marshalling user data in register, ", err)
		} else {
			fmt.Println("Error marshalling user data in register, ", err)
		}
		return
	}
	err = ioutil.WriteFile(filepath, res, 0644)
	if err != nil {
		if logToFile {
			logging.Error("Error writing user data in register, ", err)
		} else {
			fmt.Println("Error writing user data in register, ", err)
		}
		return
	}
}

func writeToLoggedInFileStore(email string) {
	filename := DBPATH + LoggedInUsersFile
	f, err := os.OpenFile(filename, os.O_WRONLY, 0600)
	// Empty contents
	emptyContentsForFile(filename)
	
	if err != nil {
		if logToFile {
			logging.Error("Cannot open file, ", err)
		} else {
			fmt.Println("Cannot open file, ", err)
		}
	}
	
	defer f.Close()
	
	if _, err = f.WriteString(email); err != nil {
		if logToFile {
			logging.Error("Cannot write to  file, ", err)
		} else {
			fmt.Println("Cannot write to  file, ", err)
		}
	}
}

func emptyContentsForFile(filename string) {
	err := os.Truncate(filename, 0)
	if err != nil {
		if logToFile {
			logging.Error("Error truncating file, ",err)
		} else {
			fmt.Println("Error truncating file, ",err)
		}
	}
}

func checkIfFileContains(filepath string, str string) bool {
	b, err := ioutil.ReadFile(filepath)
    if err != nil {
		if logToFile {
			logging.Error("Cannot read file", err)
		} else {
			fmt.Println("Cannot read file", err)
		}
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
		if logToFile {
			logging.Error("Cannot view file ", err)
		} else {
			fmt.Println(err)
		}
		return
    }

    str := string(b)

    fmt.Println(str)
}


func getAccessLevelsFromAccessString(accessString string) string {
	res := ""
	if strings.Contains(accessString, "READ") == true {
		res = res + "R"
	}
	if strings.Contains(accessString, "WRITE") == true {
		res = res + "W"
	}
	if strings.Contains(accessString, "DELETE") == true {
		res = res + "D"
	}
	return res
}

func getFileContents(filepath string) string {
	b, err := ioutil.ReadFile(filepath)
    if err != nil {
		if logToFile {
			logging.Error("Cannot read file ", err)
		} else {
			fmt.Println("Cannot read file", err)
		}
		return ""
	}
	return string(b)
}

func showUsersWithPromp(usersList []UserDetails) {
	var ans string
	for i := 0; i < len(usersList); i++ {
		fmt.Println("Request Access - " + usersList[i].RequestPending + " for User: ")
		fmt.Printf("%+v\n",usersList[i])
		fmt.Println("Enter Y or y to grant else any key for skipping.")
		_, err := fmt.Scanf("%s", &ans)
		if err != nil {
			if logToFile {
				logging.Error("Unable to read input ", err)
			} else {
				fmt.Println("Unable to read input ", err)
			}
			return
		}
		if ans == "Y" || ans == "y" {
			grantAccessToUser(usersList[i].Email)
			fmt.Println("Granted access.")
		} else {
			fmt.Println("Request skipped for user with Email: "+ usersList[i].Email)
		}
	}
	fmt.Println("Finished List. Cheers!")
}

func grantAccessToUser(email string) {
	usersList := getUnmarshalledUsersList()
	for i := 0; i < len(usersList); i++ {
		if usersList[i].Email == email {
			accessLevelSlice := getAccessLevelFromAbbreviation(usersList[i].RequestPending)
			usersList[i].RequestPending = ""
			for _, c := range accessLevelSlice {
				usersList[i].Access += "," + c
			}
			// Check if all access levels
			if strings.Contains(usersList[i].Access, "READ") && strings.Contains(usersList[i].Access, "WRITE") && strings.Contains(usersList[i].Access, "DELETE") {
				usersList[i].Role = SUPERUSER
			}
		}
	}
	updateUsersList(usersList)
}

func getAccessLevelFromAbbreviation(str string) []string {
	var accessString []string
	for _, c := range str {
		switch string(c) {
			case "R":
				accessString = append(accessString, "READ")
			case "W": 
				accessString = append(accessString, "WRITE")
			case "D":
				accessString = append(accessString, "DELETE")
		}
	}
	return accessString
}

func appendToFile(filename string, textToWrite string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		if logToFile {
			logging.Error("Unable to open file, ",err)
		} else {
			fmt.Println("Unable to open file, ",err)
		}
		return
	}
	fmt.Println("Coming here")

	defer f.Close()

	if _, err = f.WriteString(textToWrite); err != nil {
		if logToFile {
			logging.Error("Unable to write to resource file, ",err)
		} else {
			fmt.Println("Unable to write to resource file, ",err)
		}
		return
	}
}

func checkIfAccessPresent(str string, accessLevelToCheckFor string) bool {
	if strings.Contains(str, accessLevelToCheckFor) == true {
		return true
	}
	return false
}