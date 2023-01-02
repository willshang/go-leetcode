# 基本算法

## 0001.股神(1)

- 题目

```
题目描述
有股神吗？
有，小赛就是！
经过严密的计算，小赛买了一支股票，他知道从他买股票的那天开始，股票会有以下变化：
第一天不变，以后涨一天，跌一天，涨两天，跌一天，涨三天，跌一天...依此类推。
为方便计算，假设每次涨和跌皆为1，股票初始单价也为1，请计算买股票的第n天每股股票值多少钱？
输入 输入包括多组数据；每行输入一个n，1<=n<=10^9 。
样例输入
1
2
3
4
5
输出 请输出他每股股票多少钱，对于每组数据，输出一行。
样例输出
1
2
1
2
3
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 遍历 | O(n^1/2)   | O(1)       |

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	res := getResult(n)
	fmt.Println(res)
}

func getResult(n int) int {
	if n <= 2 {
		return n
	}
	count := 0 // 统计下跌多少次
	target := n
	// n-2-3-4-...-
	// 每减去1次，碰到的下跌次数+1
	for i := 2; i <= n; i++ {
		if target-i > 0 {
			target = target - i
			count++
		} else {
			break
		}
	}
	return n - 2*count
}
```

## 0002.翻转数组(1)

- 题目

```
题目描述
给定一个长度为n的整数数组a，元素均不相同，问数组是否存在这样一个片段，
只将该片段翻转就可以使整个数组升序排列。其中数组片段[l,r]表示序列a[l], a[l+1], ..., a[r]。
原始数组为
a[1], a[2], ..., a[l-2], a[l-1], a[l], a[l+1], ..., a[r-1], a[r],
a[r+1], a[r+2], ..., a[n-1], a[n]，
将片段[l,r]反序后的数组是
a[1], a[2], ..., a[l-2], a[l-1], a[r], a[r-1], ..., a[l+1], a[l], 
a[r+1], a[r+2], ..., a[n-1], a[n]。
输入 第一行数据是一个整数：n (1≤n≤105)，表示数组长度。
第二行数据是n个整数a[1], a[2], ..., a[n] (1≤a[i]≤109)。
样例输入
4
2 1 3 4
输出 输出“yes”，如果存在；否则输出“no”，不用输出引号。
样例输出 yes
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
	fmt.Scanf("%d\n", &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		var temp int
		fmt.Scanf("%d", &temp)
		arr[i] = temp
	}
	res := getResult(arr)
	if res == true {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

func getResult(arr []int) bool {
	temp := make([]int, len(arr))
	copy(temp, arr)
	sort.Ints(temp)
	left, right := 0, len(arr)-1
	for left < len(arr) && arr[left] == temp[left] {
		left++
	}
	if left == len(arr) {
		return false
	}
	for 0 <= right && arr[right] == temp[right] {
		right--
	}
	for i := left; i < right; i++ {
		if arr[i] < arr[i+1] {
			return false
		}
	}
	return true
}
```

## 0003.约德尔测试(1)

- 题目

```
题目描述
兰博和提莫闲聊之后，回归到了他们的正题，约德尔人的未来。
说起约德尔人的未来，黑默丁格曾经提出了一个约德尔测试，将约德尔人的历史的每个阶段都用一个字符表达出来。
(包括可写字符,不包括空格。)。然后将这个字符串转化为一个01串。
转化规则是如果这个字符如果是字母或者数字，这个字符变为1,其它变为0。
然后将这个01串和黑默丁格观测星空得到的01串做比较，得到一个相似率。相似率越高,则约德尔的未来越光明。
请问:相似率为多少？
输入 每组输入数据为两行，第一行为有关约德尔人历史的字符串，第二行是黑默丁格观测星空得到的字符串。
(两个字符串的长度相等,字符串长度不小于1且不超过1000。)
样例输入
@!%12dgsa
010111100
输出 输出一行，在这一行输出相似率。用百分数表示。
(相似率为相同字符的个数/总个数,精确到百分号小数点后两位。printf("%%");输出一个%。)
样例输出
66.67%
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 遍历 | O(n)       | O(1)       |

```go
package main

import "fmt"

func main() {
	var a, b string
	fmt.Scanf("%s\n%s", &a, &b)
	res := getResult(a, b)
	fmt.Printf("%.2f%%\n", res)
}

