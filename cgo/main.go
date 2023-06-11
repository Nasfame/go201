package main

/*
typedef int foo;

typedef char* stringer;

*/
import "C"

import "fmt"

type foo = C.foo

type stringer = C.stringer

func (foo) method() int { return 123 }

func (f *foo) method1() {
	fmt.Println("hey")
}

func main() {
	var x foo
	println(x.method()) // "123"
	x.method1()
}
