package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		n, _, _ := reader.ReadLine()
		if len(n) <= 0 {
			break
		}
		num, _ := strconv.Atoi(string(n))
		temp := make([]bool, 1024)
		for i := 0; i < num; i++ {
			data, _, _ := reader.ReadLine()
			value, _ := strconv.Atoi(string(data))
			temp[value] = true
		}
		for i := 0; i < len(temp); i++ {
			if temp[i] == true {
				fmt.Println(i)
			}
		}
	}
}
