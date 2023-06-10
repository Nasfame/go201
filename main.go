package main

import (
	"examples/linkname/customtime"

	_ "examples/linkname/time"
)

func main() {
	customtime.Now()
}
