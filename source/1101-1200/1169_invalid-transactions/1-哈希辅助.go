package main

import (
	"sort"
	"strconv"
	"strings"
)

func main() {

}

type Node struct {
	Name   string
	Time   int
	Amount int
	City   string
	Index  int
	Flag   bool
}

func invalidTransactions(transactions []string) []string {
	res := make([]string, 0)
	m := make(map[string][]Node)
	n := len(transactions)
	for i := 0; i < n; i++ {
		arr := strings.Split(transactions[i], ",")
		name, city := arr[0], arr[3]
		t, _ := strconv.Atoi(arr[1])
		amount, _ := strconv.Atoi(arr[2])
		m[name] = append(m[name], Node{Name: name, Time: t,
			Amount: amount, City: city, Index: i,
		})
	}
	for _, v := range m {
		sort.Slice(v, func(i, j int) bool {
			return v[i].Time < v[j].Time
		})
		for i := 0; i < len(v); i++ {
			if v[i].Amount > 1000 {
				v[i].Flag = true
			}
			for j := i + 1; j < len(v); j++ {
				if v[j].Time-v[i].Time <= 60 {
					if v[i].City != v[j].City {
						v[i].Flag = true
						v[j].Flag = true
					}
				} else {
					break
				}
			}
		}
	}
	for _, v := range m {
		for i := 0; i < len(v); i++ {
			if v[i].Flag == true {
				res = append(res, transactions[v[i].Index])
			}
		}
	}
	return res
}
