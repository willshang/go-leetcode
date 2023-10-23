package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 剑指OfferII053.二叉搜索树中的中序后继
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var res *TreeNode
	for root != nil {
		if root.Val > p.Val {
			res = root
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return res
}
