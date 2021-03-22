package main

func main() {

}

// leetcode1790_仅执行一次字符串交换能否使两个字符串相等
func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if len(s1) != len(s2) {
		return false
	}
	arr := make([]int, 0)
	count := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			arr = append(arr, i)
			count++
		}
	}
	if count != 2 {
		return false
	}
	if s1[arr[0]] == s2[arr[1]] && s1[arr[1]] == s2[arr[0]] {
		return true
	}
	return false
}
