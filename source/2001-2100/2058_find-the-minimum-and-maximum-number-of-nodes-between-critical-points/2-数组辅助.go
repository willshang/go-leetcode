package main

import "math"

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode2058_找出临界点之间的最小和最大距离
func nodesBetweenCriticalPoints(head *ListNode) []int {
	a, b, c := head, head.Next, head.Next.Next // 题目保证数量范围>=2
	arr := make([]int, 0)
	index := 1
	for c != nil {
		if (a.Val < b.Val && b.Val > c.Val) ||
			(a.Val > b.Val && b.Val < c.Val) {
			arr = append(arr, index)
		}
		a, b, c = b, c, c.Next
		index++
	}
	if len(arr) < 2 {
		return []int{-1, -1}
	}
	minDis := math.MaxInt32
	for i := 0; i < len(arr)-1; i++ {
		minDis = min(minDis, arr[i+1]-arr[i])
	}
	return []int{minDis, arr[len(arr)-1] - arr[0]}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
