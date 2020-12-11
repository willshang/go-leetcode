package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	var root *TreeNode
	for i := 0; i < len(preorder); i++ {
		root = insert(root, &TreeNode{
			Val:   preorder[i],
			Left:  nil,
			Right: nil,
		})
	}
	return root
}

func insert(root *TreeNode, node *TreeNode) *TreeNode {
	if root == nil {
		return node
	}
	if root.Val < node.Val {
		root.Right = insert(root.Right, node)
	} else if root.Val > node.Val {
		root.Left = insert(root.Left, node)
	} else {
		root.Val = node.Val
	}
	return root
}
