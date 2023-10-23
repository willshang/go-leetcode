package main

import "fmt"

func main() {
	arr := []string{"42", "10", "O", "t", "y", "p", "g", "B", "96", "H", "5", "v", "P", "52", "25",
		"96", "b", "L", "Y", "z", "d", "52", "3", "v", "71", "J", "A", "0", "v", "51", "E", "k", "H",
		"96", "21", "W", "59", "I", "V", "s", "59", "w", "X", "33", "29", "H", "32", "51", "f", "i",
		"58", "56", "66", "90", "F", "10", "93", "53", "85", "28", "78", "d", "67", "81", "T", "K", "S",
		"l", "L", "Z", "j", "5", "R", "b", "44", "R", "h", "B", "30", "63", "z", "75", "60", "m", "61", "a",
		"5", "S", "Z", "D", "2", "A", "W", "k", "84", "44", "96", "96", "y", "M"}
	fmt.Println(findLongestSubarray(arr))
	//fmt.Println(findLongestSubarray([]string{"A", "1", "B", "C", "D", "2", "3", "4", "E", "5", "F", "G", "6", "7", "H", "I", "J", "K", "L", "M"}))
	//fmt.Println(findLongestSubarray([]string{"A", "1"}))
}

// 程序员面试金典17.05_字母与数字
func findLongestSubarray(array []string) []string {
	m := make(map[int]int)
	m[0] = 0
	res := 0
	begin := 0
	total := 0
	for i := 0; i < len(array); i++ {
		if '0' <= array[i][0] && array[i][0] <= '9' {
			total++
		} else {
			total--
		}
		if total == 0 {
			begin = 0
			res = i + 1
		} else if index, ok := m[total]; ok {
			if i-index > res {
				res = i - index
				begin = index + 1
			}
		} else {
			m[total] = i
		}
	}
	return array[begin : begin+res]
}
