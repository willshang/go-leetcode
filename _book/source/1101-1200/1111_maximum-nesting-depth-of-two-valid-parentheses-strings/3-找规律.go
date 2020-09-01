package main

import "fmt"

func main() {
	fmt.Println(maxDepthAfterSplit("(()())"))
}

// leetcode1111_有效括号的嵌套深度
func maxDepthAfterSplit(seq string) []int {
	res := make([]int, 0)
	a, b := 0, 0
	for i := 0; i < len(seq); i++ {
		if seq[i] == '(' {
			// 谁少给谁
			if a <= b {
				a++
				res = append(res, 0)
			} else {
				b++
				res = append(res, 1)
			}
		} else {
			// 谁多减谁
			if a > b {
				a--
				res = append(res, 0)
			} else {
				b--
				res = append(res, 1)
			}
		}
	}
	return res
}
