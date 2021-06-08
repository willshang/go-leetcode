package main

func main() {

}

func minKBitFlips(nums []int, k int) int {
	n := len(nums)
	res := 0
	count := 0 // 翻转奇数偶数次
	for i := 0; i < n; i++ {
		if i >= k && nums[i-k] > 1 {
			count = 1 - count
		}
		if (nums[i]+count)%2 == 0 { // 是1翻转奇数次变为1，是0翻转偶数次为0
			if i+k > n {
				return -1
			}
			res++
			count = 1 - count
			nums[i] = nums[i] + 2 // 标记
		}
	}
	return res
}
