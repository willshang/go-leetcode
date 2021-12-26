package main

func main() {

}

func targetIndices(nums []int, target int) []int {
	res := make([]int, 0)
	start := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			count++
		} else if nums[i] < target {
			start++
		}
	}
	for i := start; i < start+count; i++ {
		res = append(res, i)
	}
	return res
}
