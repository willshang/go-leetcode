package main

func main() {

}

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	m1, m2, m3 := make(map[int]int), make(map[int]int), make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		m1[nums1[i]] = 1
	}
	for i := 0; i < len(nums2); i++ {
		m2[nums2[i]] = 1
	}
	for i := 0; i < len(nums3); i++ {
		m3[nums3[i]] = 1
	}
	res := make([]int, 0)
	for i := 1; i <= 300; i++ {
		a := m1[i] + m2[i] + m3[i]
		if a >= 2 {
			res = append(res, i)
		}
	}
	return res
}
