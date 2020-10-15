package main

import (
	"flag"
	"fmt"
	"os"
)

// flags
var (
	logical bool
	// physical bool
)

func init() {
	flag.BoolVar(&logical, "L", true, "Display the logical current directory")
	// TODO: physical
}

func main() {
	res, err := os.Getwd()
	if err != nil {
		panic(err) // TODO: testing coverage here
	}
	fmt.Println(res)
}
