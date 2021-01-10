// +build !windows
// +build !plan9

package main

import (
	"os"
	"os/exec"
)

func chroot(root string, args []string) error {
	if len(args) == 0 {
		args = []string{shellPath, "-i"}
	}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.SysProcAttr.Chroot = root

	return cmd.Run()
}
