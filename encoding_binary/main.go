package main

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"log"
)

type MyStruct struct {
	Name string
}

func main() {
	var a *MyStruct
	hash := md5.New()

	fmt.Printf("a is %p; a==nil is \n",a,a==nil)
	err := binary.Write(bytes.NewBuffer(nil), binary.LittleEndian, a)
	if err!=nil{
		log.Fatal("main ended with panic", err)
	}
	fmt.Printf("%x", hash.Sum(nil))
}
