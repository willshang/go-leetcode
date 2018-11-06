package main

import "fmt"

func main() {
	first := TreeNode{}
	first.Val = 19
	second := TreeNode{}
	second.Val = 5
	third := TreeNode{}
	third.Val = 19
	fourth := TreeNode{}
	fourth.Val = 5

	first.Right = &second
	second.Left = &third
	second.Right = &fourth

	fmt.Println(findMode(&first))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	r := map[int]int{}
	search(root,r)

	max := -1
	res := []int{}
	for n, v := range r {
		if max <= v {
			if max < v{
				max = v
				res = res[0:0]
			}
			res = append(res,n)
		}
	}
	return res
}

func search(root *TreeNode,rec map[int]int)  {
	if root == nil{
		return
	}

	rec[root.Val]++
	search(root.Left,rec)
	search(root.Right,rec)
}
