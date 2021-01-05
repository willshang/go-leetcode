package main

func main() {

}

// leetcode1390_四因数
func sumFourDivisors(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		count := 0
		for j := 1; j*j <= nums[i]; j++ {
			if nums[i]%j == 0 {
				count++
				sum = sum + j
				if j*j != nums[i] {
					count++
					sum = sum + nums[i]/j
				}
			}
		}
		if count == 4 {
			res = res + sum
		}
	}
	return res
}
