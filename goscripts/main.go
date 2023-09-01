//go:build ignore

package main

import "fmt"

type C struct {
}

func Generic[T int | float32]( v T) {
	return 
}


func main() {
	fmt.Println("main.go")
	
}
