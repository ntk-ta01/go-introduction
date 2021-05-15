package main

func is_even(num int) bool {
	return num%2 == 0
}

func swap(x, y int) (int, int) {
	return y, x
}

func swap2(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	for i := 1; i <= 100; i++ {
		print(i)
		if is_even(i) {
			println("-偶数")
		} else {
			println("-奇数")
		}
	}
	n, m := swap(10, 20)
	println(n, m)
	swap2(&n, &m)
	println(n, m)
}
