package main

func main() {

}

func countBadPairs(nums []int) int64 {
	res := int64(0)
	n := int64(len(nums))
	m := make(map[int]int64)
	for i := 0; i < len(nums); i++ {
		diff := nums[i] - i
		res = res + m[diff] // 累计相等
		m[diff]++
	}
	return n*(n-1)/2 - res // 总数-不满足条件的数量
}
