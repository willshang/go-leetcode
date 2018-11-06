package main

import (
	"fmt"
	"math"
)

func main() {
	first := TreeNode{Val:4}
	second := TreeNode{Val:6}
	third := TreeNode{Val:2}

	first.Left = &third
	first.Right = &second

	fmt.Println(minDiffInBST(&first))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
/*func minDiffInBST(root *TreeNode) int {
	res := math.MaxInt64
	pre := 1 >> 63
	null := pre
	var helper func(*TreeNode)
	helper = func(root *TreeNode) {
		if root.Left != nil{
			helper(root.Left)
		}
		if pre != null{
			res = min(res,root.Val-pre)
		}
		pre = root.Val

		if root.Right != nil{
			helper(root.Right)
		}
	}
	helper(root)
	return res
}
func min(a, b int)int  {
	if a < b{
		return a
	}
	return b
}*/
func minDiffInBST(root *TreeNode) int {
	res := []int{}
	BST(root,&res)
	//sort.Ints(res)
	min := math.MaxInt32
	for i := 0; i < len(res)-1;i++{
		if abs(res[i],res[i+1]) < min{
			min = abs(res[i],res[i+1])
		}
	}
	return min
}
func abs(a,b int)  int{
	if a > b{
		return a - b
	}
	return b - a
}
func BST(root *TreeNode, res *[]int )  {
	if root == nil{
		return
	}
	BST(root.Left,res)
	*res = append(*res,root.Val)
	BST(root.Right,res)
}