func getResult(a, b string) float64 {
	n := len(a)
	count := 0
	for i := 0; i < n; i++ {
		c := 0
		d := int(b[i] - '0')
		if ('0' <= a[i] && a[i] <= '9') ||
			('a' <= a[i] && a[i] <= 'z') ||
			('A' <= a[i] && a[i] <= 'Z') {
			c = 1
		}
		if c == d {
			count++
		}
	}
	return float64(count) / float64(n) * 100
}
```

## 0004.路灯(1)

- 题目

```
题目描述
V先生有一天工作到很晚，回家的时候要穿过一条长l的笔直的街道，这条街道上有n个路灯。
假设这条街起点为0，终点为l，第i个路灯坐标为ai。
路灯发光能力以正数d来衡量，其中d表示路灯能够照亮的街道上的点与路灯的最远距离，所有路灯发光能力相同。
为了让V先生看清回家的路，路灯必须照亮整条街道，又为了节省电力希望找到最小的d是多少？
输入 输入两行数据，
第一行是两个整数：路灯数目n (1≤n≤1000)，街道长度l (1 ≤l≤109)。
第二行有n个整数ai (0 ≤ ai≤ l)，表示路灯坐标，多个路灯可以在同一个点，也可以安放在终点位置。
样例输入
7 15
15 5 3 7 9 14 0
输出 输出能够照亮整个街道的最小d，保留两位小数。
样例输出
2.50
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
	var n, l int
	fmt.Scanf("%d %d\n", &n, &l)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		var temp int
		fmt.Scanf("%d", &temp)
		arr[i] = temp
	}
	res := getResult(arr, l)
	fmt.Printf("%.2f\n", res)
}

func getResult(arr []int, l int) float64 {
	sort.Ints(arr)
	maxValue := 0
	for i := 1; i < len(arr); i++ {
		maxValue = max(maxValue, arr[i]-arr[i-1])
	}
	left := arr[0]
	right := l - arr[len(arr)-1]
	res := max(maxValue, left*2)
	res = max(res, right*2)
	return float64(res) / float64(2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

## 0005. 计算器的新功能

### 题目

```
题目描述
当你学一些可视化程序设计语言时，老师经常会让你设计并且编程做出一个计算器，
这时也许你会仿照windows系统自带的计算器外观和功能去设计，但是现在老师要你多做出一个有新功能的计算器，
实现当输入一个数时，能够将这个数分解成一个或多个素因子乘积的形式，并按素因子的大小排列显示出来。
大家对计算器中数的表示应该很清楚的。
下面显示出了0 — 9这十个数字的表示形式。每个数字都占据5 * 3大小的字符区域
你能实现这个新功能吗？试试看吧！
输入 输入有多组测试数据，每组包括一个正整数n（1 < n <= 1000000)。
样例输入
10
2
输出 对于每个数，将它分解成若干个素数乘积的形式，并按从小到大的顺序输出，素因子之间用“ * ”的形式连接。
样例输出
 -     -

  |   |

 -  *  -

|       |

 -     -
 -
  |
 -
|
 -
```

### 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 排序遍历 | O(nlog(n)) | O(n)       |

```go

```

## 0006.公交车乘客(1)

- 题目

```
题目描述
一个公交车经过n个站点，乘客从前门上车，从后门下车。
现在统计了在第i个站，下车人数a[i]，以及上车人数b[i]。
问公交车运行时候车上最多有多少乘客
输入 第一行读入一个整数n(1<=n<=100)，表示有n个站点
接下来n行，每行两个数值，分别表示在第i个站点下车人数和上车人数
样例输入
4
0 3
2 5
4 2
4 0
输出 每组输出车上最多的乘客数目
样例输出
6
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 遍历 | O(n)       | O(n)       |

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arrA := make([]int, n)
	arrB := make([]int, n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		arrA[i] = a
		arrB[i] = b
	}
	res := getResult(arrA, arrB)
	fmt.Println(res)
}

func getResult(a, b []int) int {
	sum := 0
	res := 0
	for i := 0; i < len(a); i++ {
		sum = sum + b[i] - a[i]
		if sum > res {
			res = sum
		}
	}
	return res
}
```

## 0007.分苹果(2)

- 题目

```
题目描述
果园里有堆苹果，N（1＜N＜9）只熊来分。第一只熊把这堆苹果平均分为N份，多了一个，
它把多的一个扔了，拿走了一份。第二只熊把剩下的苹果又平均分成N份，又多了一个，
它同样把多的一个扔了，拿走了一份，第三、第四直到第N只熊都是这么做的，问果园里原来最少有多少个苹果？
输入 输入1个整数，表示熊的个数。它的值大于1并且小于9。
样例输入
5
输出 为1个数字，表示果园里原来有的苹果个数。
样例输出
3121
```

- 解题思路

| No.  | 思路   | 时间复杂度 | 空间复杂度 |
| ---- | ------ | ---------- | ---------- |
| 01   | 数学   | O(n)       | O(1)       |
| 02   | 暴力法 | O(n^2)     | O(1)       |

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	res := getResult(n)
	fmt.Println(res)
}

// res = n^n-n+1
func getResult(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res = res * n
	}
	return res - n + 1
}

