package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res string

func smallestFromLeaf(root *TreeNode) string {
	res = ""
	dfs(root, make([]byte, 0))
	return res
}

func dfs(root *TreeNode, arr []byte) {
	if root == nil {
		return
	}
	arr = append([]byte{byte('a' + root.Val)}, arr...)
	if root.Left == nil && root.Right == nil {
		if string(arr) < res || res == "" {
			res = string(arr)
		}
		return
	}
	dfs(root.Left, arr)
	dfs(root.Right, arr)
}
