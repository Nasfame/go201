package main

func f[T ~bool](x *T) *T {
	return x
}

func f1[T ~bool, P *T](x P) P { return x }

// func isPointer[T any](v T) bool {
// 	return true
// }

func main() {

	x := true
	f(&x)
	f1(&x)

}
