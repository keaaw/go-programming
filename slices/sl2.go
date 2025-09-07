package main

import "fmt"

func main() {
	a := [4]int{1,2,3,4}

	s := a[:]
	sliceinfo(s)

	s = append(s, 7)
	sliceinfo(s)

	s2 := append(s, 8, 9, 10)
	sliceinfo(s2)

	s2 = append(s2, 11)
	sliceinfo(s2)

	copy(s2[1:], s)
	sliceinfo(s2)

	s2 = append(s2, s2...)
	sliceinfo(s2)
}

func sliceinfo [T any] (s []T) {
	fmt.Printf("slice data: %v\n", s)
	fmt.Printf("len: %v, cap: %v, &s[0]: %v,k &s[len-1]: %v\n", len(s), cap(s), &s[0], &s[len(s)-1])
}
