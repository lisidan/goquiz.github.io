package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var b = []byte("123")
	s := __ //A

	b[1] = '4'
	fmt.Printf("%+v\n", s) //print 143
}
