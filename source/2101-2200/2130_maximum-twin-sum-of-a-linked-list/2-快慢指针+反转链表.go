package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	res := 0
	slow, fast := head, head.Next
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	slow = reverseList(slow.Next)
	fast = head
	for slow != nil {
		res = max(res, fast.Val+slow.Val)
		fast = fast.Next
		slow = slow.Next
	}
	return res
}

func reverseList(head *ListNode) *ListNode {
	var result *ListNode
	var temp *ListNode
	for head != nil {
		temp = head.Next
		head.Next = result
		result = head
		head = temp
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
