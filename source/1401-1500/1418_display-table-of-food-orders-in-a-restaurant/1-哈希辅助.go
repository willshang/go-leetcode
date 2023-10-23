package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	arr := [][]string{
		{"David", "3", "Ceviche"},
		{"Corina", "10", "Beef Burrito"},
		{"David", "3", "Fried Chicken"},
		{"Carla", "5", "Water"},
		{"Carla", "5", "Ceviche"},
		{"Rous", "3", "Ceviche"},
	}
	fmt.Println(displayTable(arr))
}

// leetcode1418_点菜展示表
func displayTable(orders [][]string) [][]string {
	res := make([][]string, 0)
	titles := make([]string, 0)
	idArr := make([]int, 0)
	m := make(map[string]bool)
	m2 := make(map[string]map[string]int)
	for i := 0; i < len(orders); i++ {
		m[orders[i][2]] = true
		if m2[orders[i][1]] == nil {
			m2[orders[i][1]] = make(map[string]int)
		}
		m2[orders[i][1]][orders[i][2]]++
	}
	for k := range m {
		titles = append(titles, k)
	}
	for k := range m2 {
		tableID, _ := strconv.Atoi(k)
		idArr = append(idArr, tableID)
	}
	sort.Strings(titles)
	sort.Ints(idArr)
	res = append(res, append([]string{"Table"}, titles...))
	for i := 0; i < len(idArr); i++ {
		tableStr := strconv.Itoa(idArr[i])
		temp := make([]string, 0)
		temp = append(temp, tableStr)
		for j := 0; j < len(titles); j++ {
			temp = append(temp, strconv.Itoa(m2[tableStr][titles[j]]))
		}
		res = append(res, temp)
	}
	return res
}