# 2
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	res := getResult(n)
	fmt.Println(res)
}

func getResult(n int) int {
	for i := 1; ; i++ {
		target := i
		bear := n
		for bear > 0 {
			if (target-1)%n == 0 {
				bear--
				target = target - 1 - (target-1)/n
			} else {
				break
			}
		}
		if bear == 0 {
			return i
		}
	}
	return 0
}
```

## 0008.马路上的路灯(1)

- 题目

```
题目描述
城市E的马路上有很多路灯，每两个相邻路灯之间的间隔都是1公里。
小赛是城市E的领导，为了使E城市更快更好的发展，需要在城市E的一段长度为M的主干道上的一些区域建地铁。
这些区域要是建了地铁，就需要挪走相应的路灯。
可以把长度为M的主干道看成一个数轴，一端在数轴0的位置，另一端在M的位置；数轴上的每个整数点都有一个路灯。
要建地铁的这些区域可以用它们在数轴上的起始点和终止点表示，
已知任一区域的起始点和终止点的坐标都是整数，区域之间可能有重合的部分。
现在要把这些区域中的路灯（包括区域端点处的两个路灯）移走。
你能帮助小赛计算一下，将这些路灯移走后，马路上还有多少路灯？
输入 输入文件的第一行有两个整数M（1 <= M <= 10000）和 N（1 <= N <= 100），M代表马路的长度，
N代表区域的数目，M和N之间用一个空格隔开。
接下来的N行每行包含两个不同的整数，用一个空格隔开，表示一个区域的起始点和终止点的坐标。
所有输入都为整数。且M和N的范围为上面提示范围。
样例输入
500 3
100 200
150 300
360 361
输出 输出文件包括一行，这一行只包含一个整数，表示马路上剩余路灯的数目。
样例输出
298
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 遍历 | O(n)       | O(n)       |

```go
package main

import "fmt"

func main() {
	var M, N int
	fmt.Scan(&M, &N)
	arr := make([][2]int, N)
	for i := 0; i < N; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		arr[i] = [2]int{a, b}
	}
	res := getResult(arr, M)
	fmt.Println(res)
}

func getResult(arr [][2]int, M int) int {
	temp := make([]int, M+1)
	for i := 0; i < len(arr); i++ {
		a, b := arr[i][0], arr[i][1]
		for j := a; j <= b; j++ {
			temp[j] = 1
		}
	}
	res := 0
	for i := 0; i <= M; i++ {
		if temp[i] == 0 {
			res++
		}
	}
	return res
}
```

## 0009.日期倒计时(1)

- 题目

```
题目描述
在经济、科技日益发达的今天，人们对时间的把握越来越严格，对于一个一定影响力的公司的高管来说，
他可能要将自己的行程提前安排到下个月。
对于普通人来说，他也可能将几天之后的安排已经提前做好。
请设计一个程序计算出今天距离未来的某一天还剩多少天。
假设今天是2015年10月18日。
输入 输入一个日期格式为yyyy-MM-dd，不考虑日期是否小于今天。
样例输入
2015-10-19
输出 输出一个数字表示今天（2015年10月18日）距离该日期还剩多少天。
样例输出
1
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 内置函数 | O(1)       | O(1)       |

```go
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
```

## 0010.比大小

### 题目

```
题目描述
现在有"abcdefghijkl”12个字符，将其所有的排列中按字典序排列，给出任意一种排列，
说出这个排列在所有的排列中是第几小的？
输入 第一行有一个整数n（0＜n＜=10000）;
随后有n行，每行是一个排列；
样例输入
3
abcdefghijkl
hgebkflacdji
gfkedhjblcia
输出
输出一个整数m，占一行，m表示排列是第几位；
样例输出
1
302715242
260726926
```

### 解题思路

```go

