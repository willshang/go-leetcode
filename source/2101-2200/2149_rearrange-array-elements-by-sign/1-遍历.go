package main

func main() {

}

// leetcode2149_按符号重排数组
func rearrangeArray(nums []int) []int {
	res := make([]int, len(nums))
	i, j := 0, 1
	for k := 0; k < len(nums); k++ {
		if nums[k] > 0 {
			res[i] = nums[k]
			i = i + 2
		} else {
			res[j] = nums[k]
			j = j + 2
		}
	}
	return res
}
