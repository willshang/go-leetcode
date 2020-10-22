package main

func main() {

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := make([]*Node, 0)
	if root.Left != nil {
		queue = append(queue, root.Left)
	}
	if root.Right != nil {
		queue = append(queue, root.Right)
	}
	for len(queue) > 0 {
		length := len(queue)
		i := 0
		for i = 0; i < length; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if i+1 < length {
				node.Next = queue[i+1]
			}
		}
		queue = queue[length:]
	}
	return root
}
