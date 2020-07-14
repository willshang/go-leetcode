package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _, _ := reader.ReadLine()
	char, _, _ := reader.ReadLine()
	s := strings.ToLower(string(str))
	c := strings.ToLower(string(char))
	fmt.Println(strings.Count(s, c))
	return
}
