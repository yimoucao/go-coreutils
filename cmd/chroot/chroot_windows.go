// +build windows

package main

import "errors"

func chroot(root string, args []string) error {
	return errors.New("not implemented for windows")
}
