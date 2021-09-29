# Easy

## 剑指OfferII003.前n个数字二进制中1的个数(4)

- 题目

```
给定一个非负整数 n ，请计算 0 到 n 之间的每个数字的二进制表示中 1 的个数，并输出一个数组。
示例 1:输入: n = 2 输出: [0,1,1]
解释: 0 --> 0
1 --> 1
2 --> 10
示例 2:输入: n = 5 输出: [0,1,1,2,1,2]
解释:0 --> 0
1 --> 1
2 --> 10
3 --> 11
4 --> 100
5 --> 101
说明 :0 <= n <= 105
进阶:给出时间复杂度为 O(n*sizeof(integer)) 的解答非常容易。但你可以在线性时间 O(n) 内用一趟扫描做到吗？
要求算法的空间复杂度为 O(n) 。
你能进一步完善解法吗？要求在C++或任何其他语言中不使用任何内置函数（如 C++ 中的 __builtin_popcount ）来执行此操作。
注意：本题与主站 338 题相同
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 位运算   | O(n)       | O(n)       |
| 02   | 动态规划 | O(n)       | O(n)       |
| 03   | 暴力法   | O(n)       | O(n)       |
| 04   | 内置函数 | O(n)       | O(n)       |

```go
func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 1; i <= n; i++ {
		res[i] = res[i&(i-1)] + 1
	}
	return res
}

# 2
func countBits(n int) []int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			dp[i] = dp[i/2]
		} else {
			dp[i] = dp[i-1] + 1
		}
	}
	return dp
}

# 3
func countBits(n int) []int {
	res := make([]int, 0)
	for i := 0; i <= n; i++ {
		count := 0
		value := i
		for value != 0 {
			if value%2 == 1 {
				count++
			}
			value = value / 2
		}
		res = append(res, count)
	}
	return res
}

# 4
func countBits(n int) []int {
	res := make([]int, 0)
	for i := 0; i <= n; i++ {
		count := bits.OnesCount(uint(i))
		res = append(res, count)
	}
	return res
}
```

## 剑指OfferII006.排序数组中两个数字之和(4)

- 题目

```
给定一个已按照 升序排列  的整数数组 numbers ，请你从数组中找出两个数满足相加之和等于目标数 target 。
函数应该以长度为 2 的整数数组的形式返回这两个数的下标值。numbers 的下标 从 0 开始计数 ，
所以答案数组应当满足 0 <= answer[0] < answer[1] < numbers.length 。
假设数组中存在且只存在一对符合条件的数字，同时一个数字不能使用两次。
示例 1：输入：numbers = [1,2,4,6,10], target = 8 输出：[1,3]
解释：2 与 6 之和等于目标数 8 。因此 index1 = 1, index2 = 3 。
示例 2：输入：numbers = [2,3,4], target = 6 输出：[0,2]
示例 3：输入：numbers = [-1,0], target = -1 输出：[0,1]
提示：2 <= numbers.length <= 3 * 104
-1000 <= numbers[i] <= 1000
numbers 按 递增顺序 排列
-1000 <= target <= 1000
仅存在一个有效答案
注意：本题与主站 167 题相似（下标起点不同）
```

- 解题思路

| No.      | 思路                | 时间复杂度 | 空间复杂度 |
| -------- | ------------------- | ---------- | ---------- |
| 01       | 暴力法: 2层循环遍历 | O(n^2)     | O(1)       |
| 02       | 两遍哈希遍历        | O(n)       | O(n)       |
| 03       | 一遍哈希遍历        | O(n)       | O(n)       |
| 04(最优) | 双指针              | O(n)       | O(1)       |

```go
func twoSum(numbers []int, target int) []int {
	m := make(map[int]int, len(numbers))
	for i, n := range numbers {
		if m[target-n] != 0 {
			return []int{m[target-n] - 1, i}
		}
		m[n] = i + 1
	}
	return nil
}

#2
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for k, v := range nums {
		m[v] = k
	}

	for i := 0; i < len(nums); i++ {
		b := target - nums[i]
		if num, ok := m[b]; ok && num != i {
			return []int{i, m[b]}
		}
	}
	return []int{}
}

# 3
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

# 4
func twoSum(numbers []int, target int) []int {
	first := 0
	last := len(numbers) - 1
	result := make([]int, 2)

	for {
		if numbers[first]+numbers[last] == target {
			result[0] = first
			result[1] = last
			return result
		} else if numbers[first]+numbers[last] > target {
			last--
		} else {
			first++
		}
	}
}
```

# Medium

## 剑指OfferII004.只出现一次的数字(5)

- 题目

```
给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
示例 1：输入：nums = [2,2,3,2] 输出：3
示例 2：输入：nums = [0,1,0,1,0,1,100] 输出：100
提示：1 <= nums.length <= 3 * 104
-231 <= nums[i] <= 231 - 1
nums 中，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次
进阶：你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
注意：本题与主站 137 题相同
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 哈希辅助 | O(n)       | O(n)       |
| 02   | 排序遍历 | O(nlog(n)) | O(1)       |
| 03   | 位运算   | O(n)       | O(1)       |
| 04   | 位运算   | O(n)       | O(1)       |
| 05   | 数学计算 | O(n)       | O(n)       |

```go
func singleNumber(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}

# 2
func singleNumber(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i=i+3{
		if nums[i] != nums[i+1]{
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}

# 3
func singleNumber(nums []int) int {
	var res int
	for i := 0; i < 64; i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			if (nums[j]>>i)&1 == 1 {
				count++
			}
		}
		res = res | ((count % 3) << i) // 哪一位出现求余后1次，该位置为1
	}
	return res
}

# 4
func singleNumber(nums []int) int {
	a, b := 0, 0
	for i := 0; i < len(nums); i++ {
		a = (a ^ nums[i]) & (^b) // a：保留出现1次的数
		b = (b ^ nums[i]) & (^a) // b：保留出现2次的数
	}
	return a // 最后返回只出现1次的数
}

# 5
func singleNumber(nums []int) int {
	m := make(map[int]int)
	sum := 0
	singleSum := 0
	for _, v := range nums {
		if m[v] == 0{
			singleSum = singleSum+v
		}
		m[v] = 1
		sum = sum + v
	}
	return (singleSum*3-sum)/2
}
```

# 



