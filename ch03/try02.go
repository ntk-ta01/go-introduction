package main

func main() {
	var sum int
	for _, v := range []int{19, 86, 1, 2} {
		sum += v
	}
	println(sum)
}
