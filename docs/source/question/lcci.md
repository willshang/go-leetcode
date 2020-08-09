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

## 面试题02.01.移除重复节点(3)

- 题目

```
编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。
示例1:输入：[1, 2, 3, 3, 2, 1]输出：[1, 2, 3]
示例2:输入：[1, 1, 1, 1, 2]输出：[1, 2]
提示：
    链表长度在[0, 20000]范围内。
    链表元素在[0, 20000]范围内。
进阶：如果不得使用临时缓冲区，该怎么解决？
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 哈希辅助 | O(n)       | O(n)       |
| 02   | 遍历     | O(n^2)     | O(1)       |
| 03   | 递归     | O(n)       | O(n)       |

```go
func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	m := make(map[int]bool)
	m[head.Val] = true
	temp := head
	for temp.Next != nil {
		if m[temp.Next.Val] == true {
			temp.Next = temp.Next.Next
		} else {
			m[temp.Next.Val] = true
			temp = temp.Next
		}
	}
	return head
}

# 2
func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	temp := head
	for temp != nil {
		second := temp
		for second.Next != nil {
			if second.Next.Val == temp.Val {
				second.Next = second.Next.Next
			} else {
				second = second.Next
			}
		}
		temp = temp.Next
	}
	return head
}

# 3
var m map[int]bool

func removeDuplicateNodes(head *ListNode) *ListNode {
	m = make(map[int]bool)
	return remove(head)

}

func remove(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if m[head.Val] == true {
		return remove(head.Next)
	}
	m[head.Val] = true
	head.Next = remove(head.Next)
	return head
}
```

## 面试题02.02.返回倒数第k个节点(4)

- 题目

```
实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。
注意：本题相对原题稍作改动
示例：输入： 1->2->3->4->5 和 k = 2 输出： 4
说明：给定的 k 保证是有效的。
```

- 解题思路

| No.  | 思路      | 时间复杂度 | 空间复杂度 |
| ---- | --------- | ---------- | ---------- |
| 01   | 数组辅助  | O(n)       | O(n)       |
| 02   | 快慢指针  | O(n)       | O(1)       |
| 03   | 统计+遍历 | O(n)       | O(1)       |
| 04   | 递归      | O(n)       | O(n)       |

```go
func kthToLast(head *ListNode, k int) int {
	arr := make([]*ListNode, 0)
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}
	if len(arr) >= k {
		return arr[len(arr)-k].Val
	}
	return -1
}

# 2
func kthToLast(head *ListNode, k int) int {
	fast := head
	for k > 0 && head != nil {
		fast = fast.Next
		k--
	}
	if k > 0 {
		return -1
	}
	slow := head
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow.Val
}

# 3
func kthToLast(head *ListNode, k int) int {
	temp := head
	count := 0
	for temp != nil {
		count++
		temp = temp.Next
	}
	if count < k {
		return -1
	}
	for i := 0; i < count-k; i++ {
		head = head.Next
	}
	return head.Val
}

# 4
func kthToLast(head *ListNode, k int) int {
	res, count := dfs(head, k)
	if count > 0 {
		return -1
	}
	return res.Val
}

func dfs(node *ListNode, k int) (*ListNode, int) {
	if node == nil {
		return node, k
	}
	next, nextValue := dfs(node.Next, k)
	if nextValue <= 0 {
		return next, nextValue
	}
	nextValue = nextValue - 1
	return node, nextValue
}
```

## 面试题02.03.删除中间节点(1)

- 题目

```
实现一种算法，删除单向链表中间的某个节点（即不是第一个或最后一个节点），假定你只能访问该节点。
示例：输入：单向链表a->b->c->d->e->f中的节点c
结果：不返回任何数据，但该链表变为a->b->d->e->f
```

- 解题思路

| No.  | 思路                       | 时间复杂度 | 空间复杂度 |
| ---- | -------------------------- | ---------- | ---------- |
| 01   | 把当前节点替换成下一个节点 | O(1)       | O(1)       |

```go
func deleteNode(node *ListNode) {
	// *node = *node.Next
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
```

## 面试题02.04.分割链表(2)

- 题目

```
编写程序以 x 为基准分割链表，使得所有小于 x 的节点排在大于或等于 x 的节点之前。
如果链表中包含 x，x 只需出现在小于 x 的元素之后(如下所示)。
分割元素 x 只需处于“右半部分”即可，其不需要被置于左右两部分之间。
示例:输入: head = 3->5->8->5->10->2->1, x = 5
输出: 3->1->2->10->5->5->8
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 双指针   | O(n)       | O(1)       |
| 02   | 数组辅助 | O(n)       | O(n)       |

```go
func partition(head *ListNode, x int) *ListNode {
	first := &ListNode{}
	second := &ListNode{}
	a := first
	b := second
	for head != nil {
		if head.Val < x {
			a.Next = head
			a = head
		} else {
			b.Next = head
			b = head
		}
		head = head.Next
	}
	b.Next = nil
	a.Next = second.Next
	return first.Next
}

# 2
func partition(head *ListNode, x int) *ListNode {
	a := make([]*ListNode, 0)
	b := make([]*ListNode, 0)

	for head != nil {
		if head.Val < x {
			a = append(a, head)
		} else {
			b = append(b, head)
		}
		head = head.Next
	}
	temp := &ListNode{}
	node := temp
	for i := 0; i < len(a); i++ {
		node.Next = a[i]
		node = node.Next
	}
	for i := 0; i < len(b); i++ {
		node.Next = b[i]
		node = node.Next
	}
	node.Next = nil
	return temp.Next
}
```

## 面试题02.05.链表求和(2)

- 题目

```
给定两个用链表表示的整数，每个节点包含一个数位。
这些数位是反向存放的，也就是个位排在链表首部。
编写函数对这两个整数求和，并用链表形式返回结果。
示例：输入：(7 -> 1 -> 6) + (5 -> 9 -> 2)，即617 + 295 输出：2 -> 1 -> 9，即912
进阶：假设这些数位是正向存放的，请再做一遍。
示例：输入：(6 -> 1 -> 7) + (2 -> 9 -> 5)，即617 + 295 输出：9 -> 1 -> 2，即912
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 遍历 | O(n)       | O(n)       |
| 02   | 递归 | O(n)       | O(n)       |

```go
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := res
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10 // 进位
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}
	return res.Next
}

# 2
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	sum := l1.Val + l2.Val
	res := &ListNode{Val: sum % 10}
	if sum >= 10 {
		l1.Next = addTwoNumbers(l1.Next, &ListNode{Val: 1})
	}
	res.Next = addTwoNumbers(l1.Next, l2.Next)
	return res
}
```

## 面试题02.06.回文链表(4)

- 题目

```
编写一个函数，检查输入的链表是否是回文的。
示例 1：输入： 1->2 输出： false 
示例 2：输入： 1->2->2->1 输出： true 
进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
```

- 解题思路

| No.  | 思路             | 时间复杂度 | 空间复杂度 |
| ---- | ---------------- | ---------- | ---------- |
| 01   | 数组辅助         | O(n)       | O(n)       |
| 02   | 快慢指针反转链表 | O(n)       | O(1)       |
| 03   | 栈辅助           | O(n)       | O(n)       |
| 04   | 递归             | O(n)       | O(n)       |

```go
func isPalindrome(head *ListNode) bool {
	m := make([]int, 0)
	for head != nil {
		m = append(m, head.Val)
		head = head.Next
	}
	i, j := 0, len(m)-1
	for i < j {
		if m[i] != m[j] {
			return false
		}
		i++
		j--
	}
	return true
}

