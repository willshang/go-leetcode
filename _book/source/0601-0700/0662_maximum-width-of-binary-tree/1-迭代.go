package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode662_二叉树最大宽度
func widthOfBinaryTree(root *TreeNode) int {
	res := 1
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	arr := make([]int, 0)
	arr = append(arr, 1)
	for len(queue) > 0 {
		if arr[len(arr)-1]-arr[0]+1 > res {
			res = arr[len(arr)-1] - arr[0] + 1
		}
		length := len(queue)
		for i := 0; i < length; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
				arr = append(arr, arr[i]*2)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
				arr = append(arr, arr[i]*2+1)
			}
		}
		queue = queue[length:]
		arr = arr[length:]
	}
	return res
}
