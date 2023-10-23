package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}
	temp := root
	for temp != nil {
		if temp.Val > val {
			if temp.Left == nil {
				temp.Left = &TreeNode{
					Val: val,
				}
				break
			}
			temp = temp.Left
		} else {
			if temp.Right == nil {
				temp.Right = &TreeNode{
					Val: val,
				}
				break
			}
			temp = temp.Right
		}
	}
	return root
}
