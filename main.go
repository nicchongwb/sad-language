package main

import (
	"fmt"
	"little-cp/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Welcome %s... This is SAD programming language :(\n", user.Username)
	fmt.Printf("Sadly type in commands...\n")
	repl.Start(os.Stdin, os.Stdout)
}
