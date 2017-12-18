package main

import (
	"fmt"
)

type S struct {
	v int
}

func main() {
	s := []S{{1}, {3}, {5}, {2}}
	// A
	fmt.Printf("%#v", s)
}
