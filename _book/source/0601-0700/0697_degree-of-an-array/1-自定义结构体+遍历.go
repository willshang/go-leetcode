package main

import "fmt"

func main() {
	arr := []int{1, 2, 2, 3, 1, 4, 2}
	fmt.Println(findShortestSubArray(arr))
}

type node struct {
	count int
	left  int
	right int
}

func findShortestSubArray(nums []int) int {
	m := make(map[int]*node, 0)
	for k, v := range nums {
		if nd, ok := m[v]; ok {
			nd.count = nd.count + 1
			nd.right = k
		} else {
			m[v] = &node{
				count: 1,
				left:  k,
				right: k,
			}
		}
	}
	maxNode := new(node)
	for _, v := range m {
		if v.count > maxNode.count {
			maxNode = v
		} else if v.count == maxNode.count &&
			v.right-v.left < maxNode.right-maxNode.left {
			maxNode = v
		}
	}
	return maxNode.right - maxNode.left + 1
}
