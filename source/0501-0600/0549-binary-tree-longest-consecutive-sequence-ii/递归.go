package main

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestConsecutive(root *TreeNode) int {
	in,de := make(map[*TreeNode]int),make(map[*TreeNode]int)
	res:=0
	incF(root,&in)
	decF(root,&de)

	for p,v := range in{
		tmp := v+0-1
		if d,ok:=de[p];ok{
			tmp = v+d-1
		}

		if res < tmp{
			res = tmp
		}
	}
	return res
}

func incF(root *TreeNode, inc *map[*TreeNode]int)int{
	l,r:=0,0
	if root.Left != nil{
		l = incF(root.Left,inc)
		if root.Val+1 !=root.Left.Val{
			l = 0
		}
	}
	if root.Right != nil{
		r = incF(root.Right,inc)
		if root.Val+1 !=root.Right.Val{
			r = 0
		}
	}
	if r>l{
		l=r
	}
	ttt:=*inc
	ttt[root]=l+1
	return l+1
}

func decF(root *TreeNode, dec *map[*TreeNode]int)int{
	l,r:=0,0
	if root.Left != nil{
		l = decF(root.Left,dec)
		if root.Val-1 !=root.Left.Val{
			l=0
		}
	}
	if root.Right != nil{
		r = decF(root.Right,dec)
		if root.Val-1 !=root.Right.Val{
			r=0
		}
	}
	if r>l{
		l=r
	}
	ttt:=*dec
	ttt[root]=l+1
	return l+1
}