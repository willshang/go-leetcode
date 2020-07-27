package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		n, _, _ := reader.ReadLine()
		if len(n) == 0 {
			break
		}
		str := string(n)
		arr := strings.Split(str, " ")
		for i := 0; i < len(arr)/2; i++ {
			arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
		}
		fmt.Println(strings.Join(arr, " "))
	}
}
