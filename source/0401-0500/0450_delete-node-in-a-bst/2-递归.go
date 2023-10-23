package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode450_删除二叉搜索树中的节点
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		// 找到左节点的最大值，把右节点给最大值
		maxNode := root.Left
		for maxNode.Right != nil {
			maxNode = maxNode.Right
		}
		maxNode.Right = root.Right
		root = root.Left
	}
	return root
}
