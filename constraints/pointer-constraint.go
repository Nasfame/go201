package main

import "fmt"


func f[T ~bool](x *T) *T {
	return x
}

func f1[T ~bool, P *T](x P) P { return x }

// func isPointer[T any](v T) bool {
// 	return true
// }

func isPointer[T any, P ~*T](x P) bool { return true }

func main() {

	x := true
	f(&x)
	f1(&x)
	fmt.Println(isPointer(&x))

}
