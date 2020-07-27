package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

type Node struct {
	str   string
	left  int
	right int
}

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	if n == 0 {
		return res
	}
	queue := make([]*Node, 0)
	queue = append(queue, &Node{
		str:   "",
		left:  n,
		right: n,
	})
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.left == 0 && node.right == 0 {
			res = append(res, node.str)
		}
		if node.left > 0 {
			queue = append(queue, &Node{
				str:   node.str + "(",
				left:  node.left - 1,
				right: node.right,
			})
		}
		if node.right > 0 && node.left < node.right {
			queue = append(queue, &Node{
				str:   node.str + ")",
				left:  node.left,
				right: node.right - 1,
			})
		}
	}
	return res
}
