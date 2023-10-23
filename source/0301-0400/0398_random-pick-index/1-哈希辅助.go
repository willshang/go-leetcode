package main

import (
	"math/rand"
	"time"
)

func main() {

}

type Solution struct {
	m map[int][]int
	r *rand.Rand
}

func Constructor(nums []int) Solution {
	res := Solution{
		m: make(map[int][]int),
		r: rand.New(rand.NewSource(time.Now().Unix())),
	}
	for i := 0; i < len(nums); i++ {
		res.m[nums[i]] = append(res.m[nums[i]], i)
	}
	return res
}

func (this *Solution) Pick(target int) int {
	arr := this.m[target]
	index := this.r.Intn(len(arr))
	return arr[index]
}
