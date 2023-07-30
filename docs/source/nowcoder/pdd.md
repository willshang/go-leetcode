# 牛客-拼多多机试

- https://www.nowcoder.com/ta/exam-pdd

## 1.最大乘积(1)

- 题目

```
题目描述
给定一个无序数组，包含正数、负数和0，要求从中找出3个数的乘积，使得乘积最大，
要求时间复杂度：O(n)，空间复杂度：O(1)
输入描述:输入共2行，第一行包括一个整数n，表示数组长度
第二行为n个以空格隔开的整数，分别为A1,A2, … ,An
输出描述:满足条件的最大乘积
示例1: 输入
4
3 4 1 2
输出: 24
```

- 解题思路

| No. | 思路 | 时间复杂度      | 空间复杂度 |
|:----|:---|:-----------|:------|
| 01  | 排序 | O(nlog(n)) | O(1)  |

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	for {
		a, _ := fmt.Scan(&n)
		if a == 0 {
			break
		}
		nums := make([]int, 0)
		for i := 0; i < n; i++ {
			_, _ = fmt.Scan(&m)
			nums = append(nums, m)
		}
		fmt.Println(maximumProduct(nums))
	}
}

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	return max(nums[0]*nums[1]*nums[len(nums)-1],
		nums[len(nums)-3]*nums[len(nums)-2]*nums[len(nums)-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

## 2.大整数相乘(2)

- 题目

```
题目描述: 有两个用字符串表示的非常大的大整数,算出他们的乘积，也是用字符串表示。不能用系统自带的大整数类型。
输入描述:空格分隔的两个字符串，代表输入的两个大整数
输出描述:输入的乘积，用字符串表示
示例1:
输入:72106547548473106236 982161082972751393
输出:70820244829634538040848656466105986748
```

- 解题思路

| No. | 思路   | 时间复杂度  | 空间复杂度 |
| :-----| :------| :--------| :-------|
| 01  | 模拟   | O(n^2) | O(n)  |
| 02  | 内置函数 | O(n^2) | O(n)  |

```go
package main

import (
	"fmt"
)

func main() {
	var a, b string
	for {
		n, _ := fmt.Scanf("%s %s", &a, &b)
		if n == 0 {
			break
		}
		fmt.Println(multiply(a, b))
	}
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	arr := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		a := int(num1[i] - '0')
		for j := len(num2) - 1; j >= 0; j-- {
			b := int(num2[j] - '0')
			value := a*b + arr[i+j+1]
			arr[i+j+1] = value % 10
			arr[i+j] = value/10 + arr[i+j]
		}
	}
	res := ""
	for i := 0; i < len(arr); i++ {
		if i == 0 && arr[i] == 0 {
			continue
		}
		res = res + string(arr[i]+'0')
	}
	return res
}

# 2
package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b string
	for {
		n, _ := fmt.Scanf("%s %s", &a, &b)
		if n == 0 {
			break
		}
		fmt.Println(multiply(a, b))
	}
}

func multiply(num1 string, num2 string) string {
	a, b := new(big.Int), new(big.Int)
	a.SetString(num1, 10)
	b.SetString(num2, 10)
	a.Mul(a, b)
	return a.String()
}
```

## 3.六一儿童节(2)

- 题目

```
题目描述:六一儿童节，老师带了很多好吃的巧克力到幼儿园。
每块巧克力j的重量为w[j]，对于每个小朋友i，当他分到的巧克力大小达到h[i] (即w[j]>=h[i])，
他才会上去表演节目。
老师的目标是将巧克力分发给孩子们，使得最多的小孩上台表演。
可以保证每个w[i]> 0且不能将多块巧克力分给一个孩子或将一块分给多个孩子。
输入描述:
第一行：n，表示h数组元素个数
第二行：n个h数组元素
第三行：m，表示w数组元素个数
第四行：m个w数组元素
输出描述:上台表演学生人数
示例1: 输入
3
2 2 3
2
3 1 
输出: 1
```

- 解题思路

| No. | 思路    | 时间复杂度      | 空间复杂度 |
| :-----| :-------| :------------| :-------|
| 01  | 排序遍历  | O(nlog(n)) | O(n)  |
| 02  | 排序双指针 | O(nlog(n)) | O(n)  |

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, h int
	var m, w int
	hArr := make([]int, 0)
	wArr := make([]int, 0)
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&h)
		hArr = append(hArr, h)
	}
	fmt.Scan(&m)
	for i := 0; i < m; i++ {
		fmt.Scan(&w)
		wArr = append(wArr, w)
	}
	fmt.Println(deal(hArr, wArr))
}

