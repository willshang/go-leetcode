package main

import "fmt"

func main() {
	first := TreeNode{Val: 5}
	second := TreeNode{Val: 2}
	//third := TreeNode{Val: -3}
	secondRight := TreeNode{Val: -3}

	first.Left = &second
	first.Right = &secondRight
	//second.Left = &third
	fmt.Println(averageOfLevels(&first))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	var sum, node []int
	res := make([]float64, 0)
	sum = append(sum, root.Val)
	node = append(node, 1)
	sum, node = dfs(root, sum, node, 1)
	for i := 0; i < len(sum); i++ {
		res = append(res, float64(sum[i])/float64(node[i]))
	}
	return res
}

func dfs(root *TreeNode, sum, node []int, level int) ([]int, []int) {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return sum, node
	}
	if level >= len(sum) {
		sum = append(sum, 0)
		node = append(node, 0)
	}
	if root.Left != nil {
		sum[level] += root.Left.Val
		node[level]++
	}
	if root.Right != nil {
		sum[level] += root.Right.Val
		node[level]++
	}
	sum, node = dfs(root.Left, sum, node, level+1)
	sum, node = dfs(root.Right, sum, node, level+1)
	return sum, node
}
