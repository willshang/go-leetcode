package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	m := map[uint8]string{
		'0': "ling",
		'1': "yi",
		'2': "er",
		'3': "san",
		'4': "si",
		'5': "wu",
		'6': "liu",
		'7': "qi",
		'8': "ba",
		'9': "jiu",
	}
	fmt.Scanf("%s", &str)
	sum := 0
	for k := range str {
		sum = sum + int(str[k]-'0')
	}

	toString := strconv.Itoa(sum)
	for k := range toString {
		if k != 0 {
			fmt.Print(" ")
		}
		fmt.Print(m[toString[k]])
	}
}
