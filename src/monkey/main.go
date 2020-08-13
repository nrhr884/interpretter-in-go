package main

import (
	"fmt"
	"os"
	user2 "os/user"

	"github.com/noissefnoc/book-writing-an-interpreter-in-go/monkey/repl"
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
