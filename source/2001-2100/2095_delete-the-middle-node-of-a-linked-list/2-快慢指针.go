package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode2095_删除链表的中间节点
func deleteMiddle(head *ListNode) *ListNode {
	prev := &ListNode{Next: head}
	slow := prev
	fast := prev
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return prev.Next
}
