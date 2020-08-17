package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	mid := find(head)
	if mid == head {
		return &TreeNode{Val: mid.Val}
	}
	return &TreeNode{
		Val:   mid.Val,
		Left:  sortedListToBST(head),
		Right: sortedListToBST(mid.Next),
	}
}

func find(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	if prev != nil {
		prev.Next = nil
	}
	return slow
}
