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
	m := make(map[string]int)
	for _, domains := range cpdomains {
		domain, count := parse(domains)
		isNew := true
		for isNew {
			m[domain] = m[domain] + count
			domain, isNew = cut(domain)
		}
	}
	return getResult(m)
}

func parse(s string) (string, int) {
	ss := strings.Split(s, " ")
	count, _ := strconv.Atoi(ss[0])
	return ss[1], count
}

func cut(s string) (string, bool) {
	index := strings.Index(s, ".")
	if index == -1 {
		return "", false
	}
	return s[index+1:], true
}

func getResult(m map[string]int) []string {
	res := make([]string, 0, len(m))
	for k, v := range m {
		res = append(res, fmt.Sprintf("%d %s", v, k))
	}
	return res
}
