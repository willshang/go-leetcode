package main

import "fmt"

func main() {
	fmt.Println(verifyPostorder([]int{1, 3, 2, 6, 5}))
}

func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 2 {
		return true
	}
	last := len(postorder) - 1
	for last > 0 {
		i := 0
		for postorder[i] < postorder[last] {
			i++
		}
		for postorder[i] > postorder[last] {
			i++
		}
		if i != last {
			return false
		}
		last--
	}
	return true
}
