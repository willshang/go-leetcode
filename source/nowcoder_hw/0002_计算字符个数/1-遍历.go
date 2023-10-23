package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	str := strings.ToLower(data)
	// data, _ = reader.ReadString('\n')
	// c := strings.ToLower(data)[0]
	c, _, _ := reader.ReadRune()
	res := 0
	for i := 0; i < len(str); i++ {
		if str[i] == uint8(c) {
			res++
		}
	}
	fmt.Println(res)
	return
}
