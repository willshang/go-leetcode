package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode654_最大二叉树
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	index := 0
	maxValue := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxValue {
			maxValue = nums[i]
			index = i
		}
	}
	return &TreeNode{
		Val:   maxValue,
		Left:  constructMaximumBinaryTree(nums[:index]),
		Right: constructMaximumBinaryTree(nums[index+1:]),
	}
}
