# 程序员面试金典

## 面试题01.01.判定字符是否唯一(5)

- 题目

```
实现一个算法，确定一个字符串 s 的所有字符是否全都不同。
示例 1：输入: s = "leetcode" 输出: false 
示例 2：输入: s = "abc" 输出: true
限制：
    0 <= len(s) <= 100
    如果你不使用额外的数据结构，会很加分。
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 哈希辅助 | O(n)       | O(1)       |
| 02   | 位运算   | O(n)       | O(1)       |
| 03   | 遍历     | O(n^2)     | O(1)       |
| 04   | 排序遍历 | O(nlog(n)) | O(n)       |
| 05   | 数组辅助 | O(n)       | O(1)       |

```go
func isUnique(astr string) bool {
	m := make(map[byte]bool)
	for i := 0; i < len(astr); i++ {
		if m[astr[i]] == true {
			return false
		}
		m[astr[i]] = true
	}
	return true
}

# 2
func isUnique(astr string) bool {
	value := uint32(0)
	for i := 0; i < len(astr); i++ {
		index := astr[i] - 'a'
		if value&(1<<index) == (1 << index) {
			return false
		}
		value = value ^ (1 << index)
	}
	return true
}

# 3
func isUnique(astr string) bool {
	for i := 0; i < len(astr); i++ {
		for j := i + 1; j < len(astr); j++ {
			if astr[i] == astr[j] {
				return false
			}
		}
	}
	return true
}

# 4
func isUnique(astr string) bool {
	arr := []byte(astr)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			return false
		}
	}
	return true
}

# 5
func isUnique(astr string) bool {
	arr := make([]int, 256)
	for i := 0; i < len(astr); i++ {
		if arr[astr[i]] > 0 {
			return false
		}
		arr[astr[i]] = 1
	}
	return true
}
```

## 面试题01.02.判定是否互为字符重排(2)

- 题目

```
给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。
示例 1：输入: s1 = "abc", s2 = "bca" 输出: true 
示例 2：输入: s1 = "abc", s2 = "bad" 输出: false
说明：
    0 <= len(s1) <= 100
    0 <= len(s2) <= 100 
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 内置函数 | O(nlog(n)) | O(n)       |
| 02   | 哈希辅助 | O(n)       | O(1)       |
| 03   | 数组辅助 | O(n)       | O(1)       |

```go
func CheckPermutation(s1 string, s2 string) bool {
	arr1 := strings.Split(s1, "")
	arr2 := strings.Split(s2, "")
	sort.Strings(arr1)
	sort.Strings(arr2)
	return strings.Join(arr1,"") == strings.Join(arr2,"")
	// return reflect.DeepEqual(arr1, arr2)
}

#
func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		m[s1[i]]++
		m[s2[i]]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

#
func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	arr := [256]int{}
	for i := 0; i < len(s1); i++ {
		arr[s1[i]]++
		arr[s2[i]]--
	}
	for _, v := range arr {
		if v != 0 {
			return false
		}
	}
	return true
}
```

## 面试题01.03.URL化(2)

- 题目

```
URL化。编写一种方法，将字符串中的空格全部替换为%20。假定该字符串尾部有足够的空间存放新增字符，
并且知道字符串的“真实”长度。（注：用Java实现的话，请使用字符数组实现，以便直接在数组上操作。）
示例1:输入："Mr John Smith    ", 13 输出："Mr%20John%20Smith"
示例2:输入："               ", 5 输出："%20%20%20%20%20"
提示：
    字符串长度在[0, 500000]范围内。
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 内置函数 | O(n)       | O(n)       |
| 02   | 遍历     | O(n)       | O(n)       |

```go
func replaceSpaces(S string, length int) string {
	return strings.ReplaceAll(S[:length], " ","%20")
}

#
func replaceSpaces(S string, length int) string {
	res := make([]byte,0)
	for i := 0; i < length; i++ {
		if S[i] == ' ' {
			res = append(res,'%')
			res = append(res,'2')
			res = append(res,'0')
		} else {
			res = append(res,S[i])
		}
	}
	return string(res)
}
```

## 面试题01.04.回文排列(2)

- 题目

```
给定一个字符串，编写一个函数判定其是否为某个回文串的排列之一。
回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。
回文串不一定是字典当中的单词。
示例1：输入："tactcoa" 输出：true（排列有"tacocat"、"atcocta"，等等）
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 哈希辅助 | O(n)       | O(1)       |
| 02   | 数组辅助 | O(n)       | O(1)       |

```go
func canPermutePalindrome(s string) bool {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
		if m[s[i]] == 2 {
			delete(m, s[i])
		}
	}
	return len(m) <= 1
}

#
func canPermutePalindrome(s string) bool {
	arr := [256]int{}
	for i := 0; i < len(s); i++ {
		arr[s[i]]++
	}
	count := 0
	for i := 0; i < len(arr); i++{
		if arr[i] % 2== 1{
			count++
		}
	}
	return count <= 1
}
```

