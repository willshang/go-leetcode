package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		str := scan.Text()
		arr := strings.Fields(str)
		if len(arr) == 0 {
			fmt.Println(0)
		} else {
			fmt.Println(len(arr[len(arr)-1]))
		}
	}
	return
}
