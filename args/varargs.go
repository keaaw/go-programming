package main

import (
	"fmt"
)

func f [T any] (args ...T) {
	fmt.Println("func f:")
	for idx, arg := range args {
		fmt.Println(idx, arg)
	}
}

func g (args ...any) {
	fmt.Println("func g:")
	for idx, arg := range args {
		fmt.Println(idx, arg)
	}
}

func h (sl []any) {
	sl[0] = "foop"
}

func changeArg(args ...any) {
	fmt.Println("in changeArg")
	fmt.Println(args)
	fmt.Printf("type of args: %T, type of args[0]: %T\n", args, args)
	fmt.Printf("&args[0] = %v\n", &args[0])
	args[0] = "howdy there"
	fmt.Println(args)
}

func main() {
	f(1)
	f(42, 99, -123)
	ns := []int{4,9,6,7,8,1,2,3}
	f(ns)
	//h(ns)
	sliceOfAny := []any{1, "hi", 4.5}
	h(sliceOfAny)
	fmt.Printf("slice of any: %v\n", sliceOfAny)
	changeArg(sliceOfAny)

	fmt.Printf("type of ns: %T, len=%v, cap=%v\n", ns, len(ns), cap(ns))
	slint := make([]int, 10, 100)
	fmt.Printf("type of slint: %T, len=%v, cap=%v\n", slint, len(slint), cap(slint))
	for i, v := range slint {
		fmt.Printf("slint[%v] = %v\n", i, v)
	}

	sa := []any{"hi there", 1, 5.6}
	fmt.Println(sa)
	f(sa)
	g(sa...)

	changeArg(sa)
	fmt.Println(sa)
}
