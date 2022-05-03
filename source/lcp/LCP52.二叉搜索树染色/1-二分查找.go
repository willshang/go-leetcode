package main

import "sort"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var arr []int

func getNumber(root *TreeNode, ops [][]int) int {
	res := 0 // 蓝色的数量
	n := len(arr)
	arr = make([]int, 0)
	dfs(root)
	for i := len(ops) - 1; i >= 0; i-- {
		if len(arr) == 0 {
			break
		}
		t, left, right := ops[i][0], ops[i][1], ops[i][2]
		index := sort.Search(len(arr), func(j int) bool {
			return arr[j] >= left
		})
		temp := make([]int, index)
		copy(temp, arr[:index])
		var k int
		for k = index; k < len(arr) && arr[k] <= right; k++ {
			if t == 0 { // 染蓝
				res++
			}
		}
		temp = append(temp, arr[k:]...)
		arr = temp
	}
	res = res + len(arr) // 染蓝的数量+开始蓝色的数量
	return n - res       // 红色数量
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	arr = append(arr, root.Val)
	dfs(root.Right)
}
