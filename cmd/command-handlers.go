package cmd

import (
	"fmt"
	"strings"
	"shiriff/config"
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
		Role: NORMAL,
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
				fmt.Println("Your Role - ",usersList[i].Role)
				viewResource(RESOURCEPATH)
			} else {
				fmt.Println("You shall not pass! (INCORRECT PASSWORD)")
			}
			return
		}
	}
	fmt.Println("Oops, you need to register first")
	return

}

func requestAccess(args ...string) {
	msg := "Provide email as well as non-space separated Access Level abbreciations"
	if !validateArgumentsLength(args, 2, msg) {
		return
	}
	email := args[0]
	accessAbbr := args[1]
	var r, w, d bool
	// Check if access abbreviations are correct. TODO - shopuld come from config
	for _, val := range accessAbbr {
		oneAbbr := string(val)
		if oneAbbr != "R" && oneAbbr != "W" && oneAbbr != "D" {
			fmt.Println("Unknown access abbreviation (Possible values: R, W, D (READ, WRITE, DELETE). Got "+oneAbbr)
			return
		}
		switch oneAbbr {
		case "R":
			r = true
		case "W":
			w = true
		case "D":
			d = true
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
			// TODO : Remove nested if for more readability OR change approach.
			accessString := usersList[i].Access
			abb := getAccessLevelsFromAccessString(accessString)
			// below satisfied check i.e. r is requested and is already present
			if r == true && strings.Contains(abb, "R") {
				accessAbbr = strings.Replace(accessAbbr, "R", "", -1)
				fmt.Println("R(READ) access already present")
			}
			if w == true && strings.Contains(abb, "W") {
				accessAbbr = strings.Replace(accessAbbr, "W", "", -1)
				fmt.Println("W(WRITE) access already present")
			}
			if d == true && strings.Contains(abb, "D") {
				accessAbbr = strings.Replace(accessAbbr, "D", "", -1)
				fmt.Println("D(DELETE) access already present")
			}
			// Case of requested access levels already present
			if len(accessAbbr) < 1 {
				fmt.Println("Requested access levels already present")
				return
			}
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

func grantSuperuserRoleToUser(args ...string) {
	msg := "Please enter email whom you wish to make superuser"
	if !validateArgumentsLength(args, 1, msg) {
		return
	}
	email := args[0]
	var secret string
	fmt.Println("Enter the secret")
	_, err := fmt.Scanf("%s", &secret)
	if err != nil {
		fmt.Println("Unable to read input ", err)
		return
	}
	// Check secret
	if secret != config.APPSECRET {
		fmt.Println("Incorrect Secret. Request denied.")
		return
	}
	// Get existing users and update role
	usersList := getUnmarshalledUsersList()
	for i := 0; i < len(usersList); i++ {
		if usersList[i].Email == email {
			usersList[i].Role = SUPERUSER
			// Also since user has all privileges, unset `accessRequestsPending` and update `access`
			usersList[i].RequestPending = ""
			usersList[i].Access = AccessLevelAbbrToTextMap["R"] + "," + AccessLevelAbbrToTextMap["W"] + "," + AccessLevelAbbrToTextMap["D"]
		}
	}
	updateUsersList(usersList)
	fmt.Println("User Role Updated successfully for " + email)
}

func listUsersInfo() {
	var isSuperUser bool
	// Get logged in user
	loggedInUsersFilePath := DBPATH + LoggedInUsersFile
	userEmail := getFileContents(loggedInUsersFilePath)
	usersList := getUnmarshalledUsersList()
	for i := 0; i < len(usersList); i++ {
		if usersList[i].Email == userEmail {
			if usersList[i].Role == SUPERUSER {
				isSuperUser = true
			}
		}
	}
	if isSuperUser {
		fmt.Println("Shiriff Users List")
		fmt.Printf("%+v\n", usersList)
	} else {
		fmt.Println("You are not superuser. Unable to grant request.")
	}
}

func grantAccessToPendingUserRequests() {
	var isSuperUser bool
	// Get logged in user
	loggedInUsersFilePath := DBPATH + LoggedInUsersFile
	userEmail := getFileContents(loggedInUsersFilePath)

	var pendingReqUsers []UserDetails

	usersList := getUnmarshalledUsersList()
	for i := 0; i < len(usersList); i++ {
		if usersList[i].Email == userEmail {
			if usersList[i].Role == SUPERUSER {
				isSuperUser = true
			}
		}
		if usersList[i].RequestPending != "" {
			pendingReqUsers = append(pendingReqUsers, usersList[i])
		}
	}
	if isSuperUser {
		fmt.Println("Users with pending access requests")
		showUsersWithPromp(pendingReqUsers)
	} else {
		fmt.Println("You are not superuser. Unable to grant request.")
	}

}

func updateResourceFile(args ...string) {
	msg := "Enter something to update resource file"
	if !validateArgumentsLength(args, 1, msg) {
		return
	}
	// Check for Write access.
	var writeAccess = false
	loggedInUsersFilePath := DBPATH + LoggedInUsersFile
	userEmail := getFileContents(loggedInUsersFilePath)

	usersList := getUnmarshalledUsersList()
	for i := 0; i < len(usersList); i++ {
		if usersList[i].Email == userEmail {
			writeAccess = checkIfAccessPresent(usersList[i].Access, "WRITE")
		}
	}
	if writeAccess {
		appendToFile(RESOURCEPATH, args[0])
		fmt.Println("Resource Updated")
		viewResource(RESOURCEPATH)
	} else {
		fmt.Println("Sorry you dont have correct access. Register your request for `WRITE` access to admin.")
	}
}

func deleteResourceFile() {
	// Check for DELETE access.
	var delAccess = false
	loggedInUsersFilePath := DBPATH + LoggedInUsersFile
	userEmail := getFileContents(loggedInUsersFilePath)

	usersList := getUnmarshalledUsersList()
	for i := 0; i < len(usersList); i++ {
		if usersList[i].Email == userEmail {
			delAccess = checkIfAccessPresent(usersList[i].Access, "DELETE")
		}
	}

	if delAccess {
		// Again prompt before deleting
		fmt.Println("Are you sure you want to delete the resource? Enter y or Y to confirm or any other key to abort.")
		ans := "no"
		_, err := fmt.Scanf("%s", &ans)
		if err != nil {
			fmt.Println("Unable to read input ", err)
			return
		}
		if ans == "y" || ans == "Y" {
			emptyContentsForFile(RESOURCEPATH)
		} else {
			fmt.Println("Cancelled delete command. Resource intact.")
		}
	} else {
		fmt.Println("Sorry you dont have correct access. Register your request for `DELETE` access to admin.")
	}

}