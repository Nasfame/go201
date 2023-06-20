package main

import "fmt"

func main() {
	defer func() {
		recover()
		fmt.Println("Checkpoint 1")
		//panic(1)
	}()
	defer func() {
		recover()
		fmt.Println("Checkpoint 2")
		panic(2)
	}()
	panic(999)

	//Output: doesn't panic in the end. Since all panics are caught in defer functions.
}
