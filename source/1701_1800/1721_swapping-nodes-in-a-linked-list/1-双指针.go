package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	slow, fast := head, head
	for i := 0; i < k-1; i++ {
		fast = fast.Next
	}
	a := fast
	for fast != nil && fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Val, a.Val = a.Val, slow.Val
	return head
}
