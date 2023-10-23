package main

func main() {
	root := TreeNode{Val: 2}
	rootfirst := TreeNode{Val: 3}
	rootSecond := TreeNode{Val: 1}
	root.Left = &rootfirst
	root.Right = &rootSecond
	recoverTree(&root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var arr []*TreeNode

func recoverTree(root *TreeNode) {
	arr = make([]*TreeNode, 0)
	dfs(root)
	a, b := -1, -1
	for i := 0; i < len(arr)-1; i++ {
		if arr[i].Val > arr[i+1].Val {
			b = i + 1
			if a == -1 {
				a = i
			}
		}
	}
	arr[a].Val, arr[b].Val = arr[b].Val, arr[a].Val
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	arr = append(arr, root)
	dfs(root.Right)
}
