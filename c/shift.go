package main

import "fmt"

func main() {
	x := 16
	fmt.Println(x & 2)
	fmt.Println(x | 2)
	fmt.Println(x ^ 2)
	fmt.Println(x << 2)
	fmt.Println(x >> 2)
	fmt.Println(x)
}