func deal(hArr, wArr []int) int {
	sort.Ints(hArr)
	sort.Ints(wArr)
	res := 0
	index := len(hArr) - 1
	for i := len(wArr) - 1; i >= 0; i-- {
		for j := index; j >= 0; j-- {
			if wArr[i] >= hArr[j] {
				res++
				index = j - 1
				break
			}
		}
	}
	return res
}

# 2
package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, h int
	var m, w int
	need := make([]int, 0)
	have := make([]int, 0)
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&h)
		need = append(need, h)
	}
	fmt.Scan(&m)
	for i := 0; i < m; i++ {
		fmt.Scan(&w)
		have = append(have, w)
	}
	fmt.Println(deal(need, have))
}

func deal(need, have []int) int {
	sort.Ints(need)
	sort.Ints(have)
	res := 0
	for i, j := 0, 0; i < len(have) && j < len(need); i++ {
		if have[i] >= need[j] {
			res++
			j++
		}
	}
	return res
}
```

## 4.迷宫寻路

### 题目

```
题目描述:假设一个探险家被困在了地底的迷宫之中，要从当前位置开始找到一条通往迷宫出口的路径。
迷宫可以用一个二维矩阵组成，有的部分是墙，有的部分是路。
迷宫之中有的路上还有门，每扇门都在迷宫的某个地方有与之匹配的钥匙，只有先拿到钥匙才能打开门。
请设计一个算法，帮助探险家找到脱困的最短路径。
如前所述，迷宫是通过一个二维矩阵表示的，每个元素的值的含义如下 
0-墙，1-路，2-探险家的起始位置，3-迷宫的出口，大写字母-门，小写字母-对应大写字母所代表的门的钥匙
输入描述:迷宫的地图，用二维矩阵表示。
第一行是表示矩阵的行数和列数M和N,后面的M行是矩阵的数据，每一行对应与矩阵的一行（中间没有空格）。
M和N都不超过100, 门不超过10扇。
输出描述:路径的长度，是一个整数
示例1:输入
5 5
02111
01a0A
01003
01001
01111
输出: 7
```

### 解题思路

| No. | 思路   | 时间复杂度      | 空间复杂度 |
| :-----| :------| :------------| :-------|
| 01  | 排序遍历 | O(nlog(n)) | O(n)  |

```go

