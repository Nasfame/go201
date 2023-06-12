package main

import (
	"fmt"
	"reflect"
	_ "unsafe"
)

//go:linkname printReflectValue reflect.Value.String
func printReflectValue(v reflect.Value) string

func main() {
	str := printReflectValue(reflect.ValueOf(42))
	fmt.Println(str)
}
