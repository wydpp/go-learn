package command

import (
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The flag testing!")
}

func Hello() {
	flag.Parse()
	fmt.Printf("Hello %s\n", name)
}
