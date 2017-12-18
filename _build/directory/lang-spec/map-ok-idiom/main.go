package main

func main() {
	var m map[string]int //A
	m["a"] = 1
	if v := m["b"]; v != nil { //B
		println(v)
	}
}
