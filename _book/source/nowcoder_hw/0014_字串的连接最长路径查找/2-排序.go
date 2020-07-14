package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		n, _, _ := reader.ReadLine()
		if len(n) == 0 {
			break
		}
		num, _ := strconv.Atoi(string(n))
		arr := make([]string, 0)
		for i := 0; i < num; i++ {
			data, _, _ := reader.ReadLine()
			arr = append(arr, string(data))
		}
		sort.Strings(arr)
		for i := 0; i < len(arr); i++ {
			fmt.Println(arr[i])
		}
	}
}
