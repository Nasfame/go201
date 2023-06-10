package customtime

import (
	"fmt"
	_ "test/time"
)

// Present in test/time
func Now() {
	fmt.Println("from customtime")
}
