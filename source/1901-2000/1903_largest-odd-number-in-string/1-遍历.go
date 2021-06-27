package main

func main() {

}

// leetcode1903_字符串中的最大奇数
func largestOddNumber(num string) string {
	n := len(num)
	for i := n - 1; i >= 0; i-- {
		if int(num[i]-'0')%2 == 1 {
			return num[:i+1]
		}
	}
	return ""
}
