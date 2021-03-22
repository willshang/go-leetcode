package main

import (
	"strconv"
	"strings"
)

func main() {

}

func maxDiff(num int) int {
	maxValue, minValue := num, num
	str := strconv.Itoa(num)
	for i := 0; i < len(str); i++ {
		if str[i] < '9' {
			maxValue, _ = strconv.Atoi(strings.ReplaceAll(str, string(str[i]), "9"))
			break
		}
	}
	if str[0] > '1' {
		minValue, _ = strconv.Atoi(strings.ReplaceAll(str, string(str[0]), "1"))
	} else {
		for i := 1; i < len(str); i++ {
			if str[i] > '1' && str[0] != str[i] {
				minValue, _ = strconv.Atoi(strings.ReplaceAll(str, string(str[i]), "0"))
				break
			}
		}
	}
	return maxValue - minValue
}
