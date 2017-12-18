package main

import (
	"sync"
)

const N = 100

func main() {
	m := &map[int]int{} //A
	wg := &sync.WaitGroup{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(i int) {
			(*m)[i] = i //B
			wg.Done()
		}(i)
	}
	wg.Wait()
}
