package main

func main() {

}

// leetcode2270_分割数组的方案数
func waysToSplitArray(nums []int) int {
	res := 0
	sum, temp := 0, 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
	}
	for i := 0; i < len(nums)-1; i++ {
		temp = temp + nums[i]
		sum = sum - nums[i]
		if temp >= sum {
			res++
		}
	}
	return res
}
