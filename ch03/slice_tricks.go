package main

import "fmt"

func main() {
	a := []int{10, 20, 30, 40}
	fmt.Println(a) // [10 20 30 40]
	a = append(a[:2], a[3:]...)
	fmt.Println(a) // [10 20 40]
	a = append(a[:1], a[2:]...)
	fmt.Println(a) // [10 40]
}
