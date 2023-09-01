package main

import "fmt"

type C struct {
}

func Generic[T int | float32](v T) {
	return
}

func main() {
	y := 1
	fmt.Println("main.go")
	_ = 1

}
