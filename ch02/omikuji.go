package main

import (
	"math/rand"
	"time"
)

func main() {
	t := time.Now().UnixNano()
	rand.Seed(t)
	s := rand.Intn(6)
	switch s {
	case 0:
		println("凶")
	case 1, 2:
		println("吉")
	case 3, 4:
		println("中吉")
	case 5:
		println("大吉")
	default:
		println("サイコロの目でない")
	}
}
