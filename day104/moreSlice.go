package main

func main() {
	var x *struct {
		s [][32]byte
	}

	println(len(x.s[99]))
	println(len(x.s[1]))
}
