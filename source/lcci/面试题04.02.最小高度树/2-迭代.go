package main

import (
	"fmt"
)

func main() {
	//fmt.Println(sortedArrayToBST([]int{-10, -3, 0, 5, 9}))
	root := sortedArrayToBST([]int{-10, -3, 2, 5, 9})
	fmt.Println(root, root.Right, root.Left)
	//fmt.Println(root.Val,root.Left.Val,root.Right.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type MyTreeNode struct {
	root  *TreeNode
	start int
	end   int
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	queue := make([]MyTreeNode, 0)
	root := &TreeNode{Val: 0}
	queue = append(queue, MyTreeNode{root, 0, len(nums)})
	for len(queue) > 0 {
		myRoot := queue[0]
		queue = queue[1:]
		start := myRoot.start
		end := myRoot.end
		mid := (start + end) / 2
		curRoot := myRoot.root
		curRoot.Val = nums[mid]
		if start < mid {
			curRoot.Left = &TreeNode{Val: 0}
			queue = append(queue, MyTreeNode{curRoot.Left, start, mid})
		}
		if mid+1 < end {
			curRoot.Right = &TreeNode{Val: 0}
			queue = append(queue, MyTreeNode{curRoot.Right, mid + 1, end})
		}
	}
	return root
}
