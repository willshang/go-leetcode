package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

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
		if root == nil {
			res = append(res, nil)
			continue
		}
		node := root
		for j := 1; j < a && root.Next != nil; j++ {
			root = root.Next
		}
		if b > 0 {
			root = root.Next
			b--
		}
		temp := root.Next
		root.Next = nil
		root = temp
		res = append(res, node)
	}
	return res
}
