package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	str := scan.Text()
	arr := strings.Fields(str)
	if len(arr) == 0 {
		fmt.Println(0)
		return
	}
	fmt.Println(len(arr[len(arr)-1]))
	return
}
