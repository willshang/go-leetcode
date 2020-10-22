package main

import "fmt"

func main() {
	fmt.Println(numberToWords(123))
}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	res := ""
	billion := num / 1000000000
	million := (num - billion*1000000000) / 1000000
	thousand := (num - billion*1000000000 - million*1000000) / 1000
	left := num - billion*1000000000 - million*1000000 - thousand*1000
	if billion != 0 {
		res += three(billion) + " Billion"
	}
	if million != 0 {
		if res != "" {
			res += " "
		}
		res += three(million) + " Million"
	}
	if thousand != 0 {
		if res != "" {
			res += " "
		}
		res += three(thousand) + " Thousand"
	}
	if left != 0 {
		if res != "" {
			res += " "
		}
		res += three(left)
	}
	return res
}

func three(num int) string {
	hundred := num / 100
	left := num - hundred*100
	if hundred == 0 {
		return two(num)
	}
	res := transfer[hundred] + " Hundred"
	if left != 0 {
		res += " " + two(left)
	}
	return res
}

func two(num int) string {
	if num == 0 {
		return ""
	} else if num < 10 {
		return transfer[num]
	} else if num < 20 {
		return transfer[num]
	}
	ten := num / 10
	left := num - ten*10
	ten = ten * 10
	res := transfer[ten]
	if left != 0 {
		res += " " + transfer[left]
	}
	return res
}

var transfer = map[int]string{
	0:  "Zero",
	1:  "One",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
	11: "Eleven",
	12: "Twelve",
	13: "Thirteen",
	14: "Fourteen",
	15: "Fifteen",
	16: "Sixteen",
	17: "Seventeen",
	18: "Eighteen",
	19: "Nineteen",
	20: "Twenty",
	30: "Thirty",
	40: "Forty",
	50: "Fifty",
	60: "Sixty",
	70: "Seventy",
	80: "Eighty",
	90: "Ninety",
}
