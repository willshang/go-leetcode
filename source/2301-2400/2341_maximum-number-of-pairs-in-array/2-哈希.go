package main

func main() {

}

func numberOfPairs(nums []int) []int {
	a := 0
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		if m[nums[i]]%2 == 0 {
			a++
		}
	}
	return []int{a, len(nums) - 2*a}
}
