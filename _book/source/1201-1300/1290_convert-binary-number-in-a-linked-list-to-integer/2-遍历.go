package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode1290_二进制链表转整数
func getDecimalValue(head *ListNode) int {
	res := 0
	for head != nil {
		res = res*2 + head.Val
		head = head.Next
	}
	return res
}
