package main

import "math/rand"

func main() {

}

// leetcode497_非重叠矩形中的随机点
type Solution struct {
	nums  []int   // 前缀和
	total int     // 总和
	rects [][]int // 原数组
}

func Constructor(rects [][]int) Solution {
	arr := make([]int, 0)
	total := 0
	for i := 0; i < len(rects); i++ {
		x1, y1, x2, y2 := rects[i][0], rects[i][1], rects[i][2], rects[i][3]
		total = total + (x2-x1+1)*(y2-y1+1) // x点的个数 * y点的个数：注意包含边+1
		arr = append(arr, total)
	}
	return Solution{nums: arr, total: total, rects: rects}
}

// leetcode528.按权重随机选择
func (this *Solution) Pick() []int {
	target := rand.Intn(this.total) // 目标值
	left, right := 0, len(this.nums)
	for left < right {
		mid := left + (right-left)/2
		if this.nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	temp := this.rects[left]
	x1, y1, x2, y2 := temp[0], temp[1], temp[2], temp[3]
	width := x2 - x1 + 1
	height := y2 - y1 + 1
	start := this.nums[left] - width*height // 前缀和-目标值
	return []int{x1 + (target-start)%width, y1 + (target-start)/width}
}
