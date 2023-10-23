package main

func main() {

}

// leetcode1625_执行操作后字典序最小的字符串
func findLexSmallestString(s string, a int, b int) string {
	res := s
	m := make(map[string]bool)
	queue := make([]string, 0)
	queue = append(queue, s)
	for len(queue) > 0 {
		str := queue[0]
		queue = queue[1:]
		if m[str] == true {
			continue
		}
		m[str] = true
		if str < res {
			res = str
		}
		queue = append(queue, str[b:]+str[:b])
		queue = append(queue, add(str, a))
	}
	return res
}

func add(s string, a int) string {
	res := []byte(s)
	for i := 1; i < len(s); i = i + 2 {
		res[i] = byte('0' + (int(s[i]-'0')+a)%10)
	}
	return string(res)
}
