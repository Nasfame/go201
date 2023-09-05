package main

import (
	"fmt"
	"regexp"
)

func F() {
	_ = regexp.MustCompile(`\w+`)
	fmt.Println("hey")
}

func main() {
	fmt.Println("plugin init")
}
