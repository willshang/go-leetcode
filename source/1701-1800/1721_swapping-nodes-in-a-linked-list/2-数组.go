package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode1721_交换链表中的节点
func swapNodes(head *ListNode, k int) *ListNode {
	res := make([]*ListNode, 0)
	cur := head
	for cur != nil {
		res = append(res, cur)
		cur = cur.Next
	}
	res[k-1].Val, res[len(res)-k].Val = res[len(res)-k].Val, res[k-1].Val
	return head
}