## 面试题01.05.一次编辑(2)

- 题目

```
字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。
给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。
示例 1:输入: first = "pale"second = "ple" 输出: True
示例 2:输入: first = "pales"second = "pal" 输出: False
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 遍历 | O(n)       | O(1)       |
| 02   | 遍历 | O(n)       | O(1)       |

```go
func oneEditAway(first string, second string) bool {
	if len(first)-len(second) > 1 || len(second)-len(first) > 1 {
		return false
	}
	if first == second {
		return true
	}
	i := 0
	for ; i < len(first) && i < len(second); i++ {
		if first[i] != second[i] {
			if len(first) == len(second) {
				if first[i+1:] == second[i+1:] {
					return true
				}
			} else if len(first) < len(second) {
				if first[i:] == second[i+1:] {
					return true
				}
			} else {
				if first[i+1:] == second[i:] {
					return true
				}
			}
			break
		}
	}
	if i == len(first) || i == len(second) {
		return true
	}
	return false
}

#
func oneEditAway(first string, second string) bool {
	if len(first)-len(second) > 1 || len(second)-len(first) > 1 {
		return false
	}
	if first == second {
		return true
	}
	if len(first) > len(second) {
		first, second = second, first
	}
	for i := 0; i < len(first); i++ {
		if first[i] == second[i] {
			continue
		}
		return first[i:] == second[i+1:] || first[i+1:] == second[i+1:]
	}
	return true
}
```

## 面试题01.06.字符串压缩(2)

- 题目

```
字符串压缩。利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。
比如，字符串aabcccccaaa会变为a2b1c5a3。若“压缩”后的字符串没有变短，则返回原先的字符串。
你可以假设字符串中只包含大小写英文字母（a至z）。
示例1:输入："aabcccccaaa" 输出："a2b1c5a3"
示例2:输入："abbccd" 输出："abbccd" 
解释："abbccd"压缩后为"a1b2c2d1"，比原字符串长度更长。
提示：字符串长度在[0, 50000]范围内。
```

- 解题思路

| No.  | 思路   | 时间复杂度 | 空间复杂度 |
| ---- | ------ | ---------- | ---------- |
| 01   | 遍历   | O(n)       | O(n)       |
| 02   | 双指针 | O(n)       | O(n)       |

```go
func compressString(S string) string {
	if len(S) <= 1 {
		return S
	}
	prev := S[0]
	count := 1
	res := ""
	for i := 1; i < len(S); i++ {
		if prev == S[i] {
			count++
		} else {
			res = res + string(prev) + strconv.Itoa(count)
			prev = S[i]
			count = 1
		}
	}
	res = res + string(prev) + strconv.Itoa(count)
	if len(res) >= len(S) {
		return S
	}
	return res
}

#
func compressString(S string) string {
	if len(S) <= 1 {
		return S
	}
	i := 0
	j := 0
	res := ""
	for j = 1; j < len(S); j++ {
		if S[i] != S[j] {
			res = res + string(S[i]) + strconv.Itoa(j-i)
			i = j
		}
	}
	res = res + string(S[i]) + strconv.Itoa(j-i)
	if len(res) >= len(S) {
		return S
	}
	return res
}
```

## 面试题01.07.旋转矩阵(3)

- 题目

```
给你一幅由 N × N 矩阵表示的图像，其中每个像素的大小为 4 字节。请你设计一种算法，将图像旋转 90 度。
不占用额外内存空间能否做到？
示例 1:给定 matrix = 
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],
原地旋转输入矩阵，使其变为:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
示例 2:给定 matrix =
[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
], 
原地旋转输入矩阵，使其变为:
[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 遍历     | O(n^2)     | O(1)       |
| 02   | 遍历     | O(n^2)     | O(1)       |
| 03   | 数组辅助 | O(n^2)     | O(n^2)     |

```go
func rotate(matrix [][]int) {
	n := len(matrix)
	// 同行逆置
	// [[1 2 3] [4 5 6] [7 8 9]]
	// [[3 2 1] [6 5 4] [9 8 7]]
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
	// 左下右上对角线对互换
	// [[3 2 1] [6 5 4] [9 8 7]]
	// [[7 4 1] [8 5 2] [9 6 3]]
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			matrix[i][j], matrix[n-1-j][n-1-i] = matrix[n-1-j][n-1-i], matrix[i][j]
		}
	}
}

# 2
func rotate(matrix [][]int) {
	n := len(matrix)
	for start, end := 0, n-1; start < end; {
		for s, e := start, end; s < end; {
			matrix[start][s], matrix[e][start], matrix[end][e], matrix[s][end] =
				matrix[e][start], matrix[end][e], matrix[s][end], matrix[start][s]
			s++
			e--
		}
		start++
		end--
	}
}

# 3
func rotate(matrix [][]int) {
	n := len(matrix)
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			arr[j][n-1-i] = matrix[i][j]
		}
	}
	copy(matrix, arr)
}
```

## 面试题01.08.零矩阵(4)

- 题目

```
编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零。
示例 1：输入：
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
输出：
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]
示例 2：输入：
[
  [0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]
]
输出：
[
  [0,0,0,0],
  [0,4,5,0],
  [0,3,1,0]
]
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 哈希辅助 | O(n^2)     | O(n)       |
| 02   | 暴力法   | O(n^4)     | O(1)       |
| 03   | 遍历     | O(n^2)     | O(1)       |
| 04   | 遍历     | O(n^2)     | O(1)       |

```go
func setZeroes(matrix [][]int) {
	x := make(map[int]int)
	y := make(map[int]int)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				x[i] = 1
				y[j] = 1
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if x[i] == 1 || y[j] == 1 {
				matrix[i][j] = 0
			}
		}
	}
}

