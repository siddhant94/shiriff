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

// func init() {
// 		app := cli.NewApp()
// }

func Start() {
	app := cli.NewApp()
	app.Commands = cmds
	app.Name = "SHIRIFF"
	app.Usage = "Access based auth system"
	app.Author = "Siddhant"
	err := app.Run(os.Args)
	if err!= nil {
		fmt.Println("App run failed: ", err)
	}
}

// Command helper functioons to create a command
func (command Command) AddCommandWithArgs() {
	newCommand := cli.Command{
		Name: command.Name,
		Usage: command.Description,
		Category: command.Category,
		UsageText: command.UsageText,
		Action: func(c *cli.Context) error {
			fmt.Printf("Command run for ",c)
			if c.Args().Present() {
				t := c.Args().First()
				fmt.Println("Argument", t)
				return nil
			} else {
				fmt.Println("No argument")
			}
			return nil
		},
	}
	cmds = append(cmds, newCommand)
}