# 程序员面试金典

## 面试题01.01.判定字符是否唯一

### 题目

```
实现一个算法，确定一个字符串 s 的所有字符是否全都不同。
示例 1：输入: s = "leetcode" 输出: false 
示例 2：输入: s = "abc" 输出: true
限制：
    0 <= len(s) <= 100
    如果你不使用额外的数据结构，会很加分。
```

### 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 哈希辅助 | O(n)       | O(1)       |

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

```

