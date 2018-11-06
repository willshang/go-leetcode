package main

import (
	"fmt"
)

func main() {
	fmt.Println(readBinaryWatch(1))
}

func bincount(num int) int  {
	count := []int{}
	for num != 0{
		temp := num % 2
		count = append(count,temp)
		num = num / 2
	}
	count_num := 0
	for i := 0; i < len(count); i++{
		if count[i] == 1{
			count_num++
		}
	}
	return count_num
}

func readBinaryWatch(num int) []string {
	res := []string{}
	for i := 0; i < 12; i++{
		for j := 0; j < 60; j++{
			if bincount(i) + bincount(j) == num{
				res = append(res,fmt.Sprintf("%d:%02d",i,j))
			}
		}
	}
	return res
}


/*func readBinaryWatch(num int) []string {
	res := make([]string, 0, 8)
	leds := make([]bool, 10)

	var dfs  func(int,int)
	dfs = func(idx , n int) {
		var h, m int
		if n == 0{
			m, h = get(leds[:6]), get(leds[6:])
			if h < 12 && m < 60{
				res = append(res, fmt.Sprintf("%d:%02d",h,m))
			}
			return
		}
		for i := idx; i < len(leds)-n+1; i++{
			leds[i] = true
			dfs(i+1, n-1)
			leds[i] = false
		}
	}
	dfs(0,num)

	sort.Strings(res)
	return res
}
var bs = []int{1,2,4,8,16,32}
func get(leds []bool) int  {
	var sum int
	size := len(leds)
	for i := 0; i < size; i++{
		if leds[i]{
			sum += bs[i]
		}
	}
	return sum
}*/