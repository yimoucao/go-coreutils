package main

import (
	"fmt"
	"syscall"
)

func main() {
	sysinfo := &syscall.Sysinfo_t{}
	if err := syscall.Sysinfo(sysinfo); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", sysinfo)
}
