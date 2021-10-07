package main

func main() {

}

func waysToPartition(nums []int, k int) int {
	res := 0
	n := len(nums)
	sum := 0
	prev := make(map[int]int) // 前缀和
	arr := make([]int, n+1)
	for i := 0; i < n; i++ {
		arr[i+1] = arr[i] + nums[i]
		sum = sum + nums[i]
		if i != n-1 {
			prev[sum]++
		}
	}
	if sum%2 == 0 {
		res = prev[sum/2] // 不改变的结果
	}
	suf := make(map[int]int) // 后缀和
	sufSum := 0
	temp := sum
	for i := n - 1; i >= 0; i-- { // 枚举每一位修改后的结果
		target := sum - nums[i] + k // 替换后的总和
		temp = temp - nums[i]
		prev[temp]-- // 前缀和 减一
		sufSum = sufSum + k
		suf[sufSum]++
		if target%2 == 0 {
			res = max(res, prev[target/2]+suf[target/2])
		}
		suf[sufSum]--
		sufSum = sufSum - k + nums[i]
		suf[sufSum]++
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
