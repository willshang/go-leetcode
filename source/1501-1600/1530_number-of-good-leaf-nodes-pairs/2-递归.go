package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode1530_好叶子节点对的数量
var res int

func countPairs(root *TreeNode, distance int) int {
	res = 0
	dfs(root, distance)
	return res
}

func dfs(root *TreeNode, distance int) (arr []int) {
	arr = make([]int, distance+2)
	if root == nil {
		return arr
	}
	if root.Left == nil && root.Right == nil {
		arr[1] = 1
		return arr
	}
	leftArr, rightArr := make([]int, distance+1), make([]int, distance+1)
	if root.Left != nil {
		leftArr = dfs(root.Left, distance)
	}
	if root.Right != nil {
		rightArr = dfs(root.Right, distance)
	}
	for i := 1; i <= distance; i++ {
		arr[i+1] = leftArr[i] + rightArr[i]
	}
	for i := 0; i <= distance; i++ {
		for j := 0; j <= distance-i; j++ {
			res = res + leftArr[i]*rightArr[j]
		}
	}
	return arr
}
