package main

import "fmt"

func main() {
	fmt.Println(intToRoman(4))
}

func intToRoman(num int) string {
	res := ""
	arr1 := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	arr2 := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	arr3 := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	arr4 := []string{"", "M", "MM", "MMM"}
	res = arr4[num/1000] + arr3[num%1000/100] + arr2[num%100/10] + arr1[num%10]
	return res
}
