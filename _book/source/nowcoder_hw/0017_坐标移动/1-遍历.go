package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var str string
	for {
		n, _ := fmt.Scanf("%s", &str)
		if n == 0 {
			break
		}
		arr := strings.Split(str, ";")
		x := 0
		y := 0
		for i := 0; i < len(arr); i++ {
			if len(arr[i]) <= 1 || len(arr[i]) > 3 {
				continue
			}
			value, _ := strconv.Atoi(arr[i][1:])
			switch arr[i][0] {
			case 'A':
				x = x - value
			case 'D':
				x = x + value
			case 'W':
				y = y + value
			case 'S':
				y = y - value
			}
		}
		fmt.Printf("%d,%d\n", x, y)
	}
}
