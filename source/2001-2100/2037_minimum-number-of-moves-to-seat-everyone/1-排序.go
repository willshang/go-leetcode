package main

import "sort"

func main() {

}

// leetcode2037_使每位学生都有座位的最少移动次数
func minMovesToSeat(seats []int, students []int) int {
	res := 0
	sort.Ints(seats)
	sort.Ints(students)
	for i := 0; i < len(seats); i++ {
		res = res + abs(students[i]-seats[i])
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
