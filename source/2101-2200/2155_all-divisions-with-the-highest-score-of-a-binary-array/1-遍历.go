package main

func main() {

}

func maxScoreIndices(nums []int) []int {
	res := []int{0}
	sum := 0
	maxValue := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			sum++
		} else {
			sum--
		}
		if sum > maxValue {
			maxValue = sum
			res = []int{i + 1}
		} else if sum == maxValue {
			res = append(res, i+1)
		}
	}
	return res
}