```

## 0011.约会(1)

- 题目

```
题目描述
Bob和Alice有个约会，一大早Bob就从点(0,0)出发，前往约会地点(a,b)。Bob没有一点方向感，
因此他每次都随机的向上下左右四个方向走一步。简而言之，如果Bob当前在(x,y)，
那么下一步他有可能到达(x+1,y),(x-1,y),(x,y+1), (x,y-1)。
很显然，当他到达目的地的时候，已经很晚了，Alice早已离去。
第二天，Alice质问Bob为什么放她鸽子，Bob说他昨天花了s步到达了约会地点。Alice怀疑Bob是不是说谎了。
你能否帮她验证一下？
输入 输入三个整数a,b,s (-109
样例输入
5 5 11
输出 输出“Yes”，如果Bob可能用s步到达(a,b)；否则输出“No”，不需要输出引号。
样例输出
No
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 数学 | O(1)       | O(1)       |

```go
package main

import (
	"fmt"
)

func main()  {
	var a, b, s int
	fmt.Scan(&a, &b, &s)
	res := getResult(a, b, s)
	if res == true{
		fmt.Println("Yes")
	}else {
		fmt.Println("No")
	}
}

func getResult(a,b,s int) bool  {
	a = abs(a)
	b = abs(b)
	if  a+b <= s && (s-a-b) % 2== 0{ // 注意2个条件
		return true
	}
	return false
}

func abs(a int) int  {
	if a < 0{
		return -a
	}
	return a
}
```

## 0012.研究生考试(1)

- 题目

```
题目描述
欢迎大家参加奇虎360 2016校招在线招聘考试，首先预祝大家都有个好成绩!
我相信参加本次在线招聘考试的有不少研究生同学。我们知道，就计算机相关专业来说，
考研有4门科目，分别是政治（满分100分），英语（满分100分），数学（满分150分）和专业课（满分150分）。
某校计算机专业今年录取研究生的要求是：政治、英语每门课成绩不低于60分，数学和专业课不低于90分，
总成绩不低于310分。
并且规定：在满足单科以及总成绩最低要求的基础上，350分以上（含350分）为公费（Gongfei），
310分~349分为自费(Zifei)。
请编程判断考生的录取情况。
输入 输入数据首先包括一个正整数N，表示有N组测试数据。
每组数据包含4个正整数，分别表示考生的四门课成绩（顺序为：政治、英语、数学、专业课），
你可以假设所有的分数数据都合法。
样例输入
3
61 62 100 120
80 80 120 100
55 90 130 130
输出 请输出每组数据对应考生的录取情况（Fail/Zifei/Gongfei）。
样例输出
Zifei
Gongfei
Fail
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 数学 | O(1)       | O(1)       |

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var a, b, c, d int
		fmt.Scan(&a, &b, &c, &d)
		sum := a + b + c + d
		if a < 60 || b < 60 || c < 90 || d < 90 || sum < 310 {
			fmt.Println("Fail")
		} else if sum >= 310 && sum <= 349 {
			fmt.Println("Zifei")
		} else {
			fmt.Println("Gongfei")
		}
	}
}
```

## 0013.行编辑器(1)

- 题目

```
题目描述
你知道行编辑器吗？不知道也没关系，现在我会告诉你：
1.如果你收到一个
输入 第一行是一个整数T，代表有T组数据。
每组数据的开始时一个字符串，字符串长度小于100，每个字符一定是(
样例输入
3
whli##ilr#e(s#*s)
outcha@putchar(*s=#++)
returnWA##A!!##C
输出 每组数据输出一行经过行编辑器编辑过的字符串，具体可以看样例。
样例输出
while(*s)
putchar(*s++)
returnAC
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 栈   | O(n)       | O(n)       |

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var str string
		fmt.Scan(&str)
		res := getResult(str)
		fmt.Println(res)
	}
}

func getResult(str string) string {
	res := make([]byte, 0)
	for i := 0; i < len(str); i++ {
		if str[i] == '@' {
			res = make([]byte, 0)
		} else if str[i] == '#' {
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
		} else {
			res = append(res, str[i])
		}
	}
	return string(res)
}
```

## 0014.接金币(1)

- 题目

```
题目描述
小赛非常喜欢玩游戏，最近喜欢上了一个接金币的游戏。
在游戏中，使用帽子左右移动接金币，金币接的越多越好，但是金币掉到地上就不能再接了。
为了方便问题的描述，我们把电脑屏幕分成11格，帽子每次能左右移动一格。
现在给电脑屏幕如图标上坐标：
也就是说在游戏里，金币都掉落在0-10这11个位置。
开始时帽子刚开始在5这个位置，因此在第一秒，帽子只能接到4,5,6这三个位置中其中一个位置上的金币。
问小赛在游戏中最多可能接到多少个金币？（假设帽子可以容纳无穷多个金币）。
输入: 输入数据有多组。每组数据的第一行为以正整数n (0＜n＜100000)，表示有n个金币掉在屏幕上上。
在结下来的n行中，每行有两个整数x,T (0＜T＜100000),表示在第T秒有一个金币掉在x点上。
同一秒钟在同一点上可能掉下多个金币。n=0时输入结束。输入数据以空格隔开
样例输入
7
6 3
8 2
9 3
7 1
6 2
5 1
7 2
输出 每一组输入数据对应一行输出。输出一个整数m，表示帽子最多可能接到m个金币。
样例输出
3
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 动态规划 | O(n)       | O(n)       |

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([][2]int, 0)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		arr = append(arr, [2]int{a, b})
	}
	res := getResult(arr)
	fmt.Println(res)
}

func getResult(arr [][2]int) int {
	maxValue := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		maxValue = max(maxValue, arr[i][0])
	}
	// dp[i][j] => 最后在时间i位置j 能获取到的金币数
	dp := make([][11]int, maxValue+1)
	for i := 0; i < n; i++ {
		a, b := arr[i][0], arr[i][1]
		dp[b][a]++
	}
	// 倒推
	for i := maxValue - 1; i >= 0; i-- {
		for j := 0; j < 11; j++ {
			if j == 0 {
				dp[i][j] = dp[i][j] + max(dp[i+1][j], dp[i+1][j+1])
			} else if j == 10 {
				dp[i][j] = dp[i][j] + max(dp[i+1][j], dp[i+1][j-1])
			} else {
				dp[i][j] = dp[i][j] + max(dp[i+1][j], max(dp[i+1][j-1], dp[i+1][j+1]))
			}
		}
	}
	return dp[0][5]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

## 0015.文艺青年爱文学(1)

- 题目

```
题目描述
小赛是一名文艺的程序员，他十分热爱文学。乘车去公司应聘的路上，小赛又在构想自己的诗歌了——
"啊！小赛啊小赛！帅啊很帅可帅了！
啊！小赛啊小赛！棒啊很棒可棒了！
啊！小赛啊小赛！啊啊啊啊啊啊啊！"
尽管小赛的诗歌——额……有那么一点——（啊啊别拦我——让我掐死这只小赛！）……
但是，小赛自己还是深深陶醉其中的。
小赛现在想要创作一首恰好为一定字数（共有n个能满足要求的字数，达到任一皆可）的新诗歌……
他会构想m种长度的短句（如上面那首“诗歌”，有长度为1和7的短句），构想每种长度的短句所耗费的时间是相同的。
现在让你帮忙算下，小赛最少需要多少时间，才能达成自己的目标呢？如果没有办法实现，请输出"It is not true!"（没有引号）。
输入 第一行一个整数n，表示小赛想创作诗歌的字数的集合大小。
接下来n行，其中第i行为一个数a[i]，表示这首诗歌可以是a[i]个字。
第二行一个整数m，表示小赛可以构想出m种不同长度短句的个数。
接下来m行，其中第i行为两个整数b[i],c[i]，其中b[i]表示小赛可以创作出长度为b[i]的短句，
对应的c[i]表示创作这种长度短句所消耗的时间。
数据保证——a[]中的数各不相同，b[]中的数各不相同，c[]不超过10000.
数据保证——
对于30%的测试点，n=1，1<=m<=2，1<=a[],b[]<=20，
对于70%的测试点，1<=n,m<=5，1<=a[],b[]<=100;
对于100%的测试点，1<=n,m<=100，1<=a[],b[]<=10000;
样例输入
2
7
11
2
3 1
2 100
输出 输出一行，如果有办法达成目标，则输出一个整数，表示达成目标所最少需要的创作时间。
如果没有办法达成目标，则输出"It is not true!"（没有引号）
样例输出
103
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 动态规划 | O(n^2)     | O(n)       |

```go
package main

import (
	"fmt"
	"math"
	"sort"
)

var maxConstValue = math.MaxInt32 / 10

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		var tempA int
		fmt.Scan(&tempA)
		a[i] = tempA
	}
	var m int
	fmt.Scan(&m)
	b, c := make([]int, m), make([]int, m)
	for i := 0; i < m; i++ {
		var tempB, tempC int
		fmt.Scan(&tempB, &tempC)
		b[i] = tempB
		c[i] = tempC
	}
	res := getResult(a, b, c)
	fmt.Println(res)
}

func getResult(a, b, c []int) string {
	sort.Ints(a)
	maxValue := a[len(a)-1]
	// dp[i] => 构造长度为i的诗歌长度耗时
	dp := make([]int, maxValue+1)
	for i := 0; i <= maxValue; i++ {
		dp[i] = maxConstValue
	}
	dp[0] = 0
	// 对应的诗歌b[i]需要c[i]的时间,其中b中的数各不相同
	for i := 0; i < len(b); i++ {
		if b[i] <= maxValue {
			dp[b[i]] = c[i]
		}
	}
	for i := 1; i <= maxValue; i++ {
		for j := 0; j < len(b); j++ {
			if b[j] <= i { // 长度i的耗时=长度j的耗时+长度i-j的耗时，找到最小的
				target := i - b[j]
				dp[i] = min(dp[i], dp[target]+c[j])
			}
		}
	}
	res := maxConstValue
	for i := 0; i < len(a); i++ {
		res = min(res, dp[a[i]])
	}
	if res == maxConstValue {
		return "It is not true!"
	}
	return fmt.Sprintf("%d", res)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
```

## 0016.下起楼来我最快(1)

- 题目

```
题目描述
小赛是一名机智的程序员，他的机智主要表现在他下楼的速度特别快( > c < )。
小赛的家住在第n层，他可以选择从电梯下楼（假设只有小赛一个人会用电梯）或者走楼梯下楼。
当前电梯停在第m层，如果他从电梯下到第1层，
需要：电梯先到达这一层->开门->关门->电梯再到达第一层->开门（最后的开门时间也要计算在内）。
现在告诉你——
小赛的家在楼层n，当前电梯停在的楼层m，
以及电梯每经过一层楼的时间t1，开门时间t2，关门时间t3，还有小赛每下一层楼的时间t4，
让你帮小赛计算一下，他最快到达第1层的时间。
输入 第一行两个整数n，m，其中n表示小赛家在的楼层，m表示当前电梯停在的楼层，
第二行四个整数，t1，t2，t3，t4，
其中t1表示电梯每经过一层楼的时间，t2表示开门时间，t3表示关门时间，t4表示小赛每下一层楼的时间。
数据保证——
对于80%的测试点，1<=n,m<=10，1<=t1,t2,t3,t4<=100
对于100%的测试点，1<=n,m<=100000，1<=t1,t2,t3,t4<=100000
样例输入
5 10
1 5 5 4
输出 输出一行，含有一个整数，表示小赛最快到达第1层的时间。
样例输出
16
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 数学 | O(1)       | O(1)       |

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	var t1, t2, t3, t4 int
	fmt.Scan(&t1, &t2, &t3, &t4)
	res := getResult(n, m, t1, t2, t3, t4)
	fmt.Println(res)
}

func getResult(n, m, t1, t2, t3, t4 int) int {
	a := t4 * (n - 1)                    // 纯爬楼梯
	dis := int(math.Abs(float64(n - m))) // 楼层差距
	b := t2*2 + t3 + dis*t1 + (n-1)*t1   // 先爬楼梯到电梯所在层，即m层，然后坐电梯=>开门x2+关门+爬楼梯+电梯运行
	if a < b {
		return a
	}
	return b
}
```

## 0017.回文串(1)

- 题目

```
题目描述
给定一个字符串，问是否能够通过添加一个字母将其变成“回文串”。 
“回文串”是指正着和反着读都一样的字符串。如：”aa”,”bob”,”testset”是回文串，”alice”,”time”都不是回文串。
输入 一行一个有小写字母构成的字符串，字符串长度不超过10。
样例输入
coco
输出 如果输入字符串可以通过添加一个字符，则输出”YES”，否则输出”NO”。
样例输出
YES
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 遍历 | O(n^2)     | O(n)       |

```go
package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	res := getResult(str)
	fmt.Println(res)
}

