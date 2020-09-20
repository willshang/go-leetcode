package main

import "fmt"

func main() {
	origin := []int{1, 4, 8, 10, 5, 2, 16} // 开始的数组
	d := make([]int, len(origin)+1)        // 多1位，在做d[end+1]不会超出范围
	// 1.计算差分数组
	d[0] = origin[0] // 第一项
	for i := 1; i < len(origin); i++ {
		d[i] = origin[i] - origin[i-1]
	}
	fmt.Println(origin) // [1 4 8 10 5 2 16]
	fmt.Println(d)      // [1 3 4 2 -5 -3 14 0]
	// 验证 origin[i] = 求和(d[0]+..+d[i]
	arr := make([]int, 0)
	total := 0
	for i := 0; i < len(d)-1; i++ {
		total = total + d[i]
		arr = append(arr, total)
	}
	fmt.Println(arr) // [1 4 8 10 5 2 16]

	// 2.操作
	// 开始坐标，结束坐标，操作次数
	operation := [][]int{
		{1, 2, 1},
		{2, 5, 1},
		{0, 6, 1},
	}
	for i := 0; i < len(operation); i++ {
		start := operation[i][0]
		end := operation[i][1]
		count := operation[i][2]
		// 差分操作
		d[start] = d[start] + count
		d[end+1] = d[end+1] - count
		// 正常模拟操作
		for j := start; j <= end; j++ {
			arr[j] = arr[j] + count
		}
	}
	fmt.Println(arr) // 正常模拟操作的结果: [2 6 11 12 7 4 17]
	res := make([]int, 0)
	total = 0
	for i := 0; i < len(d)-1; i++ {
		total = total + d[i]
		res = append(res, total)
	}
	fmt.Println(res) // 差分计算的结果：[2 6 11 12 7 4 17]
}
