package main

import "fmt"

func main() {
	fmt.Println(intToRoman(4))
}

// leetcode12_整数转罗马数字
func intToRoman(num int) string {
	m := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}
	arr := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	result := ""
	for i := 0; i < len(arr); i++ {
		if num == 0 {
			break
		}
		value := num / arr[i]
		for j := 0; j < value; j++ {
			result = result + m[arr[i]]
		}
		num = num - value*arr[i]
	}
	return result
}
