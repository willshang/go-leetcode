package main

func main() {

}

func zeroFilledSubarray(nums []int) int64 {
	res := int64(0)
	count := int64(0)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count++
			res = res + count
		} else {
			count = 0
		}
	}
	return res
}
