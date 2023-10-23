package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode971_翻转二叉树以匹配先序遍历
func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	index := 0
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}
		if node.Val != voyage[index] {
			return []int{-1}
		}
		index++
		if node.Left != nil && node.Right != nil && node.Left.Val != voyage[index] {
			res = append(res, node.Val)
			stack = append(stack, node.Left) // 翻转
			stack = append(stack, node.Right)
		} else {
			stack = append(stack, node.Right)
			stack = append(stack, node.Left)
		}
	}
	return res
}
