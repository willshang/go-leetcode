package main

func main() {

}

// leetcode1750_删除字符串两端相同字符后的最短长度
func minimumLength(s string) int {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			break
		}
		temp := s[left]
		for left <= right && s[left] == temp {
			left++
		}
		for left <= right && s[right] == temp {
			right--
		}
	}
	return right - left + 1
}
