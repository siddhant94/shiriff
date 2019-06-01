package command

import "fmt"

type Command struct {
	Name        string
	Description string
	Category    string
	UsageText   string
}

func init() {
	fmt.Println("Command pkg")
}

// Command helper functioons to create a command