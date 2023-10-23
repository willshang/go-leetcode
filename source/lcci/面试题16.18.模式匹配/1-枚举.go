package main

import "fmt"

func main() {
	fmt.Println(patternMatching("bbb", "xxxxxx"))
	fmt.Println(patternMatching("aaa", "dogcatcatdog"))
}

// 程序员面试金典16.18_模式匹配
func patternMatching(pattern string, value string) bool {
	countA := 0
	for i := 0; i < len(pattern); i++ {
		if pattern[i] == 'a' {
			countA++
		}
	}
	countB := len(pattern) - countA
	if value == "" {
		return countA == 0 || countB == 0
	}
	if countA < countB { // 令a>b
		countA, countB = countB, countA
		str := ""
		for i := 0; i < len(pattern); i++ {
			if pattern[i] == 'a' {
				str = str + "b"
			} else {
				str = str + "a"
			}
		}
		pattern = str
	}
	for a := 0; a <= len(value)/countA; a++ { // 枚举
		if judge(pattern, value, a, countA, countB) {
			return true
		}
	}
	return false
}

func judge(pattern string, value string, a, countA, countB int) bool {
	left := len(value) - a*countA
	if (countB == 0 && left == 0) || (countB > 0 && left%countB == 0) {
		var b int
		if countB > 0 {
			b = left / countB
		}
		var strA, strB string
		index := 0
		flag := true
		for i := 0; i < len(pattern); i++ {
			if pattern[i] == 'a' {
				str := value[index : index+a]
				if strA == "" {
					strA = str
				} else if str != strA {
					flag = false
					break
				}
				index = index + a
			} else {
				str := value[index : index+b]
				if strB == "" {
					strB = str
				} else if str != strB {
					flag = false
					break
				}
				index = index + b
			}
		}
		if flag == true && strA != strB {
			return true
		}
	}
	return false
}
