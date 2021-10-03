package main

func main() {

}

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	res := &Node{}
	cur := res
	for root != nil {
		cur.Next = root
		root.Prev = cur
		cur = cur.Next
		root = root.Next
		// 处理子节点
		if cur.Child != nil {
			ch := flatten(cur.Child)
			cur.Child = nil
			cur.Next = ch
			ch.Prev = cur
			// 指针移动
			for cur.Next != nil {
				cur = cur.Next
			}
		}
	}
	res.Next.Prev = nil
	return res.Next
}
