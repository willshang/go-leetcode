package main

import "fmt"

func main() {
	arr := []int{1, 2, 2, 3, 1, 4, 2}
	fmt.Println(findShortestSubArray(arr))
}

type node struct {
	count int
	head  int
	tail  int
}

func findShortestSubArray(nums []int) int {
	m := make(map[int]*node, 0)
	for k, v := range nums {
		if nd, ok := m[v]; ok {
			nd.count = nd.count + 1
			nd.tail = k
		} else {
			m[v] = &node{
				count: 1,
				head:  k,
				tail:  k,
			}
		}
	}

	maxNode := new(node)
	for _, v := range m {
		if v.count > maxNode.count {
			maxNode = v
		} else if v.count == maxNode.count && v.tail-v.head < maxNode.tail-maxNode.head {
			maxNode = v
		}
	}

	return maxNode.tail - maxNode.head + 1
}

/*func findShortestSubArray(nums []int) int {
	size := len(nums)
	if size < 2{
		return size
	}

	first := make(map[int]int,size)
	count := make(map[int]int,size)
	maxCount := 1
	minLen := size
	for i, n := range nums{
		count[n]++
		if count[n] == 1{
			first[n] = i
		}else {
			l := i - first[n]+1
			if maxCount < count[n] ||
				(maxCount == count[n] && minLen > l){
					maxCount = count[n]
					minLen = l
			}
		}
	}
	if len(count) == size{
		return 1
	}
	return minLen
}*/
