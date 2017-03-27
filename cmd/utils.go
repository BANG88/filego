package cmd

import (
	"fmt"
)

// filego
func log(format string, a ...interface{}) {
	fmt.Printf(fmt.Sprintf("FILEGO: %s\n", format), a...)
}
