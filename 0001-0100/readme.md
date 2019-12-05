# 0001-0100
[TOC]
## 1. 两数之和
### 题目
``` 
给定一个整数数组 nums 和一个目标值 target，
请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:
给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
```
### 解答思路
| No.      | 思路                | 时间复杂度 | 空间复杂度 |
| -------- | ------------------- | ---------- | ---------- |
| 01       | 暴力法: 2层循环遍历 | O(n^2)     | O(1)       |
| 02       | 两遍哈希遍历        | O(n)       | O(n)       |
| 03(最优) | 一遍哈希遍历        | O(n)       | O(n)       |

```go
# 暴力法: 2层循环遍历
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

# 两遍哈希遍历
func twoSum(nums []int, target int) []int {
	m := make(map[int]int,len(nums))
	for k, v := range nums{
		m[v] = k
	}

	for i := 0; i < len(nums); i++{
		b := target - nums[i]
		if num, ok := m[b]; ok && num != i{
			return []int{i,m[b]}
		}
	}
	return []int{}
}

# 一遍哈希遍历
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, b := range nums {
		if j, ok := m[target-b]; ok {
			return []int{j, i}
		}
		m[b] = i
	}
	return nil
}
```

##  7. 整数反转

### 题目

```
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:
输入: 123
输出: 321

示例 2:
输入: -123
输出: -321

示例 3:
输入: 120
输出: 21

注意:假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。
请根据这个假设，如果反转后整数溢出那么就返回 0。
```

### 解答思路
| No.      | 思路                                                  | 时间复杂度 | 空间复杂度 |
| -------- | ----------------------------------------------------- | ---------- | ---------- |
| 01       | 使用符号标记，转成正数，循环得到%10的余数，再加上符号 | O(log(x))  | O(1)       |
| 02(最优) | 对x进行逐个%10取个位，一旦溢出，直接跳出循环          | O(log(x))  | O(1)       |

```go
// 使用符号标记，转成正数，循环得到%10的余数，再加上符号
func reverse(x int) int {
	flag := 1
	if x < 0 {
		flag = -1
		x = -1 * x
	}

	result := 0
	for x > 0 {
		temp := x % 10
		x = x / 10

		result = result*10 + temp
	}

	result = flag * result
	if result > math.MaxInt32 || result < math.MinInt32 {
		result = 0
	}
	return result
}

// 对x进行逐个%10取个位，一旦溢出，直接跳出循环
func reverse(x int) int {
	result := 0
	for x != 0 {
		temp := x % 10
		result = result*10 + temp
		if result > math.MaxInt32 || result < math.MinInt32 {
			return 0
		}
		x = x / 10
	}
	return result
}
```



## 9.回文数

### 题目

```
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:	输入: 121	输出: true

示例 2:输入: -121	输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。

示例 3:输入: 10  	输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。

进阶:
你能不将整数转为字符串来解决这个问题吗？
```

### 解答思路

| No.      | 思路                                               | 时间复杂度 | 空间复杂度 |
| -------- | -------------------------------------------------- | ---------- | ---------- |
| 01(最优) | 数学解法，取出后半段数字进行翻转，然后判断是否相等 | O(log(x))  | O(1)       |
| 02       | 转成字符串，依次判断                               | O(log(x))  | O(log(x))  |
| 03       | 转成byte数组，依次判断，同2                        | O(log(x))  | O(log(x))  |

```go
// 数学解法，取出后半段数字进行翻转，然后判断是否相等
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		temp := x % 10
		revertedNumber = revertedNumber*10 + temp
		x = x / 10
	}
	// for example:
	// x = 1221  => x = 12 revertedNumber = 12
	// x = 12321 => x = 12 revertedNumber = 123
	return x == revertedNumber || x == revertedNumber/10
}

// 转成字符串，依次判断
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

// 转成byte数组，依次判断，同2
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	arrs := []byte(strconv.Itoa(x))
	Len := len(arrs)
	for i := 0; i < Len/2; i++ {
		if arrs[i] != arrs[Len-i-1] {
			return false
		}
	}
	return true
}
```



