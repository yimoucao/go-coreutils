package main

import (
	"fmt"
	"time"
)

/*
environment variables:
TZ='America/Los_Angles' date

read date
date -d
date --date

*/

func main() {
	fmt.Println(time.Now().Format(time.UnixDate))
}
