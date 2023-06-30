package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
)

func F[T any](v ...T) {
	var zero T
	fmt.Printf("T:%T\n", &zero)
}

func main() {
	fmt.Printf("go-version: %s\n", runtime.Version())
	F(io.Discard)
	F(os.Stdout)
	F[*os.File]()
	F(io.Discard, os.Stdout)
	F(os.Stdout, io.Discard)
}
