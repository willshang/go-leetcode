package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	arr  []int
	root *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	arr := make([]int, 0)
	inorder(root, &arr)
	return BSTIterator{
		arr:  arr,
		root: root,
	}
}

func inorder(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, nums)
	*nums = append(*nums, root.Val)
	inorder(root.Right, nums)
}

func (this *BSTIterator) Next() int {
	if len(this.arr) == 0 {
		return -1
	}
	res := this.arr[0]
	this.arr = this.arr[1:]
	return res
}

func (this *BSTIterator) HasNext() bool {
	if len(this.arr) > 0 {
		return true
	}
	return false
}
