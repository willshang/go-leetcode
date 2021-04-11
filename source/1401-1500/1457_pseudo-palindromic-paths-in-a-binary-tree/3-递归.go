package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode1457_二叉树中的伪回文路径
var res int

func pseudoPalindromicPaths(root *TreeNode) int {
	res = 0
	dfs(root, 0)
	return res
}

func dfs(root *TreeNode, value int) {
	if root == nil {
		return
	}
	temp := value ^ (1 << root.Val)
	if root.Left == nil && root.Right == nil {
		if temp == 0 || (temp&(temp-1)) == 0 {
			res++
		}
		return
	}
	dfs(root.Left, temp)
	dfs(root.Right, temp)
}
