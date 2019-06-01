package main

import (
	"fmt"
	"shiriff/cmd"
	"shiriff/config"
	"shiriff/internal/command"
)

// Hello - Returns Greeting message
func Hello() string {
	return "Hi, Welcome to Shiriff"
}

func main() {
	fmt.Println(Hello())
	fmt.Println(cmd.SayHello())
	fmt.Printf("%+v",config.GetConfig())
}