# 2
func isPalindrome(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	var pre *ListNode
	cur := slow
	for cur != nil{
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	for pre != nil{
		if head.Val != pre.Val{
			return false
		}
		pre = pre.Next
		head = head.Next
	}
	return true
}

# 3
func isPalindrome(head *ListNode) bool {
	m := make([]int, 0)
	temp := head
	for temp != nil {
		m = append(m, temp.Val)
		temp = temp.Next
	}
	for head != nil {
		val := m[len(m)-1]
		m = m[:len(m)-1]
		if head.Val != val {
			return false
		}
		head = head.Next
	}
	return true
}

# 4
var p *ListNode
func isPalindrome(head *ListNode) bool {
	if head == nil{
		return true
	}
	if p == nil{
		p = head
	}
	if isPalindrome(head.Next) && (p.Val == head.Val){
		p = p.Next
		return true
	}
	p = nil
	return false
}
```

## 面试题02.07.链表相交(4)

- 题目

```
给定两个（单向）链表，判定它们是否相交并返回交点。请注意相交的定义基于节点的引用，而不是基于节点的值。
换句话说，如果一个链表的第k个节点与另一个链表的第j个节点是同一节点（引用完全相同），则这两个链表相交。
示例 1：
输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
输出：Reference of the node with value = 8
输入解释：相交节点的值为 8 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，
链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1,8,4,5]。
在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。
示例 2：输入：intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
输出：Reference of the node with value = 2
输入解释：相交节点的值为 2 （注意，如果两个列表相交则不能为 0）。
从各自的表头开始算起，链表 A 为 [0,9,1,2,4]，链表 B 为 [3,2,4]。
在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。
示例 3：输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
输出：null
输入解释：从各自的表头开始算起，链表 A 为 [2,6,4]，链表 B 为 [1,5]。
由于这两个链表不相交，所以 intersectVal 必须为 0，而 skipA 和 skipB 可以是任意值。
解释：这两个链表不相交，因此返回 null。
注意：
    如果两个链表没有交点，返回 null 。
    在返回结果后，两个链表仍须保持原有的结构。
    可假定整个链表结构中没有循环。
    程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 对齐比较 | O(n)       | O(1)       |
| 02   | 交换比较 | O(n)       | O(1)       |
| 03   | 暴力法   | O(n^2)     | O(1)       |
| 04   | 哈希辅助 | O(n)       | O(n)       |

```go
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	ALength := 0
	A := headA
	for A != nil {
		ALength++
		A = A.Next
	}
	BLength := 0
	B := headB
	for B != nil {
		BLength++
		B = B.Next
	}

	pA := headA
	pB := headB
	if ALength > BLength {
		n := ALength - BLength
		for n > 0 {
			pA = pA.Next
			n--
		}
	} else {
		n := BLength - ALength
		for n > 0 {
			pB = pB.Next
			n--
		}
	}

	for pA != pB {
		pA = pA.Next
		pB = pB.Next
	}
	return pA
}

# 2
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	A, B := headA, headB
	for A != B {
		if A != nil {
			A = A.Next
		} else {
			A = headB
		}
		if B != nil {
			B = B.Next
		} else {
			B = headA
		}
	}
	return A
}

# 3
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    A, B := headA, headB
    for A != nil {
        for B != nil {
            if A == B {
                return A
            }
            B = B.Next
        }
        A = A.Next
        B = headB
    }
    return nil
}

# 4
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]bool)
	for headA != nil {
		m[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if _, ok := m[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
```

## 面试题02.08.环路检测(3)

- 题目

```
给定一个链表，如果它是有环链表，实现一个算法返回环路的开头节点。
有环链表的定义：在链表中某个节点的next元素指向在它前面出现过的节点，则表明该链表存在环路。
示例 1：输入：head = [3,2,0,-4], pos = 1 输出：tail connects to node index 1
解释：链表中有一个环，其尾部连接到第二个节点。
示例 2：输入：head = [1,2], pos = 0 输出：tail connects to node index 0
解释：链表中有一个环，其尾部连接到第一个节点。
示例 3：输入：head = [1], pos = -1 输出：no cycle
解释：链表中没有环。
进阶：你是否可以不用额外空间解决此题？
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 哈希辅助 | O(n)       | O(n)       |
| 02   | 快慢指针 | O(n)       | O(1)       |
| 03   | 遍历标记 | O(n)       | O(1)       |

```go
func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]bool)
	for head != nil {
		if m[head] {
			return head
		}
		m[head] = true
		head = head.Next
	}
	return nil
}

# 2
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	slow = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

# 3
func detectCycle(head *ListNode) *ListNode {
	for head != nil {
		if head.Val == math.MaxInt32 {
			return head
		}
		head.Val = math.MaxInt32
		head = head.Next
	}
	return head
}
```

## 面试题03.01.三合一

### 题目

```

```

### 解题思路

```go

```

## 面试题03.02.栈的最小值(2)

- 题目

```
请设计一个栈，除了常规栈支持的pop与push函数以外，还支持min函数，该函数返回栈元素中的最小值。
执行push、pop和min操作的时间复杂度必须为O(1)。
示例：MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.
```

- 解题思路

| No.  | 思路   | 时间复杂度 | 空间复杂度 |
| ---- | ------ | ---------- | ---------- |
| 01   | 结构体 | O(n)       | O(n)       |
| 02   | 双栈   | O(n)       | O(n)       |

```go
type item struct {
	min, x int
}

type MinStack struct {
	stack []item
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	min := x
	if len(this.stack) > 0 && this.GetMin() < x {
		min = this.GetMin()
	}
	this.stack = append(this.stack, item{
		min: min,
		x:   x,
	})
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1].x
}

func (this *MinStack) GetMin() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1].min
}

# 2
type MinStack struct {
	data []int
	min  []int
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{}}
}

func (this *MinStack) Push(x int) {
	if len(this.data) == 0 || x <= this.GetMin() {
		this.min = append(this.min, x)
	}
	this.data = append(this.data, x)
}

