package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := Constructor(1000000000, []int{1, 2, 3, 440404, 999999998})
	//a := Constructor(4, []int{0, 1, 2})
	fmt.Println(a)
	for i := 0; i < 10; i++ {
		fmt.Println(a.Pick())
	}
}

// leetcode710_黑名单中的随机数
type Solution struct {
	m map[int]int
	N int
}

func Constructor(N int, blacklist []int) Solution {
	m := make(map[int]int)
	temp := make(map[int]bool)
	for i := 0; i < len(blacklist); i++ {
		temp[blacklist[i]] = true
	}
	length := N - len(blacklist)
	arr := make([]int, 0) // 需要替换为较大数
	for i := 0; i < len(blacklist); i++ {
		if blacklist[i] < length {
			arr = append(arr, blacklist[i])
		}
	}
	a := make([]int, 0) // 没有使用过的较大数
	for i := length; i < N; i++ {
		if temp[i] == false {
			a = append(a, i)
		}
	}

	for i := 0; i < len(a); i++ {
		m[arr[i]] = a[i]
	}
	return Solution{
		m: m,
		N: length,
	}
}

func (this *Solution) Pick() int {
	index := rand.Intn(this.N)
	if value, ok := this.m[index]; ok {
		return value
	}
	return index
}
