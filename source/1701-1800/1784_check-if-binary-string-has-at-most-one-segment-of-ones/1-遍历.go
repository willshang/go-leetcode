package main

func main() {

}

// leetcode1784_检查二进制字符串字段
func checkOnesSegment(s string) bool {
	flag := true
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			flag = false
		}
		if flag == false && s[i] == '1' {
			return false
		}
	}
	return true
}