func (this *MinStack) Pop() {
	x := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	if x == this.GetMin() {
		this.min = this.min[:len(this.min)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.data) == 0 {
		return 0
	}
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}
```

## 面试题04.04.检查平衡性(3)

- 题目

```
实现一个函数，检查二叉树是否平衡。在这个问题中，平衡树的定义如下：
任意一个节点，其两棵子树的高度差不超过 1。
示例 1:给定二叉树 [3,9,20,null,null,15,7]
    3
   / \
  9  20
    /  \
   15   7
返回 true 。
示例 2:给定二叉树 [1,2,2,3,3,null,null,4,4]
      1
     / \
    2   2
   / \
  3   3
 / \
4   4
返回 false 。
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 递归 | O(n)       | O(log(n))  |
| 02   | 递归 | O(n)       | O(log(n))  |
| 03   | 递归 | O(n)       | O(log(n))  |

```go
func isBalanced(root *TreeNode) bool {
	_, isBalanced := dfs(root)
	return isBalanced
}

func dfs(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	leftDepth, leftIsBalanced := dfs(root.Left)
	if leftIsBalanced == false {
		return 0, false
	}
	rightDepth, rightIsBalanced := dfs(root.Right)
	if rightIsBalanced == false {
		return 0, false
	}

	if -1 <= leftDepth-rightDepth &&
		leftDepth-rightDepth <= 1 {
		return max(leftDepth, rightDepth) + 1, true
	}
	return 0, false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

# 2
func isBalanced(root *TreeNode) bool {
	return dfs(root) != -1
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	if left != -1 && right != -1 &&
		abs(left, right) <= 1 {
		return max(left, right) + 1
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

# 3
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if math.Abs(dfs(root.Left)-dfs(root.Right)) <= 1 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
	return false
}

func dfs(root *TreeNode) float64 {
	if root == nil {
		return 0
	}
	return math.Max(dfs(root.Left), dfs(root.Right)) + 1
}
```

## 面试题04.05.合法二叉搜索树(5)

- 题目

```
实现一个函数，检查一棵二叉树是否为二叉搜索树。
示例 1:输入:
    2
   / \
  1   3
输出: true
示例 2:输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。根节点的值为 5 ，但是其右子节点值为 4 。
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 递归 | O(n)       | O(log(n))  |
| 02   | 递归 | O(n)       | O(n)       |
| 03   | 迭代 | O(n)       | O(n)       |
| 04   | 迭代 | O(n)       | O(n)       |
| 05   | 递归 | O(n)       | O(log(n))  |

```go
func isValidBST(root *TreeNode) bool {
	return dfs(root, math.MinInt64, math.MaxInt64)
}

func dfs(root *TreeNode, left, right int) bool {
	if root == nil {
		return true
	}
	if left >= root.Val || right <= root.Val {
		return false
	}
	return dfs(root.Left, left, root.Val) && dfs(root.Right, root.Val, right)
}

# 2
var res []int

func isValidBST(root *TreeNode) bool {
	res = make([]int, 0)
	dfs(root)
	for i := 0; i < len(res)-1; i++ {
		if res[i] >= res[i+1] {
			return false
		}
	}
	return true
}

func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Left)
		res = append(res, root.Val)
		dfs(root.Right)
	}
}

# 3
func isValidBST(root *TreeNode) bool {
    if root == nil {
        return true
    }
    stack := make([]*TreeNode, 0)
    res := make([]int, 0)
    for len(stack) > 0 || root != nil {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        last := len(stack) - 1
        res = append(res, stack[last].Val)
        root = stack[last].Right
        stack = stack[:last]
    }
    for i := 0; i < len(res)-1; i++ {
        if res[i] >= res[i+1] {
            return false
        }
    }
    return true
}

# 4
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	stack := make([]*TreeNode, 0)
	pre := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		last := len(stack) - 1
		if stack[last].Val <= pre {
			return false
		}
		pre = stack[last].Val
		root = stack[last].Right
		stack = stack[:last]
	}
	return true
}

# 5
var pre int

func isValidBST(root *TreeNode) bool {
	pre = math.MinInt64
	return dfs(root)
}

func dfs(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if dfs(root.Left) == false {
		return false
	}
	if root.Val <= pre {
		return false
	}
	pre = root.Val
	return dfs(root.Right)
}
```

## 面试题04.08.首个共同祖先(2)

- 题目

```
设计并实现一个算法，找出二叉树中某两个节点的第一个共同祖先。不得将其他的节点存储在另外的数据结构中。
注意：这不一定是二叉搜索树。
例如，给定如下二叉树: root = [3,5,1,6,2,0,8,null,null,7,4]
    3
   / \
  5   1
 / \ / \
6  2 0  8
  / \
 7   4
示例 1:输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1输出: 3
解释: 节点 5 和节点 1 的最近公共祖先是节点 3。
示例 2:输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4 输出: 5
解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。
说明:所有节点的值都是唯一的。p、q 为不同节点且均存在于给定的二叉树中。
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 递归 | O(n)       | O(log(n))  |
| 02   | 递归 | O(n)       | O(n)       |

```go
func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

# 2
func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	m = make(map[int]*TreeNode)
	dfs(root)
	visited := make(map[int]bool)
	for p != nil {
		visited[p.Val] = true
		p = m[p.Val]
	}
	for q != nil {
		if visited[q.Val] == true {
			return q
		}
		q = m[q.Val]
	}
	return nil
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		m[root.Left.Val] = root
		dfs(root.Left)
	}
	if root.Right != nil {
		m[root.Right.Val] = root
		dfs(root.Right)
	}
}
```

## 面试题04.10.检查子树(2)

- 题目

```go
检查子树。你有两棵非常大的二叉树：T1，有几万个节点；T2，有几万个节点。
设计一个算法，判断 T2 是否为 T1 的子树。
如果 T1 有这么一个节点 n，其子树与 T2 一模一样，则 T2 为 T1 的子树，
也就是说，从节点 n 处把树砍断，得到的树与 T2 完全相同。
示例1:输入：t1 = [1, 2, 3], t2 = [2] 输出：true
示例2:输入：t1 = [1, null, 2, 4], t2 = [3, 2] 输出：false
提示：树的节点数目范围为[0, 20000]。
```

- 解题思路

| No.  | 思路            | 时间复杂度 | 空间复杂度 |
| ---- | --------------- | ---------- | ---------- |
| 01   | 递归            | O(n^2)     | O(log(n))  |
| 02   | 递归+字符串辅助 | O(n)       | O(log(n))  |
| 03   | 栈辅助(超时)    | O(n)       | O(n)       |

```go
func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil {
		return false
	}
	return isSame(t1, t2) || checkSubTree(t1.Left, t2) || checkSubTree(t1.Right, t2)
}

func isSame(s *TreeNode, t *TreeNode) bool {
	if s == nil || t == nil {
		return t == s
	}
	return isSame(s.Left, t.Left) && isSame(s.Right, t.Right) && s.Val == t.Val
}

# 2
func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	sStr := dfs(t1, "")
	tStr := dfs(t2, "")
	return strings.Contains(sStr, tStr)
}

func dfs(s *TreeNode, pre string) string {
	if s == nil {
		return pre
	}
	return fmt.Sprintf("#%d%s%s", s.Val, dfs(s.Left, "l"), dfs(s.Right, "r"))
}

# 3
func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	sStr := preOrder(t1)
	tStr := preOrder(t2)
	return strings.Contains(sStr, tStr)
}

func preOrder(root *TreeNode) string {
	if root == nil {
		return ""
	}
	res := "!"
	stack := make([]*TreeNode, 0)
	temp := root
	for {
		for temp != nil {
			res += strconv.Itoa(temp.Val)
			res += "!"
			stack = append(stack, temp)
			temp = temp.Left
		}
		res += "#!"
		if len(stack) > 0 {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			temp = node.Right
		} else {
			break
		}
	}
	return res
}
```

## 面试题04.12.求和路径(4)

- 题目

```
给定一棵二叉树，其中每个节点都含有一个整数数值(该值或正或负)。设计一个算法，
打印节点数值总和等于某个给定值的所有路径的数量。
注意，路径不一定非得从二叉树的根节点或叶节点开始或结束，但是其方向必须向下(只能从父节点指向子节点方向)。
示例:给定如下二叉树，以及目标和 sum = 22，
              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:3
解释：和为 22 的路径有：[5,4,11,2], [5,8,4,5], [4,11,7]
提示：节点总数 <= 10000
```

- 解题思路

| No.  | 思路      | 时间复杂度 | 空间复杂度 |
| ---- | --------- | ---------- | ---------- |
| 01   | 递归      | O(n^2)     | O(n)       |
| 02   | 递归      | O(n^2)     | O(n)       |
| 03   | 迭代+递归 | O(n^2)     | O(n)       |
| 04   | 递归+路径 | O(n^2)     | O(n)       |

