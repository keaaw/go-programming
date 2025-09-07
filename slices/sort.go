package main

import (
	"fmt"
	"math/rand"
	"time"
	"slices"
)

var (
	size     = 10000000
	nWorkers = 4
)

func divideWorkAndWait(s []uint32, doWork func([]uint32, chan<- int)) {
	workerSize := size / nWorkers
	c := make(chan int)
	for w := 0; w < nWorkers; w, s = w + 1, s[workerSize:] {
		// account for rounding by making sure the last worker covers the very end of the slice
		top := workerSize
		if w == nWorkers - 1 {
			top = len(s)
		}
		go doWork(s[:top], c)
	}
	// wait for all workers to complete
	for w := 0; w < nWorkers; w++ {
		<-c
	}
}

func fillRandom(s []uint32, c chan<- int) {
	for i := range s {
		s[i] = rand.Uint32()
	}
	c <- 1
}

func sort(s []uint32, c chan<- int) {
	slices.Sort(s)
	c <- 1
}

func main() {
	s := make([]uint32, size)
	start := time.Now()
	divideWorkAndWait(s, fillRandom)
	divideWorkAndWait(s, sort)
	end := time.Now()
	duration := end.Sub(start)
	fmt.Println(s[:10], "\n", duration)
}
