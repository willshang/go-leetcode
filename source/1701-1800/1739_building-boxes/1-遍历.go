package main

import "fmt"

func main() {
	fmt.Println(minimumBoxes(4))
}

// leetcode1739_放置盒子
func minimumBoxes(n int) int {
	res := 0
	k := 1 // 第几层
	total := 0
	for {
		count := k * (k + 1) / 2 // 第几层数量
		if total+count > n {     // 大于n就不加了
			break
		}
		total = total + count
		k++
	}
	k--
	res = k * (k + 1) / 2 // 最底层
	k = 1
	for total < n { // 底层再从一条边上往墙角叠加
		total = total + k
		k++
		res++
	}
	return res
}
