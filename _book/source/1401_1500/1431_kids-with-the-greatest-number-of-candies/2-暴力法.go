package main

import "fmt"

func main() {
	fmt.Println(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		flag := true
		for j := 0; j < len(candies); j++ {
			if candies[i]+extraCandies < candies[j] {
				flag = false
				res[i] = false
				break
			}
		}
		if flag == true {
			res[i] = true
		}
	}
	return res
}
