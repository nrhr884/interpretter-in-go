package main

import (
	"fmt"
	"monkey/repl"
	"os"
	user2 "os/user"
)

var (
	Version  string
	Revision string
)

func main() {
	user, err := user2.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programing language!\n",
		user.Username)
	fmt.Printf("version: %s(%s)\n", Version, Revision)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
