package main

func main() {

}

// leetcode1807_替换字符串中的括号内容
func evaluate(s string, knowledge [][]string) string {
	m := make(map[string]string)
	for i := 0; i < len(knowledge); i++ {
		a, b := knowledge[i][0], knowledge[i][1]
		m[a] = b
	}
	res := ""
	left := -1
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left = i
		} else if s[i] == ')' {
			str := s[left+1 : i]
			if v, ok := m[str]; ok {
				res = res + v
			} else {
				res = res + "?"
			}
			left = -1
		} else if left == -1 {
			res = res + string(s[i])
		}
	}
	return res
}
