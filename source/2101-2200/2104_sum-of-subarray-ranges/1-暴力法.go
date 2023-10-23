package main

func main() {

}

// leetcode2104_子数组范围和
func subArrayRanges(nums []int) int64 {
	res := int64(0)
	for i := 0; i < len(nums); i++ {
		minValue, maxValue := nums[i], nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > maxValue {
				maxValue = nums[j]
			}
			if nums[j] < minValue {
				minValue = nums[j]
			}
			res = res + int64(maxValue-minValue)
		}
	}
	return res
}
