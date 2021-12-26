package main

func main() {

}

// leetcode2100_适合打劫银行的日子
func goodDaysToRobBank(security []int, time int) []int {
	n := len(security)
	if 2*time >= n {
		return nil
	}
	left, right := make([]int, n), make([]int, n)
	for i := 1; i < n; i++ {
		if security[i-1] >= security[i] {
			left[i] = left[i-1] + 1
		}
	}
	for i := n - 2; i >= 0; i-- {
		if security[i] <= security[i+1] {
			right[i] = right[i+1] + 1
		}
	}
	res := make([]int, 0)
	for i := time; i < n-time; i++ {
		if left[i] >= time && right[i] >= time {
			res = append(res, i)
		}
	}
	return res
}
