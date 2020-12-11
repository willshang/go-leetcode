package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int

func sumEvenGrandparent(root *TreeNode) int {
	res = 0
	if root == nil {
		return res
	}
	dfs(root, false, false)
	return res
}

func dfs(root *TreeNode, grandfather, father bool) {
	if root == nil {
		return
	}
	if grandfather == true {
		res = res + root.Val
	}
	flag := true
	if root.Val%2 == 1 {
		flag = false
	}
	dfs(root.Left, father, flag)
	dfs(root.Right, father, flag)
}
