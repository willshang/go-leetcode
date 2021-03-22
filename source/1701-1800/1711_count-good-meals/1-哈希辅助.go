package main

func main() {

}

// leetcode1711_大餐计数
func countPairs(deliciousness []int) int {
	res := 0
	m := make(map[int]int)
	for i := 0; i < len(deliciousness); i++ {
		total := 1
		for j := 0; j <= 21; j++ {
			if m[total-deliciousness[i]] > 0 {
				res = res + m[total-deliciousness[i]]
			}
			total = total * 2
		}
		m[deliciousness[i]]++
	}
	return res % 1000000007
}
