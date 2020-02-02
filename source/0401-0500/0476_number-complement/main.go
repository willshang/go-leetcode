package main

import "fmt"

func main() {
	//fmt.Println(findComplement(5))
	fmt.Println(findComplement(2))
	fmt.Println(findComplement(5))
	fmt.Println(findComplement(8))
}

/*func findComplement(num int) int {
	temp := num
	res := 0
	for temp > 0{
		temp = temp >> 1
		res = res << 1
		res++
	}
	return res ^ num
}*/

func findComplement(num int) int {
	res := 0
	if num == 0 {
		return 1
	}
	if num == 1 {
		return 0
	}

	exp := 1
	for num > 0 {
		temp := num % 2
		if temp == 0 {
			res = res + exp
			exp = exp * 2
		} else {
			exp = exp * 2
		}

		num = num / 2
	}
	return res
}
