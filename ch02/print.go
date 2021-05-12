package main

import "fmt"

func main() {
	const N = 1000000000000000000000 / 1000000000000000000000
	var m = N
	s := "Hello, World!"
	const (
		a = iota
		b
		c
	)
	const (
		d = 1 << iota
		e
	)
	fmt.Println(s, m, 1+4i, a, b, c, d, e)
}
