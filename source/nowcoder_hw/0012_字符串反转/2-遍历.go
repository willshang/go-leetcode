package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		n, _, _ := reader.ReadLine()
		if len(n) == 0 {
			break
		}
		str := n
		for i := len(str) - 1; i >= 0; i-- {
			fmt.Print(string(str[i]))
		}
		fmt.Println()
	}
}
