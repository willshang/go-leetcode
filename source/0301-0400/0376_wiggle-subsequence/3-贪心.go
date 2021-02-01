package main

func main() {

}

func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	res := 1
	diff := nums[1] - nums[0]
	if diff != 0 {
		res = 2
	}
	for i := 2; i < n; i++ {
		newDiff := nums[i] - nums[i-1]
		if (diff >= 0 && newDiff < 0) ||
			(diff <= 0 && newDiff > 0) {
			res++
			diff = newDiff
		}
	}
	return res
}
