package main

import "sort"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var m map[int][][2]int

func verticalTraversal(root *TreeNode) [][]int {
	m = make(map[int][][2]int)
	res := make([][]int, 0)
	bfs(root, 0, 0)
	arr := make([]int, 0)
	for k := range m {
		arr = append(arr, k)
	}
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		temp := m[arr[i]]
		sort.Slice(temp, func(i, j int) bool {
			if temp[i][1] == temp[j][1] {
				return temp[i][0] < temp[j][0]
			}
			return temp[i][1] < temp[j][1]
		})
		tempArr := make([]int, 0)
		for j := 0; j < len(temp); j++ {
			tempArr = append(tempArr, temp[j][0])
		}
		res = append(res, tempArr)
	}
	return res
}

func bfs(root *TreeNode, x, y int) {
	if root == nil {
		return
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	queueArr := make([][2]int, 0)
	queueArr = append(queueArr, [2]int{0, 0})
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[i]
			x, y := queueArr[i][0], queueArr[i][1]
			m[x] = append(m[x], [2]int{node.Val, y})
			if node.Left != nil {
				queue = append(queue, node.Left)
				queueArr = append(queueArr, [2]int{x - 1, y + 1})
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				queueArr = append(queueArr, [2]int{x + 1, y + 1})
			}
		}
		queue = queue[length:]
		queueArr = queueArr[length:]
	}
}
