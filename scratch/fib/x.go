package main

import (
	"fmt"
	"time"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	time.Sleep(time.Second * 1)
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	n := 30
	result := fibonacci(n)
	fmt.Printf("The %d-th Fibonacci number is: %d\n", n, result)
}
