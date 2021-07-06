package main

import "sort"

func main() {

}

// leetcode_lcs02_完成一半题目
func halfQuestions(questions []int) int {
	res := 0
	n := len(questions)
	arr := make([]int, 1001)
	for i := 0; i < n; i++ {
		arr[questions[i]]++
	}
	sort.Ints(arr)
	count := n / 2
	for i := 1000; i >= 0; i-- {
		res++
		count = count - arr[i]
		if count <= 0 {
			break
		}
	}
	return res
}
