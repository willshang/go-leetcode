package main

func main() {

}

// leetcode517_超级洗衣机
func findMinMoves(machines []int) int {
	n := len(machines)
	sum := 0
	for i := 0; i < n; i++ {
		sum = sum + machines[i]
	}
	if sum%n != 0 { // 先判断
		return -1
	}
	per := sum / n // 最终每个洗衣机里面的衣服数
	for i := 0; i < n; i++ {
		machines[i] = machines[i] - per // 计算每个洗衣机需要拿出或者需要放入的衣服数
	}
	maxValue := 0
	curSum := 0
	res := 0
	// 注意：选择任意m台，不要求连续
	// 2种情况：
	// 1、数组的最大值：是取出，每次一件，会有最大值的次数
	// 2、前缀和的最大绝对值：前面多余或者需要的数量
	for i := 0; i < n; i++ {
		curSum = curSum + machines[i]              // 前缀和：需要移动curSum件衣服到当前节点
		maxValue = max(maxValue, abs(curSum))      // 需要移动的最大值
		res = max(res, max(maxValue, machines[i])) // 取：数组的最大值和数组前缀和的绝对值的最大值中的较大值
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
