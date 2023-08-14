package main

import (
	"fmt"
	"runtime"
)

type S1 struct {
	a, b, c []int
	i       int
}

type S2 struct {
	a, b []int
	m    map[int]int
}

func F(i int, f func(S1, S2, int) int) int {
	return f(
		S1{},
		S2{m: map[int]int{}},
		1>>i,
	)
}

func main() {
	println(runtime.Version())
	println(F(10, func(s1 S1, s2 S2, i int) int {
		fmt.Println(s1, s2, i)
		return i
	}))
	println("Hello, Hiro")
}
