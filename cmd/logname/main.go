package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		fmt.Fprintln(os.Stderr, "logname: no login name")
		os.Exit(1)
	}
	fmt.Println(usr.Username)
}