func getResult(str string) string {
	for i := 0; i < len(str); i++ {
		if isPalindrome(str[0:i]+str[i+1:]) == true {
			return "Yes" // 注意跟题目大小写不一样
		}
	}
	return "No" // 注意跟题目大小写不一样
}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
```

## 0018.小赛的升级之路(1)

- 题目

```
题目描述
小赛经常沉迷于网络游戏。有一次，他在玩一个打怪升级的游戏，他的角色的初始能力值为a。
在接下来的一段时间内，他将会依次遇见n个怪物，每个怪物的防御力为b1,b2,b3,…bn。
如果遇到的怪物防御力bi小于等于小赛的当前能力值c，那么他就能轻松打败怪物，并且使得自己的能力值增加bi；
如果bi大于c，那他也能打败怪物，但他的能力值只能增加bi与c的最大公约数。
那么问题来了，在一系列的锻炼后，小赛的最终能力值为多少？
输入 对于每组数据，第一行是两个整数n(1<=n<=100000)表示怪物的数量和a表示小赛的初始能力值，
第二行n个整数，b1,b2..bn.(1<=bi<=n)表示每个怪物的防御力
数据保证——
50%的n<=100,
80%的n<=1000,
90%的n<=10000,
100%的n<=100000.
样例输入
3 50
50 105 200
5 20
30 20 15 40 100
输出 对于每组数据，输出一行。每行仅包含一个整数，表示小赛的最终能力值。
样例输出
110
205
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 遍历 | O(nlog(n)) | O(1)       |

