package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/addsict/interpreter-in-go/repl"
)

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s!\nThis is REPL for Monkey language.\n", u.Username)
	repl.Start(os.Stdin, os.Stdout)
}
