package main

import "fmt"

func main() {
	fmt.Println(convertToBase7(10))
	fmt.Println(convertToBase7(-7))
}
func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	minus := ""
	if num < 0 {
		minus = "-"
		num = -1 * num
	}

	s := ""
	for num > 0 {
		s = fmt.Sprintf("%d", num%7) + s
		num = num / 7
	}
	return minus + s
}
