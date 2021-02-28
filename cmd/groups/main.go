package main

import (
	"flag"
	"fmt"
	osuser "os/user"
	"strings"

	"github.com/yimoucao/go-coreutils/pkg/user"
)

// TODO: --version, --help

func main() {
	flag.Parse()
	users := flag.Args()

	if len(users) == 0 {
		current, err := osuser.Current()
		if err != nil {
			fmt.Println(err)
			return
		}
		groups, err := user.GroupNamesOfUser(current)
		if err != nil {
			fmt.Printf("groups: %v", err)
			return
		}
		fmt.Println(strings.Join(groups, " "))
	}

	for _, u := range users {
		groups, err := user.GroupNamesOfUsername(u)
		if err != nil {
			fmt.Printf("groups: '%s': %v", u, err)
			continue
		}
		fmt.Println(u, ":", strings.Join(groups, " "))
	}
}
