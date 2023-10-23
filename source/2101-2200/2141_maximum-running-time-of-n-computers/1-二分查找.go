package main

func main() {

}

// leetcode2141_同时运行N台电脑的最长时间
func maxRunTime(n int, batteries []int) int64 {
	sum := 0
	for i := 0; i < len(batteries); i++ {
		sum = sum + batteries[i]
	}
	left, right := 1, sum/n
	res := 0
	for left <= right {
		mid := left + (right-left)/2
		total := 0
		for j := 0; j < len(batteries); j++ {
			// 假设让n台电脑运行mid分钟
			// 电量>=mid的电池，只能被使用mid分钟，仅能给1台电脑充电
			total = total + min(batteries[j], mid)
		}
		if total >= n*mid {
			res = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return int64(res)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
