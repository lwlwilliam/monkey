package main

import (
	"fmt"
	"github.com/lwlwilliam/monkey/repl"
	"os"
	"os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		u.Username)
	fmt.Printf("Fell free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
