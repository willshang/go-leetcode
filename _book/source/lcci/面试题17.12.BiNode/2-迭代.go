package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBiNode(root *TreeNode) *TreeNode {
	head := &TreeNode{}
	cur := head
	stack := make([]*TreeNode, 0)
	node := root
	for node != nil || len(stack) > 0 {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		} else {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			node.Left = nil
			cur.Right = node
			cur = node
			node = node.Right
		}
	}
	return head.Right
}
