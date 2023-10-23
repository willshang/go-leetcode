package main

import (
	"strconv"
	"strings"
)

func main() {

}

func sortSentence(s string) string {
	arr := strings.Split(s, " ")
	temp := make([]string, len(arr)+1)
	for i := 0; i < len(arr); i++ {
		index, _ := strconv.Atoi(string(arr[i][len(arr[i])-1]))
		temp[index] = arr[i][:len(arr[i])-1]
	}
	return strings.Join(temp[1:], " ")
}
