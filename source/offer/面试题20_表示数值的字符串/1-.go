package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(isNumber("-123"))
	//fmt.Println(isNumber("e"))
	//fmt.Println(isNumber("."))
	fmt.Println(isNumber("0e"))
}

var str string
func isNumber(s string) bool {
	str = strings.Trim(s," ")
	if s == "" || len(s) == 0 || len(str) == 0{
		return false
	}
	numeric := scanInteger()
	i := 0
	if len(str) > 0 && str[i] == '.'{
		str = str[1:]
		numeric = scanUnsignedInteger() || numeric
	}
	fmt.Println(str)
	if len(str) > 0 && (str[i] == 'e' || str[i] == 'E'){
		str = str[1:]
		fmt.Println(str)
		numeric = numeric && scanInteger()
	}

	return numeric && len(str) == 0
}

func scanInteger() bool  {
	if str[0] == '+' || str[0] == '-'{
		str = str[1:]
	}
	return scanUnsignedInteger()
}

func scanUnsignedInteger() bool {
	i := 0
	for i < len(str){
		if str[i] >= '0' && str[i] <= '9'{
			i++
		}else {
			if i == 0{
				return false
			}
			break
		}
	}
	if i >= len(str){
		str = ""
	}else {
		str = str[i:]
	}
	return i > 0
}