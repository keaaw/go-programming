package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	fmt.Printf("%v args\n", len(os.Args))
}
