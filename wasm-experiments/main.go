package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("CPU", runtime.NumCPU())
	mainBad()
	fmt.Println("end")
}

func mainBad() {
	go func() {
		print("in go routine 1\n")
		for {

		}
	}()
	//uncomment me to see the print above in the output
	//runtime.Gosched()
	go func() {
		print("in go routine 2\n")
		for {

		}
	}()
	for {
	}

}
