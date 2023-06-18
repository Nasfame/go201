package main

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
)

type MyStruct struct {
	Name string
}

func main() {
	var a *MyStruct
	fmt.Println("go-", runtime.Version())

	fmt.Printf("a is %p; a==nil is %t \n\n", a, a == nil)

	checkNil(a)
	checkNil(nil)
}

func checkNil(this any) {

	fmt.Printf("%v %p\n", this, this)

	if reflect.ValueOf(this).IsNil() {
		log.Println("reflect says nil", this)
	}

	if this == nil {
		log.Println("this is nil")
	} else {
		log.Println("this is not nil", this)
	}
}
