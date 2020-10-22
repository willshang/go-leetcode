package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		var str string
		fmt.Scanf("%s", &str)
		fmt.Println(judge(str))
	}
}

func judge(str string) string {
	length := len(str)
	left, right := 0, 0
	charMap := make(map[uint8]int)
	for i := 0; i < length; i++ {
		charMap[str[i]]++
		if str[i] == 'P' {
			left = i
		}
		if str[i] == 'T' {
			right = i
		}
	}
	l := left
	m := right - left - 1
	r := length - right - 1
	// 2.任意形如 xPATx 的字符串都可以获得“答案正确”，其中 x 或者是空字符串，或者是仅由字母 A 组成的字符串；
	// 3.如果 aPbTc 是正确的，那么 aPbATca 也是正确的，其中 a、 b、 c 均或者是空字符串，或者是仅由字母 A 组成的字符串。
	// P的个数1， T的个数1， A的个数不为0 只有3种数据 左边 * 中间 = 右边
	// =>就是P和T中间每增加一个A，需要将P之前的内容复制到字符串末尾，得到的新字符串就也是正确的。
	if charMap['P'] == 1 && charMap['T'] == 1 && charMap['A'] != 0 && len(charMap) == 3 &&
		right-length != 1 && l*m == r {
		return "YES"
	}
	return "NO"
}
