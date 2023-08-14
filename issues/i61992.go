package main


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
		i<<2,
	)
}

func main() {
	println("Hello, Hiro")
}
