package main

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"fmt"
)

type MyStruct struct {
	Name string
}

func main() {
	var a *MyStruct
	hash := md5.New()
	binary.Write(bytes.NewBuffer(nil), binary.LittleEndian, a)
	fmt.Printf("%x", hash.Sum(nil))
}
