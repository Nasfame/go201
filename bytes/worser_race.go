package main

import "runtime"

// No race report (because it calls runtime.memequal):
var x = [1024]byte{}

// Has race report (because the equality is generated in-line):
//var x = [3]byte{}

var ch = make(chan bool)

func main() {
	go func() {
		runtime.Gosched()
		var y = [len(x)]byte{}
		eq := x == y
		ch <- eq
	}()
	runtime.Gosched()
	for k := range x {
		x[k]++
	}
	println(<-ch)
}

// https://github.com/golang/go/issues/61204#issuecomment-1624130980
