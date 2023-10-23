package main

func main() {

}

// leetcode1946_子字符串突变后可能得到的最大整数
func maximumNumber(num string, change []int) string {
	res := []byte(num)
	flag := false
	for i := 0; i < len(num); i++ {
		value := int(num[i] - '0')
		if value < change[value] {
			flag = true
			res[i] = byte(change[value] + '0')
		} else if value == change[value] {
			res[i] = byte(change[value] + '0')
		} else {
			if flag == true {
				break
			}
		}
	}
	return string(res)
}
