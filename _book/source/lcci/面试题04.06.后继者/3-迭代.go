package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var res *TreeNode
	cur := root
	for cur != nil {
		if p.Val >= cur.Val {
			cur = cur.Right
		} else {
			res = cur
			cur = cur.Left
		}
	}
	return res
}
