package main

func main() {

}

func maximumGain(s string, x int, y int) int {
	if x > y {
		x, y = y, x
		s = reverse(s)
	}
	res := 0
	stack := make([]byte, 0)
	// 先处理ba,y分多
	for i := 0; i < len(s); i++ {
		if s[i] != 'a' {
			stack = append(stack, s[i])
		} else {
			if len(stack) > 0 && stack[len(stack)-1] == 'b' {
				stack = stack[:len(stack)-1]
				res = res + y
			} else {
				stack = append(stack, s[i])
			}
		}
	}
	// 处理ab
	temp := make([]byte, 0)
	for len(stack) > 0 {
		c := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if c != 'a' {
			temp = append(temp, c)
		} else {
			if len(temp) > 0 && temp[len(temp)-1] == 'b' {
				temp = temp[:len(temp)-1]
				res = res + x
			} else {
				temp = append(temp, c)
			}
		}
	}
	return res
}

func reverse(s string) string {
	arr := []byte(s)
	for i := 0; i < len(s)/2; i++ {
		arr[i], arr[len(s)-1-i] = arr[len(s)-1-i], arr[i]
	}
	return string(arr)
}
