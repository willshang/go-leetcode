package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumEvenGrandparent(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}
	quque := make([]*TreeNode, 0)
	quque = append(quque, root)
	for len(quque) > 0 {
		length := len(quque)
		for i := 0; i < length; i++ {
			node := quque[i]
			if node.Val%2 == 0 {
				if node.Left != nil && node.Left.Left != nil {
					res = res + node.Left.Left.Val
				}
				if node.Left != nil && node.Left.Right != nil {
					res = res + node.Left.Right.Val
				}
				if node.Right != nil && node.Right.Left != nil {
					res = res + node.Right.Left.Val
				}
				if node.Right != nil && node.Right.Right != nil {
					res = res + node.Right.Right.Val
				}
			}
			if node.Left != nil {
				quque = append(quque, node.Left)
			}
			if node.Right != nil {
				quque = append(quque, node.Right)
			}
		}
		quque = quque[length:]
	}
	return res
}
