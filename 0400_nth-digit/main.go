package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findNthDigit(9999))
	fmt.Println(findNthDigit(13))
	fmt.Println(findNthDigit(11))
}
/**
 * 这里是找第n个数字(这里的数和数字有区别，数字可以理解为将所有数拼合成一个字符串后的第n为对应的数字（0-9)）
 * 这里首先分析一下位数和规律
 * 个位数：1-9，一共9个,共计9个数字
 * 2位数：10-99,一共90个，共计180个数字
 * 3位数：100-999，一共900个，共计270个数字
 * 4位数，1000-9999，一共9000个，共计36000个数字
 * 以此类推，
 * 这样我们就可以首先定位到是哪个数，再找到其对应的数字
 * */
func findNthDigit(n int) int {
	digitType := 1
	digitNum := 9
	for n > digitNum*digitType{
		n = n - digitNum * digitType
		digitType++
		digitNum = digitNum * 10
	}

	indexInSubRange := (n-1) / digitType
	indexInNum := (n-1) % digitType
	num := int(math.Pow10(digitType-1)) + indexInSubRange
	arr := []int{}
	for num > 0{
		arr = append(arr,num%10)
		num = num / 10
	}
	return arr[len(arr)-1-indexInNum]
}