package main

func main() {

}

// 剑指OfferII011.0和1个数相同的子数组
func findMaxLength(nums []int) int {
	res := 0
	m := make(map[int]int)
	m[0] = -1
	total := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			total--
		} else {
			total++
		}
		if first, ok := m[total]; !ok {
			m[total] = i
		} else {
			if i-first > res {
				res = i - first
			}
		}
	}
	return res
}
