package main

import "fmt"

func main() {
	fmt.Println(hasGroupsSizeX([]int{1, 2, 3, 4, 4, 3, 2, 1}))
}

func hasGroupsSizeX(deck []int) bool {
	if len(deck) < 2 {
		return false
	}
	m := make(map[int]int)
	for i := 0; i < len(deck); i++ {
		m[deck[i]]++
	}
	for i := 2; i <= len(deck); i++ {
		flag := true
		if len(deck)%i == 0 {
			for _, v := range m {
				if v%i != 0 {
					flag = false
					break
				}
			}
			if flag == true {
				return true
			}
		}
	}
	return false
}
