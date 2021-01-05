package main

func main() {

}

func reverseParentheses(s string) string {
	arr := []byte(s)
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			reverse(arr)
		}
	}
}
