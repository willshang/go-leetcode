package main

func main() {

}

// leetcode1814_统计一个数组中好对子的数目
func countNicePairs(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]-reverse(nums[i])]++
	}
	res := 0
	for _, v := range m {
		res = (res + v*(v-1)/2) % 1000000007
	}
	return res
}

func reverse(num int) int {
	res := 0
	for num > 0 {
		res = res*10 + num%10
		num = num / 10
	}
	return res
}
