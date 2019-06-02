package cmd

import (
	"fmt"
)

func registerUser(args ...string) {
	msg := "Please provide all the details - username, email and password"
	if !validateArgumentsLength(args, 3, msg) {
		return
	}

	userName := args[0]
	email := args[1]
	password := args[2]

	oneUser := UserDetails{
		UserName: userName,
		Email: email,
		Password: password,
		Access: AccessLevelAbbrToTextMap["R"], // By default READ Access given
	}

	// Get existing users and append new user to the list.
	usersList := getUnmarshalledUsersList()
	// Add new user to existing list.
	usersList = append(usersList, oneUser)

	updateUsersList(usersList)
	fmt.Println("Write Successful. Successfully registered "+ userName)
	return
}

func loginUser(args ...string) {
	msg := "Please provide all the details - email and password"
	if !validateArgumentsLength(args, 2, msg) {
		return
	}
	email := args[0]
	password:= args[1]
	// Check if user logged in already
	loggedInUsersFilePath := DBPATH + LoggedInUsersFile
	res := checkIfFileContains(loggedInUsersFilePath, email)
	if res == true {
		fmt.Println("You are already logged in with "+email)
		viewResource(RESOURCEPATH)
		return
	}
	// Check users list for credentials
	usersList := getUnmarshalledUsersList()
	for i := 0; i < len(usersList); i++ {
		if usersList[i].Email == email {
			if usersList[i].Password == password {
				//Write to logged in users store. TODO - Add check if write was successful
				writeToLoggedInFileStore(usersList[i].Email)
				fmt.Println("Yay! you are now logged in!")
				fmt.Println("Your Access Levels - ",usersList[i].Access)
				viewResource(RESOURCEPATH)
				return
			}
			fmt.Println("You shall not pass! (INCORRECT PASSWORD)")
		} else {
			fmt.Println("Oops, you need to register first")
			return
		}
	}

}

func requestAccess(args ...string) {
	msg := "Provide email as well as non-space separated Access Level abbreciations"
	if !validateArgumentsLength(args, 2, msg) {
		return
	}
	email := args[0]
	accessAbbr := args[1]
	// Check if access abbreviations are correct. TODO - shopuld come from config
	for _, val := range accessAbbr {
		oneAbbr := string(val)
		if oneAbbr != "R" && oneAbbr != "W" && oneAbbr != "D" {
			fmt.Println("Unknown access abbreviation (Possible values: R- READ, W-WRITE, D-DELETE). Got "+oneAbbr)
			return
		}
	}
	// Check if user is logged in
	loggedInUsersFilePath := DBPATH + LoggedInUsersFile
	res := checkIfFileContains(loggedInUsersFilePath, email)
	if res == false {
		fmt.Println("You need to login to request access for "+email)
		return
	}
	// Add it to request pending in user state.
	usersList := getUnmarshalledUsersList()
	for i := 0; i < len(usersList); i++ {
		if usersList[i].Email == email {
			usersList[i].RequestPending = accessAbbr
		}
	}
	updateUsersList(usersList)
	fmt.Println("Request registered successfully for - ", accessAbbr)
}

func checkUserAccessLevels(args ...string) {
	msg := "Email required for checking access levels"
	if !validateArgumentsLength(args, 1, msg) {
		return
	}
	email := args[0]
	loggedInUsersFilePath := DBPATH + LoggedInUsersFile
	res := checkIfFileContains(loggedInUsersFilePath, email)
	if res == false {
		fmt.Println("You need to login to check access levels for "+email)
		return
	}
	accessLevels := getAccessLevelsForAUser(email)
	fmt.Println("Access Levels for "+ email + " is " + accessLevels)
}