```go
package main

import "fmt"

func main() {
	var n, a int
	for {
		v, _ := fmt.Scan(&n, &a)
		if v == 0 {
			break
		} else {
			for i := 0; i < n; i++ {
				var temp int
				fmt.Scan(&temp)
				if a >= temp {
					a = a + temp
				} else {
					a = a + gcd(a, temp)
				}
			}
			fmt.Println(a)
		}
	}
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	if a%b == 0 {
		return b
	}
	return gcd(a%b, b)
}
```

## 0019.装载乘客(1)

- 题目

```
题目描述
X学校最近组织了一场春游踏青活动，向Y公司租赁汽车运输学生。
这次参加活动的总共有n个班级，第i班总共有ai名学生，每辆车最大乘车人数为m，满足m>a1, a2, ..., an。
乘车时必须按照班级排列顺序进行乘车，不能调整班级顺序进行拼车。
为保证同一个班级的学生在同一辆车上，如果当前汽车装完上一个班级后，下一个班级所有同学无法装下，那么当前车开走使用下一辆车。
问最少需要多少辆车才能把所有学生运完？
输入 第一行数据是两个整数：n, m (1≤n,m≤100)，n表示班级数目，m表示汽车最大装载人数。
接下来n行是数据表示每个班级的人数数字a1, a2, ..., an (1≤ai≤m)。
样例输入
4 3
2 3 2 1
输出 输出需要的汽车数目。
样例输出
3
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 遍历 | O(n)       | O(n)       |

```go
package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		var temp int
		fmt.Scan(&temp)
		arr[i] = temp
	}
	res := getResult(arr, m)
	fmt.Println(res)
}

