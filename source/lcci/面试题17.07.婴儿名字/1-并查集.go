package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(trulyMostPopular([]string{"John(15)", "Jon(12)", "Chris(13)", "Kris(4)", "Christopher(19)"},
		[]string{"(Jon,John)", "(John,Johnny)", "(Chris,Kris)", "(Chris,Christopher)"}))
}

// 程序员面试金典17.07_婴儿名字
func trulyMostPopular(names []string, synonyms []string) []string {
	res := make([]string, 0)
	node = Node{}
	nameArr := make([]string, 0)
	countArr := make([]int, 0)
	m := make(map[string]int)
	for i := 0; i < len(names); i++ {
		arr := strings.Split(names[i], "(")
		nameArr = append(nameArr, arr[0])
		tempArr := strings.Split(arr[1], ")")
		count, _ := strconv.Atoi(tempArr[0])
		countArr = append(countArr, count)
		m[arr[0]] = i
	}
	Init(nameArr, countArr)

	for i := 0; i < len(synonyms); i++ {
		str := strings.TrimLeft(synonyms[i], "(")
		str = strings.TrimRight(str, ")")
		arr := strings.Split(str, ",")
		a := m[arr[0]]
		b := m[arr[1]]
		union(a, b)
	}
	for i := 0; i < len(node.fa); i++ {
		if node.fa[i] < 0 {
			temp := node.names[i] + "(" + strconv.Itoa(node.count[i]) + ")"
			res = append(res, temp)
		}
	}
	return res
}

var node Node

type Node struct {
	fa    []int
	names []string
	count []int
}

// 初始化
func Init(names []string, count []int) {
	node.fa = make([]int, len(names))
	for i := 0; i < len(names); i++ {
		node.fa[i] = -1
	}
	node.names = names
	node.count = count
}

// 查询
func find(x int) int {
	if node.fa[x] < 0 {
		return x
	}
	res := find(node.fa[x])
	node.fa[x] = res
	return res
}

// 合并
func union(i, j int) {
	x, y := find(i), find(j)
	if x == y {
		return
	}
	if node.names[x] <= node.names[y] {
		node.fa[y] = x
		node.count[x] = node.count[x] + node.count[y]
	} else {
		node.fa[x] = y
		node.count[y] = node.count[y] + node.count[x]
	}
}
