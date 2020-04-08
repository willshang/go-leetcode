package main

import "fmt"

func main() {
	//fmt.Println(findPairs([]int{1,3,1,5,4},1))
	fmt.Println(findPairs([]int{1, 3, 1, 5, 4}, 2))
}
func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	record := make(map[int]int)
	for _, num := range nums {
		record[num]++
	}

	res := 0

	if k == 0 {
		for _, count := range record {
			if count > 1 {
				res++
			}
		}
		return res
	}

	for n := range record {
		//判断小于k是否有值
		if record[n-k] > 0 {
			res++
		}
	}
	return res
}
