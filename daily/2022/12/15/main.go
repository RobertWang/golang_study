package main

func main() {
	var x *struct {
		s [][32]byte
	}

	println(len(x.s[99]))
}
