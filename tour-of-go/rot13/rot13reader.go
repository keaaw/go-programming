package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

var count = 0

func rot13(b byte) (r13b byte) {
	var base byte
	r13b = b
	if b >= 'A' && b <= 'Z' {
		base = 'A'
	} else if b >= 'a' && b <= 'z' {
		base = 'a'
	}
	if base != 0 {
		r13b = base + ((b - base + 13) % 26)
	}
	return r13b
}

func (r13 rot13Reader) Read(dst []byte) (n int, err error) {
	count++
	if len(dst) == 0 {
		return 0, nil
	}
	n, err = r13.r.Read(dst)
	if n == 0 {
		return 0, err
	}
	for i := 0; i < n; i++ {
		dst[i] = rot13(dst[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}

	//dst := make([]byte, 3)
	//n, err := r.r.Read(dst)
	//fmt.Println(n, err)
	//fmt.Println(string(dst))
	//return
	//fmt.Println("begin")
	io.Copy(os.Stdout, &r)
	//fmt.Println("end")
	fmt.Println("\ncount =", count)
}
