package time

import (
	"fmt"
	_ "unsafe"
)

//go:linkname timenow examples/linkname/customtime.Now
func timenow() {
	fmt.Println("linked:time")
}
