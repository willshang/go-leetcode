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

## 面试题01.02.判定是否互为字符重排

### 题目

```
给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。
示例 1：输入: s1 = "abc", s2 = "bca" 输出: true 
示例 2：输入: s1 = "abc", s2 = "bad" 输出: false
说明：
    0 <= len(s1) <= 100
    0 <= len(s2) <= 100 
```

### 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 哈希辅助 | O(n)       | O(1)       |

```go

```

