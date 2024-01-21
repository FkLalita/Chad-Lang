package main

import (
	"fmt"
	"os"
	"os/user"
	"github.com/FkLalita/Chad-Lang/repl"
)


func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello, %s....Welcome To The Chad Programming Language\n", user.Username)
	fmt.Printf("What would you like to do today\n Enter Command\n")
	repl.START(os.Stdin, os.Stdout)

}