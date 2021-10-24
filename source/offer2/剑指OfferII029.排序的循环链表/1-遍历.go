package main

func main() {

}

type Node struct {
	Val  int
	Next *Node
}

// 剑指OfferII029.排序的循环链表
func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		res := &Node{Val: x}
		res.Next = res
		return res
	}
	cur := aNode
	for cur.Next != aNode {
		if (cur.Val <= x && x <= cur.Next.Val) || // a<x<b
			(cur.Val > cur.Next.Val && (cur.Val <= x || x <= cur.Next.Val)) { // 插入最大值或者最小值
			break
		}
		cur = cur.Next
	}
	cur.Next = &Node{
		Val:  x,
		Next: cur.Next,
	}
	return aNode
}
