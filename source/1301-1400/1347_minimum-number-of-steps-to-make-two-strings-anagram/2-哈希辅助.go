package main

func main() {

}

// leetcode1347_制造字母异位词的最小步骤数
func minSteps(s string, t string) int {
	res := 0
	m := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
		m[t[i]]--
	}
	for _, v := range m {
		if v > 0 {
			res = res + v
		}
	}
	return res
}
