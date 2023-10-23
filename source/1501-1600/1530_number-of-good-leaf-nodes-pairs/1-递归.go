package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countPairs(root *TreeNode, distance int) int {
	_, res := dfs(root, distance)
	return res
}

func dfs(root *TreeNode, distance int) (arr []int, count int) {
	arr = make([]int, distance+2)
	if root == nil {
		return arr, 0
	}
	if root.Left == nil && root.Right == nil {
		arr[1] = 1
		return arr, 0
	}
	leftArr, rightArr := make([]int, distance+1), make([]int, distance+1)
	leftCount, rightCount := 0, 0
	if root.Left != nil {
		leftArr, leftCount = dfs(root.Left, distance)
	}
	if root.Right != nil {
		rightArr, rightCount = dfs(root.Right, distance)
	}
	for i := 1; i <= distance; i++ {
		arr[i+1] = leftArr[i] + rightArr[i]
	}
	count = 0
	for i := 0; i <= distance; i++ {
		for j := 0; j <= distance-i; j++ {
			count = count + leftArr[i]*rightArr[j]
		}
	}
	return arr, count + leftCount + rightCount
}
