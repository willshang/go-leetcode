package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 剑指OfferII043.往完全二叉树添加节点
type CBTInserter struct {
	root *TreeNode
	arr  []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	arr := make([]*TreeNode, 0)
	queue := make([]*TreeNode, 0)
	arr = append(arr, root)
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
				arr = append(arr, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
				arr = append(arr, queue[i].Right)
			}
		}
		queue = queue[length:]
	}
	return CBTInserter{root: root, arr: arr}
}

func (this *CBTInserter) Insert(v int) int {
	newNode := &TreeNode{Val: v}
	this.arr = append(this.arr, newNode)
	n := len(this.arr)
	target := this.arr[n/2-1]
	if n%2 == 0 {
		target.Left = newNode
	} else {
		target.Right = newNode
	}
	return target.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}
