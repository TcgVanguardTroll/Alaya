package main

import (
	"Alaya/main/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	current, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome to the Alaya Programming Language!\n", current.Username)
	repl.Start(os.Stdin, os.Stdout)
}
