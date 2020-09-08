package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//fmt.Println(fractionToDecimal(1, 2))
	fmt.Println(fractionToDecimal(2, 3))
}

// leetcode166_分数到小数
func fractionToDecimal(numerator int, denominator int) string {
	res := make([]string, 0)
	if numerator == 0 {
		return "0"
	}
	flag := false
	if numerator < 0 {
		flag = !flag
		numerator = -numerator
	}
	if denominator < 0 {
		flag = !flag
		denominator = -denominator
	}
	if flag == true {
		res = append(res, "-")
	}
	a, b := numerator/denominator, numerator%denominator
	res = append(res, strconv.Itoa(a))
	if b == 0 {
		return strings.Join(res, "")
	}
	res = append(res, ".")
	m := make(map[int]int)
	last, index := -1, len(res)
	for b > 0 {
		b = b * 10
		if v, ok := m[b]; ok {
			last = v
			break
		} else {
			m[b] = index
		}
		index++
		a, b = b/denominator, b%denominator
		res = append(res, strconv.Itoa(a))
	}
	if last != -1 {
		res = append(res[:last], append([]string{"("}, res[last:]...)...)
		res = append(res, ")")
	}
	return strings.Join(res, "")
}
