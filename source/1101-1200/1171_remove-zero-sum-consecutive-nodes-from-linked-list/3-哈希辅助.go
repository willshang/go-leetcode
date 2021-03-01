package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode1171_从链表中删去总和值为零的连续节点
func removeZeroSumSublists(head *ListNode) *ListNode {
	res := &ListNode{Next: head}
	m := make(map[int]*ListNode)
	sum := 0
	for cur := res; cur != nil; cur = cur.Next {
		sum = sum + cur.Val
		m[sum] = cur
	}
	sum = 0
	for cur := res; cur != nil; cur = cur.Next {
		sum = sum + cur.Val
		cur.Next = m[sum].Next
	}
	return res.Next
}
