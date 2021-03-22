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

// leetcode1604_警告一小时内使用相同员工卡大于等于三次的人
func alertNames(keyName []string, keyTime []string) []string {
	m := make(map[string][]int)
	for i := 0; i < len(keyName); i++ {
		m[keyName[i]] = append(m[keyName[i]], strToInt(keyTime[i]))
	}
	res := make([]string, 0)
	for k, v := range m {
		sort.Ints(v)
		for i := 0; i < len(v)-2; i++ {
			if v[i+2]-v[i] <= 60 {
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
