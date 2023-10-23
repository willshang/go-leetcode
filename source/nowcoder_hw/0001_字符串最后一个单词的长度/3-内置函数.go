package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	//data, _ := reader.ReadBytes('\n')
	data, _ := reader.ReadString('\n')
	str := string(data)
	arr := strings.Fields(str)
	if len(arr) == 0 {
		fmt.Println(0)
		return
	}
	fmt.Println(len(arr[len(arr)-1]))
	return
}
