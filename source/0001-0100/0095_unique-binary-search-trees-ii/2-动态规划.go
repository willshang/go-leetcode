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

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	dp := make([][]*TreeNode, n+1)
	dp[1] = append(dp[1], &TreeNode{Val: 1})
	for i := 2; i <= n; i++ {
		for _, node := range dp[i-1] {
			root := &TreeNode{Val: i}
			root.Left = node
			dp[i] = append(dp[i], copyTree(root))
			root = node
			temp := root
			newNode := &TreeNode{Val: i}
			for temp != nil {
				newNode.Left = temp.Right
				temp.Right = newNode
				dp[i] = append(dp[i], copyTree(root))
				temp.Right = newNode.Left
				newNode.Left = nil
				temp = temp.Right
			}
		}
	}
	return dp[n]
}

func copyTree(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	newNode := &TreeNode{Val: node.Val}
	newNode.Left = copyTree(node.Left)
	newNode.Right = copyTree(node.Right)
	return newNode
}
