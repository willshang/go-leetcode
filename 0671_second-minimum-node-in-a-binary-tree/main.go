package main

import (
	"fmt"
	"math"
)

func main() {
	first := TreeNode{Val:2}
	second := TreeNode{Val:2}
	third := TreeNode{Val:2}

	first.Left = &third
	first.Right = &second

	fmt.Println(findSecondMinimumValue(&first))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	var arr  []int
	helper(root,&arr)
	min, second := math.MaxInt32,math.MaxInt32
	for i := 0; i < len(arr); i++{
		if arr[i] < min{
			second = min
			min = arr[i]
		}else if min < arr[i] && arr[i] < second{
			second = arr[i]
		}
	}
	if second == math.MaxInt32{
		return -1
	}
	return second
}

func helper(root *TreeNode,arr *[]int)  {
	if root == nil{
		return
	}
	*arr = append(*arr,root.Val)
	helper(root.Left,arr)
	helper(root.Right,arr)
}
/*const intMax  = 1 << 63 -1
func findSecondMinimumValue(root *TreeNode) int {
	res := intMax
	helper(root,root.Val, &res)
	if res == intMax{
		return -1
	}
	return res
}

func helper(root *TreeNode,lo int, hi *int)  {
	if root == nil{
		return
	}
	if lo < root.Val && root.Val < *hi{
		*hi = root.Val
	}

	helper(root.Left,lo,hi)
	helper(root.Right,lo,hi)
}
*/