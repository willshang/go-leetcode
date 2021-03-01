package main

func main() {

}

// leetcode930_和相同的二元子数组
func numSubarraysWithSum(A []int, S int) int {
	m := make(map[int]int)
	res := 0
	sum := 0
	m[0] = 1
	for i := 0; i < len(A); i++ {
		sum = sum + A[i]
		res = res + m[sum-S]
		m[sum]++
	}
	return res
}
