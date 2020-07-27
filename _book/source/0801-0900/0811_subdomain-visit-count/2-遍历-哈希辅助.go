package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	arr := []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}
	fmt.Println(subdomainVisits(arr))
}

// leetcode811_子域名访问计数
func subdomainVisits(cpdomains []string) []string {
	m := make(map[string]int)
	for _, domains := range cpdomains {
		arr := strings.Split(domains, " ")
		count, _ := strconv.Atoi(arr[0])
		tempArr := getSubdomains(arr[1])
		for i := 0; i < len(tempArr); i++ {
			m[tempArr[i]] += count
		}
	}
	res := make([]string, 0)
	for k, v := range m {
		res = append(res, strconv.Itoa(v)+" "+k)
	}
	return res
}

func getSubdomains(s string) []string {
	res := make([]string, 0)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			res = append(res, s[i+1:])
		}
	}
	res = append(res, s)
	return res
}
