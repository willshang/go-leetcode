package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 剑指OfferII055.二叉搜索树迭代器
type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	res := BSTIterator{}
	res.left(root)
	return res
}

func (this *BSTIterator) left(root *TreeNode) {
	for root != nil {
		this.stack = append(this.stack, root)
		root = root.Left
	}
}

func (this *BSTIterator) Next() int {
	node := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if node.Right != nil {
		this.left(node.Right)
	}
	return node.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}