```go
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	res := 0
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		sum = sum - node.Val
		// 路径不需要从根节点开始，也不需要在叶子节点结束
		if sum == 0 {
			res++
		}
		dfs(node.Left, sum)
		dfs(node.Right, sum)
	}
	dfs(root, sum)
	return res + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

# 2
func dfs(node *TreeNode, sum int) int {
	if node == nil {
		return 0
	}
	sum = sum - node.Val
	res := 0
	if sum == 0 {
		res = 1
	}
	return res + dfs(node.Left, sum) + dfs(node.Right, sum)
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return dfs(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

# 3
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	res := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		tempSum := 0
		res += dfs(node, sum, tempSum)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return res
}

func dfs(node *TreeNode, sum int, curSum int) int {
	res := 0
	curSum = curSum + node.Val
	if curSum == sum {
		res++
	}
	if node.Left != nil {
		res += dfs(node.Left, sum, curSum)
	}
	if node.Right != nil {
		res += dfs(node.Right, sum, curSum)
	}
	return res
}

# 4
func pathSum(root *TreeNode, sum int) int {
	return dfs(root, sum, make([]int, 1001), 0)
}

func dfs(node *TreeNode, sum int, path []int, level int) int {
	if node == nil {
		return 0
	}
	res := 0
	if sum == node.Val {
		res = 1
	}
	temp := node.Val
	for i := level - 1; i >= 0; i-- {
		temp = temp + path[i]
		if temp == sum {
			res++
		}
	}
	path[level] = node.Val
	return res + dfs(node.Left, sum, path, level+1) +
		dfs(node.Right, sum, path, level+1)
}
```

## 面试题05.03.翻转数位

### 题目

```
给定一个32位整数 num，你可以将一个数位从0变为1。请编写一个程序，找出你能够获得的最长的一串1的长度。
示例 1：输入: num = 1775(110111011112)输出: 8
示例 2：输入: num = 7(01112)输出: 4
```

### 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 内置函数 | O(1)       | O(1)       |

```go

```

## 面试题05.06.整数转换(4)

- 题目

```
整数转换。编写一个函数，确定需要改变几个位才能将整数A转成整数B。
示例1:输入：A = 29 （或者0b11101）, B = 15（或者0b01111）输出：2
示例2:输入：A = 1，B = 2 输出：2
提示: A，B范围在[-2147483648, 2147483647]之间
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 内置函数 | O(1)       | O(1)       |
| 02   | 位运算   | O(1)       | O(1)       |
| 03   | 位运算   | O(1)       | O(1)       |
| 04   | 位运算   | O(1)       | O(1)       |

```go
func convertInteger(A int, B int) int {
	C := uint32(A) ^ uint32(B)
	return bits.OnesCount(uint(C))
}

# 2
func convertInteger(A int, B int) int {
	C := uint32(A) ^ uint32(B)
	res := 0
	for C != 0 {
		if C&1 == 1 {
			res++
		}
		C = C >> 1
	}
	return res
}

# 3
func convertInteger(A int, B int) int {
	C := uint32(A) ^ uint32(B)
	res := 0
	for C != 0 {
		res++
		C = C & (C - 1)
	}
	return res
}

# 4
func convertInteger(A int, B int) int {
	C := A ^ B
	res := 0
	for i := 0; i < 32; i++{
		if C & 1 ==1{
			res++
		}
		C = C >> 1
	}
	return res
}
```



## 面试题08.01.三步问题

### 题目

```
三步问题。有个小孩正在上楼梯，楼梯有n阶台阶，小孩一次可以上1阶、2阶或3阶。
实现一种方法，计算小孩有多少种上楼梯的方式。结果可能很大，你需要对结果模1000000007。
示例1:输入：n = 3 输出：4
说明: 有四种走法
示例2:输入：n = 5输出：13
提示:n范围在[1, 1000000]之间
```

### 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 遍历 | O(n)       | O(1)       |

```go

```

## 面试题08.03.魔术索引(2)

- 题目

```
魔术索引。 在数组A[0...n-1]中，有所谓的魔术索引，满足条件A[i] = i。
给定一个有序整数数组，编写一种方法找出魔术索引，若有的话，在数组A中找出一个魔术索引，如果没有，则返回-1。
若有多个魔术索引，返回索引值最小的一个。
示例1:输入：nums = [0, 2, 3, 4, 5] 输出：0 说明: 0下标的元素为0
示例2:输入：nums = [1, 1, 1] 输出：1
说明:
    nums长度在[1, 1000000]之间
    此题为原书中的 Follow-up，即数组中可能包含重复元素的版本
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 遍历 | O(n)       | O(1)       |
| 02   | 递归 | O(n)       | O(n)       |

```go
func findMagicIndex(nums []int) int {
	for i := 0; i < len(nums); i++{
		if nums[i] == i{
			return i
		}
	}
	return -1
}

# 2
func findMagicIndex(nums []int) int {
	return search(nums, 0, len(nums)-1)
}

func search(nums []int, left, right int) int {
	if left > right {
		return -1
	}
	mid := left + (right-left)/2
	res := search(nums, left, mid-1)
	if res != -1 {
		return res
	} else if nums[mid] == mid {
		return mid
	}
	return search(nums, mid+1, right)
}
```

## 面试题08.07.无重复字符串的排列组合(3)

- 题目

```
无重复字符串的排列组合。编写一种方法，计算某字符串的所有排列组合，字符串每个字符均不相同。
示例1:输入：S = "qwe"输出：["qwe", "qew", "wqe", "weq", "ewq", "eqw"]
示例2:输入：S = "ab"输出：["ab", "ba"]
提示:字符都是英文字母。
    字符串长度在[1, 9]之间。
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 回溯 | O(n^n)     | O(n*n!)    |
| 02   | 递归 | O(n!)      | O(n*n!)    |
| 03   | 回溯 | O(n!)      | O(n*n!)    |

```go
var res []string

func permutation(S string) []string {
	res = make([]string, 0)
	nums := []byte(S)
	visited := make(map[int]bool)
	dfs(nums, 0, "", visited)
	return res
}

func dfs(nums []byte, index int, str string, visited map[int]bool) {
	if index == len(nums) {
		res = append(res, str)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] == false {
			str = str + string(nums[i])
			visited[i] = true
			dfs(nums, index+1, str, visited)
			str = str[:len(str)-1]
			visited[i] = false
		}
	}
}

# 2
func permutation(S string) []string {
	if len(S) == 1 {
		return []string{S}
	}
	res := make([]string, 0)
	for i := 0; i < len(S); i++ {
		str := S[:i] + S[i+1:]
		arr := permutation(str)
		for _, v := range arr {
			res = append(res, v+string(S[i]))
		}
	}
	return res
}

# 3
var res []string

func permutation(S string) []string {
	res = make([]string, 0)
	nums := []byte(S)
	dfs(nums, 0, "")
	return res
}

