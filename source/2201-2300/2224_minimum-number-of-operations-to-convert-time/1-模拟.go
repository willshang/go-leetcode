package main

func main() {

}

// leetcode2224_转化时间需要的最少操作数
func convertTime(current string, correct string) int {
	diff := parseTime(correct) - parseTime(current)
	res := 0
	arr := []int{60, 15, 5, 1}
	for i := 0; i < len(arr); i++ {
		res = res + diff/arr[i]
		diff = diff % arr[i]
	}
	return res
}

func parseTime(s string) int {
	a, b := int(s[0]-'0'), int(s[1]-'0')
	c, d := int(s[3]-'0'), int(s[4]-'0')
	return (a*10+b)*60 + (c*10 + d)
}
