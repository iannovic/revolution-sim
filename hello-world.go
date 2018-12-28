package main

import (
	"fmt"

	flag "github.com/ogier/pflag"
)

//flags
var (
	user string
)

func main() {
	fmt.Println("hello world")
	flag.Parse()
}

func init() {
	flag.StringVarP(&user, "user", "u", "", "Search Users")
}
