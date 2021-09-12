package main

func main() {

}

// leetcode1991_找到数组的中间位置
func findMiddleIndex(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
	}
	count := 0
	for i := 0; i < len(nums); i++ {
		if 2*count == sum-nums[i] {
			return i
		}
		count = count + nums[i]
	}
	return -1
}
