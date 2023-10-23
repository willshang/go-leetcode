package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p.Val >= root.Val {
		return inorderSuccessor(root.Right, p)
	}
	res := inorderSuccessor(root.Left, p)
	if res == nil {
		return root
	}
	return res
}
