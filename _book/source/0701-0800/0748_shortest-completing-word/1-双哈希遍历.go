package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(shortestCompletingWord("1s3 PSt", []string{"step", "steps", "stripe", "stepple"}))
	// fmt.Println(shortestCompletingWord("1s3 456", []string{"looks","pest","stew","show"}))
	fmt.Println(shortestCompletingWord("AN87005", []string{"participant", "individual", "start", "exist", "above", "already", "easy", "attack", "player", "important"}))
}

func shortestCompletingWord(licensePlate string, words []string) string {
	m := make(map[byte]int)
	licensePlate = strings.ToLower(licensePlate)
	for i := 0; i < len(licensePlate); i++ {
		if licensePlate[i] >= 'a' && licensePlate[i] <= 'z' {
			m[licensePlate[i]]++
		}
	}
	res := ""
	for i := 0; i < len(words); i++ {
		if len(words[i]) >= len(res) && res != "" {
			continue
		}
		tempM := make(map[byte]int)
		for j := 0; j < len(words[i]); j++ {
			tempM[words[i][j]]++
		}
		flag := true
		for k := range m {
			if tempM[k] < m[k] {
				flag = false
				break
			}
		}
		if flag == true {
			res = words[i]
		}
	}
	return res
}
