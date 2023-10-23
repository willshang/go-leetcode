package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var num string
	for {
		data, _, _ := reader.ReadLine()
		if len(data) == 0 {
			break
		}
		num = string(data[2:])
		value, _ := strconv.ParseInt(num, 16, 32)
		fmt.Println(value)
	}
}
