package main

func main() {
	var a, b, c bool
	if a && b || !c {
		// 5通りtrue
		println("true")
	} else {
		println("false")
	}
}
