package main

import (
	"fmt"
	"shiriff/cmd"
	"shiriff/config"
)

func main() {
	// fmt.Println(cmd.SayHello())
	fmt.Printf("%+v",config.GetConfig())
	fmt.Printf("\n")
	cmd.SetCommands()
	cmd.StartApp()
}
