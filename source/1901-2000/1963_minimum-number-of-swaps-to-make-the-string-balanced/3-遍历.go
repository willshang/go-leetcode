package main

func main() {

}

func minSwaps(s string) int {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			stack = append(stack, '[')
		} else {
			if len(stack) > 0 && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack)/2 + len(stack)%2
}
