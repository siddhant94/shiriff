package main

import (
	"fmt"
	"shiriff/cmd"
	"shiriff/config"
)

// Hello - Returns Greeting message
func Hello() string {
	return "Hi, Welcome to Shiriff"
}

func main() {
	// fmt.Println(cmd.SayHello())
	fmt.Printf("%+v",config.GetConfig())
	fmt.Printf("\n")
	cmd.SetCommands()
	cmd.StartApp()
}
