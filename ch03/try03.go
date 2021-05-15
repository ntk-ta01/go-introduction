package main

import "fmt"

type Score struct {
	// 解答例のstruct
	UserID string
	GameID int
	Point  int
}

func main() {
	type User string
	type GameResults map[User][]int

	gameresults := GameResults{}
	var user1 User = "Bob"
	var user2 User = "Alice"
	gameresults[user1] = []int{90, 60, 88}
	gameresults[user2] = []int{49, 58, 36}
	fmt.Println(gameresults[user1][0])

	// 確かに解答例のstructのが使いやすい…
	score := Score{UserID: "Bob", GameID: 0, Point: 90}
	fmt.Println(score)
	scores := []Score{score}
	fmt.Println(scores)
}
