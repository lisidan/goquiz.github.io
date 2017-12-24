package main

import "flag"
import "fmt"

var ip string
var port int

func init() {
	// A
	// B
}

func main() {
	flag.Parse()
	fmt.Printf("%s:%d", ip, port)
}
