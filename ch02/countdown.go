package main

func main() {
	var i = 5
LOOP:
	for {
		println(i)
		if i == 0 {
			break LOOP
		}
		i--
	}
	println("end!")
}