```

## 7.数三角形

### 题目

```
题目描述: 给出平面上的n个点，现在需要你求出，在这n个点里选3个点能构成一个三角形的方案有几种。
输入描述:第一行包含一个正整数n，表示平面上有n个点（n <= 100)
第2行到第n + 1行，每行有两个整数，表示这个点的x坐标和y坐标。
(所有坐标的绝对值小于等于100，且保证所有坐标不同）
输出描述:输出一个数，表示能构成三角形的方案数。
示例1: 输入
4
0 0
0 1
1 0
1 1
输出 4
说明4个点中任意选择3个都能构成三角形
```

### 解题思路

| No. | 思路  | 时间复杂度  | 空间复杂度 |
| :-----| :-----| :--------| :-------|
| 01  | 暴力法 | O(n^3) | O(1)  |

```go
package main

import "fmt"

func main() {
	var n int
	for {
		a, _ := fmt.Scan(&n)
		if a == 0 {
			break
		}
		arr := make([][2]int, n)
		for i := 0; i < n; i++ {
			var x, y int
			fmt.Scan(&x, &y)
			arr[i] = [2]int{x, y}
		}
		fmt.Println(count(arr))
	}
}

func count(arr [][2]int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				// 斜率判断
				l1 := (arr[i][0] - arr[j][0]) * (arr[i][1] - arr[k][1])
				l2 := (arr[i][1] - arr[j][1]) * (arr[i][0] - arr[k][0])
				if l1 != l2 {
					res++
				}
			}
		}
	}
	return res
}
```

## 8.最大乘积(1)

- 题目

```
题目描述给定一个无序数组，包含正数、负数和0，要求从中找出3个数的乘积，使得乘积最大，
要求时间复杂度：O(n)，空间复杂度：O(1)。
n<=1e5。
输入描述:第一行是数组大小n，第二行是无序整数数组A[n]
输出描述:满足条件的最大乘积
示例1 输入
4
3 4 1 2
输出 24
```

- 解题思路

| No. | 思路 | 时间复杂度      | 空间复杂度 |
| :-----| :----| :------------| :-------|
| 01  | 排序 | O(nlog(n)) | O(1)  |

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	for {
		a, _ := fmt.Scan(&n)
		if a == 0 {
			break
		}
		nums := make([]int, 0)
		for i := 0; i < n; i++ {
			_, _ = fmt.Scan(&m)
			nums = append(nums, m)
		}
		fmt.Println(maximumProduct(nums))
	}
}

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	return max(nums[0]*nums[1]*nums[len(nums)-1],
		nums[len(nums)-3]*nums[len(nums)-2]*nums[len(nums)-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

## 13.回合制游戏(1)

- 题目

```
题目描述
你在玩一个回合制角色扮演的游戏。现在你在准备一个策略，以便在最短的回合内击败敌方角色。
在战斗开始时，敌人拥有HP格血量。当血量小于等于0时，敌人死去。
一个缺乏经验的玩家可能简单地尝试每个回合都攻击。但是你知道辅助技能的重要性。
在你的每个回合开始时你可以选择以下两个动作之一：聚力或者攻击。
    聚力会提高你下个回合攻击的伤害。
    攻击会对敌人造成一定量的伤害。如果你上个回合使用了聚力，那这次攻击会对敌人造成buffedAttack点伤害。
    否则，会造成normalAttack点伤害。
给出血量HP和不同攻击的伤害，buffedAttack和normalAttack，返回你能杀死敌人的最小回合数。
输入描述:
第一行是一个数字HP
第二行是一个数字normalAttack
第三行是一个数字buffedAttack
1 <= HP,buffedAttack,normalAttack <= 10^9
输出描述:输出一个数字表示最小回合数
示例1: 输入
13
3
5
输出: 5
```

- 解题思路

| No. | 思路   | 时间复杂度  | 空间复杂度 |
| :-----| :------| :--------| :-------|
| 01  | 数学计算 | O(n^3) | O(1)  |

```go
package main

import "fmt"

func getCount(hp, normalAttack, buffedAttack int) int {
	res := 0
	if normalAttack*2 < buffedAttack {
		res = res + 2*(hp/buffedAttack)
		hp = hp - hp/buffedAttack*buffedAttack
		if hp > 0 {
			if hp <= normalAttack {
				res = res + 1
			} else {
				res = res + 2
			}
		}
	} else {
		res = res + hp/normalAttack
		if hp%normalAttack > 0 {
			res = res + 1
		}
	}
	return res
}

func main() {
	var hp, normalAttack, buffedAttack int
	for {
		// 三行数字用fmt.Scan
		a, _ := fmt.Scan(&hp, &normalAttack, &buffedAttack)
		if a == 0 {
			break
		}
		fmt.Println(getCount(hp, normalAttack, buffedAttack))
	}
}
```