package main

import "fmt"

func main() {
	printMe("hi")
	printMe(0)
	printMe(make([]byte, 3))
}

func printMe(arg any) {
	fmt.Println(arg)
	arg = 3
	fmt.Printf("%T\n", arg)
	arg = "hi"
	fmt.Printf("%T\n", arg)
	arg = make([]any, 2)
	fmt.Printf("%T\n", arg)
}
