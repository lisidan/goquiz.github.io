package main

import (
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("file")
	if err != nil {
		return
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	println(string(b))
}
