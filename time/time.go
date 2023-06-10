package time

import (
	"fmt"
	_ "unsafe"
)

//go:linkname timenow test/customtime.Now
func timenow() {
	fmt.Println("time")
}
