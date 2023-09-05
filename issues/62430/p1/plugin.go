package main

import (
	"regexp"
)

func F() {
	_ = regexp.MustCompile(`\w+`)
}

func main() {}