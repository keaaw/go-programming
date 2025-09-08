
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i, arg := range os.Args {
		s += sep + arg
		if i == 0 {
			sep = " "
		}
	}
	fmt.Println(s)
}
