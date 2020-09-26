package main

import (
	"fmt"
	"math"
)

func main() {
	first := TreeNode{Val: 2}
	second := TreeNode{Val: 2}
	third := TreeNode{Val: 5}
	first.Left = &third
	first.Right = &second
	fmt.Println(findSecondMinimumValue(&first))
	fmt.Println(math.MaxInt32)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var arr []int

func findSecondMinimumValue(root *TreeNode) int {
	arr = make([]int, 0)
	dfs(root)
	min, second := math.MaxInt32, math.MaxInt32
	flag := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			second = min
			min = arr[i]
		} else if min < arr[i] && arr[i] <= second {
			flag = 1
			second = arr[i]
		}
	}
	if second == math.MaxInt32 && flag == 0 {
		return -1
	}
	return second
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	arr = append(arr, root.Val)
	dfs(root.Left)
	dfs(root.Right)
}
