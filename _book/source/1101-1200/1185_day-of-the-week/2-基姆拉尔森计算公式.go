package main

import (
	"fmt"
)

func main() {
	fmt.Println(dayOfTheWeek(31, 8, 2019))
	fmt.Println(dayOfTheWeek(29, 2, 2016))
}

// 蔡勒公式
// 基姆拉尔森计算公式
// https://baike.baidu.com/item/%E8%94%A1%E5%8B%92%E5%85%AC%E5%BC%8F
// https://www.cnblogs.com/SeekHit/p/7498408.html
// Week = (y+y/4-y/100+y/400+2*m+3*(m+1)/5+d) mod 7；
func dayOfTheWeek(day int, month int, year int) string {
	arr := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	if month == 1 || month == 2 {
		month = month + 12
		year--
	}
	week := (year + year/4 - year/100 + year/400 + 2*month + 3*(month+1)/5 + day) % 7
	return arr[week]
}
