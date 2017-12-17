package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}
	ss := s[1:]
	for i := range ss {
		ss[i] += 10
	}
	fmt.Println(s)
	ss = append(ss, 4)
	for i := range ss {
		ss[i] += 10
	}
	fmt.Println(s)
}