# 2
func setZeroes(matrix [][]int) {
	m := make(map[[2]int]bool)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == math.MinInt32 {
				m[[2]int{i, j}] = true
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				for k := 0; k < len(matrix); k++ {
					for l := 0; l < len(matrix[k]); l++ {
						if (k == i || l == j) && matrix[k][l] != 0 {
							delete(m, [2]int{k, l})
							matrix[k][l] = math.MinInt32
						}
					}
				}
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == math.MinInt32 && m[[2]int{i, j}] == false {
				matrix[i][j] = 0
			}
		}
	}
}

# 3
func setZeroes(matrix [][]int) {
	flag := false
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			flag = true
		}
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	// 第一行处理
	if matrix[0][0] == 0 {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}
	// 第一列处理
	if flag == true {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}

# 4
func setZeroes(matrix [][]int) {
	flag := false
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			flag = true
		}
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := len(matrix) - 1; i >= 0; i-- {
		for j := len(matrix[i]) - 1; j >= 1; j-- {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	// 第一列处理
	if flag == true {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}
```

## 面试题01.09.字符串轮转(2)

- 题目

```
字符串轮转。给定两个字符串s1和s2，请编写代码检查s2是否为s1旋转而成（
比如，waterbottle是erbottlewat旋转后的字符串）。
示例1: 输入：s1 = "waterbottle", s2 = "erbottlewat" 输出：True
示例2:输入：s1 = "aa", s2 = "aba" 输出：False
提示：字符串长度在[0, 100000]范围内。
说明: 你能只调用一次检查子串的方法吗？
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 内置函数 | O(n)       | O(1)       |
| 02   | 遍历     | O(n)       | O(1)       |

```go
func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2){
		return false
	}
	return strings.Contains(s1+s1, s2)
}

#
func isFlipedString(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		s1 = s1[1:] + string(s1[0])
		if s1 == s2 {
			return true
		}
	}
	return false
}
```

## 面试题02.01.移除重复节点

### 题目

```
编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。
示例1:输入：[1, 2, 3, 3, 2, 1]输出：[1, 2, 3]
示例2:输入：[1, 1, 1, 1, 2]输出：[1, 2]
提示：
    链表长度在[0, 20000]范围内。
    链表元素在[0, 20000]范围内。
进阶：如果不得使用临时缓冲区，该怎么解决？
```

### 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 内置函数 | O(n)       | O(1)       |

```go
func oneEditAway(first string, second string) bool {
	if len(first)-len(second) > 1 || len(second)-len(first) > 1 {
		return false
	}
	if first == second {
		return true
	}
	i := 0
	for ; i < len(first) && i < len(second); i++ {
		if first[i] != second[i] {
			if len(first) == len(second) {
				if first[i+1:] == second[i+1:] {
					return true
				}
			} else if len(first) < len(second) {
				if first[i:] == second[i+1:] {
					return true
				}
			} else {
				if first[i+1:] == second[i:] {
					return true
				}
			}
			break
		}
	}
	if i == len(first) || i == len(second) {
		return true
	}
	return false
}
#
func oneEditAway(first string, second string) bool {
	if len(first)-len(second) > 1 || len(second)-len(first) > 1 {
		return false
	}
	if first == second {
		return true
	}
	if len(first) > len(second) {
		first, second = second, first
	}
	for i := 0; i < len(first); i++ {
		if first[i] == second[i] {
			continue
		}
		return first[i:] == second[i+1:] || first[i+1:] == second[i+1:]
	}
	return true
}
```