func dfs(nums []byte, index int, str string) {
	if index == len(nums) {
		res = append(res, str)
		return
	}
	for i := index; i < len(nums); i++ {
		str = str + string(nums[i])
		nums[i], nums[index] = nums[index], nums[i]
		dfs(nums, index+1, str)
		nums[i], nums[index] = nums[index], nums[i]
		str = str[:len(str)-1]
	}
}
```

## 面试题08.08.有重复字符串的排列组合(3)

- 题目

```
有重复字符串的排列组合。编写一种方法，计算某字符串的所有排列组合。
示例1:输入：S = "qqe" 输出：["eqq","qeq","qqe"]
示例2:输入：S = "ab"输出：["ab", "ba"]
提示:
    字符都是英文字母。
    字符串长度在[1, 9]之间。
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 回溯 | O(n!)      | O(n*n!)    |
| 02   | 回溯 | O(n!)      | O(n*n!)    |
| 03   | 回溯 | O(n!)      | O(n*n!)    |

```go
var res []string

func permutation(S string) []string {
	res = make([]string, 0)
	nums := []byte(S)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	dfs(nums, 0, make([]int, len(nums)), "")
	return res
}

func dfs(nums []byte, index int, visited []int, str string) {
	if len(nums) == index {
		res = append(res, str)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] == 1 {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && visited[i-1] == 0 {
			continue
		}
		str = str + string(nums[i])
		visited[i] = 1
		dfs(nums, index+1, visited, str)
		visited[i] = 0
		str = str[:len(str)-1]
	}
}

# 2
var res []string

func permutation(S string) []string {
	res = make([]string, 0)
	nums := []byte(S)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	dfs(nums, 0)
	return res
}

func dfs(nums []byte, index int) {
	if index == len(nums) {
		res = append(res, string(nums))
		return
	}
	m := make(map[byte]int)
	for i := index; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			continue
		}
		m[nums[i]] = 1
		nums[i], nums[index] = nums[index], nums[i]
		dfs(nums, index+1)
		nums[i], nums[index] = nums[index], nums[i]
	}
}

# 3
var res []string

func permutation(S string) []string {
	res = make([]string, 0)
	nums := []byte(S)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	dfs(nums, "")
	return res
}

func dfs(nums []byte, str string) {
	if len(nums) == 0 {
		res = append(res, str)
		return
	}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		str = str + string(nums[i])
		arr := append([]byte{}, nums[:i]...)
		arr = append(arr, nums[i+1:]...)
		dfs(arr, str)
		str = str[:len(str)-1]
	}
}
```



## 面试题08.09.括号(3)

- 题目

```
括号。设计一种算法，打印n对括号的所有合法的（例如，开闭一一对应）组合。
说明：解集不能包含重复的子集。
例如，给出 n = 3，生成结果为：
[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
```

- 解题思路

| No.  | 思路         | 时间复杂度     | 空间复杂度     |
| ---- | ------------ | -------------- | -------------- |
| 01   | 全排列-递归  | O(4^n/n^(1/2)) | O(4^n/n^(1/2)) |
| 02   | 动态规划     | O(4^n/n^(1/2)) | O(4^n/n^(1/2)) |
| 03   | 广度优先搜索 | O(4^n/n^(1/2)) | O(4^n/n^(1/2)) |

```go
var res []string

func generateParenthesis(n int) []string {
	res = make([]string, 0)
	dfs(0, 0, n, "")
	return res
}

func dfs(left, right, max int, str string) {
	if left == right && left == max {
		res = append(res, str)
		return
	}
	if left < max {
		dfs(left+1, right, max, str+"(")
	}
	if right < left {
		dfs(left, right+1, max, str+")")
	}
}

# 2
/*
dp[i]表示n=i时括号的组合
dp[i]="(" + dp[j] + ")"+dp[i-j-1] (j<i)
dp[0] = ""
*/
func generateParenthesis(n int) []string {
	dp := make([][]string, n+1)
	dp[0] = make([]string, 0)
	if n == 0 {
		return dp[0]
	}
	dp[0] = append(dp[0], "")
	for i := 1; i <= n; i++ {
		dp[i] = make([]string, 0)
		for j := 0; j < i; j++ {
			for _, a := range dp[j] {
				for _, b := range dp[i-j-1] {
					str := "(" + a + ")" + b
					dp[i] = append(dp[i], str)
				}
			}
		}
	}
	return dp[n]
}

# 3
type Node struct {
	str   string
	left  int
	right int
}

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	if n == 0 {
		return res
	}
	queue := make([]*Node, 0)
	queue = append(queue, &Node{
		str:   "",
		left:  n,
		right: n,
	})
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.left == 0 && node.right == 0 {
			res = append(res, node.str)
		}
		if node.left > 0 {
			queue = append(queue, &Node{
				str:   node.str + "(",
				left:  node.left - 1,
				right: node.right,
			})
		}
		if node.right > 0 && node.left < node.right {
			queue = append(queue, &Node{
				str:   node.str + ")",
				left:  node.left,
				right: node.right - 1,
			})
		}
	}
	return res
}
```



## 面试题08.10.颜色填充(2)

- 题目

```
编写函数，实现许多图片编辑软件都支持的「颜色填充」功能。
待填充的图像用二维数组 image 表示，元素为初始颜色值。初始坐标点的横坐标为 sr 纵坐标为 sc。
需要填充的新颜色为 newColor 。
「周围区域」是指颜色相同且在上、下、左、右四个方向上存在相连情况的若干元素。
请用新颜色填充初始坐标点的周围区域，并返回填充后的图像。
示例：输入： image = [[1,1,1],[1,1,0],[1,0,1]]  sr = 1, sc = 1, newColor = 2
输出：[[2,2,2],[2,2,0],[2,0,1]]
解释: 初始坐标点位于图像的正中间，坐标 (sr,sc)=(1,1) 。
初始坐标点周围区域上所有符合条件的像素点的颜色都被更改成 2 。
注意，右下角的像素没有更改为 2 ，因为它不属于初始坐标点的周围区域。
提示：
    image 和 image[0] 的长度均在范围 [1, 50] 内。
    初始坐标点 (sr,sc) 满足 0 <= sr < image.length 和 0 <= sc < image[0].length 。
    image[i][j] 和 newColor 表示的颜色值在范围 [0, 65535] 内。
```

- 解题思路

| No.  | 思路         | 时间复杂度 | 空间复杂度 |
| ---- | ------------ | ---------- | ---------- |
| 01   | 广度优先搜索 | O(n^2)     | O(n^2)     |
| 02   | 深度优先搜索 | O(n^2)     | O(n^2)     |

```go
var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := image[sr][sc]
	if oldColor == newColor {
		return image
	}
	m, n := len(image), len(image[0])
	list := make([][]int, 1)
	list[0] = []int{sr, sc}

	for len(list) > 0 {
		node := list[0]
		list = list[1:]
		image[node[0]][node[1]] = newColor
		for i := 0; i < 4; i++ {
			x := node[0] + dx[i]
			y := node[1] + dy[i]
			if 0 <= x && x < m && 0 <= y && y < n &&
				image[x][y] == oldColor {
				list = append(list, []int{x, y})
			}
		}
	}
	return image
}

# 2
var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if sr < 0 || sc < 0 || sr >= len(image) ||
		sc >= len(image[sr]) || image[sr][sc] == newColor {
		return image
	}
	oldColor := image[sr][sc]
	image[sr][sc] = newColor
	for i := 0; i < 4; i++ {
		x := sr + dx[i]
		y := sc + dy[i]
		if 0 <= x && x < len(image) && 0 <= y && y < len(image[x]) &&
			image[x][y] == oldColor {
			floodFill(image, x, y, newColor)
		}
	}
	return image
}
```

## 面试题10.01.合并排序的数组(3)

- 题目