func getResult(arr []int, m int) int {
	res := 1
	count := 0
	for i := 0; i < len(arr); i++ {
		if count+arr[i] <= m {
			count = count + arr[i]
			continue
		} else {
			count = arr[i]
			res++
		}
	}
	return res
}
```

## 0020.搬圆桌(1)

- 题目

```
题目描述
小A有一张半径为r的圆桌，其中心位于(x,y)，现在他想把圆桌的中心移到(x1, y1)。
每次移动一步，小A都得在圆桌边界上固定一个点，然后将圆桌绕这个点旋转。 
问最少需要几步才能把圆桌移到目标位置？
输入 一行五个整数r,x,y,x1,y1( 1 ≤ r ≤ 100000,  - 100000 ≤ x, y, x1, y1 ≤ 100000)。
样例输入
2 0 0 0 4
输出 一个整数，表示最少需要移动的步数。
样例输出
1
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 几何 | O(1)       | O(1)       |

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var r, x, y, x1, y1 int
	fmt.Scan(&r, &x, &y, &x1, &y1)
	res := getResult(r, x, y, x1, y1)
	fmt.Println(res)
}

func getResult(r, x, y, x1, y1 int) int {
	v := (x-x1)*(x-x1) + (y-y1)*(y-y1)
	dis := math.Sqrt(float64(v))
	res := int(dis / float64(2*r)) // 除以直径
	if float64(res) < dis/float64(2*r) {
		res = res + 1
	}
	return res
}
```

## 0021.喷水装置(1)

- 题目

```
题目描述
小赛家有一块草坪，长为20米，宽为2米，妈妈要他给草坪浇水，在草坪上放置半径为Ri的喷水装置，
每个喷水装置可以给以它为中心的半径为实数Ri(1＜Ri＜15)的圆形区域浇水。
他家有充足的喷水装置i（1＜i＜600)个，并且一定能把草坪全部湿润。
你能帮他计算一下，把整个草坪全部湿润，最少需要几个喷水装置。
输入 输入第一个数字为喷水装置的个数n，后面n个数字分别为n个喷水装置的半径r，r表示该喷水装置能覆盖的圆的半径。
喷水装置i的范围为：1＜i＜600，半径的范围为：1＜Ri＜15。
样例输入
5
2 3.2 4 4.5 6
输出 输出所用装置的个数。
样例输出
2
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 几何 | O(n)       | O(n)       |

