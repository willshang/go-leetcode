package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode114_二叉树展开为链表
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	res := make([]*TreeNode, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		res = append(res, node)
		stack = stack[:len(stack)-1]
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	for i := 1; i < len(res); i++ {
		res[i-1].Left = nil
		res[i-1].Right = res[i]
	}
	res[len(res)-1].Left = nil
}
