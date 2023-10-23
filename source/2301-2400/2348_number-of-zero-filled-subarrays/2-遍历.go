package main

func main() {

}

func zeroFilledSubarray(nums []int) int64 {
	res := int64(0)
	count := int64(0)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count++
		} else {
			res = res + count*(count+1)/2
			count = 0
		}
	}
	res = res + count*(count+1)/2
	return res
}
