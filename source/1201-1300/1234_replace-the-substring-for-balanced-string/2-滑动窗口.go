package main

func main() {

}

func balancedString(s string) int {
	n := len(s)
	target := n / 4
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	if m['Q'] == target && m['W'] == target && m['E'] == target && m['R'] == target {
		return 0
	}
	res := n
	left := 0
	right := 0
	for left < n {
		if m['Q'] > target || m['W'] > target || m['E'] > target || m['R'] > target {
			if right < n {
				m[s[right]]--
				right++
			} else {
				break
			}
		} else {
			res = min(res, right-left)
			m[s[left]]++
			left++
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
