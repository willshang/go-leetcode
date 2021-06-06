package main

import "sort"

func main() {

}

// leetcode732_我的日程安排表III
type MyCalendarThree struct {
	m map[int]int
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{m: make(map[int]int)}
}

func (this *MyCalendarThree) Book(start int, end int) int {
	this.m[start]++
	this.m[end]--
	arr := make([]int, 0)
	for k := range this.m {
		arr = append(arr, k)
	}
	sort.Ints(arr)
	res := 0
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum = sum + this.m[arr[i]]
		if sum > res {
			res = sum
		}
	}
	return res
}
