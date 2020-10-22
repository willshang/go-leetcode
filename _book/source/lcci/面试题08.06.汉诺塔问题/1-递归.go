package main

import "fmt"

func main() {
	fmt.Println(hanota([]int{2, 1, 0}, []int{}, []int{}))
}

// 程序员面试金典08.06_汉诺塔问题
func hanota(A []int, B []int, C []int) []int {
	if A == nil {
		return nil
	}
	move(len(A), &A, &B, &C)
	return C
}

func move(num int, A, B, C *[]int) {
	if num < 0 {
		return
	}
	if num == 1 {
		*C = append(*C, (*A)[len(*A)-1])
		*A = (*A)[:len(*A)-1]
		return
	}
	move(num-1, A, C, B)
	move(1, A, B, C)
	move(num-1, B, A, C)
}
