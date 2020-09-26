package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(reconstructQueue([][]int{
		{7, 0},
		{4, 4},
		{7, 1},
		{5, 0},
		{6, 1},
		{5, 2},
	}))
}

// leetcode406_根据身高重建队列
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1] // k 递增
		}
		return people[i][0] > people[j][0] // 升高 递减
	})
	for i := 0; i < len(people); i++ {
		index := people[i][1]
		p := people[i]
		for j := i; j > index; j-- {
			people[j] = people[j-1]
		}
		people[index] = p
	}
	return people
}
