package main

type S struct {
	m string
}

func f() *S {
	return __ //A
}

func main() {
	p := __    //B
	print(p.m) //print "foo"
}
