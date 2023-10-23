package main

func main() {

}

// leetcode1862_向下取整数对和
func sumOfFlooredPairs(nums []int) int {
	n := len(nums)
	count := make([]int, 200001)
	for i := 0; i < n; i++ {
		count[nums[i]]++
	}
	arr := make([]int, 200001+1) // 前缀和
	for i := 0; i < len(count); i++ {
		arr[i+1] = arr[i] + count[i]
	}
	res := 0
	// a/b = c
	for i := 0; i < len(count); i++ { // 枚举b
		if count[i] > 0 { // i个数大于0
			for j := 1; i*j <= 100000; j++ { // 枚举c
				// b的个数 X c X 目标范围内的数字个数
				total := count[i] * j * (arr[i*(j+1)] - arr[i*j])
				res = (res + total) % 1000000007
			}
		}
	}
	return res
}
