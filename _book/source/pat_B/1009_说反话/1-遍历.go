package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var str string
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	str = string(data)
	arr := strings.Split(str, " ")
	for k := len(arr) - 1; k >= 0; k-- {
		fmt.Print(arr[k])
		if k != 0 {
			fmt.Print(" ")
		}
	}
}
