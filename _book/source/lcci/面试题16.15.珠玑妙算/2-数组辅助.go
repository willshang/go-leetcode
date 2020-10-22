package main

import "fmt"

func main() {
	fmt.Println(masterMind("RGBY", "GGRR"))
}

func masterMind(solution string, guess string) []int {
	arr := [256]int{}
	a, b := 0, 0
	for i := 0; i < len(solution); i++ {
		if solution[i] == guess[i] {
			a++
		} else {
			arr[solution[i]]++
		}
	}
	for i := 0; i < len(guess); i++ {
		if solution[i] != guess[i] {
			if arr[guess[i]] > 0 {
				b++
				arr[guess[i]]--
			}
		}
	}
	return []int{a, b}
}
