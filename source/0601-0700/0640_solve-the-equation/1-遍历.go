package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(solveEquation("x+5-3+x=6+x-2"))
}

// leetcode640_求解方程
func solveEquation(equation string) string {
	arr := strings.Split(equation, "=")
	left, right := split(arr[0]), split(arr[1])
	l, r := getValue(left) // l左边x的系数，r右边值
	a, b := getValue(right)
	l, r = l-a, r-b
	if l == r && l == 0 {
		return "Infinite solutions"
	} else if l == 0 && r != 0 {
		return "No solution"
	}
	return "x=" + fmt.Sprintf("%d", r/l)
}

func getValue(arr []string) (l, r int) {
	for i := 0; i < len(arr); i++ {
		s := arr[i]
		if strings.Contains(s, "x") == true {
			s = strings.ReplaceAll(s, "x", "")
			if s == "" || s == "+" || s == "-" {
				s = s + "1"
			}
			value, _ := strconv.Atoi(s)
			l = l + value
		} else {
			value, _ := strconv.Atoi(s)
			r = r - value
		}
	}
	return l, r
}

func split(str string) (res []string) {
	prev := ""
	for i := 0; i < len(str); i++ {
		if str[i] == '+' || str[i] == '-' {
			if len(prev) > 0 {
				res = append(res, prev)
				prev = ""
			}
		}
		prev = prev + string(str[i])
	}
	res = append(res, prev)
	return res
}