```go
package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]float64, n)
	for i := 0; i < n; i++ {
		var temp float64
		fmt.Scan(&temp)
		arr[i] = temp
	}
	res := getResult(arr, 20)
	fmt.Println(res)
}

func getResult(arr []float64, length float64) int {
	res := 0
	count := float64(0)
	sort.Sort(sort.Reverse(sort.Float64Slice(arr))) // 倒序排序
	for i := 0; i < len(arr); i++ {
		count = count + getValue(arr[i])
		res++
		if count >= length {
			return res
		}
	}
	return res
}

// 计算半径为r的圆可以覆盖的宽度
func getValue(r float64) float64 {
	if r <= 1 {
		return 0
	}
	// 放在中心轴，形成的直角三角形：r*r = 1*1 + x*x
	// 求 2x => 2 * sqrt(r*r-1)
	return 2 * math.Sqrt(r*r-1)
}
```

## 0022.投篮游戏(1)

- 题目

```
题目描述
小赛最近迷上了篮球，报名参加一个投篮游戏。球场有p个篮筐，编号为0, 1, ..., p-1，每个篮筐下面有个袋子，
每个袋子最多能装入一个篮球。
现在有n个篮球，第i个篮球有一个数字xi，投篮规则是将数字为xi的篮球，投入篮筐编号为xi除以p所得的余数。
如果袋子里面已经有球，那么篮球就会弹出，投篮游戏结束，输出i；
否则重复进行将篮球投完，游戏结束，输出-1。问小赛会在何时结束游戏？
输入 第一行数据是两个整数：p, n (2≤p,n≤300)，p表示篮筐数目，n表示篮球数目。接着n行数据表示篮球上的数字xi (0≤xi≤109)。
样例输入
10 5
0
21
53
41
53
输出 输出投篮游戏结束时输出结果。
样例输出
4
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 数组 | O(n)       | O(n)       |

```go
package main

import "fmt"

func main()  {
	var p, n int
	fmt.Scan(&p, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++{
		var temp int
		fmt.Scan(&temp)
		arr[i] = temp
	}
	res := getResult(arr, p)
	fmt.Println(res)
}

func getResult(arr []int, p int) int  {
	m := make(map[int]bool)
	for i := 0; i < len(arr); i++{
		target := arr[i]%p
		if m[target] == true{
			return i+1
		}
		m[target] = true
	}
	return -1
}
```

## 0023.三分线

### 题目

```
题目描述
小赛很喜欢看A队和B队的篮球比赛。
众所周知，篮球每回合根据投篮远近可以得2分或3分。
如果投篮距离小于d那么得2分，大于等于d得3分。我们将d记为三分线。
每次小赛都喜欢通过改变三分线的大小来让自己支持的A队获取更大的优势。
现给出两个队伍投篮得分的距离，小赛希望你能够帮他选择一个合理的三分线，使得A队优势最大。
输入 输入数据包含两行。
第一行第一个数为n(1≤n≤2*105)， 表示A队进球数，接下来n个正整数表示A队每次进球的投篮位置ai(1≤ai≤2*109)。 
第二行第一个数为m(1≤m≤2*105)，表示B队进球数，接下来m个正整数表示B队每次进球的投篮位置bi(1≤bi≤2*109)。
样例输入
3 1 2 3
2 5 6
输出 一个整数，表示A队得分减去B队得分的最大值 
样例输出
3
```

### 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 数组 | O(n)       | O(n)       |

```go
```

