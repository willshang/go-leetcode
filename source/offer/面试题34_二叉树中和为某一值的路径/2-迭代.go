package main

import "fmt"

func main() {
	first := TreeNode{Val: 5}
	second := TreeNode{Val: 4}
	third := TreeNode{Val: 8}
	first.Left = &second
	first.Right = &third

	secondLeft := TreeNode{Val: 11}
	second.Left = &secondLeft
	value7 := TreeNode{Val: 7}
	value2 := TreeNode{Val: 2}
	secondLeft.Left = &value7
	secondLeft.Right = &value2

	value13 := TreeNode{Val: 13}
	value4 := TreeNode{Val: 4}
	value5 := TreeNode{Val: 5}
	value1 := TreeNode{Val: 1}

	third.Left = &value13
	third.Right = &value4
	value4.Left = &value5
	value4.Right = &value1
	fmt.Println(pathSum(&first, 22))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	temp := make([]int, 0)
	stack := make([]*TreeNode, 0)
	visited := make(map[*TreeNode]bool)
	curSum := 0
	for root != nil || len(stack) > 0 {
		for root != nil {
			temp = append(temp, root.Val)
			curSum = curSum + root.Val
			visited[root] = true
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		if node.Right == nil || visited[node.Right] {
			if node.Left == nil && node.Right == nil && curSum == sum {
				tmp := make([]int, len(temp))
				copy(tmp, temp)
				res = append(res, tmp)
			}
			stack = stack[:len(stack)-1]
			temp = temp[:len(temp)-1]
			curSum = curSum - node.Val
			root = nil
		} else {
			root = node.Right
		}
	}
	return res
}
