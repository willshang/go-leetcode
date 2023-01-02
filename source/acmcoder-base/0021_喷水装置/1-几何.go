package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]float64, n)
	for i := 0; i < n; i++ {
		var temp float64
		fmt.Scan(&temp)
		arr[i] = temp
	}
	res := getResult(arr, 20)
	fmt.Println(res)
}

func getResult(arr []float64, length float64) int {
	res := 0
	count := float64(0)
	sort.Sort(sort.Reverse(sort.Float64Slice(arr))) // 倒序排序
	for i := 0; i < len(arr); i++ {
		count = count + getValue(arr[i])
		res++
		if count >= length {
			return res
		}
	}
	return res
}

// 计算半径为r的圆可以覆盖的宽度
func getValue(r float64) float64 {
	if r <= 1 {
		return 0
	}
	// 放在中心轴，形成的直角三角形：r*r = 1*1 + x*x
	// 求 2x => 2 * sqrt(r*r-1)
	return 2 * math.Sqrt(r*r-1)
}
