package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	var end string
	fmt.Scan(&end)
	res := getResult("2015-10-18", end)
	fmt.Println(res)
}

func getResult(start, end string) int {
	arr := strings.Split(end, "-") // 转换格式 2016-1-1 => 2016-01-01
	end = fmt.Sprintf("%04d-%02d-%02d", transfer(arr[0]), transfer(arr[1]), transfer(arr[2]))
	startTime, _ := time.Parse("2006-01-02", start)
	endTime, _ := time.Parse("2006-01-02", end)
	return int(endTime.Unix()-startTime.Unix()) / (24 * 60 * 60)
}

func transfer(s string) int {
	res, _ := strconv.Atoi(s)
	return res
}
