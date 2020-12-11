package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode508_出现次数最多的子树元素和
var m map[int]int

func findFrequentTreeSum(root *TreeNode) []int {
	m = make(map[int]int)
	res := make([]int, 0)
	dfs(root)
	maxValue := 0
	for k, v := range m {
		if v > maxValue {
			maxValue = v
			res = []int{k}
		} else if maxValue == v {
			res = append(res, k)
		}
	}
	return res
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := root.Val + dfs(root.Left) + dfs(root.Right)
	m[sum]++
	return sum
}
