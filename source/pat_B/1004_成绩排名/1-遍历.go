package main

import "fmt"

type Student struct {
	Name  string
	Num   string
	Grade int
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	m := make(map[string]Student)
	var max, low int
	var maxName, lowName string
	for i := 0; i < n; i++ {
		var s Student
		fmt.Scanf("%s %s %d", &s.Name, &s.Num, &s.Grade)
		m[s.Name] = s
		if i == 0 {
			max = s.Grade
			maxName = s.Name
			low = s.Grade
			lowName = s.Name
		} else {
			if s.Grade > max {
				max = s.Grade
				maxName = s.Name
			}
			if s.Grade < low {
				low = s.Grade
				lowName = s.Name
			}
		}
	}
	fmt.Println(m[maxName].Name, m[maxName].Num)
	fmt.Println(m[lowName].Name, m[lowName].Num)
}
