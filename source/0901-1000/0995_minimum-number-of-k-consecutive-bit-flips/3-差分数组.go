package main

func main() {

}

func minKBitFlips(nums []int, k int) int {
	n := len(nums)
	arr := make([]int, n+1) // 差分数组
	res := 0
	count := 0 // 翻转奇数偶数次
	for i := 0; i < n; i++ {
		count = (count + arr[i]) % 2
		if (count+nums[i])%2 == 0 { // 当前位置翻转count次后是0
			if i+k > n { // 剩下达不到k个一组
				return -1
			}
			count = 1 - count // 翻转+1
			arr[i]++          // 差分数组i 加1
			arr[i+k]--        // 差分数组i+k 减1
			res++
		}
	}
	return res
}
