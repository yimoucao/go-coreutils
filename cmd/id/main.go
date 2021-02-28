package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yimoucao/go-coreutils/pkg/user"
)

// TODO: --version, --help, options

func main() {
	flag.Parse()
	users := flag.Args()
	if len(users) > 1 {
		fmt.Fprintf(os.Stderr, "id: extra operand '%s'", users[1])
		return
	}

	var u *user.User
	var err error
	if len(users) == 0 {
		u, err = user.Current()
	} else {
		u, err = user.Lookup(users[0])
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	groups, err := user.GroupsOfUser(u.User)
	if err != nil {
		fmt.Printf("groups: %v", err)
		return
	}

	fmt.Printf("uid=%s(%s)", u.Uid, u.Username)
	fmt.Printf(" gid=%s(%s)", u.Gid, u.Group.Name)
	fmt.Print(" groups=")
	for i, g := range groups {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Printf("%s(%s)", g.Gid, g.Name)
	}
	fmt.Println()
	return
}
