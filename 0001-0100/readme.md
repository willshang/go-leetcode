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

| No.  | 思路                | 时间复杂度 | 空间复杂度 |
| ---- | ------------------- | ---------- | ---------- |
| 01   | 暴力法: 2层循环遍历 | O(n^2)     | O(1)       |
| 02   | 两遍哈希遍历        | O(n)       | O(n)       |
| 03   | 一遍哈希遍历        | O(n)       | O(n)       |

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

