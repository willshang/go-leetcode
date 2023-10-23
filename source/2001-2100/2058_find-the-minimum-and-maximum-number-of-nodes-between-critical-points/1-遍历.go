package main

import "math"

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	a, b, c := head, head.Next, head.Next.Next // 题目保证数量范围>=2
	maxDis, minDis := math.MinInt32, math.MaxInt32
	index := 1
	prev := 0  // 前一个位置
	first := 0 // 第一个位置
	for c != nil {
		if (a.Val < b.Val && b.Val > c.Val) ||
			(a.Val > b.Val && b.Val < c.Val) {
			if first == 0 {
				first = index // 第一次出现
			}
			maxDis = max(maxDis, index-first) // 当前位置减去第一次位置
			if 0 < prev {
				minDis = min(minDis, index-prev) // 当前位置减去前一次位置
			}
			prev = index
		}
		a, b, c = b, c, c.Next
		index++
	}
	if minDis == math.MaxInt32 {
		return []int{-1, -1}
	}
	return []int{minDis, maxDis}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
