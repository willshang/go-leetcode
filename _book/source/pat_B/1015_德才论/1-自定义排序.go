package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Student struct {
	id      int
	d, c, t int
	sortNum int
}

type Students []Student

func (s Students) Len() int      { return len(s) }
func (s Students) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Students) Less(i, j int) bool {
	if s[i].sortNum == s[j].sortNum {
		if s[i].t == s[j].t {
			if s[i].d == s[j].d {
				return s[i].id < s[j].id
			}
			return s[i].d > s[j].d
		}
		return s[i].t > s[j].t
	}
	return s[i].sortNum > s[j].sortNum
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var N, L, H int
	_, _ = fmt.Scanf("%d %d %d", &N, &L, &H)
	arr := make(Students, N)
	count := 0
	for i := 0; i < N; i++ {
		var id int
		var d, c int
		str, _ := reader.ReadString('\n')
		strArray := strings.Fields(str)
		id, _ = strconv.Atoi(strArray[0])
		d, _ = strconv.Atoi(strArray[1])
		c, _ = strconv.Atoi(strArray[2])
		if c < L || d < L {
			continue
		}
		total := d + c
		arr[i].t = total
		arr[i].id = id
		arr[i].d = d
		arr[i].c = c
		sortNum := 0
		if c >= H && d >= H {
			sortNum = 4
		} else if d >= H && c < H {
			sortNum = 3
		} else if d < H && c < H && d >= c {
			sortNum = 2
		} else {
			sortNum = 1
		}
		arr[i].sortNum = sortNum
		count++
	}
	fmt.Println(count)
	sort.Sort(arr)
	for i := 0; i < count; i++ {
		fmt.Printf("%d %d %d\n", arr[i].id, arr[i].d, arr[i].c)
	}
}
