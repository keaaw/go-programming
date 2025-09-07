package main

import "fmt"

func doit(f func([]int), s []int) {
	f(s)
}

func work(s []int) {
	if len(s) == 0 {
		return
	}
	s[0] = 123
}

func main() {
	myslice := []int{1,2,3}
	doit(work, myslice)
	fmt.Println(myslice[0])
}
