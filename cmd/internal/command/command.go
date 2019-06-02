package command

import (
	"fmt"
	"os"
	"github.com/urfave/cli"
)

// Command - Describes structure for a single command
type Command struct {
	Name        string
	Description string
	Category    string
	UsageText   string
}

var (
	cmds []cli.Command
	// app *cli.App
)

type actionFunctionArgs func(...string)
type actionFunction func()

// Start - starts the app
func Start() {
	app := cli.NewApp()
	app.Commands = cmds
	app.Name = "SHIRIFF"
	app.Usage = "Access based auth system"
	app.Version = "1.0.0"
	app.Author = "Siddhant"
	err := app.Run(os.Args)
	if err!= nil {
		fmt.Println("App run failed: ", err)
	}
}

// AddCommandWithArgs - Command helper functioons to create a command
func (command Command) AddCommandWithArgs(fn actionFunctionArgs) {
	newCommand := cli.Command{
		Name: command.Name,
		Usage: command.Description,
		Category: command.Category,
		UsageText: command.UsageText,
		Action: func(c *cli.Context) error {
			fn(c.Args()...)
			return nil
		},
	}
	cmds = append(cmds, newCommand)
}

// AddCommand - Command helper functioons to create a command
func (command Command) AddCommand(fn actionFunction) {
	newCommand := cli.Command{
		Name: command.Name,
		Usage: command.Description,
		Category: command.Category,
		UsageText: command.UsageText,
		Action: func(c *cli.Context) error {
			fn()
			return nil
		},
	}
	cmds = append(cmds, newCommand)
}