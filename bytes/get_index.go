package main

import (
	"bytes"
	"runtime"
)

var x = []byte("1234567890")

var ch = make(chan int)

func main() {
	go func() {
		runtime.Gosched()
		// No race report:
		i := bytes.IndexByte(x, '5')

		// Has race report:
		//i := IndexByte(x, '5')
		ch <- i
	}()
	runtime.Gosched()
	for k := range x {
		x[k]++
	}
	println(<-ch)
}

func IndexByte(b []byte, c byte) int {
	for i, x := range b {
		if x == c {
			return i
		}
	}
	return -1
}
