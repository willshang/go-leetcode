package main

import "math"

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

var count = 0

func getDecimalValue(head *ListNode) int {
	if head == nil {
		count = 0
		return 0
	}
	value := getDecimalValue(head.Next)
	res := head.Val*int(math.Pow(2, float64(count))) + value
	count++
	return res
}
