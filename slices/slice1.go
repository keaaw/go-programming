package main

import "fmt"

func main() {
	a := [4]int{1,2,3,4}
	s := a[:]
	sliceinfo(s)

	s = make([]int, 4, 10)
	sliceinfo(s)

	s = s[:]
	sliceinfo(s)

	s = s[:6]
	sliceinfo(s)

	s2 := s[3:]
	sliceinfo(s2)

	s3 := s2[:1]
	sliceinfo(s3)

	s3 = s3[:cap(s3)]
	sliceinfo(s3)

	ss := []string{"now", "is", "the", "time"}
	sliceinfo(ss)
}

func sliceinfo [T any] (s []T) {
	fmt.Printf("slice data: %v\n", s)
	fmt.Printf("len: %v, cap: %v, &s[0]: %v,k &s[len-1]: %v\n", len(s), cap(s), &s[0], &s[len(s)-1])
}
