package main

import "fmt"

func main() {
	first := ListNode{Val: 1}
	second := ListNode{Val: 2}
	third := ListNode{Val: 3}
	first.Next = &second
	second.Next = &third
	fmt.Println(getKthFromEnd(&first, 2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	res, count := dfs(head, k)
	if count > 0 {
		return nil
	}
	return res
}

func dfs(node *ListNode, k int) (*ListNode, int) {
	if node == nil {
		return node, k
	}
	next, nextValue := dfs(node.Next, k)
	if nextValue <= 0 {
		return next, nextValue
	}
	nextValue = nextValue - 1
	return node, nextValue
}
