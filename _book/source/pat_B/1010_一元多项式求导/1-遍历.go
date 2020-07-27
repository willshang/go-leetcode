package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var str string
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	str = string(data)
	var flag = false
	arr := strings.Fields(str)
	for k := 0; k < len(arr)-1; k = k + 2 {
		a, _ := strconv.Atoi(arr[k])
		b, _ := strconv.Atoi(arr[k+1])
		if a == 0 && b == 0 {
			fmt.Print("0 0")
			flag = true
		}
		if b != 0 {
			if flag == true {
				fmt.Print(" ")
			}
			fmt.Print(a*b, " ", b-1)
			flag = true
		}
	}
	if flag == false {
		fmt.Print("0 0")
	}
}
