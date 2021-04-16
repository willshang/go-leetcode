package main

func main() {

}

// leetcode1822_数组元素积的符号
func arraySign(nums []int) int {
	res := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			res = -res
		} else if nums[i] == 0 {
			return 0
		}
	}
	return res
}
