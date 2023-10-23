package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	sum := 0
	for cur := head; cur != nil; cur = cur.Next {
		sum = sum + cur.Val
		if sum == 0 {
			return removeZeroSumSublists(cur.Next)
		}
	}
	head.Next = removeZeroSumSublists(head.Next)
	return head
}
