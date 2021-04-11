package main

func main() {

}

// leetcode1234_替换子串得到平衡字符串
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
	for right := 0; right < n; right++ {
		m[s[right]]--
		for left < n && left <= right && // [left,right]之间子串都是需要替换的（减去），使得m[x]<=target
			m['Q'] <= target && m['W'] <= target &&
			m['E'] <= target && m['R'] <= target {
			res = min(res, right-left+1)
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
