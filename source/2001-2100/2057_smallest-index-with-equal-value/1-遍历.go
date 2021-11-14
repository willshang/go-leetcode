package main

func main() {

}

// leetcode2057_值相等的最小索引
func smallestEqual(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if i%10 == nums[i] {
			return i
		}
	}
	return -1
}
