package main

import (
	"strconv"
	"strings"
)

func main() {

}

// leetcode1169_查询无效交易
type Node struct {
	Name   string
	Time   int
	Amount int
	City   string
}

func invalidTransactions(transactions []string) []string {
	res := make([]string, 0)
	arr := make([]Node, 0)
	n := len(transactions)
	for i := 0; i < n; i++ {
		temp := strings.Split(transactions[i], ",")
		name, city := temp[0], temp[3]
		t, _ := strconv.Atoi(temp[1])
		amount, _ := strconv.Atoi(temp[2])
		arr = append(arr, Node{Name: name, Time: t, Amount: amount, City: city})
	}
	for i := 0; i < n; i++ {
		if arr[i].Amount > 1000 {
			res = append(res, transactions[i])
			continue
		}
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if arr[i].Name == arr[j].Name && arr[i].City != arr[j].City &&
				abs(arr[i].Time-arr[j].Time) <= 60 {
				res = append(res, transactions[i])
				break
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
