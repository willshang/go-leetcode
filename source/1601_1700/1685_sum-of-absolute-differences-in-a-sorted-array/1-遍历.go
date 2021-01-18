package main

func main() {

}

func getSumAbsoluteDifferences(nums []int) []int {
	n := len(nums)
	res := make([]int, 0)
	right := 0 // 右边和
	left := 0  // 左边和
	for i := 1; i < n; i++ {
		right = right + (nums[i] - nums[0])
	}
	res = append(res, right)
	for i := 1; i < n; i++ {
		diff := nums[i] - nums[i-1]
		left = left + diff*i
		right = right - diff*(n-i)
		res = append(res, left+right)
	}
	return res
}
