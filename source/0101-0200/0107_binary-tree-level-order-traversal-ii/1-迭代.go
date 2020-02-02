package main

import "fmt"

func main() {
	root := TreeNode{}
	root.Val = 1

	left := TreeNode{}
	left.Val = 2

	right := TreeNode{}
	right.Val = 2

	root.Left = &left
	root.Right = &right
	res := levelOrderBottom(&root)

	for _, v := range res {
		fmt.Println(v)
	}

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	out := make([][]int, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		l := len(queue)
		arr := make([]int, 0)
		for i := 0; i < l; i++ {
			pop := queue[i]
			arr = append(arr, pop.Val)
			if pop.Left != nil {
				queue = append(queue, pop.Left)
			}
			if pop.Right != nil {
				queue = append(queue, pop.Right)
			}
		}
		out = append(out, arr)
		queue = queue[l:]
	}

	out2 := make([][]int, len(out))
	for i := 0; i < len(out); i++ {
		out2[len(out)-1-i] = out[i]
	}

	return out2
}
