package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	for {
		scan.Scan()
		n := scan.Text()
		if len(n) <= 0 {
			break
		}
		num, _ := strconv.Atoi(string(n))
		temp := make([]bool, 1024)
		for i := 0; i < num; i++ {
			scan.Scan()
			data := scan.Text()
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
