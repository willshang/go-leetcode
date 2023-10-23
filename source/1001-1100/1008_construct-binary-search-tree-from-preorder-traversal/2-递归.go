package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode1008_前序遍历构造二叉搜索树
func bstFromPreorder(preorder []int) *TreeNode {
	length := len(preorder)
	if length == 0 {
		return nil
	}
	index := length
	for i := 1; i < length; i++ {
		if preorder[i] > preorder[0] {
			index = i
			break
		}
	}
	return &TreeNode{
		Val:   preorder[0],
		Left:  bstFromPreorder(preorder[1:index]),
		Right: bstFromPreorder(preorder[index:]),
	}
}
