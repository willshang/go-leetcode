package main

func main() {

}

// leetcode2155_分组得分最高的所有下标
func maxScoreIndices(nums []int) []int {
	res := []int{0}
	left, right := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			right++
		}
	}
	maxValue := left + right
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			left++
		} else {
			right--
		}
		if left+right > maxValue {
			maxValue = left + right
			res = []int{i + 1}
		} else if left+right == maxValue {
			res = append(res, i+1)
		}
	}
	return res
}