```
给定两个排序后的数组 A 和 B，其中 A 的末端有足够的缓冲空间容纳 B。 编写一个方法，将 B 合并入 A 并排序。
初始化 A 和 B 的元素数量分别为 m 和 n。
示例:输入:
A = [1,2,3,0,0,0], m = 3
B = [2,5,6],       n = 3
输出: [1,2,2,3,5,6]
说明:A.length == n + m
```

- 解题思路

| No.  | 思路       | 时间复杂度 | 空间复杂度 |
| ---- | ---------- | ---------- | ---------- |
| 01   | 合并后排序 | O(nlog(n)) | O(1)       |
| 02   | 双指针法   | O(n)       | O(1)       |
| 03   | 数组辅助   | O(n)       | O(n)       |

```go
func merge(A []int, m int, B []int, n int) {
	A = A[:m]
	A = append(A, B[:n]...)
	sort.Ints(A)
}

# 2
func merge(A []int, m int, B []int, n int) {
	for m > 0 && n > 0 {
		if A[m-1] < B[n-1] {
			A[m+n-1] = B[n-1]
			n--
		} else {
			A[m+n-1] = A[m-1]
			m--
		}
	}
	if m == 0 && n > 0 {
		for n > 0 {
			A[n-1] = B[n-1]
			n--
		}
	}
}

# 3
func merge(A []int, m int, B []int, n int) {
	temp := make([]int, m)
	copy(temp, A)

	if n == 0 {
		return
	}
	first, second := 0, 0
	for i := 0; i < len(A); i++ {
		if second >= n {
			A[i] = temp[first]
			first++
			continue
		}
		if first >= m {
			A[i] = B[second]
			second++
			continue
		}
		if temp[first] < B[second] {
			A[i] = temp[first]
			first++
		} else {
			A[i] = B[second]
			second++
		}
	}
}
```

## 面试题10.03.搜索旋转数组(2)

- 题目

```
搜索旋转数组。给定一个排序后的数组，包含n个整数，但这个数组已被旋转过很多次了，次数不详。
请编写代码找出数组中的某个元素，假设数组元素原先是按升序排列的。若有多个相同元素，返回索引值最小的一个。
示例1: 输入: arr = [15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14], target = 5
输出: 8（元素5在该数组中的索引）
示例2:输入：arr = [15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14], target = 11
输出：-1 （没有找到）
提示:arr 长度范围在[1, 1000000]之间
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 二分查找 | O(log(n))  | O(1)       |
| 02   | 遍历     | O(n)       | O(1)       |

```go
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[left] < nums[mid] { // 左边升序的情况
			if nums[left] <= target && target <= nums[mid] {
				right = mid
			} else {
				left = mid + 1
			}
		} else if nums[left] > nums[mid] { // 右边升序
			if nums[mid] < target && target <= nums[right] && nums[left] > nums[right] {
				left = mid + 1
			} else {
				right = mid
			}
		} else if nums[left] == nums[mid] {
			if nums[left] != target {
				left++
			} else {
				return left
			}
		}
	}
	if nums[left] == target {
		return left
	}
	return -1
}

#
func search(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if target == nums[i] {
			return i
		}
	}
	return -1
}
```

## 面试题10.09.排序矩阵查找(6)

- 题目

```
给定M×N矩阵，每一行、每一列都按升序排列，请编写代码找出某元素。
示例:现有矩阵 matrix 如下：
[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
给定 target = 5，返回 true。
给定 target = 20，返回 false。
```

- 解题思路

| No.  | 思路        | 时间复杂度 | 空间复杂度 |
| ---- | ----------- | ---------- | ---------- |
| 01   | 暴力法      | O(n^2)     | O(1)       |
| 02   | 暴力法-优化 | O(n^2)     | O(1)       |
| 03   | 二分查找    | O(nlog(n)) | O(1)       |
| 04   | 左下角查找  | O(n)       | O(1)       |
| 05   | 右上角查找  | O(n)       | O(1)       |
| 06   | 内置函数    | O(n^2)     | O(1)       |

```go
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	if len(matrix[0]) == 0 {
		return false
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}

# 2
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	if len(matrix[0]) == 0 {
		return false
	}
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] <= target && matrix[i][len(matrix[i])-1] >= target {
			for j := 0; j < len(matrix[i]); j++ {
				if matrix[i][j] == target {
					return true
				}
			}
		}
	}
	return false
}

# 3
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	if len(matrix[0]) == 0 {
		return false
	}
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] <= target && matrix[i][len(matrix[i])-1] >= target {
			res := binarySearch(matrix[i], target)
			if res == true {
				return true
			}
		}
	}
	return false
}

func binarySearch(arr []int, target int) bool {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return true
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

# 4
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	if len(matrix[0]) == 0 {
		return false
	}
	i := len(matrix) - 1
	j := 0
	for i >= 0 && j < len(matrix[0]) {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			i--
		} else {
			j++
		}
	}
	return false
}

# 5
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	if len(matrix[0]) == 0 {
		return false
	}
	i := 0
	j := len(matrix[0]) - 1
	for j >= 0 && i < len(matrix) {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return false
}

# 6
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	if len(matrix[0]) == 0 {
		return false
	}
	for i := 0; i < len(matrix); i++ {
		index := sort.SearchInts(matrix[i], target)
		if index < len(matrix[i]) && target == matrix[i][index] {
			return true
		}
	}
	return false
}
```

## 面试题16.05.阶乘尾数(1)

- 题目

```
设计一个算法，算出 n 阶乘有多少个尾随零。
示例 1:输入: 3 输出: 0 解释: 3! = 6, 尾数中没有零。
示例 2:输入: 5输出: 1 解释: 5! = 120, 尾数中有 1 个零.
说明: 你算法的时间复杂度应为 O(log n) 。
```

- 解题思路

| No.  | 思路         | 时间复杂度 | 空间复杂度 |
| ---- | ------------ | ---------- | ---------- |
| 01   | 数学，找规律 | O(log(n))  | O(1)       |

```go
// N!有多少个后缀0，即N!有多少个质因数5。
// N!有多少个质因数5，即N可以划分成多少组5个数字一组，
// 加上划分成多少组25个数字一组，加上划分多少组成125个数字一组，等等
// Ans = N/5 + N/(5^2) + N/(5^3) + ...
func trailingZeroes(n int) int {
	result := 0
	for n >= 5 {
		n = n / 5
		result = result + n
	}
	return result
}
```

## 面试题16.11.跳水板(2)

- 题目

```
你正在使用一堆木板建造跳水板。有两种类型的木板，其中长度较短的木板长度为shorter，
长度较长的木板长度为longer。你必须正好使用k块木板。编写一个方法，生成跳水板所有可能的长度。
返回的长度需要从小到大排列。
示例 1 输入：shorter = 1 longer = 2 k = 3 输出： [3,4,5,6]
解释：可以使用 3 次 shorter，得到结果 3；使用 2 次 shorter 和 1 次 longer，得到结果 4 。
以此类推，得到最终结果。
提示：
    0 < shorter <= longer
    0 <= k <= 100000
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 遍历 | O(n)       | O(n)       |
| 02   | 遍历 | O(n)       | O(n)       |

