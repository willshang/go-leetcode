package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	stack := make([]*TreeNode, 0)
	temp := root
	sum := 0
	for {
		if temp != nil {
			stack = append(stack, temp)
			temp = temp.Right
		} else if len(stack) != 0 {
			temp = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			temp.Val = temp.Val + sum
			sum = temp.Val
			temp = temp.Left
		} else {
			break
		}
	}
	return root
}
