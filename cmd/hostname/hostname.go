package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Println(hostname)
	names, err := net.LookupCNAME(hostname)
	if err != nil {
		panic(err)
	}
	fmt.Println(names)
}
