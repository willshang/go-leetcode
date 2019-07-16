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

/*func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, level int) {
		if root == nil{
			return
		}

		if level >= len(res){
			res = append([][]int{[]int{}},res...)
		}

		n := len(res)
		res[n-level-1] = append(res[n-level-1],root.Val)
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
	dfs(root,0)
	return res
}
*/

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{}
	out := [][]int{}
	queue = append(queue, root)

	for len(queue) != 0 {
		l := len(queue)
		arr := []int{}
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
