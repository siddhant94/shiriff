package command

import (
	"fmt"
	"os"
	"github.com/urfave/cli"
)

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

// func init() {
// 		app := cli.NewApp()
// }

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

// Command helper functioons to create a command
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