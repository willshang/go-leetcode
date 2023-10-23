package main

func main() {

}

// leetcode1957_删除字符使字符串变好
func makeFancyString(s string) string {
	res := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if len(res) >= 2 && res[len(res)-1] == s[i] && res[len(res)-2] == s[i] {
			continue
		}
		res = append(res, s[i])
	}
	return string(res)
}
