package main

import "fmt"

func main() {
	arr := []int{1, 0, 0}
	fmt.Println(isOneBitCharacter(arr))
	fmt.Println(isOneBitCharacter([]int{1, 1, 1, 1, 0}))
}

func isOneBitCharacter(bits []int) bool {
	n := len(bits)
	count := 0
	// 统计末尾1的个数，偶数正确，奇数错误
	for i := n - 2; i >= 0; i-- {
		if bits[i] == 0 {
			break
		} else {
			count++
		}
	}
	// return count & 1 == 0
	return count%2 == 0
}