```go
func divingBoard(shorter int, longer int, k int) []int {
	res := make([]int, 0)
	if k == 0 {
		return res
	}
	if shorter == longer {
		return []int{shorter * k}
	}
	for i := 0; i <= k; i++ {
		res = append(res, shorter*(k-i)+longer*i)
	}
	return res
}

#
func divingBoard(shorter int, longer int, k int) []int {
	res := make([]int, 0)
	if k == 0 {
		return res
	}
	if shorter == longer {
		return []int{shorter * k}
	}
	start := shorter * k
	diff := longer - shorter
	for i := 0; i <= k; i++ {
		res = append(res, start+i*diff)
	}
	return res
}
```

## 面试题16.17.连续数列(5)

- 题目

```
给定一个整数数组，找出总和最大的连续数列，并返回总和。
示例：输入： [-2,1,-3,4,-1,2,1,-5,4]输出： 6
解释： 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 贪心法   | O(n)       | O(1)       |
| 02   | 暴力法   | O(n^2)     | O(1)       |
| 03   | 动态规划 | O(n)       | O(n)       |
| 04   | 动态规划 | O(n)       | O(1)       |
| 05   | 分治     | O(nlog(n)) | O(log(n))  |

```go
func maxSubArray(nums []int) int {
	result := nums[0]
	sum := 0
	for i := 0; i < len(nums); i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		if sum > result {
			result = sum
		}
	}
	return result
}

# 2
func maxSubArray(nums []int) int {
	result := math.MinInt32
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum > result {
				result = sum
			}
		}
	}
	return result
}

# 3
// dp[i] = max(dp[i-1]+nums[i], nums[i])
// res = max(dp[i], res)
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1]+nums[i] > nums[i] {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > result {
			result = dp[i]
		}
	}
	return result
}

# 4
func maxSubArray(nums []int) int {
	dp := nums[0]
	result := dp
	for i := 1; i < len(nums); i++ {
		if dp+nums[i] > nums[i] {
			dp = dp + nums[i]
		} else {
			dp = nums[i]
		}

		if dp > result {
			result = dp
		}
	}
	return result
}

# 5
func maxSubArray(nums []int) int {
	result := maxSubArr(nums, 0, len(nums)-1)
	return result
}

func maxSubArr(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}

	mid := (left + right) / 2
	leftSum := maxSubArr(nums, left, mid)        // 最大子序在左边
	rightSum := maxSubArr(nums, mid+1, right)    // 最大子序在右边
	midSum := findMaxArr(nums, left, mid, right) // 跨中心
	result := max(leftSum, rightSum)
	result = max(result, midSum)
	return result
}

func findMaxArr(nums []int, left, mid, right int) int {
	leftSum := math.MinInt32
	sum := 0
	// 从右到左
	for i := mid; i >= left; i-- {
		sum += nums[i]
		leftSum = max(leftSum, sum)
	}
	rightSum := math.MinInt32
	sum = 0
	// 从左到右
	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		rightSum = max(rightSum, sum)
	}
	return leftSum + rightSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

## 面试题17.01.不用加号的加法(2)

- 题目

```
设计一个函数把两个数字相加。不得使用 + 或者其他算术运算符。
示例:输入: a = 1, b = 1 输出: 2
提示： a, b 均可能是负数或 0
    结果不会溢出 32 位整数
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 迭代 | O(1)       | O(1)       |
| 02   | 递归 | O(1)       | O(1)       |

```go
/*
0 + 0 = 0
0 + 1 = 1
1 + 0 = 1
1 + 1 = 0（进位 1）
异或的一个重要特性是无进位加法
(a 和 b 的无进位结果) + (a 和 b 的进位结果)
*/
func add(a int, b int) int {
	for b != 0 {
		a, b = a^b, (a&b)<<1
	}
	return a
}

# 
func add(a int, b int) int {
	if b == 0 {
		return a
	}
	return add(a^b, (a&b)<<1)
}
```

## 面试题17.04.消失的数字(5)

- 题目

```
数组nums包含从0到n的所有整数，但其中缺了一个。请编写代码找出那个缺失的整数。你有办法在O(n)时间内完成吗？
注意：本题相对书上原题稍作改动
示例 1：输入：[3,0,1]输出：2
示例 2：输入：[9,6,4,2,3,5,7,0,1] 输出：8
```

- 解题思路

| No.  | 思路               | 时间复杂度 | 空间复杂度 |
| ---- | ------------------ | ---------- | ---------- |
| 01   | 数学计算           | O(n)       | O(1)       |
| 02   | 排序遍历           | O(nlog(n)) | O(1)       |
| 03   | 异或-位运算        | O(n)       | O(1)       |
| 04   | 交换排序(就地排序) | O(n)       | O(1)       |
| 05   | 哈希辅助           | O(n)       | O(n)       |

```go
func missingNumber(nums []int) int {
	n := len(nums)
	sum := n * (n + 1) / 2
	for i := 0; i < n; i++ {
		sum = sum - nums[i]
	}
	return sum
}

# 2
func missingNumber(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	return len(nums)
}

# 3
func missingNumber(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res = res ^ (i+1) ^ nums[i]
	}
	return res
}

# 4
func missingNumber(nums []int) int {
	n := len(nums)
	index := n
	for i := 0; i < n; {
		if nums[i] == n{
			index = i
			i++
			continue
		}
		if i == nums[i]{
			i++
			continue
		}
		nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
	}
	return index
}

# 5
func missingNumber(nums []int) int {
	m := make(map[int]bool)
	for i := range nums {
		m[nums[i]] = true
	}
	for i := 0; i <= len(nums); i++ {
		if m[i] == false {
			return i
		}
	}
	return 0
}
```

## 面试题17.09.第k个数(1)

- 题目

```
有些数的素因子只有 3，5，7，请设计一个算法找出第 k 个数。
注意，不是必须有这些素因子，而是必须不包含其他的素因子。
例如，前几个数按顺序应该是 1，3，5，7，9，15，21。
示例 1:输入: k = 5输出: 9
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 动态规划 | O(n)       | O(n)       |

```go
func getKthMagicNumber(k int) int {
	dp := make([]int, k)
	dp[0] = 1
	// *3或5或7之后得到
	idx3, idx5, idx7 := 0, 0, 0
	for i := 1; i < k; i++ {
		dp[i] = min(dp[idx3]*3, min(dp[idx5]*5, dp[idx7]*7))
		if dp[i] == dp[idx3]*3 {
			idx3++
		}
		if dp[i] == dp[idx5]*5 {
			idx5++
		}
		if dp[i] == dp[idx7]*7 {
			idx7++
		}
	}
	return dp[k-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
```



## 面试题17.10.主要元素(5)

- 题目

```
数组中占比超过一半的元素称之为主要元素。给定一个整数数组，找到它的主要元素。若没有，返回-1。
示例 1：输入：[1,2,5,9,5,9,5,5,5]输出：5
示例 2：输入：[3,2]输出：-1
示例 3：输入：[2,2,1,1,1,2,2]输出：2
说明：你有办法在时间复杂度为 O(N)，空间复杂度为 O(1) 内完成吗？
```

- 解题思路

| No.  | 思路                | 时间复杂度 | 空间复杂度 |
| ---- | ------------------- | ---------- | ---------- |
| 01   | 哈希辅助            | O(n)       | O(n)       |
| 02   | Boyer-Moore投票算法 | O(n)       | O(1)       |
| 03   | 排序                | O(nlog(n)) | O(1)       |
| 04   | 位运算              | O(n)       | O(1)       |
| 05   | 分治法              | O(nlog(n)) | O(log(n))  |

