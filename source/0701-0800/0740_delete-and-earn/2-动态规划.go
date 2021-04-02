package main

func main() {

}

func deleteAndEarn(nums []int) int {
	count := make([]int, 10001)
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}
	a, b := 0, 0 // a使用i，b不使用i
	prev := -1
	for i := 0; i < 10001; i++ {
		if count[i] > 0 {
			maxValue := max(a, b)
			if prev != i-1 { // 不等于上一个，使用最大值
				a = i*count[i] + maxValue
				b = maxValue
			} else { // 等于上一个，使用b
				a = i*count[i] + b
				b = maxValue
			}
			prev = i
		}
	}
	return max(a, b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
