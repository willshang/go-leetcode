package main

func main() {

}

// leetcode2215_找出两数组的不同
func findDifference(nums1 []int, nums2 []int) [][]int {
	m1, m2 := make(map[int]bool), make(map[int]bool)
	a, b := make([]int, 0), make([]int, 0)
	for i := 0; i < len(nums1); i++ {
		m1[nums1[i]] = true
	}
	for i := 0; i < len(nums2); i++ {
		m2[nums2[i]] = true
	}
	for k := range m1 {
		if m2[k] == false {
			a = append(a, k)
		}
	}
	for k := range m2 {
		if m1[k] == false {
			b = append(b, k)
		}
	}
	return [][]int{a, b}
}
