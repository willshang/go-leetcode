package main

func main() {

}

func reverseParentheses(s string) string {
	stack := make([]int, 0)
	m := make(map[int]int)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			last := stack[len(stack)-1]
			m[i] = last
			m[last] = i
			stack = stack[:len(stack)-1]
		}
	}
	dir := 1
	res := make([]byte, 0)
	for i := 0; i < len(s); i = i + dir {
		if s[i] == '(' || s[i] == ')' {
			i = m[i]   // 从反方向遍历
			dir = -dir // 变换方向，+1/-1
		} else {
			res = append(res, s[i])
		}
	}
	return string(res)
}
