package main

func main() {

}

// leetcode1726_同积元组
func tupleSameProduct(nums []int) int {
	res := 0
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			total := nums[i] * nums[j]
			m[total]++
		}
	}
	for _, v := range m {
		if v >= 2 {
			res = res + v*(v-1)/2*8
		}
	}
	return res
}
