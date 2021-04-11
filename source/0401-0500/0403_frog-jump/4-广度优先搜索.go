package main

func main() {

}

type Node struct {
	index int
	size  int
}

func canCross(stones []int) bool {
	n := len(stones)
	m := make(map[int]bool)
	for i := 0; i < n; i++ {
		m[stones[i]] = true
	}
	isVisited := make(map[Node]bool)

	stack := make([]Node, 0)
	node := Node{
		index: 0,
		size:  0,
	}
	stack = append(stack, node)
	isVisited[node] = true
	for len(stack) > 0 {
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.index == stones[n-1] {
			return true
		}
		for k := node.size - 1; k <= node.size+1; k++ {
			next := node.index + k
			if k > stones[n-1] {
				continue
			}
			temp := Node{
				index: next,
				size:  k,
			}
			if k > 0 && m[next] == true && isVisited[temp] == false {
				isVisited[temp] = true
				stack = append(stack, temp)
			}
		}
	}
	return false
}
