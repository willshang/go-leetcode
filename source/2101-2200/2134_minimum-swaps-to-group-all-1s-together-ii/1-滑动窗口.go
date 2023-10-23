package main

func main() {

}

func minSwaps(nums []int) int {
	n := len(nums)
	count := 0 // 1的个数=>滑动窗口的大小
	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			count++
		}
	}
	res := 0 // 枚举所有起点，统计滑动窗口内0的个数
	for i := 0; i < count; i++ {
		if nums[i] == 0 {
			res++
		}
	}
	temp := res
	for i := 0; i < n-1; i++ {
		if nums[i] == 0 { // 左边
			temp--
		}
		if nums[(i+count)%n] == 0 { // 右边
			temp++
		}
		res = min(res, temp)
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
