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
	cur := root
	stack := make([]*Node, 0)
	for cur != nil {
		// 处理child
		if cur.Child != nil {
			if cur.Next != nil {
				stack = append(stack, cur.Next)
			}
			cur.Child.Prev = cur
			cur.Next = cur.Child
			cur.Child = nil
			continue
		}
		if cur.Next != nil {
			cur.Child = nil
			cur = cur.Next
			continue
		}
		if len(stack) == 0 {
			break
		}
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur.Next = last
		last.Prev = cur
		cur = last
	}
	return root
}
