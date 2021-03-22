package main

func main() {

}

func waysToMakeFair(nums []int) int {
	res := 0
	a := 0
	b := 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			a = a + nums[i]
		} else {
			b = b + nums[i]
		}
	}
	x, y := 0, 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			a = a - nums[i]
		} else {
			b = b - nums[i]
		}
		even := x + b
		odd := y + a
		if even == odd {
			res++
		}
		if i%2 == 0 {
			x = x + nums[i]
		} else {
			y = y + nums[i]
		}
	}
	return res
}
