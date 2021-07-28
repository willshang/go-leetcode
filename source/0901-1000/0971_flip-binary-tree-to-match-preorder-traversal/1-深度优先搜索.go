package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int
var arr []int

func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	res = make([]int, 0)
	arr = make([]int, len(voyage))
	copy(arr, voyage)
	dfs(root)
	return res
}

func dfs(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Val != arr[0] { // 根不满足直接返回
		res = []int{-1}
		return false
	}
	if root.Left != nil && root.Right != nil && root.Right.Val == arr[1] { // 交换
		res = append(res, root.Val)
		root.Left, root.Right = root.Right, root.Left
	}
	arr = arr[1:]
	if dfs(root.Left) == false {
		return false
	}
	return dfs(root.Right)
}
