package main

import "fmt"

func main() {
	fmt.Println(masterMind("RGBY", "GGRR"))
}

// 程序员面试金典16.15_珠玑妙算
func masterMind(solution string, guess string) []int {
	m := make(map[byte]int)
	a, b := 0, 0
	for i := 0; i < len(solution); i++ {
		if solution[i] == guess[i] {
			a++
		} else {
			m[solution[i]]++
		}
	}
	for i := 0; i < len(guess); i++ {
		if solution[i] != guess[i] {
			if m[guess[i]] > 0 {
				b++
				m[guess[i]]--
			}
		}
	}
	return []int{a, b}
}