```go
func majorityElement(nums []int) int {
	m := make(map[int]int)
	result := -1
	for _, v := range nums{
		if _,ok := m[v];ok{
			m[v]++
		}else {
			m[v]=1
		}
		if m[v] > (len(nums)/2){
			result = v
		}
	}
	return result
}

# 2
func majorityElement(nums []int) int {
	result, count := 0, 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			result = nums[i]
			count++
		} else if result == nums[i] {
			count++
		} else {
			count--
		}
	}
	total := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == result {
			total++
		}
	}
	if total <= len(nums)/2 {
		return -1
	}
	return result
}

# 3
func majorityElement(nums []int) int {
	sort.Ints(nums)
	for i := 0; i <= len(nums)/2; i++ {
		if nums[i] == nums[i+len(nums)/2] {
			return nums[i]
		}
	}
	return -1
}

# 4
func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	result := int32(0)
	mask := int32(1)
	for i := 0; i < 32; i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			if mask&int32(nums[j]) == mask {
				count++
			}
		}
		if count > len(nums)/2 {
			result = result | mask
		}
		mask = mask << 1
	}
	total := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == int(result) {
			total++
		}
	}
	if total <= len(nums)/2 {
		return -1
	}
	return int(result)
}

# 5
func majorityElement(nums []int) int {
	res := majority(nums, 0, len(nums)-1)
	total := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == res {
			total++
		}
	}
	if total <= len(nums)/2 {
		return -1
	}
	return res
}

func count(nums []int, target int, start int, end int) int {
	countNum := 0
	for i := start; i <= end; i++ {
		if nums[i] == target {
			countNum++
		}
	}
	return countNum
}

func majority(nums []int, start, end int) int {
	if start == end {
		return nums[start]
	}
	mid := (start + end) / 2
	left := majority(nums, start, mid)
	right := majority(nums, mid+1, end)
	if left == right {
		return left
	}
	leftCount := count(nums, left, start, end)
	rightCount := count(nums, right, start, end)
	if leftCount > rightCount {
		return left
	}
	return right
}
```

## 面试题17.16.按摩师(4)

- 题目

```
一个有名的按摩师会收到源源不断的预约请求，每个预约都可以选择接或不接。
在每次预约服务之间要有休息时间，因此她不能接受相邻的预约。
给定一个预约请求序列，替按摩师找到最优的预约集合（总预约时间最长），返回总的分钟数。
注意：本题相对原题稍作改动
示例 1：输入： [1,2,3,1] 输出： 4
解释： 选择 1 号预约和 3 号预约，总时长 = 1 + 3 = 4。
示例 2：输入： [2,7,9,3,1] 输出： 12
解释： 选择 1 号预约、 3 号预约和 5 号预约，总时长 = 2 + 9 + 1 = 12。
示例 3：输入： [2,1,4,5,3,1,1,3] 输出： 12
解释： 选择 1 号预约、 3 号预约、 5 号预约和 8 号预约，总时长 = 2 + 4 + 3 + 3 = 12。
```

- 解题思路

| No.  | 思路              | 时间复杂度 | 空间复杂度 |
| ---- | ----------------- | ---------- | ---------- |
| 01   | 动态规划          | O(n)       | O(1)       |
| 02   | 动态规划+一维数组 | O(n)       | O(n)       |
| 03   | 动态规划+二维数组 | O(n)       | O(n)       |
| 04   | 奇偶法            | O(n)       | O(1)       |

```go
func massage(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	a := nums[0]
	b := max(a, nums[1])

	for i := 2; i < len(nums); i++ {
		a, b = b, max(a+nums[i], b)
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

# 2
func massage(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	if nums[0] > nums[1] {
		dp[1] = nums[0]
	} else {
		dp[1] = nums[1]
	}
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

# 3
func massage(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	n := len(nums)
	dp := make([][]int, n)
	for n := range dp {
		dp[n] = make([]int, 2)
	}
	dp[0][0], dp[0][1] = 0, nums[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]
	}
	return max(dp[n-1][0], dp[n-1][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

# 4
func massage(nums []int) int {
	var a, b int
	for i, v := range nums {
		if i%2 == 0 {
			a = max(a+v, b)
		} else {
			b = max(a, b+v)
		}
	}
	return max(a, b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

## 面试题17.21.直方图的水量(4)

- 题目

```
给定一个直方图(也称柱状图)，假设有人从上面源源不断地倒水，最后直方图能存多少水量?直方图的宽度为 1。
上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的直方图，
在这种情况下，可以接 6 个单位的水（蓝色部分表示水）。 感谢 Marcos 贡献此图。
示例:输入: [0,1,0,2,1,0,1,3,2,1,2,1] 输出: 6
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 暴力法   | O(n^2)     | O(1)       |
| 02   | 数组辅助 | O(n)       | O(n)       |
| 03   | 栈辅助   | O(n)       | O(n)       |
| 04   | 双指针   | O(n)       | O(1)       |

```go
func trap(height []int) int {
	res := 0
	for i := 0; i < len(height); i++ {
		left, right := 0, 0
		for j := i; j >= 0; j-- {
			left = max(left, height[j])
		}
		for j := i; j < len(height); j++ {
			right = max(right, height[j])
		}
		// 当前坐标形成的面积=(min(左边最高，右边最高)-当前高度) * 宽度(1,可省略)
		area := min(left, right) - height[i]
		res = res + area
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

# 2
func trap(height []int) int {
	res := 0
	if len(height) == 0 {
		return 0
	}
	left := make([]int, len(height))
	right := make([]int, len(height))
	left[0] = height[0]
	right[len(right)-1] = height[len(height)-1]
	for i := 1; i < len(height); i++ {
		left[i] = max(height[i], left[i-1])
	}
	for i := len(height) - 2; i >= 0; i-- {
		right[i] = max(height[i], right[i+1])
	}
	for i := 0; i < len(height); i++ {
		// 当前坐标形成的面积=(min(左边最高，右边最高)-当前高度) * 宽度(1,可省略)
		area := min(left[i], right[i]) - height[i]
		res = res + area
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

# 3
func trap(height []int) int {
	res := 0
	stack := make([]int, 0)
	for i := 0; i < len(height); i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			bottom := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				prev := stack[len(stack)-1]
				// 横着的面积=长(min(height[i], height[prev])-bottom)*宽(i-prev-1)
				h := min(height[i], height[prev]) - bottom
				w := i - prev - 1
				area := h * w
				res = res + area
			}
		}
		stack = append(stack, i)
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

# 4
func trap(height []int) int {
	res := 0
	if len(height) == 0 {
		return 0
	}
	left := 0
	right := len(height) - 1
	leftMax := 0  // 左边的最大值
	rightMax := 0 // 右边的最大值
	for left < right {
		// 当前坐标形成的面积=(min(左边最高，右边最高)-当前高度) * 宽度(1,可省略)
		// 选择高度低的一边处理并求最大值, 说明当前侧最大值小于另一侧
		if height[left] < height[right] {
			// 也可以写成这样
			// leftMax = max(leftMax, height[left])
			// res = res + leftMax - height[left]
			if height[left] >= leftMax { // 递增无法蓄水
				leftMax = height[left]
			} else {
				res = res + leftMax - height[left]
			}
			left++
		} else {
			// 也可以写成这样
			// rightMax = max(rightMax, height[right])
			// res = res + rightMax - height[right]
			if height[right] >= rightMax { // 递减无法蓄水
				rightMax = height[right]
			} else {
				res = res + rightMax - height[right]
			}
			right--
		}
	}
	return res
}
```

