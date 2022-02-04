package main

import "sort"

func main() {

}

// leetcode2136_全部开花的最早一天
func earliestFullBloom(plantTime []int, growTime []int) int {
	n := len(plantTime)
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = []int{plantTime[i], growTime[i]}
	}
	sort.Slice(arr, func(i, j int) bool { // 按生长日期降序排序
		return arr[i][1] > arr[j][1]
	})
	res := 0
	sum := 0
	// 播种的总时间不变（串行），生长时间尽可能小（并行）
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i][0]
		if sum+arr[i][1] > res {
			res = sum + arr[i][1]
		}
	}
	return res
}
