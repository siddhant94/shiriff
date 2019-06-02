package cmd

import (
	"fmt"
)

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
	if len(args)!= 2 {
		fmt.Println("Provide email as well as non-space separated Access Level abbreciations")
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
	if len(args) != 1 {
		fmt.Println("Email required for checking access levels")
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