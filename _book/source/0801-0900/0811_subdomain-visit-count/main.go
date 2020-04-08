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
func subdomainVisits(cpdomains []string) []string {
	m := make(map[string]int, len(cpdomains))

	for _, domain := range cpdomains {
		d, n := parse(domain)
		isNew := true
		for isNew {
			m[d] = m[d] + n
			d, isNew = cut(d)
		}
	}
	return GetResult(m)
}

func parse(s string) (string, int) {
	ss := strings.Split(s, " ")
	n, _ := strconv.Atoi(ss[0])
	return ss[1], n
}
func cut(s string) (string, bool) {
	index := strings.Index(s, ".")
	if index == -1 {
		return "", false
	}
	return s[index+1:], true
}

func GetResult(m map[string]int) []string {
	res := make([]string, 0, len(m))
	for k, v := range m {
		res = append(res, fmt.Sprintf("%d %s", v, k))
	}
	return res
}
