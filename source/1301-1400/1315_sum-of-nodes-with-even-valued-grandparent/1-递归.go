package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode1315_祖父节点值为偶数的节点和
func sumEvenGrandparent(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}
	if root.Val%2 == 0 {
		if root.Left != nil && root.Left.Left != nil {
			res = res + root.Left.Left.Val
		}
		if root.Left != nil && root.Left.Right != nil {
			res = res + root.Left.Right.Val
		}
		if root.Right != nil && root.Right.Left != nil {
			res = res + root.Right.Left.Val
		}
		if root.Right != nil && root.Right.Right != nil {
			res = res + root.Right.Right.Val
		}
	}
	res = res + sumEvenGrandparent(root.Left)
	res = res + sumEvenGrandparent(root.Right)
	return res
}
