package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode889_根据前序和后序遍历构造二叉树
func constructFromPrePost(pre []int, post []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: pre[0],
	}
	if len(pre) == 1 {
		return root
	}
	index := len(pre)
	for i := 0; i < len(post); i++ {
		if post[i] == pre[1] {
			index = i
			break
		}
	}
	root.Left = constructFromPrePost(pre[1:index+2], post[:index+1])
	root.Right = constructFromPrePost(pre[index+2:], post[index+1:])
	return root
}
