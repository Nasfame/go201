package main

import (
	"fmt"
	"reflect"
)

type S struct{

}

func main(){
	s := S{}

	if reflect.ValueOf(s).IsZero(){
		fmt.Println("hey")
	}
}