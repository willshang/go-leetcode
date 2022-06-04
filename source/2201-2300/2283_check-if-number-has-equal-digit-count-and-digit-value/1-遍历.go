package main

func main() {

}

// leetcode2283_判断一个数的数字计数是否等于数位的值
func digitCount(num string) bool {
	m := make(map[int]int)
	for i := 0; i < len(num); i++ {
		m[int(num[i]-'0')]++
	}
	for i := 0; i < len(num); i++ {
		if m[i] != int(num[i]-'0') {
			return false
		}
	}
	return true
}
