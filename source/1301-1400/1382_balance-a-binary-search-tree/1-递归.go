package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode1382_将二叉搜索树变平衡
var arr []int

func balanceBST(root *TreeNode) *TreeNode {
	arr = make([]int, 0)
	dfs(root) // 先转换成数组
	return build(0, len(arr)-1)
}

func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Left)
		arr = append(arr, root.Val)
		dfs(root.Right)
	}
}

// leetcode108.将有序数组转换为二叉搜索树
func build(left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := left + (right-left)/2
	return &TreeNode{
		Val:   arr[mid],
		Left:  build(left, mid-1),
		Right: build(mid+1, right),
	}
}
