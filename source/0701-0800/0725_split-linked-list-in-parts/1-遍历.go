package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode725_分隔链表
func splitListToParts(root *ListNode, k int) []*ListNode {
	res := make([]*ListNode, 0)
	cur := root
	length := 0
	for cur != nil {
		length++
		cur = cur.Next
	}
	a, b := length/k, length%k
	for i := 0; i < k; i++ {
		node := &ListNode{Next: nil}
		temp := node
		for j := 0; j < a; j++ {
			temp.Next = &ListNode{
				Val:  root.Val,
				Next: nil,
			}
			temp = temp.Next
			root = root.Next
		}
		if b > 0 {
			temp.Next = &ListNode{
				Val:  root.Val,
				Next: nil,
			}
			temp = temp.Next
			root = root.Next
			b = b - 1
		}
		res = append(res, node.Next)
	}
	return res
}
