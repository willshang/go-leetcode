package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode988_从叶结点开始的最小字符串
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
	arr = append(arr, byte('a'+root.Val))
	if root.Left == nil && root.Right == nil {
		for i := 0; i < len(arr)/2; i++ {
			arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
		}
		if string(arr) < res || res == "" {
			res = string(arr)
		}
		for i := 0; i < len(arr)/2; i++ {
			arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
		}
		return
	}
	dfs(root.Left, arr)
	dfs(root.Right, arr)
}
