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
		arr := strings.Split(str, " ")
		for i := 0; i < len(arr)/2; i++ {
			arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
		}
		fmt.Println(strings.Join(arr, " "))
	}
}
