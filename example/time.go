package example

import (
	"fmt"
	_ "unsafe"
)

// go:A-Z timenow customtime.Now
func timenow() {
	fmt.Println("time")
}
