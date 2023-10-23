package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int

func pseudoPalindromicPaths(root *TreeNode) int {
	res = 0
	dfs(root, [10]int{})
	return res
}

func dfs(root *TreeNode, arr [10]int) {
	arr[root.Val]++
	if root.Left == nil && root.Right == nil {
		count := 0
		for i := 1; i <= 9; i++ {
			if arr[i]%2 == 1 {
				count++
			}
		}
		if count <= 1 {
			res++
		}
		return
	}
	if root.Left != nil {
		dfs(root.Left, arr)
	}
	if root.Right != nil {
		dfs(root.Right, arr)
	}
}
