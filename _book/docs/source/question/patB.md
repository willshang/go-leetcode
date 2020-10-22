# PAT (Basic Level) Practice 乙级

* [PAT (Basic Level) Practice 乙级](#pat-basic-level-practice-%E4%B9%99%E7%BA%A7)
  * [1001\.害死人不偿命的(3n\+1)猜想(1)](#1001%E5%AE%B3%E6%AD%BB%E4%BA%BA%E4%B8%8D%E5%81%BF%E5%91%BD%E7%9A%843n1%E7%8C%9C%E6%83%B31)
  * [1002\.写出这个数(1)](#1002%E5%86%99%E5%87%BA%E8%BF%99%E4%B8%AA%E6%95%B01)
  * [1003\.我要通过\!(1)](#1003%E6%88%91%E8%A6%81%E9%80%9A%E8%BF%871)
  * [1004\.成绩排名(1)](#1004%E6%88%90%E7%BB%A9%E6%8E%92%E5%90%8D1)
  * [1005\.继续(3n\+1)猜想(1)](#1005%E7%BB%A7%E7%BB%AD3n1%E7%8C%9C%E6%83%B31)
  * [1006\.换个格式输出整数(1)](#1006%E6%8D%A2%E4%B8%AA%E6%A0%BC%E5%BC%8F%E8%BE%93%E5%87%BA%E6%95%B4%E6%95%B01)
  * [1007\.素数对猜想(1)](#1007%E7%B4%A0%E6%95%B0%E5%AF%B9%E7%8C%9C%E6%83%B31)
  * [1008\.数组元素循环右移问题(1)](#1008%E6%95%B0%E7%BB%84%E5%85%83%E7%B4%A0%E5%BE%AA%E7%8E%AF%E5%8F%B3%E7%A7%BB%E9%97%AE%E9%A2%981)
  * [1009\.说反话(1)](#1009%E8%AF%B4%E5%8F%8D%E8%AF%9D1)
  * [1010\.一元多项式求导(1)](#1010%E4%B8%80%E5%85%83%E5%A4%9A%E9%A1%B9%E5%BC%8F%E6%B1%82%E5%AF%BC1)
  * [1011\.A\+B 和 C(1)](#1011ab-%E5%92%8C-c1)
  * [1012\.数字分类(1)](#1012%E6%95%B0%E5%AD%97%E5%88%86%E7%B1%BB1)
  * [1013\.数素数(1)](#1013%E6%95%B0%E7%B4%A0%E6%95%B01)
  * [1014\.福尔摩斯的约会(1)](#1014%E7%A6%8F%E5%B0%94%E6%91%A9%E6%96%AF%E7%9A%84%E7%BA%A6%E4%BC%9A1)
  * [1015\.德才论(1)](#1015%E5%BE%B7%E6%89%8D%E8%AE%BA1)

## 1001.害死人不偿命的(3n+1)猜想(1)

- 题目

```
卡拉兹(Callatz)猜想：
对任何一个正整数 n，如果它是偶数，那么把它砍掉一半；如果它是奇数，那么把 (3n+1) 砍掉一半。
这样一直反复砍下去，最后一定在某一步得到 n=1。卡拉兹在 1950 年的世界数学家大会上公布了这个猜想，
传说当时耶鲁大学师生齐动员，拼命想证明这个貌似很傻很天真的命题，结果闹得学生们无心学业，
一心只证 (3n+1)，以至于有人说这是一个阴谋，卡拉兹是在蓄意延缓美国数学界教学与科研的进展……
我们今天的题目不是证明卡拉兹猜想，而是对给定的任一不超过 1000 的正整数 n，
简单地数一下，需要多少步（砍几下）才能得到 n=1？
输入格式：每个测试输入包含 1 个测试用例，即给出正整数 n 的值。
输出格式：输出从 n 计算到 1 需要的步数。
输入样例：3
输出样例：5
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 遍历模拟 | O(log(n))  | O(1)       |

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d ", &n)
	var res = 0
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
			res++
		} else {
			n = (3*n + 1) / 2
			res++
		}
	}
	fmt.Println(res)
}
```

## 1002.写出这个数(1)

- 题目

```
读入一个正整数 n，计算其各位数字之和，用汉语拼音写出和的每一位数字。
输入格式：
每个测试输入包含 1 个测试用例，即给出自然数 n 的值。这里保证 n 小于 10^100。
输出格式：
在一行内输出 n 的各位数字之和的每一位，拼音数字间有 1 空格，但一行中最后一个拼音数字后没有空格。
输入样例：1234567890987654321123456789
输出样例：yi san wu
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 遍历 | O(log(n))  | O(1)       |

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	m := map[uint8]string{
		'0': "ling",
		'1': "yi",
		'2': "er",
		'3': "san",
		'4': "si",
		'5': "wu",
		'6': "liu",
		'7': "qi",
		'8': "ba",
		'9': "jiu",
	}
	fmt.Scanf("%s", &str)
	sum := 0
	for k := range str {
		sum = sum + int(str[k]-'0')
	}

	toString := strconv.Itoa(sum)
	for k := range toString {
		if k != 0 {
			fmt.Print(" ")
		}
		fmt.Print(m[toString[k]])
	}
}
```

## 1003.我要通过!(1)

- 题目

```
“答案正确”是自动判题系统给出的最令人欢喜的回复。
本题属于 PAT 的“答案正确”大派送 —— 只要读入的字符串满足下列条件，
系统就输出“答案正确”，否则输出“答案错误”。
得到“答案正确”的条件是：
    字符串中必须仅有 P、 A、 T这三种字符，不可以包含其它字符；
    任意形如 xPATx 的字符串都可以获得“答案正确”，其中 x 或者是空字符串，或者是仅由字母 A 组成的字符串；
    如果 aPbTc 是正确的，那么 aPbATca 也是正确的，其中 a、 b、 c 均或者是空字符串，
    或者是仅由字母 A 组成的字符串。
现在就请你为 PAT 写一个自动裁判程序，判定哪些字符串是可以获得“答案正确”的。
输入格式：
每个测试输入包含 1 个测试用例。第 1 行给出一个正整数 n (<10)，是需要检测的字符串个数。
接下来每个字符串占一行，字符串长度不超过 100，且不包含空格。
输出格式：每个字符串的检测结果占一行，如果该字符串可以获得“答案正确”，则输出 YES，否则输出 NO。
输入样例：
8
PAT
PAAT
AAPATAA
AAPAATAAAA
xPATx
PT
Whatever
APAAATAA
输出样例：
YES
YES
YES
YES
NO
NO
NO
NO
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 哈希辅助 | O(n)       | O(n)       |

```go
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
```

## 1004.成绩排名(1)

- 题目

```
读入 n（>0）名学生的姓名、学号、成绩，分别输出成绩最高和成绩最低学生的姓名和学号。
输入格式：每个测试输入包含 1 个测试用例，格式为
第 1 行：正整数 n
第 2 行：第 1 个学生的姓名 学号 成绩
第 3 行：第 2 个学生的姓名 学号 成绩
  ... ... ...
第 n+1 行：第 n 个学生的姓名 学号 成绩
其中姓名和学号均为不超过 10 个字符的字符串，成绩为 0 到 100 之间的一个整数，
这里保证在一组测试用例中没有两个学生的成绩是相同的。
输出格式：
对每个测试用例输出 2 行，
第 1 行是成绩最高学生的姓名和学号，第 2 行是成绩最低学生的姓名和学号，字符串间有 1 空格。
输入样例：
3
Joe Math990112 89
Mike CS991301 100
Mary EE990830 95

输出样例：
Mike CS991301
Joe Math990112
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 哈希辅助 | O(n)       | O(n)       |

```go
package main

import "fmt"

type Student struct {
	Name  string
	Num   string
	Grade int
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	m := make(map[string]Student)
	var max, low int
	var maxName, lowName string
	for i := 0; i < n; i++ {
		var s Student
		fmt.Scanf("%s %s %d", &s.Name, &s.Num, &s.Grade)
		m[s.Name] = s
		if i == 0 {
			max = s.Grade
			maxName = s.Name
			low = s.Grade
			lowName = s.Name
		} else {
			if s.Grade > max {
				max = s.Grade
				maxName = s.Name
			}
			if s.Grade < low {
				low = s.Grade
				lowName = s.Name
			}
		}
	}
	fmt.Println(m[maxName].Name, m[maxName].Num)
	fmt.Println(m[lowName].Name, m[lowName].Num)
}
```

## 1005.继续(3n+1)猜想(1)

- 题目

```
卡拉兹(Callatz)猜想已经在1001中给出了描述。在这个题目里，情况稍微有些复杂。
当我们验证卡拉兹猜想的时候，为了避免重复计算，可以记录下递推过程中遇到的每一个数。
例如对 n=3 进行验证的时候，我们需要计算 3、5、8、4、2、1，则当我们对 n=5、8、4、2 进行验证的时候，
就可以直接判定卡拉兹猜想的真伪，而不需要重复计算，因为这 4 个数已经在验证3的时候遇到过了，
我们称 5、8、4、2 是被 3“覆盖”的数。我们称一个数列中的某个数 n 为“关键数”，
如果 n 不能被数列中的其他数字所覆盖。
现在给定一系列待验证的数字，我们只需要验证其中的几个关键数，就可以不必再重复验证余下的数字。
你的任务就是找出这些关键数字，并按从大到小的顺序输出它们。
输入格式：
每个测试输入包含 1 个测试用例，第 1 行给出一个正整数 K (<100)，
第 2 行给出 K 个互不相同的待验证的正整数 n (1<n≤100)的值，数字间用空格隔开。
输出格式：
每个测试用例的输出占一行，按从大到小的顺序输出关键数字。数字间用 1 个空格隔开，
但一行中最后一个数字后没有空格。
输入样例：
6
3 5 6 7 8 11
输出样例：
7 6
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 排序遍历 | O(nlog(n)) | O(n)       |

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	arr := make([]int, 10001)
	maps := make(map[int]int, n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scanf("%d", &num)
		maps[num] = num // map保存输入的值
		if num == 1 {
			continue
		}
		for {
			if num%2 != 0 {
				num = 3*num + 1
			}
			num = num / 2
			if num == 1 {
				break
			}
			//fmt.Println("num: 出现:\t",num,arr[num] == 1)
			arr[num] = 1
		}
	}
	tempArr := make([]int, 0)
	for k := range maps {
		tempArr = append(tempArr, maps[k])
	}
	sort.Ints(tempArr)
	var flag = false
	for i := len(tempArr) - 1; i >= 0; i-- {
		if arr[tempArr[i]] == 0 {
			if flag == true {
				fmt.Print(" ")
			}
			fmt.Print(tempArr[i])
			flag = true
		}
	}
}
```

## 1006.换个格式输出整数(1)

- 题目

```
让我们用字母 B 来表示“百”、字母 S 表示“十”，用 12...n 来表示不为零的个位数字 n（<10），
换个格式来输出任一个不超过 3 位的正整数。例如 234 应该被输出为 BBSSS1234，
因为它有 2 个“百”、3 个“十”、以及个位的 4。
输入格式：每个测试输入包含 1 个测试用例，给出正整数 n（<1000）。
输出格式：每个测试用例的输出占一行，用规定的格式输出 n。
输入样例 1：234
输出样例 1：BBSSS1234
输入样例 2：23
输出样例 2：SS123
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 遍历 | O(log(n))  | O(1)       |

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	arr := make([]int, 3)
	i := 0
	for {
		if n == 0 {
			break
		}
		arr[i] = n % 10
		n = n / 10
		i++
	}
	for k := 0; k < arr[2]; k++ {
		fmt.Print("B")
	}
	for k := 0; k < arr[1]; k++ {
		fmt.Print("S")
	}
	for k := 0; k < arr[0]; k++ {
		fmt.Print(k + 1)
	}
}
```

## 1007.素数对猜想(1)

- 题目

```
让我们定义dn为：dn=pn+1−pn，其中pi是第i个素数。
显然有d1=1，且对于n>1有dn是偶数。“素数对猜想”认为“存在无穷多对相邻且差为2的素数”。

现给定任意正整数N(<10^5)，请计算不超过N的满足猜想的素数对的个数。
输入格式: 输入在一行给出正整数N。
输出格式:在一行中输出不超过N的满足猜想的素数对的个数。
输入样例:20
输出样例:4
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 遍历 | O(n^3/2)   | O(1)       |

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	count := 0
	// 2 3 5 7
	for i := 5; i <= n; i = i + 2 {
		if isPrime(i) && isPrime(i-2) {
			count++
		}
	}
	fmt.Println(count)
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
```

## 1008.数组元素循环右移问题(1)

- 题目

```
一个数组A中存有N（>0）个整数，在不允许使用另外数组的前提下，将每个整数循环向右移M（≥0）个位置，
即将A中的数据由（A0A1⋯AN−1）变换为（AN−M⋯AN−1A0A1⋯AN−M−1）
（最后M个数循环移至最前面的M个位置）。
如果需要考虑程序移动数据的次数尽量少，要如何设计移动的方法？
输入格式:
每个输入包含一个测试用例，第1行输入N（1≤N≤100）和M（≥0）；第2行输入N个整数，之间用空格分隔。
输出格式:在一行中输出循环右移M位以后的整数序列，之间用空格分隔，序列结尾不能有多余空格。
输入样例:
6 2
1 2 3 4 5 6
输出样例:5 6 1 2 3 4
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 遍历交换 | O(n)       | O(n)       |

```go
package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scanf("%d %d", &N, &M)

	M = M % N
	arr := make([]int, 0)
	for i := 0; i < N; i++ {
		var num int
		fmt.Scanf("%d", &num)
		arr = append(arr, num)
	}
	if M != 0 {
		reverse(arr)
		reverse(arr[:M])
		reverse(arr[M:])
	}
	for i := 0; i < len(arr)-1; i++ {
		fmt.Print(arr[i])
		fmt.Print(" ")
	}
	fmt.Print(arr[len(arr)-1])
}

func reverse(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}
```

## 1009.说反话(1)

- 题目

```
给定一句英语，要求你编写程序，将句中所有单词的顺序颠倒输出。
输入格式：
测试输入包含一个测试用例，在一行内给出总长度不超过 80 的字符串。字符串由若干单词和若干空格组成，
其中单词是由英文字母（大小写有区分）组成的字符串，单词之间用 1 个空格分开，输入保证句子末尾没有多余的空格。
输出格式：每个测试用例的输出占一行，输出倒序后的句子。
输入样例：Hello World Here I Come
输出样例：Come I Here World Hello
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 遍历 | O(n)       | O(n)       |

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var str string
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	str = string(data)
	arr := strings.Split(str, " ")
	for k := len(arr) - 1; k >= 0; k-- {
		fmt.Print(arr[k])
		if k != 0 {
			fmt.Print(" ")
		}
	}
}
```

## 1010.一元多项式求导(1)

- 题目

```
设计函数求一元多项式的导数。（注：xn（n为整数）的一阶导数为nxn−1。）
输入格式:
以指数递降方式输入多项式非零项系数和指数（绝对值均为不超过 1000 的整数）。数字间以空格分隔。
输出格式:
以与输入相同的格式输出导数多项式非零项的系数和指数。数字间以空格分隔，但结尾不能有多余空格。
注意“零多项式”的指数和系数都是 0，但是表示为 0 0。
输入样例:3 4 -5 2 6 1 -2 0
输出样例:12 3 -10 1 6 0
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 遍历 | O(n)       | O(1)       |

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var str string
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	str = string(data)
	var flag = false
	arr := strings.Fields(str)
	for k := 0; k < len(arr)-1; k = k + 2 {
		a, _ := strconv.Atoi(arr[k])
		b, _ := strconv.Atoi(arr[k+1])
		if a == 0 && b == 0 {
			fmt.Print("0 0")
			flag = true
		}
		if b != 0 {
			if flag == true {
				fmt.Print(" ")
			}
			fmt.Print(a*b, " ", b-1)
			flag = true
		}
	}
	if flag == false {
		fmt.Print("0 0")
	}
}
```

## 1011.A+B 和 C(1)

- 题目

```
给定区间 [−2^31,2^31] 内的 3 个整数 A、B 和 C，请判断 A+B 是否大于 C。
输入格式：
输入第 1 行给出正整数 T (≤10)，是测试用例的个数。随后给出 T 组测试用例，每组占一行，
顺序给出 A、B 和 C。整数间以空格分隔。
输出格式：
对每组测试用例，在一行中输出 Case #X: true 如果 A+B>C，
否则输出 Case #X: false，其中 X 是测试用例的编号（从 1 开始）。
输入样例：
4
1 2 3
2 3 4
2147483647 0 2147483646
0 -2147483648 -2147483647
输出样例：
Case #1: false
Case #2: true
Case #3: true
Case #4: false
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 计算 | O(1)       | O(1)       |

```go
package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		var a, b, c int
		_, _ = fmt.Scanf("%d %d %d", &a, &b, &c)
		var result int
		if a > b {
			result = a - c + b
		} else {
			result = b - c + a
		}
		if result > 0 {
			fmt.Print("Case #", i+1, ": ", true, "\n")
		} else {
			fmt.Print("Case #", i+1, ": ", false, "\n")
		}
	}
}
```

## 1012.数字分类(1)

- 题目

```
给定一系列正整数，请按要求对数字进行分类，并输出以下 5 个数字：
    A1 = 能被 5 整除的数字中所有偶数的和；
    A2 = 将被 5 除后余 1 的数字按给出顺序进行交错求和，即计算 n1−n2+n3−n4⋯；
    A3 = 被 5 除后余 2 的数字的个数；
    A4 = 被 5 除后余 3 的数字的平均数，精确到小数点后 1 位；
    A5 = 被 5 除后余 4 的数字中最大数字。
输入格式：
每个输入包含 1 个测试用例。每个测试用例先给出一个不超过 1000 的正整数 N，
随后给出 N 个不超过 1000 的待分类的正整数。数字间以空格分隔。
输出格式：
对给定的 N 个正整数，按题目要求计算 A1~A5 并在一行中顺序输出。
数字间以空格分隔，但行末不得有多余空格。
若其中某一类数字不存在，则在相应位置输出 N。
输入样例 1：13 1 2 3 4 5 6 7 8 9 10 20 16 18
输出样例 1：30 11 2 9.7 9
输入样例 2：8 1 2 4 5 6 7 9 16
输出样例 2：N 11 2 N 9
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 遍历 | O(n)       | O(1)       |

```go
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	var A1, A2, A5 = 0, 0, 0
	var A4 = 0
	mapArr := make(map[int][]int)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scanf("%d", &num)
		mapArr[num%5] = append(mapArr[num%5], num)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < len(mapArr[i]); j++ {
			if i == 0 && mapArr[i][j]%2 == 0 {
				A1 += mapArr[i][j]
			}
			if i == 1 && j%2 == 0 {
				A2 += mapArr[i][j]
			}
			if i == 1 && j%2 == 1 {
				A2 -= mapArr[i][j]
			}
			if i == 3 {
				A4 += mapArr[i][j]
			}
			if i == 4 && mapArr[i][j] > A5 {
				A5 = mapArr[i][j]
			}
		}
	}
	for i := 0; i < 5; i++ {
		if i != 0 {
			fmt.Print(" ")
		}
		if i == 0 && A1 == 0 || i != 0 && len(mapArr[i]) == 0 {
			fmt.Print("N")
			continue
		}
		if i == 0 {
			fmt.Print(A1)
		} else if i == 1 {
			fmt.Print(A2)
		} else if i == 2 {
			fmt.Print(len(mapArr[2]))
		} else if i == 3 {
			fmt.Print(fmt.Sprintf("%.1f", float64(A4)/float64(len(mapArr[i]))))
		} else if i == 4 {
			fmt.Print(A5)
		}
	}
}
```

## 1013.数素数(1)

- 题目

```
令 Pi 表示第 i 个素数。现任给两个正整数 M≤N≤10^4，请输出 PM 到 PN 的所有素数。
输入格式：输入在一行中给出 M 和 N，其间以空格分隔。
输出格式：输出从 PM 到 PN 的所有素数，每 10 个数字占 1 行，其间以空格分隔，但行末不得有多余空格。
输入样例：5 27
输出样例：
11 13 17 19 23 29 31 37 41 43
47 53 59 61 67 71 73 79 83 89
97 101 103
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 遍历 | O(n^3/2)   | O(n)       |

```go
package main

import "fmt"

func main() {
	var N, M int
	var num = 2
	var count int
	fmt.Scanf("%d %d", &M, &N)
	result := make([]int, 0)
	for {
		if count < N {
			if isPrime(num) {
				count++
				if count >= M {
					result = append(result, num)
				}
			}
			num++
		} else {
			break
		}
	}

	for i := 0; i < len(result); i++ {
		if i%10 != 0 {
			fmt.Print(" ")
		}
		fmt.Print(result[i])
		if i%10 == 9 {
			fmt.Println()
		}
	}
}

func isPrime(a int) bool {
	for i := 2; i*i <= a; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}
```

## 1014.福尔摩斯的约会(1)

- 题目

```
大侦探福尔摩斯接到一张奇怪的字条：我们约会吧！ 
3485djDkxh4hhGE 2984akDfkkkkggEdsb s&hgsfdk d&Hyscvnm。
大侦探很快就明白了，字条上奇怪的乱码实际上就是约会的时间星期四 14:04，
因为前面两字符串中第 1 对相同的大写英文字母（大小写有区分）是第 4 个字母 D，代表星期四；
第 2 对相同的字符是 E ，那是第 5 个英文字母，
代表一天里的第 14 个钟头（于是一天的 0 点到 23 点由数字 0 到 9、以及大写字母 A 到 N 表示）；
后面两字符串第 1 对相同的英文字母 s 出现在第 4 个位置（从 0 开始计数）上，代表第 4 分钟。
现给定两对字符串，请帮助福尔摩斯解码得到约会的时间。
输入格式：输入在 4 行中分别给出 4 个非空、不包含空格、且长度不超过 60 的字符串。
输出格式：
在一行中输出约会的时间，格式为 DAY HH:MM，其中 DAY 是某星期的 3 字符缩写，即 MON 表示星期一，
TUE 表示星期二，WED 表示星期三，THU 表示星期四，FRI 表示星期五，SAT 表示星期六，SUN 表示星期日。
题目输入保证每个测试存在唯一解。
输入样例：
3485djDkxh4hhGE 
2984akDfkkkkggEdsb 
s&hgsfdk 
d&Hyscvnm
输出样例：THU 14:04
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 遍历 | O(n)       | O(1)       |

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

var week = []string{
	"MON",
	"TUE",
	"WED",
	"THU",
	"FRI",
	"SAT",
	"SUN",
}

func main() {
	var a, b, c, d string
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	a = string(data)
	data, _, _ = reader.ReadLine()
	b = string(data)
	data, _, _ = reader.ReadLine()
	c = string(data)
	data, _, _ = reader.ReadLine()
	d = string(data)
	len0 := len(a)
	len2 := len(c)
	if len(b) > len0 {
		len0 = len(b)
	}
	if len(d) > len2 {
		len2 = len(d)
	}
	flag := true
	for i := 0; i < len0; i++ {
		if a[i] == b[i] {
			if flag == true {
				if a[i] >= 'A' && a[i] <= 'G' {
					fmt.Printf("%s ", week[a[i]-'A'])
					flag = false
				}
			} else {
				if a[i] >= 'A' && a[i] <= 'N' {
					fmt.Printf("%02d:", int(a[i]-'A')+10)
					break
				} else if a[i] >= '0' && a[i] <= '9' {
					fmt.Printf("%02d:", int(a[i]-'0'))
					break
				}
			}
		}
	}
	for i := 0; i < len2; i++ {
		if c[i] == d[i] && ((c[i] >= 'A' && c[i] <= 'Z') || (c[i] >= 'a' && c[i] <= 'z')) {
			fmt.Printf("%02d", i)
			break
		}
	}
}
```

## 1015.德才论(1)

- 题目

```
宋代史学家司马光在《资治通鉴》中有一段著名的“德才论”：
“是故才德全尽谓之圣人，才德兼亡谓之愚人，德胜才谓之君子，才胜德谓之小人。凡取人之术，苟不得圣人，
君子而与之，与其得小人，不若得愚人。”
现给出一批考生的德才分数，请根据司马光的理论给出录取排名。
输入格式：
输入第一行给出 3 个正整数，分别为：
N（≤10^5），即考生总数；L（≥60），为录取最低分数线，即德分和才分均不低于 L 的考生才有资格被考虑录取；
H（<100），为优先录取线——德分和才分均不低于此线的被定义为“才德全尽”，此类考生按德才总分从高到低排序；
才分不到但德分到线的一类考生属于“德胜才”，也按总分排序，但排在第一类考生之后；
德才分均低于 H，但是德分不低于才分的考生属于“才德兼亡”但尚有“德胜才”者，按总分排序，但排在第二类考生之后；
其他达到最低线 L 的考生也按总分排序，但排在第三类考生之后。

随后 N 行，每行给出一位考生的信息，
包括：准考证号 德分 才分，其中准考证号为 8 位整数，德才分为区间 [0, 100] 内的整数。数字间以空格分隔。
输出格式：
输出第一行首先给出达到最低分数线的考生人数 M，随后 M 行，每行按照输入格式输出一位考生的信息，
考生按输入中说明的规则从高到低排序。
当某类考生中有多人总分相同时，按其德分降序排列；若德分也并列，则按准考证号的升序输出。
输入样例：
14 60 80
10000001 64 90
10000002 90 60
10000011 85 80
10000003 85 80
10000004 80 85
10000005 82 77
10000006 83 76
10000007 90 78
10000008 75 79
10000009 59 90
10000010 88 45
10000012 80 100
10000013 90 99
10000014 66 60
输出样例：
12
10000013 90 99
10000012 80 100
10000003 85 80
10000011 85 80
10000004 80 85
10000007 90 78
10000006 83 76
10000005 82 77
10000002 90 60
10000014 66 60
10000008 75 79
10000001 64 90
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 排序 | O(nlog(n)) | O(n)       |

```go
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
```

