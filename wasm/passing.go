package main

import "fmt"
import "time"

func main() {
	fmt.Println("begin")
	middle()
	fmt.Println("end")

}


func middle(){
	go func() {
		fmt.Println("In goroutine 1")
		for {
			// Perform some work or computation here
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		fmt.Println("In goroutine 2")
		for {
			// Perform some work or computation here
			time.Sleep(1 * time.Second)
		}
	}()

	// Keep the main goroutine running indefinitely
	select {}

}