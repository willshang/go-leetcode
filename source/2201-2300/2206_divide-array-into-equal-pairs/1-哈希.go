package main

func main() {

}

// leetcode2206_将数组划分成相等数对
func divideArray(nums []int) bool {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for _, v := range m {
		if v%2 == 1 {
			return false
		}
	}
	return true
}
