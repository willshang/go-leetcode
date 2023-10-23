package main

func main() {

}

// leetcode327_区间和的个数
func countRangeSum(nums []int, lower int, upper int) int {
	n := len(nums)
	arr := make([]int, n+1) // 前缀和
	for i := 0; i < n; i++ {
		arr[i+1] = arr[i] + nums[i]
	}
	return countSum(arr, lower, upper)
}

func countSum(nums []int, lower int, upper int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	a := append([]int{}, nums[:n/2]...)
	b := append([]int{}, nums[n/2:]...)
	res := countSum(a, lower, upper) + countSum(b, lower, upper)
	left, right := 0, 0
	for i := 0; i < len(a); i++ {
		for left < len(b) && b[left]-a[i] < lower {
			left++
		}
		for right < len(b) && b[right]-a[i] <= upper {
			right++
		}
		res = res + right - left
	}
	i, j := 0, 0
	for k := 0; k < len(nums); k++ { // 2个有序数组合并
		if i < len(a) && (j == len(b) || a[i] <= b[j]) {
			nums[k] = a[i]
			i++
		} else {
			nums[k] = b[j]
			j++
		}
	}
	return res
}
