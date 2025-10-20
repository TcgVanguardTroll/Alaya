package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/TcgVanguardTroll/Alaya/main/alaya_repl"
)

func main() {
	current, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome to the Alaya Programming Language!\n", current.Username)
	alaya_repl.Start(os.Stdin, os.Stdout)
}
