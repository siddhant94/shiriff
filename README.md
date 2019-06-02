# Shiriff
CLI based auth with access management built in Go.

## Build Environment
### Server
Distributor ID:	Ubuntu
Description:	Ubuntu 18.04.2 LTS
Release:	18.04
Codename:	bionic

### Go
go1.11.4 linux/amd64

### Setup
Download the source project in $GOPATH/src
For GOPATH, check this link - https://golang.org/cmd/go/#hdr-GOPATH_environment_variable
Further on what value to set can be read from - https://github.com/golang/go/wiki/SettingGOPATH
There are 3 directories inside GOPATH - src, bin, pkg

1) Go to project root (in src) and run command
```go get``` to install project dependencies.
2) For compiling run
```go install```
3) For building binary run (It will create binary in root directory itself)
```go build```

Voila! Setup done.
On executing the binary the CLI tool will run displaying various sub commands it supports.
```shiriff --help``` or ```shiriff -h``` would also list out the sub commands it supports.
To get Usage help for a sub-command (register) run
```shiriff help register```
(```shiriff``` refers to the binary we just created)

### Packages Used
For creating command line utility - https://github.com/urfave/cli

### Application Flow
Below are the commands available for Shiriff. It's divided into 3 categories that are `Access control`, `Authentication`,
`Resource`. The description of the respective commands are written next to it.
  Access Control:
     requestAccess     Request for Access for a registered user
     checkAccessLevel  Check Access Levels for a registered user.
     makeSuperuser     Would grant superuser role to the user specified. Accepts secret as input.

   Authentication:
     register  Add client as a user
     login     Log in as an existing user

   Resource:
     updateResource  Update resource file by appending text to it.
     listUsers       List the users for Shiriff.
     deleteResource  Delete resource file contents.
     grantAccess     Grant access to pending access-requests for users.

##### Storage
  Folder `shiriffDB` in the project directory stores application data in file system.
It has 3 files in use namely -
users.json - used to store all redistered users on shiriff.
logged-in-users.txt - used to store logged-in user
resource.txt - is the resource file for which access control is in present.

##### Logs
There is a file called `shiriff.logs` which will have persistent logs.

##### Config
`shiriff/config/config.go` project path has config file. It has the `APPSECRET` which is used for making a user 'superuser' i.e. admin. According to shiriff, there are 2 roles a user has 'normal' and 'superuser'. superuser would have access to all action_types i.e. READ, WRITE, DELETE. By default, normal user would have 'READ' access.
Also in this config file you can set other constants such as `ENVIRONMENT` to DEBUG for dev env or PRODUCTION for production env and also `LOGPATH` if you wish to have some other file than shiriff.logs.

Resource Commands listUsers, grantAccess are only accessible by superuser. Others such as updateResource and deleteResource are accessible by any user with proper access rights i.e. with WRITE and DELETE access respectively.

AuthenticationCommands - register takes in username, email and password to register a user.
login - takes in email and password and matches against registered users to grant access.

Access Control Commands - 
requestAccess - It enables a logged in user to request for additional access rights, say for WRITE and DELETE
checkAccessLevel - It takes in email and informs what access levels user owns.
makeSuperuser - It takes in email id whom we need to make superuser and then prompts for `APPSECRET` which is in config.






