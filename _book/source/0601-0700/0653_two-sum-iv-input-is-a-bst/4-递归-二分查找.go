package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var arr []int

func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	arr = make([]int, 0)
	dfs(root)
	i := 0
	j := len(arr) - 1
	for i < j {
		if arr[i]+arr[j] == k {
			return true
		} else if arr[i]+arr[j] > k {
			j--
		} else {
			i++
		}
	}
	return false
}

func dfs(node *TreeNode) {
	if node == nil {
		return
	}
	dfs(node.Left)
	arr = append(arr, node.Val)
	dfs(node.Right)
}
