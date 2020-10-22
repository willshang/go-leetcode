package main

import "fmt"

func main() {
	fmt.Println(game([]int{1, 2, 3}, []int{1, 2, 3}))
}

func game(guess []int, answer []int) int {
	res := 0
	for i := 0; i < len(guess); i++ {
		if guess[i]^answer[i] == 0 {
			res++
		}
	}
	return res
}
