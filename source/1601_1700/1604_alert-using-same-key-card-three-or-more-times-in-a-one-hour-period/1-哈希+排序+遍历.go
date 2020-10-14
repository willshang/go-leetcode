package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(alertNames([]string{"leslie", "leslie", "clare", "clare", "clare", "clare", "clare"}, []string{
		"13:00", "13:20", "14:00", "18:00", "18:51", "19:30", "19:49",
	}))
}

func alertNames(keyName []string, keyTime []string) []string {
	m := make(map[string][]int)
	for i := 0; i < len(keyName); i++ {
		m[keyName[i]] = append(m[keyName[i]], strToInt(keyTime[i]))
	}
	res := make([]string, 0)
	for k, v := range m {
		sort.Ints(v)
		first := v[0]
		second := v[0]
		count := 1
		for i := 1; i < len(v); i++ {
			if v[i] > first && v[i]-first <= 60 {
				second = v[i]
				count++
			} else {
				first = second
				second = v[i]
				count = 2
			}
			if count >= 3 {
				res = append(res, k)
				break
			}
		}
	}
	sort.Strings(res)
	return res
}

func strToInt(str string) int {
	arr := strings.Split(str, ":")
	hour, _ := strconv.Atoi(arr[0])
	minute, _ := strconv.Atoi(arr[1])
	return hour*60 + minute
}
