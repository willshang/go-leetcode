package main

import "fmt"

func main() {
	first := TreeNode{Val: 1}
	second := TreeNode{Val: 2}
	third := TreeNode{Val: 3}
	first.Left = &second
	first.Right = &third

	first1 := TreeNode{Val: 1}
	second1 := TreeNode{Val: 2}
	third1 := TreeNode{Val: 3}
	first1.Left = &second1
	first1.Right = &third1
	fmt.Println(leafSimilar(&first, &first1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode872_叶子相似的树
var a1, a2 []int

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	a1 = make([]int, 0)
	a2 = make([]int, 0)
	dfs(root1, &a1)
	dfs(root2, &a2)
	if len(a1) != len(a2) {
		return false
	}
	for i := 0; i < len(a1); i++ {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}

func dfs(root *TreeNode, arr *[]int) {
	if root != nil {
		if root.Left == nil && root.Right == nil {
			*arr = append(*arr, root.Val)
			return
		}
		dfs(root.Left, arr)
		dfs(root.Right, arr)
	}
}
