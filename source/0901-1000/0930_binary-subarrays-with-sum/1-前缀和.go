package main

func main() {

}

func numSubarraysWithSum(A []int, S int) int {
	m := make(map[int]int)
	res := 0
	sum := 0
	for i := 0; i < len(A); i++ {
		sum = sum + A[i]
		if sum == S {
			res++
		}
		res = res + m[sum-S]
		m[sum]++
	}
	return res
}
