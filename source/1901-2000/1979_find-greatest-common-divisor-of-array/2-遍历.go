package main

func main() {

}

func findGCD(nums []int) int {
	a, b := nums[0], nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] > a {
			a = nums[i]
		}
		if nums[i] < b {
			b = nums[i]
		}
	}
	return gcd(a, b)
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
