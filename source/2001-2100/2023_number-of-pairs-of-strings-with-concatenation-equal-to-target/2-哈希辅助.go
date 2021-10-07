package main

func main() {

}

func numOfPairs(nums []string, target string) int {
	res := 0
	n := len(nums)
	m := make(map[string]int)
	for i := 0; i < n; i++ {
		m[nums[i]]++
	}
	for i := 1; i < len(target); i++ {
		a, b := target[:i], target[i:]
		if a == b {
			res = res + m[a]*(m[a]-1)
		} else {
			res = res + m[a]*m[b]
		}
	}
	return res
}
