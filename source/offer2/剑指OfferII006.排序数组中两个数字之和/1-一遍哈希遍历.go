package main

import "fmt"

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(numbers, target))
}

func twoSum(numbers []int, target int) []int {
	m := make(map[int]int, len(numbers))
	for i, n := range numbers {
		if m[target-n] != 0 {
			return []int{m[target-n] - 1, i}
		}
		m[n] = i + 1
	}
	return nil
}
