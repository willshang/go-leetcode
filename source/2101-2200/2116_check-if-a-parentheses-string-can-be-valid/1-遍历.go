package main

func main() {

}

// leetcode2116_判断一个括号字符串是否有效
func canBeValid(s string, locked string) bool {
	if len(s)%2 == 1 {
		return false
	}
	var arr []byte
	for i := 0; i < len(locked); i++ {
		if locked[i] == '1' {
			arr = append(arr, s[i])
		} else {
			arr = append(arr, '*')
		}
	}
	return checkValidString(string(arr))
}

// leetcode678.有效的括号字符串
func checkValidString(s string) bool {
	// 第1次把星号当左括号看
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			right++
		} else {
			left++
		}
		if right > left {
			return false
		}
	}
	// 第2次把星号当右括号看
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left > right {
			return false
		}
	}
	return true
}
