package main

func main() {

}

// leetcode2109_向字符串添加空格
func addSpaces(s string, spaces []int) string {
	res := make([]byte, 0)
	if len(spaces) == 0 {
		return s
	}
	j := 0
	for i := 0; i < len(s); i++ {
		if j < len(spaces) && i == spaces[j] {
			res = append(res, ' ')
			j++
		}
		res = append(res, s[i])
	}
	return string(res)
}
