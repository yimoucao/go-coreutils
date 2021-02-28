// +build !windows
// +build !plan9

package user

import "os/user"

type User struct {
	*user.User
	// Group is the user's primary group
	Group *user.Group
}

func Current() (*User, error) {
	u, err := user.Current()
	if err != nil {
		return nil, err
	}
	g, err := user.LookupGroupId(u.Gid)
	if err != nil {
		return nil, err
	}
	return &User{User: u, Group: g}, nil
}

func Lookup(username string) (*User, error) {
	u, err := user.Lookup(username)
	if err != nil {
		return nil, err
	}
	g, err := user.LookupGroupId(u.Gid)
	if err != nil {
		return nil, err
	}
	return &User{User: u, Group: g}, nil
}

func GroupsOfUsername(username string) ([]*user.Group, error) {
	u, err := user.Lookup(username)
	if err != nil {
		return nil, err
	}
	return GroupsOfUser(u)
}

func GroupsOfUser(u *user.User) ([]*user.Group, error) {
	gids, err := u.GroupIds()
	if err != nil {
		return nil, err
	}
	var grps []*user.Group
	for _, gid := range gids {
		g, err := user.LookupGroupId(gid)
		if err != nil {
			continue
		}
		grps = append(grps, g)
	}
	return grps, nil
}

func GroupNamesOfUsername(username string) ([]string, error) {
	grps, err := GroupsOfUsername(username)
	if err != nil {
		return nil, err
	}
	var gnames []string
	for _, g := range grps {
		gnames = append(gnames, g.Name)
	}
	return gnames, nil
}

func GroupNamesOfUser(u *user.User) ([]string, error) {
	grps, err := GroupsOfUser(u)
	if err != nil {
		return nil, err
	}
	var gnames []string
	for _, g := range grps {
		gnames = append(gnames, g.Name)
	}
	return gnames, nil
}
