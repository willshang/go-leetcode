package main

func main() {

}

func arithmeticTriplets(nums []int, diff int) int {
	res := 0
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if m[nums[i]-diff] && m[nums[i]-2*diff] {
			res++
		}
		m[nums[i]] = true
	}
	return res
}
