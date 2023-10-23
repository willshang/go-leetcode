package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode2181_合并零之间的节点
func mergeNodes(head *ListNode) *ListNode {
	res := &ListNode{}
	t := res
	sum := 0
	for cur := head.Next; cur != nil; cur = cur.Next {
		if cur.Val == 0 {
			node := &ListNode{
				Val: sum,
			}
			t.Next = node
			t = t.Next
			sum = 0
		} else {
			sum = sum + cur.Val
		}
	}
	return res.Next
}
