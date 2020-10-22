package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(numberToWords(123))
}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	return strings.Trim(dfs(num), " ")
}

func dfs(n int) string {
	if n < 20 {
		return transfer[n]
	}
	if n < 100 {
		return transfer[n/10*10] + dfs(n%10)
	}
	if n < 1000 {
		return transfer[n/100] + "Hundred " + dfs(n%100)
	}
	if n < 1000000 {
		return dfs(n/1000) + "Thousand " + dfs(n%1000)
	}
	if n < 1000000000 {
		return dfs(n/1000000) + "Million " + dfs(n%1000000)
	}
	return dfs(n/1000000000) + "Billion " + dfs(n%1000000000)
}

var transfer = map[int]string{
	1:  "One ",
	2:  "Two ",
	3:  "Three ",
	4:  "Four ",
	5:  "Five ",
	6:  "Six ",
	7:  "Seven ",
	8:  "Eight ",
	9:  "Nine ",
	10: "Ten ",
	11: "Eleven ",
	12: "Twelve ",
	13: "Thirteen ",
	14: "Fourteen ",
	15: "Fifteen ",
	16: "Sixteen ",
	17: "Seventeen ",
	18: "Eighteen ",
	19: "Nineteen ",
	20: "Twenty ",
	30: "Thirty ",
	40: "Forty ",
	50: "Fifty ",
	60: "Sixty ",
	70: "Seventy ",
	80: "Eighty ",
	90: "Ninety ",
}
