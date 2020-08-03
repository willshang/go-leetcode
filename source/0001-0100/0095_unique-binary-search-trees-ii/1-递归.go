package main

import "fmt"

func main() {
	fmt.Println(generateTrees(3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode95_不同的二叉搜索树II
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return dfs(1, n)
}

func dfs(left, right int) []*TreeNode {
	if left > right {
		return []*TreeNode{nil}
	}
	if left == right {
		return []*TreeNode{
			&TreeNode{Val: left},
		}
	}
	arr := make([]*TreeNode, 0)
	for i := left; i <= right; i++ {
		leftTree := dfs(left, i-1)
		rightTree := dfs(i+1, right)
		for j := 0; j < len(leftTree); j++ {
			for k := 0; k < len(rightTree); k++ {
				node := &TreeNode{Val: i}
				node.Left = leftTree[j]
				node.Right = rightTree[k]
				arr = append(arr, node)
			}
		}
	}
	return arr
}
