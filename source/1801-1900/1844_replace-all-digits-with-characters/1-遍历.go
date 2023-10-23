package main

func main() {

}

// leetcode1844_将所有数字用字符替换
func replaceDigits(s string) string {
	res := []byte(s)
	for i := 1; i < len(s); i = i + 2 {
		res[i] = s[i-1] + s[i] - '0'
	}
	return string(res)
}
