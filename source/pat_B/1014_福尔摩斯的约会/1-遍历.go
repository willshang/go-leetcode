package main

import (
	"bufio"
	"fmt"
	"os"
)

var week = []string{
	"MON",
	"TUE",
	"WED",
	"THU",
	"FRI",
	"SAT",
	"SUN",
}

func main() {
	var a, b, c, d string
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	a = string(data)
	data, _, _ = reader.ReadLine()
	b = string(data)
	data, _, _ = reader.ReadLine()
	c = string(data)
	data, _, _ = reader.ReadLine()
	d = string(data)
	len0 := len(a)
	len2 := len(c)
	if len(b) > len0 {
		len0 = len(b)
	}
	if len(d) > len2 {
		len2 = len(d)
	}
	flag := true
	for i := 0; i < len0; i++ {
		if a[i] == b[i] {
			if flag == true {
				if a[i] >= 'A' && a[i] <= 'G' {
					fmt.Printf("%s ", week[a[i]-'A'])
					flag = false
				}
			} else {
				if a[i] >= 'A' && a[i] <= 'N' {
					fmt.Printf("%02d:", int(a[i]-'A')+10)
					break
				} else if a[i] >= '0' && a[i] <= '9' {
					fmt.Printf("%02d:", int(a[i]-'0'))
					break
				}
			}
		}
	}
	for i := 0; i < len2; i++ {
		if c[i] == d[i] && ((c[i] >= 'A' && c[i] <= 'Z') || (c[i] >= 'a' && c[i] <= 'z')) {
			fmt.Printf("%02d", i)
			break
		}
	}
}
