# Easy

## 剑指OfferII001.整数除法(2)

- 题目

```
给定两个整数 a 和 b ，求它们的除法的商 a/b ，要求不得使用乘号 '*'、除号 '/' 以及求余符号 '%' 。
注意：整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231, 231−1]。本题中，如果除法结果溢出，则返回 231 − 1
示例 1：输入：a = 15, b = 2 输出：7
解释：15/2 = truncate(7.5) = 7
示例 2：输入：a = 7, b = -3 输出：-2
解释：7/-3 = truncate(-2.33333..) = -2
示例 3：输入：a = 0, b = 1 输出：0
示例 4：输入：a = 1, b = 1 输出：1
提示:-231 <= a, b <= 231 - 1
b != 0
注意：本题与主站 29 题相同
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 遍历 | O(log(n))  | O(1)       |
| 02   | 计算 | O(1)       | O(1)       |

```go
func divide(a int, b int) int {
	if b == 0 || a == 0 {
		return 0
	}
	if b == 1 {
		return a
	}
	flag, count := 1, 1
	if a < 0 {
		flag = -flag
		a = -a
	}
	if b < 0 {
		flag = -flag
		b = -b
	}
	x, y, z := a, b, 0
	temp := y
	for x-y >= 0 {
		for x-y >= 0 {
			x = x - y
			z = z + count
			y = y + y
			count = count + count
		}
		y = temp
		count = 1
	}
	if z > math.MaxInt32 {
		return math.MaxInt32
	}
	if flag < 0 {
		return -z
	}
	return z
}

# 2
func divide(a int, b int) int {
	res := a / b
	if res > math.MaxInt32 {
		return math.MaxInt32
	}
	return res
}
```

## 剑指OfferII002.二进制加法(2)

- 题目

```
给定两个 01 字符串 a 和 b ，请计算它们的和，并以二进制字符串的形式输出。
输入为 非空 字符串且只包含数字 1 和 0。
示例 1:输入: a = "11", b = "10" 输出: "101"
示例 2:输入: a = "1010", b = "1011" 输出: "10101"
提示：每个字符串仅由字符 '0' 或 '1' 组成。
1 <= a.length, b.length <= 10^4
字符串如果不是 "0" ，就都不含前导零。
注意：本题与主站 67 题相同
```

- 解题思路

| No.      | 思路 | 时间复杂度 | 空间复杂度 |
| -------- | ---- | ---------- | ---------- |
| 01       | 数组 | O(n)       | O(n)       |
| 02(最优) | 模拟 | O(n)       | O(n)       |

```go
func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	length := len(a)

	A := transToInt(a, length)
	B := transToInt(b, length)

	return makeString(add(A, B))
}

func transToInt(s string, length int) []int {
	result := make([]int, length)
	ls := len(s)
	for i, b := range s {
		result[length-ls+i] = int(b - '0')
	}
	return result
}

func add(a, b []int) []int {
	length := len(a) + 1
	result := make([]int, length)
	for i := length - 1; i >= 1; i-- {
		temp := result[i] + a[i-1] + b[i-1]
		result[i] = temp % 2
		result[i-1] = temp / 2
	}
	i := 0
	for i < length-1 && result[i] == 0 {
		i++
	}
	return result[i:]
}

func makeString(nums []int) string {
	bytes := make([]byte, len(nums))
	for i := range bytes {
		bytes[i] = byte(nums[i]) + '0'
	}
	return string(bytes)
}

// 2
func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	result := ""
	flag := 0
	current := 0

	for i >= 0 || j >= 0 {
		intA, intB := 0, 0
		if i >= 0 {
			intA = int(a[i] - '0')
		}
		if j >= 0 {
			intB = int(b[j] - '0')
		}
		current = intA + intB + flag
		flag = 0
		if current >= 2 {
			flag = 1
			current = current - 2
		}
		cur := strconv.Itoa(current)
		result = cur + result
		i--
		j--
	}
	if flag == 1 {
		result = "1" + result
	}
	return result
}
```

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

## 剑指OfferII012.左右两边子数组的和相等(2)

- 题目

```
给你一个整数数组 nums ，请计算数组的 中心下标 。
数组 中心下标 是数组的一个下标，其左侧所有元素相加的和等于右侧所有元素相加的和。
如果中心下标位于数组最左端，那么左侧数之和视为 0 ，因为在下标的左侧不存在元素。这一点对于中心下标位于数组最右端同样适用。
如果数组有多个中心下标，应该返回 最靠近左边 的那一个。如果数组不存在中心下标，返回 -1 。
示例 1：输入：nums = [1,7,3,6,5,6] 输出：3
解释：中心下标是 3 。
左侧数之和 sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11 ，
右侧数之和 sum = nums[4] + nums[5] = 5 + 6 = 11 ，二者相等。
示例 2：输入：nums = [1, 2, 3] 输出：-1
解释：数组中不存在满足此条件的中心下标。
示例 3：输入：nums = [2, 1, -1]输出：0
解释：中心下标是 0 。
左侧数之和 sum = 0 ，（下标 0 左侧不存在元素），
右侧数之和 sum = nums[1] + nums[2] = 1 + -1 = 0 。
提示：1 <= nums.length <= 104
-1000 <= nums[i] <= 1000
注意：本题与主站 724 题相同：
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 遍历 | O(n)       | O(1)       |
| 02   | 遍历 | O(n)       | O(n)       |

```go
func pivotIndex(nums []int) int {
	sum := 0
	for i := range nums {
		sum = sum + nums[i]
	}
	left := 0
	for i := range nums {
		if left*2+nums[i] == sum {
			return i
		}
		left = left + nums[i]
	}
	return -1
}

# 2
func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	arr := make([]int, len(nums))
	arr[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		arr[i] = arr[i-1] + nums[i]
	}
	for i := 0; i < len(nums); i++ {
		var left, right int
		if i == 0 {
			left = 0
		} else {
			left = arr[i-1]
		}
		r := i + 1
		if r > len(nums)-1 {
			right = 0
		} else {
			right = arr[len(nums)-1] - arr[i]
		}
		if left == right {
			return i
		}
	}
	return -1
}
```

## 剑指OfferII018.有效的回文(2)

- 题目

```
给定一个字符串 s ，验证 s 是否是 回文串 ，只考虑字母和数字字符，可以忽略字母的大小写。
本题中，将空字符串定义为有效的 回文串 。
示例 1:输入: s = "A man, a plan, a canal: Panama" 输出: true
解释："amanaplanacanalpanama" 是回文串
示例 2:输入: s = "race a car" 输出: false
解释："raceacar" 不是回文串
提示：1 <= s.length <= 2 * 105
字符串 s 由 ASCII 字符组成
 注意：本题与主站 125 题相同
```

- 解题思路

| No.       | 思路     | 时间复杂度 | 空间复杂度 |
| --------- | -------- | ---------- | ---------- |
| 01( 最优) | 双指针法 | O(n)       | O(1)       |
| 02        | 双指针法 | O(n)       | O(n)       |

```go
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1

	for i < j {
		for i < j && !isChar(s[i]) {
			i++
		}
		for i < j && !isChar(s[j]) {
			j--
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isChar(c byte) bool {
	if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
		return true
	}
	return false
}

# 2
func isPalindrome(s string) bool {
	str := ""
	s = strings.ToLower(s)
	for _, value := range s {
		if (value >= '0' && value <= '9') || (value >= 'a' && value <= 'z') {
			str += string(value)
		}
	}
	if len(str) == 0 {
		return true
	}
	i := 0
	j := len(str) - 1
	for i <= j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}
```

## 剑指OfferII019.最多删除一个字符得到回文(2)

- 题目

```
给定一个非空字符串 s，请判断如果 最多 从字符串中删除一个字符能否得到一个回文字符串。
示例 1:输入: s = "aba" 输出: true
示例 2:输入: s = "abca" 输出: true
解释: 可以删除 "c" 字符 或者 "b" 字符
示例 3:输入: s = "abc" 输出: false
提示:1 <= s.length <= 105
s 由小写英文字母组成
注意：本题与主站 680 题相同
```

- 解题思路

| No.  | 思路   | 时间复杂度 | 空间复杂度 |
| ---- | ------ | ---------- | ---------- |
| 01   | 双指针 | O(n)       | O(1)       |
| 02   | 递归   | O(n)       | O(n)       |

```go
func validPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return isPalindrome(s, i, j-1) || isPalindrome(s, i+1, j)
		}
		i++
		j--
	}
	return true
}

func isPalindrome(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

# 2
func validPalindrome(s string) bool {
	length := len(s)
	if length < 2 {
		return true
	}
	if s[0] == s[length-1] {
		return validPalindrome(s[1 : length-1])
	}
	return isPalindrome(s[0:length-1]) || isPalindrome(s[1:length])
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

## 剑指OfferII023.两个链表的第一个重合节点(4)

- 题目

```
给定两个单链表的头节点 headA 和 headB ，请找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。
图示两个链表在节点 c1 开始相交：
题目数据 保证 整个链式结构中不存在环。
注意，函数返回结果后，链表必须 保持其原始结构 。
示例 1：输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
输出：Intersected at '8'
解释：相交节点的值为 8 （注意，如果两个链表相交则不能为 0）。
从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1,8,4,5]。
在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。
示例 2：输入：intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
输出：Intersected at '2'
解释：相交节点的值为 2 （注意，如果两个链表相交则不能为 0）。
从各自的表头开始算起，链表 A 为 [0,9,1,2,4]，链表 B 为 [3,2,4]。
在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。
示例 3：输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2 输出：null
解释：从各自的表头开始算起，链表 A 为 [2,6,4]，链表 B 为 [1,5]。
由于这两个链表不相交，所以 intersectVal 必须为 0，而 skipA 和 skipB 可以是任意值。
这两个链表不相交，因此返回 null 。
提示：listA 中节点数目为 m
listB 中节点数目为 n
0 <= m, n <= 3 * 104
1 <= Node.val <= 105
0 <= skipA <= m
0 <= skipB <= n
如果 listA 和 listB 没有交点，intersectVal 为 0
如果 listA 和 listB 有交点，intersectVal == listA[skipA + 1] == listB[skipB + 1]
进阶：能否设计一个时间复杂度 O(n) 、仅用 O(1) 内存的解决方案？
注意：本题与主站 160 题相同
```

- 解题思路

| No.      | 思路                       | 时间复杂度 | 空间复杂度 |
| -------- | -------------------------- | ---------- | ---------- |
| 01       | 计算长度后，对齐长度再比较 | O(n)       | O(1)       |
| 02(最优) | 交换后相连，再比较         | O(n)       | O(1)       |
| 03       | 暴力法                     | O(n^2)     | O(1)       |
| 04       | 哈希法                     | O(n)       | O(n)       |

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

## 剑指OfferII024.反转链表(4)

- 题目

```
给定单链表的头节点 head ，请反转链表，并返回反转后的链表的头节点。
示例 1：输入：head = [1,2,3,4,5] 输出：[5,4,3,2,1]
示例 2：输入：head = [1,2] 输出：[2,1]
示例 3：输入：head = [] 输出：[]
提示：链表中节点的数目范围是 [0, 5000]
-5000 <= Node.val <= 5000
进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？
注意：本题与主站 206 题相同：
```

- 解题思路

| No.  | 思路          | 时间复杂度 | 空间复杂度 |
| ---- | ------------- | ---------- | ---------- |
| 01   | 递归          | O(n)       | O(n)       |
| 02   | 迭代          | O(n)       | O(1)       |
| 03   | 数组辅助      | O(n)       | O(n)       |
| 04   | 迭代-新建节点 | O(n)       | O(1)       |

```go
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	result := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return result
}

# 2
func reverseList(head *ListNode) *ListNode {
	var result *ListNode
	var temp *ListNode
	for head != nil {
		temp = head.Next
		head.Next = result
		result = head
		head = temp
	}
	return result
}

# 3
func reverseList(head *ListNode) *ListNode {
	result := &ListNode{}
	arr := make([]*ListNode, 0)
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}
	temp := result
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i].Next = nil
		temp.Next = arr[i]
		temp = temp.Next
	}
	return result.Next
}

# 4
func reverseList(head *ListNode) *ListNode {
	var res *ListNode
	for {
		if head == nil {
			break
		}
		res = &ListNode{head.Val, res}
		head = head.Next
	}
	return res
}
```

## 剑指OfferII027.回文链表(4)

- 题目

```
给定一个链表的 头节点 head ，请判断其是否为回文链表。
如果一个链表是回文，那么链表节点序列从前往后看和从后往前看是相同的。
示例 1：输入: head = [1,2,3,3,2,1] 输出: true
示例 2：输入: head = [1,2] 输出: false
提示：链表 L 的长度范围为 [1, 105]
0 <= node.val <= 9
进阶：能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
注意：本题与主站 234 题相同
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

## 剑指OfferII032.有效的变位词(2)

- 题目

```
给定两个字符串 s 和 t ，编写一个函数来判断它们是不是一组变位词（字母异位词）。
注意：若 s 和 t 中每个字符出现的次数都相同且字符顺序不完全相同，则称 s 和 t 互为变位词（字母异位词）。
示例 1:输入: s = "anagram", t = "nagaram" 输出: true
示例 2:输入: s = "rat", t = "car" 输出: false
示例 3:输入: s = "a", t = "a" 输出: false
提示:1 <= s.length, t.length <= 5 * 104
s and t 仅包含小写字母
进阶: 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
注意：本题与主站 242 题相似（字母异位词定义不同）
```

- 解题思路

| 01   | 哈希辅助 | O(n)       | O(n) |
| ---- | -------- | ---------- | ---- |
| 01   | 哈希辅助 | O(n)       | O(1) |
| 02   | 排序比较 | O(nlog(n)) | O(n) |

```go
func isAnagram(s string, t string) bool {
	if len(s) != len(t) || s == t {
		return false
	}

	sr := []rune(s)
	tr := []rune(t)

	rec := make(map[rune]int, len(sr))
	for i := range sr {
		rec[sr[i]]++
		rec[tr[i]]--
	}

	for _, n := range rec {
		if n != 0 {
			return false
		}
	}
	return true
}

# 2
func isAnagram(s string, t string) bool {
	if len(s) != len(t) || s == t {
		return false
	}
	sArr := make([]int, len(s))
	tArr := make([]int, len(t))
	for i := 0; i < len(s); i++ {
		sArr[i] = int(s[i] - 'a')
		tArr[i] = int(t[i] - 'a')
	}
	sort.Ints(sArr)
	sort.Ints(tArr)
	for i := 0; i < len(s); i++ {
		if sArr[i] != tArr[i] {
			return false
		}
	}
	return true
}
```

## 剑指OfferII034.外星语言是否排序(2)

- 题目

```
某种外星语也使用英文小写字母，但可能顺序 order 不同。字母表的顺序（order）是一些小写字母的排列。
给定一组用外星语书写的单词 words，以及其字母表的顺序 order，只有当给定的单词在这种外星语中按字典序排列时，返回 true；
否则，返回 false。
示例 1：输入：words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz" 输出：true
解释：在该语言的字母表中，'h' 位于 'l' 之前，所以单词序列是按字典序排列的。
示例 2：输入：words = ["word","world","row"], order = "worldabcefghijkmnpqstuvxyz" 输出：false
解释：在该语言的字母表中，'d' 位于 'l' 之后，那么 words[0] > words[1]，因此单词序列不是按字典序排列的。
示例 3：输入：words = ["apple","app"], order = "abcdefghijklmnopqrstuvwxyz" 输出：false
解释：当前三个字符 "app" 匹配时，第二个字符串相对短一些，然后根据词典编纂规则 "apple" > "app"，
因为 'l' > '∅'，其中 '∅' 是空白字符，定义为比任何其他字符都小（更多信息）。
提示：1 <= words.length <= 100
1 <= words[i].length <= 20
order.length == 26
在 words[i] 和 order 中的所有字符都是英文小写字母。
注意：本题与主站 953 题相同： 
```

- 解题思路

| No.  | 思路          | 时间复杂度 | 空间复杂度 |
| ---- | ------------- | ---------- | ---------- |
| 01   | 哈希辅助-替换 | O(n)       | O(n)       |
| 02   | 遍历比较      | O(n)       | O(1)       |

```go
func isAlienSorted(words []string, order string) bool {
	newWords := make([]string, len(words))
	m := make(map[byte]int)
	for i := 0; i < len(order); i++ {
		m[order[i]] = i
	}
	for i := 0; i < len(words); i++ {
		str := ""
		for j := 0; j < len(words[i]); j++ {
			str = str + string(m[words[i][j]]+'a')
		}
		newWords[i] = str
	}
	for i := 0; i < len(newWords)-1; i++ {
		if newWords[i] > newWords[i+1] {
			return false
		}
	}
	return true
}

#
func isAlienSorted(words []string, order string) bool {
	m := make(map[byte]int)
	for i := 0; i < len(order); i++ {
		m[order[i]] = i
	}

	for i := 0; i < len(words)-1; i++ {
		length := len(words[i])
		if len(words[i+1]) < length {
			length = len(words[i+1])
		}
		for j := 0; j < length; j++ {
			if m[words[i][j]] < m[words[i+1][j]] {
				break
			}
			if m[words[i][j]] > m[words[i+1][j]] {
				return false
			}
			if j == length-1 {
				if len(words[i]) > len(words[i+1]) {
					return false
				}
			}
		}
	}
	return true
}
```

## 剑指OfferII041.滑动窗口的平均值

### 题目

```

```

### 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 哈希辅助 | O(n)       | O(n)       |

```go

```

## 剑指OfferII042.最近请求次数(2)

- 题目

```
写一个 RecentCounter 类来计算特定时间范围内最近的请求。
请实现 RecentCounter 类：
RecentCounter() 初始化计数器，请求数为 0 。
int ping(int t) 在时间 t 添加一个新请求，其中 t 表示以毫秒为单位的某个时间，
并返回过去 3000 毫秒内发生的所有请求数（包括新请求）。确切地说，返回在 [t-3000, t] 内发生的请求数。
保证 每次对 ping 的调用都使用比之前更大的 t 值。
示例：输入： inputs = ["RecentCounter", "ping", "ping", "ping", "ping"]
inputs = [[], [1], [100], [3001], [3002]]
输出：[null, 1, 2, 3, 3]
解释：RecentCounter recentCounter = new RecentCounter();
recentCounter.ping(1);     // requests = [1]，范围是 [-2999,1]，返回 1
recentCounter.ping(100);   // requests = [1, 100]，范围是 [-2900,100]，返回 2
recentCounter.ping(3001);  // requests = [1, 100, 3001]，范围是 [1,3001]，返回 3
recentCounter.ping(3002);  // requests = [1, 100, 3001, 3002]，范围是 [2,3002]，返回 3
提示：1 <= t <= 109
保证每次对 ping 调用所使用的 t 值都 严格递增
至多调用 ping 方法 104 次
注意：本题与主站 933 题相同：
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 数组操作 | O(n)       | O(n)       |
| 02   | 数组操作 | O(n)       | O(n)       |

```go
type RecentCounter struct {
	arr []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		arr: make([]int, 0),
	}
}

func (r *RecentCounter) Ping(t int) int {
	r.arr = append(r.arr, t)
	res := 1
	for i := len(r.arr) - 2; i >= 0; i-- {
		if t-r.arr[i] <= 3000 {
			res++
		} else {
			r.arr = r.arr[i+1:]
			break
		}
	}
	return res
}

#
type RecentCounter struct {
	arr []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		arr: make([]int, 0),
	}
}

func (r *RecentCounter) Ping(t int) int {
	r.arr = append(r.arr, t)
	start := t - 3000
	for len(r.arr) > 0 && r.arr[0] < start {
		r.arr = r.arr[1:]
	}
	return len(r.arr)
}
```

## 剑指OfferII052.展平二叉搜索树(3)

- 题目

```
给你一棵二叉搜索树，请 按中序遍历 将其重新排列为一棵递增顺序搜索树，
使树中最左边的节点成为树的根节点，并且每个节点没有左子节点，只有一个右子节点。
示例 1：输入：root = [5,3,6,2,4,null,8,1,null,null,null,7,9]
输出：[1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]
示例 2：输入：root = [5,1,7] 输出：[1,null,5,null,7]
提示：树中节点数的取值范围是 [1, 100]
0 <= Node.val <= 1000
注意：本题与主站 897 题相同： 
```

- 解题思路

| No.  | 思路          | 时间复杂度 | 空间复杂度 |
| ---- | ------------- | ---------- | ---------- |
| 01   | 递归-数组辅助 | O(n)       | O(n)       |
| 02   | 递归          | O(n)       | O(log(n))  |
| 03   | 迭代          | O(n)       | O(n)       |

```go
func increasingBST(root *TreeNode) *TreeNode {
	arr := make([]int, 0)
	dfs(root, &arr)
	if len(arr) == 0 {
		return root
	}
	newRoot := &TreeNode{Val: arr[0]}
	cur := newRoot
	for i := 1; i < len(arr); i++ {
		cur.Right = &TreeNode{Val: arr[i]}
		cur = cur.Right
	}
	return newRoot
}

func dfs(node *TreeNode, arr *[]int) {
	if node == nil {
		return
	}
	dfs(node.Left, arr)
	*arr = append(*arr, node.Val)
	dfs(node.Right, arr)
}

# 2
var prev *TreeNode

func increasingBST(root *TreeNode) *TreeNode {
	prev = &TreeNode{}
	head := prev
	dfs(root)
	return head.Right
}

func dfs(node *TreeNode) {
	if node == nil {
		return
	}
	dfs(node.Left)
	node.Left = nil
	prev.Right = node
	prev = node
	dfs(node.Right)
}

# 3
func increasingBST(root *TreeNode) *TreeNode {
	stack := make([]*TreeNode, 0)
	newRoot := &TreeNode{}
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		if node.Right != nil {
			stack = append(stack, node.Right)
			node.Right = nil
			continue
		}
		stack = stack[:len(stack)-1]
		node.Right = newRoot.Right
		newRoot.Right = node
		if node.Left != nil {
			stack = append(stack, node.Left)
			node.Left = nil
		}
	}
	return newRoot.Right
}
```

## 剑指OfferII056.二叉搜索树中两个节点之和(4)

- 题目

```
给定一个二叉搜索树的 根节点 root 和一个整数 k , 请判断该二叉搜索树中是否存在两个节点它们的值之和等于 k 。
假设二叉搜索树中节点的值均唯一。
示例 1：输入: root = [8,6,10,5,7,9,11], k = 12 输出: true
解释: 节点 5 和节点 7 之和等于 12
示例 2：输入: root = [8,6,10,5,7,9,11], k = 22 输出: false
解释: 不存在两个节点值之和为 22 的节点
提示：二叉树的节点个数的范围是  [1, 104].
-104 <= Node.val <= 104
root 为二叉搜索树
-105 <= k <= 105
注意：本题与主站 653 题相同： 
```

- 解题思路

| No.  | 思路          | 时间复杂度 | 空间复杂度 |
| ---- | ------------- | ---------- | ---------- |
| 01   | 递归+哈希辅助 | O(n)       | O(n)       |
| 02   | 递归          | O(nlog(n)) | O(log(n))  |
| 03   | 迭代          | O(n)       | O(n)       |
| 04   | 递归+二分查找 | O(n)       | O(n)       |

```go
func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	m := map[int]int{}
	return dfs(root, k, m)
}

func dfs(node *TreeNode, k int, m map[int]int) bool {
	if node == nil {
		return false
	}
	if _, ok := m[k-node.Val]; ok {
		return true
	}
	m[node.Val] = node.Val
	return dfs(node.Left, k, m) || dfs(node.Right, k, m)
}

# 2
func dfs(root, searchRoot *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	found := findNode(searchRoot, k-root.Val)
	if found != nil && found != root {
		return true
	}
	return dfs(root.Left, searchRoot, k) ||
		dfs(root.Right, searchRoot, k)
}

func findNode(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == target {
		return root
	}
	if root.Val < target {
		return findNode(root.Right, target)
	}
	return findNode(root.Left, target)
}

# 3
func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	m := make(map[int]int)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if _, ok := m[k-node.Val]; ok {
			return true
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		m[node.Val] = 1
	}
	return false
}

# 4
var arr []int

func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	arr = make([]int, 0)
	dfs(root)
	i := 0
	j := len(arr) - 1
	for i < j {
		if arr[i]+arr[j] == k {
			return true
		} else if arr[i]+arr[j] > k {
			j--
		} else {
			i++
		}
	}
	return false
}

func dfs(node *TreeNode) {
	if node == nil {
		return
	}
	dfs(node.Left)
	arr = append(arr, node.Val)
	dfs(node.Right)
}
```

## 剑指OfferII059.数据流的第K大数值(2)

- 题目

```
设计一个找到数据流中第 k 大元素的类（class）。注意是排序后的第 k 大元素，不是第 k 个不同的元素。
请实现 KthLargest 类：
KthLargest(int k, int[] nums) 使用整数 k 和整数流 nums 初始化对象。
int add(int val) 将 val 插入数据流 nums 后，返回当前数据流中第 k 大的元素。
示例：输入： ["KthLargest", "add", "add", "add", "add", "add"]
[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
输出：[null, 4, 5, 5, 8, 8]
解释：KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
kthLargest.add(3);   // return 4
kthLargest.add(5);   // return 5
kthLargest.add(10);  // return 5
kthLargest.add(9);   // return 8
kthLargest.add(4);   // return 8
提示：1 <= k <= 104
0 <= nums.length <= 104
-104 <= nums[i] <= 104
-104 <= val <= 104
最多调用 add 方法 104 次
题目数据保证，在查找第 k 大元素时，数组中至少有 k 个元素
注意：本题与主站 703 题相同：
```

- 解题思路

| No.  | 思路            | 时间复杂度 | 空间复杂度 |
| ---- | --------------- | ---------- | ---------- |
| 01   | 最小堆+内置heap | O(nlog(n)) | O(n)       |
| 02   | 小根堆          | O(nlog(n)) | O(n)       |

```go
type KthLargest struct {
	k    int
	heap intHeap
}

func Constructor(k int, nums []int) KthLargest {
	h := intHeap(nums)
	heap.Init(&h)

	for len(h) > k {
		heap.Pop(&h)
	}
	return KthLargest{
		k:    k,
		heap: h,
	}
}

func (k *KthLargest) Add(val int) int {
	heap.Push(&k.heap, val)
	if len(k.heap) > k.k {
		heap.Pop(&k.heap)
	}
	return k.heap[0]
}

// 内置heap，实现接口
/*
type Interface interface {
	sort.Interface
	Push(x interface{}) // add x as element Len()
	Pop() interface{}   // remove and return element Len() - 1.
}
*/
type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

# 2
type KthLargest struct {
	nums []int
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	if k < len(nums) {
		sort.Ints(nums)
		nums = nums[len(nums)-k:]
	}
	// 向上调整
	Up(nums)
	return KthLargest{
		nums: nums,
		k:    k,
	}
}

func (k *KthLargest) Add(val int) int {
	if k.k > len(k.nums) {
		k.nums = append(k.nums, val)
		Up(k.nums)
	} else {
		if val > k.nums[0] {
			// 在堆顶，向下调整
			k.nums[0] = val
			Down(k.nums, 0)
		}
	}
	return k.nums[0]
}

func Down(nums []int, index int) {
	length := len(nums)
	minIndex := index
	for {
		left := 2*index + 1
		right := 2*index + 2
		if left < length && nums[left] < nums[minIndex] {
			minIndex = left
		}
		if right < length && nums[right] < nums[minIndex] {
			minIndex = right
		}
		if minIndex == index {
			break
		}
		swap(nums, index, minIndex)
		index = minIndex
	}
}

func Up(nums []int) {
	length := len(nums)
	for i := length/2 - 1; i >= 0; i-- {
		minIndex := i
		left := 2*i + 1
		right := 2*i + 2
		if left < length && nums[left] < nums[minIndex] {
			minIndex = left
		}
		if right < length && nums[right] < nums[minIndex] {
			minIndex = right
		}
		if i != minIndex {
			swap(nums, i, minIndex)
		}
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
```

## 剑指OfferII068.查找插入位置(3)

- 题目

```
给定一个排序的整数数组 nums 和一个整数目标值 target ，请在数组中找到 target ，并返回其下标。
如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
请必须使用时间复杂度为 O(log n) 的算法。
示例 1:输入: nums = [1,3,5,6], target = 5 输出: 2
示例 2:输入: nums = [1,3,5,6], target = 2 输出: 1
示例 3:输入: nums = [1,3,5,6], target = 7 输出: 4
示例 4:输入: nums = [1,3,5,6], target = 0 输出: 0
示例 5:输入: nums = [1], target = 0 输出: 0
提示:1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums 为无重复元素的升序排列数组
-104 <= target <= 104
注意：本题与主站 35 题相同：
```

- 解题思路

| No.      | 思路     | 时间复杂度 | 空间复杂度 |
| -------- | -------- | ---------- | ---------- |
| 01(最优) | 二分查找 | O(log(n))  | O(1)       |
| 02       | 顺序查找 | O(n)       | O(1)       |
| 03       | 顺序查找 | O(n)       | O(1)       |

```go
// 二分查找
func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) / 2
		switch {
		case nums[mid] < target:
			low = mid + 1
		case nums[mid] > target:
			high = mid - 1
		default:
			return mid
		}
	}
	return low
}

// 顺序查找
func searchInsert(nums []int, target int) int {
	i := 0
	for i < len(nums) && nums[i] < target {
		if nums[i] == target {
			return i
		}
		i++
	}
	return i
}

// 顺序查找
func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}
```

## 剑指OfferII069.山峰数组的顶部

### 题目

```
符合下列属性的数组 arr 称为 山峰数组（山脉数组） ：
arr.length >= 3
存在 i（0 < i < arr.length - 1）使得：
arr[0] < arr[1] < ... arr[i-1] < arr[i]
arr[i] > arr[i+1] > ... > arr[arr.length - 1]
给定由整数组成的山峰数组 arr ，返回任何满足 
arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1] 的下标 i ，即山峰顶部。
示例 1：输入：arr = [0,1,0] 输出：1
示例 2：输入：arr = [1,3,5,4,2] 输出：2
示例 3：输入：arr = [0,10,5,2] 输出：1
示例 4：输入：arr = [3,4,5,1] 输出：2
示例 5：输入：arr = [24,69,100,99,79,78,67,36,26,19] 输出：2
提示：3 <= arr.length <= 104
0 <= arr[i] <= 106
题目数据保证 arr 是一个山脉数组
进阶：很容易想到时间复杂度 O(n) 的解决方案，你可以设计一个 O(log(n)) 的解决方案吗？
注意：本题与主站 852 题相同
```

### 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 遍历     | O(n)       | O(1)       |
| 02   | 二分查找 | O(log(n))  | O(1)       |

```go

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

## 剑指OfferII005.单词长度的最大乘积(2)

- 题目

```
给定一个字符串数组 words，请计算当两个字符串 words[i] 和 words[j] 不包含相同字符时，它们长度的乘积的最大值。
假设字符串中只包含英语的小写字母。如果没有不包含相同字符的一对字符串，返回 0。
示例 1:输入: words = ["abcw","baz","foo","bar","fxyz","abcdef"] 输出: 16 
解释: 这两个单词为 "abcw", "fxyz"。它们不包含相同字符，且长度的乘积最大。
示例 2:输入: words = ["a","ab","abc","d","cd","bcd","abcd"] 输出: 4 
解释: 这两个单词为 "ab", "cd"。
示例 3:输入: words = ["a","aa","aaa","aaaa"] 输出: 0 
解释: 不存在这样的两个单词。
提示：2 <= words.length <= 1000
1 <= words[i].length <= 1000
words[i] 仅包含小写字母
注意：本题与主站 318 题相同
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 内置函数 | O(n^3)     | O(1)       |
| 02   | 位运算   | O(n^2)     | O(n)       |

```go
func maxProduct(words []string) int {
	res := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if strings.ContainsAny(words[i], words[j]) == false && 
				res < len(words[i])*len(words[j]) {
				res = len(words[i]) * len(words[j])
			}
		}
	}
	return res
}

# 2
func maxProduct(words []string) int {
	res := 0
	arr := make([]int, len(words))
	for i := 0; i < len(words); i++ {
		for _, char := range words[i] {
			// 位或 只要有1，那么就是1
			arr[i] = arr[i] | 1<<uint(char-'a')
		}
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]&arr[j] == 0 && res < len(words[i])*len(words[j]) {
				res = len(words[i]) * len(words[j])
			}
		}
	}
	return res
}
```

## 剑指OfferII007.数组中和为0的三个数(2)

- 题目

```
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a ，b ，c ，使得 a + b + c = 0 ？
请找出所有和为 0 且 不重复 的三元组。
示例 1：输入：nums = [-1,0,1,2,-1,-4] 输出：[[-1,-1,2],[-1,0,1]]
示例 2：输入：nums = [] 输出：[]
示例 3：输入：nums = [0] 出：[]
提示：0 <= nums.length <= 3000
-105 <= nums[i] <= 105
注意：本题与主站 15 题相同：
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 双指针   | O(n^2)     | O(n^2)     |
| 02   | 哈希辅助 | O(n^2)     | O(n^2)     |

```go
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		target := 0 - nums[i]
		left := i + 1
		right := len(nums) - 1
		if nums[i] > 0 || nums[i]+nums[left] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for left < right {
			if left > i+1 && nums[left] == nums[left-1] {
				left++
				continue
			}
			if right < len(nums)-2 && nums[right] == nums[right+1] {
				right--
				continue
			}
			if nums[left]+nums[right] > target {
				right--
			} else if nums[left]+nums[right] < target {
				left++
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--
			}
		}
	}
	return res
}

# 2
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	m := make(map[[2]int]int)
	p := make(map[int]int)
	sort.Ints(nums)
	for k, v := range nums {
		p[v] = k
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			sum := nums[i] + nums[j]
			if sum > 0 {
				break
			}
			if value, ok := p[-sum]; ok && value > j {
				if _, ok2 := m[[2]int{nums[i], nums[j]}]; !ok2 {
					res = append(res, []int{nums[i], nums[j], 0 - nums[i] - nums[j]})
					m[[2]int{nums[i], nums[j]}] = 1
				}
			}
		}
	}
	return res
}
```

## 剑指OfferII008.和大于等于target的最短子数组(3)

- 题目

```
给定一个含有 n 个正整数的数组和一个正整数 target 。
找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。
如果不存在符合条件的子数组，返回 0 。
示例 1：输入：target = 7, nums = [2,3,1,2,4,3] 输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。
示例 2：输入：target = 4, nums = [1,4,4] 输出：1
示例 3：输入：target = 11, nums = [1,1,1,1,1,1,1,1] 输出：0 
提示：1 <= target <= 109
1 <= nums.length <= 105
1 <= nums[i] <= 105
进阶：如果你已经实现 O(n) 时间复杂度的解法, 请尝试设计一个 O(n log(n)) 时间复杂度的解法。
注意：本题与主站 209 题相同
```

- 解题思路

| No.  | 思路            | 时间复杂度 | 空间复杂度 |
| ---- | --------------- | ---------- | ---------- |
| 01   | 暴力法          | O(n^2)     | O(1)       |
| 02   | 前缀和-二分查找 | O(nlog(n)) | O(n)       |
| 03   | 双指针          | O(n)       | O(1)       |

```go
func minSubArrayLen(target int, nums []int) int {
	res := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum = sum + nums[j]
			if sum >= target {
				if res > j-i+1 {
					res = j - i + 1
				}
				break
			}
		}
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}

# 2
func minSubArrayLen(target int, nums []int) int {
	res := math.MaxInt32
	arr := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		arr[i] = arr[i-1] + nums[i-1]
	}
	for i := 1; i <= len(nums); i++ {
		target := target + arr[i-1]
		index := sort.SearchInts(arr, target)
		if index <= len(nums) {
			if res > index-i+1 {
				res = index - i + 1
			}
		}
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}

# 3
func minSubArrayLen(s int, nums []int) int {
	res := math.MaxInt32
	i, j := 0, 0
	sum := 0
	for ; j < len(nums); j++ {
		sum = sum + nums[j]
		for sum >= s {
			if res > j-i+1 {
				res = j - i + 1
			}
			sum = sum - nums[i]
			i++
		}
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}
```

## 剑指OfferII009.乘积小于K的子数组(1)

- 题目

```
给定一个正整数数组 nums和整数 k ，请找出该数组内乘积小于 k 的连续的子数组的个数。
示例 1:输入: nums = [10,5,2,6], k = 100 输出: 8
解释: 8 个乘积小于 100 的子数组分别为: [10], [5], [2], [6], [10,5], [5,2], [2,6], [5,2,6]。
需要注意的是 [10,5,2] 并不是乘积小于100的子数组。
示例 2:输入: nums = [1,2,3], k = 0 输出: 0
提示: 1 <= nums.length <= 3 * 104
1 <= nums[i] <= 1000
0 <= k <= 106
注意：本题与主站 713 题相同
```

- 解题思路

| No.  | 思路   | 时间复杂度 | 空间复杂度 |
| ---- | ------ | ---------- | ---------- |
| 01   | 双指针 | O(n)       | O(1)       |

```go
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	res := 0
	left := 0
	total := 1
	for right := 0; right < len(nums); right++ {
		total = total * nums[right]
		for k <= total {
			total = total / nums[left]
			left++
		}
		res = res + right - left + 1
	}
	return res
}
```

## 剑指OfferII010.和为k的子数组(4)

- 题目

```
给定一个整数数组和一个整数 k ，请找到该数组中和为 k 的连续子数组的个数。
示例 1 :输入:nums = [1,1,1], k = 2 输出: 2
解释: 此题 [1,1] 与 [1,1] 为两种不同的情况
示例 2 :输入:nums = [1,2,3], k = 3 输出: 2
提示:1 <= nums.length <= 2 * 104
-1000 <= nums[i] <= 1000
-107 <= k <= 107
注意：本题与主站 560 题相同
```

- 解题思路

| No.  | 思路            | 时间复杂度 | 空间复杂度 |
| ---- | --------------- | ---------- | ---------- |
| 01   | 暴力法          | O(n^2)     | O(1)       |
| 02   | 前缀和-遍历     | O(n^2)     | O(n)       |
| 03   | 前缀和-哈希辅助 | O(n)       | O(n)       |
| 04   | 前缀和-哈希辅助 | O(n)       | O(n)       |

```go
func subarraySum(nums []int, k int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum = sum + nums[j]
			if sum == k {
				res++
			}
		}
	}
	return res
}

# 2
func subarraySum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	res := 0
	arr := make([]int, len(nums)+1)
	arr[0] = 0
	for i := 1; i <= len(nums); i++ {
		arr[i] = arr[i-1] + nums[i-1]
	}
	for i := 0; i <= len(nums); i++ {
		for j := 0; j < i; j++ {
			if arr[i]-arr[j] == k {
				res++
			}
		}
	}
	return res
}

# 3
func subarraySum(nums []int, k int) int {
	res := 0
	m := make(map[int]int)
	m[0] = 1 // 保证第一个k的存在
	sum := 0
	// sum[i:j]= sum[0:j]-sum[0:i]，把sum[i:j]设为k，
	// 于是可以转化为sum[0:j]-k=sum[0:i]
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		if _, ok := m[sum-k]; ok {
			res = res + m[sum-k]
		}
		m[sum]++
	}
	return res
}

# 4
func subarraySum(nums []int, k int) int {
	res := 0
	m := make(map[int][]int)
	m[0] = []int{-1} // 保证第一个k的存在
	sum := 0
	// sum[i:j]= sum[0:j]-sum[0:i]，把sum[i:j]设为k，
	// 于是可以转化为sum[0:j]-k=sum[0:i]
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		if _, ok := m[sum-k]; ok {
			res = res + len(m[sum-k])
            // 输出满足条件的子数组下标
			// for _, v := range m[sum-k] {
			//	fmt.Println(v+1, i)
			// }
		}
		m[sum] = append(m[sum], i)
	}
	return res
}
```

## 剑指OfferII011.0和1个数相同的子数组(1)

- 题目

```
给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。
示例 1:输入: nums = [0,1] 输出: 2
说明: [0, 1] 是具有相同数量 0 和 1 的最长连续子数组。
示例 2:输入: nums = [0,1,0] 输出: 2
说明: [0, 1] (或 [1, 0]) 是具有相同数量 0 和 1 的最长连续子数组。
提示：1 <= nums.length <= 105
nums[i] 不是 0 就是 1
注意：本题与主站 525 题相同
```

- 解题思路

| No.  | 思路   | 时间复杂度 | 空间复杂度 |
| ---- | ------ | ---------- | ---------- |
| 01   | 前缀和 | O(n)       | O(n)       |

```go
func findMaxLength(nums []int) int {
	res := 0
	m := make(map[int]int)
	m[0] = -1
	total := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			total--
		} else {
			total++
		}
		if first, ok := m[total]; !ok {
			m[total] = i
		} else {
			if i-first > res {
				res = i - first
			}
		}
	}
	return res
}
```

## 剑指OfferII013.二维子矩阵的和(1)

- 题目

```
给定一个二维矩阵 matrix，以下类型的多个请求：
计算其子矩形范围内元素的总和，该子矩阵的左上角为 (row1, col1) ，右下角为 (row2, col2) 。
实现 NumMatrix 类：NumMatrix(int[][] matrix) 给定整数矩阵 matrix 进行初始化
int sumRegion(int row1, int col1, int row2, int col2) 返回左上角 (row1, col1) 、
右下角 (row2, col2) 的子矩阵的元素总和。
示例 1：输入:  ["NumMatrix","sumRegion","sumRegion","sumRegion"]
[[[[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]],[2,1,4,3],[1,1,2,2],[1,2,2,4]]
输出: [null, 8, 11, 12]
解释:NumMatrix numMatrix = new NumMatrix([[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]]);
numMatrix.sumRegion(2, 1, 4, 3); // return 8 (红色矩形框的元素总和)
numMatrix.sumRegion(1, 1, 2, 2); // return 11 (绿色矩形框的元素总和)
numMatrix.sumRegion(1, 2, 2, 4); // return 12 (蓝色矩形框的元素总和)
提示：m == matrix.length
n == matrix[i].length
1 <= m, n <= 200
-105 <= matrix[i][j] <= 105
0 <= row1 <= row2 < m
0 <= col1 <= col2 < n
最多调用 104 次 sumRegion 方法
注意：本题与主站 304 题相同：
```

- 解题思路

| No.  | 思路   | 时间复杂度 | 空间复杂度 |
| ---- | ------ | ---------- | ---------- |
| 01   | 前缀和 | O(1)       | O(n^2)     |

```go
type NumMatrix struct {
	arr [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if matrix == nil || len(matrix) == 0 || matrix[0] == nil || len(matrix[0]) == 0 {
		arr := make([][]int, 1)
		for i := 0; i < 1; i++ {
			arr[i] = make([]int, 1)
		}
		return NumMatrix{arr: arr}
	}
	n, m := len(matrix), len(matrix[0])
	arr := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		arr[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			arr[i][j] = arr[i][j-1] + arr[i-1][j] - arr[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return NumMatrix{arr: arr}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.arr[row2+1][col2+1] - this.arr[row2+1][col1] - this.arr[row1][col2+1] + this.arr[row1][col1]
}
```

## 剑指OfferII014.字符串中的变位词(2)

- 题目

```
给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的某个变位词。
换句话说，第一个字符串的排列之一是第二个字符串的 子串 。
示例 1：输入: s1 = "ab" s2 = "eidbaooo" 输出: True
解释: s2 包含 s1 的排列之一 ("ba").
示例 2：输入: s1= "ab" s2 = "eidboaoo" 输出: False
提示：1 <= s1.length, s2.length <= 104
s1 和 s2 仅包含小写字母
注意：本题与主站 567 题相同
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 滑动窗口 | O(n)       | O(1)       |
| 02   | 滑动窗口 | O(n)       | O(1)       |

```go
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	arr1, arr2 := [26]int{}, [26]int{}
	for i := 0; i < len(s1); i++ {
		arr1[s1[i]-'a']++
		arr2[s2[i]-'a']++
	}
	for i := 0; i < len(s2)-len(s1); i++ {
		if arr1 == arr2 {
			return true
		}
		arr2[s2[i]-'a']--
		arr2[s2[i+len(s1)]-'a']++
	}
	return arr1 == arr2
}

# 2
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	m1, m2 := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		m1[s1[i]-'a']++
		m2[s2[i]-'a']++
	}
	for i := 0; i < len(s2)-len(s1); i++ {
		if compare(m1, m2) {
			return true
		}
		m2[s2[i]-'a']--
		if m2[s2[i]-'a'] == 0 {
			delete(m2, s2[i]-'a')
		}
		m2[s2[i+len(s1)]-'a']++
	}
	return compare(m1, m2)
}

func compare(m1, m2 map[byte]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k := range m1 {
		if m2[k] != m1[k] {
			return false
		}
	}
	return true
}
```

## 剑指OfferII015.字符串中的所有变位词(2)

- 题目

```
给定两个字符串 s 和 p，找到 s 中所有 p 的 变位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
变位词 指字母相同，但排列不同的字符串。
示例 1:输入: s = "cbaebabacd", p = "abc" 输出: [0,6]
解释:起始索引等于 0 的子串是 "cba", 它是 "abc" 的变位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的变位词。
示例 2:输入: s = "abab", p = "ab" 输出: [0,1,2]
解释:起始索引等于 0 的子串是 "ab", 它是 "ab" 的变位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的变位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的变位词。
提示:1 <= s.length, p.length <= 3 * 104
s 和 p 仅包含小写字母
注意：本题与主站 438 题相同：
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 滑动窗口 | O(n)       | O(1)       |
| 02   | 滑动窗口 | O(n)       | O(1)       |

```go
func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	if len(p) > len(s) {
		return res
	}
	arr1, arr2 := [26]int{}, [26]int{}
	for i := 0; i < len(p); i++ {
		arr1[p[i]-'a']++
		arr2[s[i]-'a']++
	}
	for i := 0; i < len(s)-len(p); i++ {
		if arr1 == arr2 {
			res = append(res, i)
		}
		arr2[s[i]-'a']--
		arr2[s[i+len(p)]-'a']++
	}
	if arr1 == arr2 {
		res = append(res, len(s)-len(p))
	}
	return res
}

# 2
func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	if len(p) > len(s) {
		return res
	}
	m1, m2 := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(p); i++ {
		m1[p[i]-'a']++
		m2[s[i]-'a']++
	}
	for i := 0; i < len(s)-len(p); i++ {
		if compare(m1, m2) {
			res = append(res, i)
		}
		m2[s[i]-'a']--
		if m2[s[i]-'a'] == 0 {
			delete(m2, s[i]-'a')
		}
		m2[s[i+len(p)]-'a']++
	}
	if compare(m1, m2) {
		res = append(res, len(s)-len(p))
	}
	return res
}

func compare(m1, m2 map[byte]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k := range m1 {
		if m2[k] != m1[k] {
			return false
		}
	}
	return true
}
```

## 剑指OfferII016.不含重复字符的最长子字符串(5)

- 题目

```
给定一个字符串 s ，请你找出其中不含有重复字符的 最长连续子字符串 的长度。
示例 1:输入: s = "abcabcbb" 输出: 3 
解释: 因为无重复字符的最长子字符串是 "abc"，所以其长度为 3。
示例 2:输入: s = "bbbbb" 输出: 1
解释: 因为无重复字符的最长子字符串是 "b"，所以其长度为 1。
示例 3:输入: s = "pwwkew" 输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
示例 4:输入: s = "" 输出: 0
提示：0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成
注意：本题与主站 3 题相同：
```

- 解题思路

| No.  | 思路            | 时间复杂度 | 空间复杂度 |
| ---- | --------------- | ---------- | ---------- |
| 01   | 双指针          | O(n)       | O(1)       |
| 02   | 双指针-内置函数 | O(n^2)     | O(1)       |
| 03   | 哈希辅助-双指针 | O(n)       | O(1)       |
| 04   | 动态规划        | O(n)       | O(n)       |
| 05   | 双指针          | O(n)       | O(1)       |

```go
func lengthOfLongestSubstring(s string) int {
	arr := [256]int{}
	for i := range arr {
		arr[i] = -1
	}
	max, j := 0, 0
	for i := 0; i < len(s); i++ {
		if arr[s[i]] >= j {
			j = arr[s[i]] + 1
		} else if i+1-j > max {
			max = i + 1 - j
		}
		arr[s[i]] = i
	}
	return max
}

# 2
func lengthOfLongestSubstring(s string) int {
	max, j := 0, 0
	for i := 0; i < len(s); i++ {
		index := strings.Index(s[j:i], string(s[i]))
		if index == -1 {
			continue
		}
		if i-j > max {
			max = i - j
		}
		j = j + index + 1
	}
	if len(s)-j > max {
		max = len(s) - j
	}
	return max
}

# 3
func lengthOfLongestSubstring(s string) int {
	m := make(map[uint8]int)
	max, j := 0, 0
	for i := 0; i < len(s); i++ {
		if v, ok := m[s[i]]; ok && v >= j {
			j = v + 1
		} else if i+1-j > max {
			max = i + 1 - j
		}
		m[s[i]] = i
	}
	return max
}

# 4
func lengthOfLongestSubstring(s string) int {
	if len(s) < 1 {
		return 0
	}
	dp := make([]int, len(s))
	dp[0] = 1
	res := 1
	m := make(map[byte]int)
	m[s[0]] = 0
	for i := 1; i < len(s); i++ {
		index := -1
		if value, ok := m[s[i]]; ok {
			index = value
		}
		if i-index > dp[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = i - index
		}
		m[s[i]] = i
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

# 5
func lengthOfLongestSubstring(s string) int {
	arr := [256]int{}
	for i := range arr {
		arr[i] = -1
	}
	res, j := 0, -1
	for i := 0; i < len(s); i++ {
		if arr[s[i]] > j { // 出现重复了，更新下标
			j = arr[s[i]]
		} else {
			res = max(res, i-j) // 没有重复，更新长度
		}
		arr[s[i]] = i
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

## 剑指OfferII020.回文子字符串的个数(5)

- 题目

```
给定一个字符串 s ，请计算这个字符串中有多少个回文子字符串。
具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
示例 1：输入：s = "abc" 输出：3
解释：三个回文子串: "a", "b", "c"
示例 2：输入：s = "aaa" 输出：6
解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
提示：1 <= s.length <= 1000
s 由小写英文字母组成
注意：本题与主站 647 题相同：
```

- 解题思路

| No.  | 思路         | 时间复杂度 | 空间复杂度 |
| ---- | ------------ | ---------- | ---------- |
| 01   | 中心扩展     | O(n^2)     | O(1)       |
| 02   | Manacher算法 | O(n^2)     | O(1)       |
| 03   | Manacher算法 | O(n)       | O(n)       |
| 04   | 动态规划     | O(n^2)     | O(n^2)     |
| 05   | 暴力法       | O(n^3)     | O(1)       |

```go
func countSubstrings(s string) int {
	n := len(s)
	res := 0
	for i := 0; i < 2*n-1; i++ {
		left, right := i/2, i/2+i%2
		for ; 0 <= left && right < n && s[left] == s[right]; left, right = left-1, right+1 {
			res++
		}
	}
	return res
}

# 2
func countSubstrings(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	str := add(s)
	length := len(str)
	res := 0
	for i := 0; i < length; i++ {
		curLength := search(str, i)
		res = res + curLength/2 + curLength%2
	}
	return res
}

func add(s string) string {
	var res []rune
	for _, v := range s {
		res = append(res, '#')
		res = append(res, v)
	}
	res = append(res, '#')
	return string(res)
}

func search(s string, center int) int {
	i := center - 1
	j := center + 1
	step := 0
	for ; i >= 0 && j < len(s) && s[i] == s[j]; i, j = i-1, j+1 {
		step++
	}
	return step
}

# 3
func countSubstrings(s string) int {
	var res []rune
	res = append(res, '$')
	for _, v := range s {
		res = append(res, '#')
		res = append(res, v)
	}
	res = append(res, '#')
	res = append(res, '!')
	str := string(res)
	n := len(str) - 1
	arr := make([]int, n)
	leftMax, rightMax, result := 0, 0, 0
	for i := 1; i < n; i++ {
		if i <= rightMax {
			arr[i] = min(rightMax-i+1, arr[2*leftMax-i])
		} else {
			arr[i] = 1
		}
		for str[i+arr[i]] == str[i-arr[i]] {
			arr[i]++
		}
		if i+arr[i]-1 > rightMax {
			leftMax = i
			rightMax = i + arr[i] - 1
		}
		result = result + arr[i]/2
	}
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

# 4
func countSubstrings(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	dp := make([][]bool, len(s))
	res := 0
	for r := 0; r < len(s); r++ {
		dp[r] = make([]bool, len(s))
		dp[r][r] = true
		res++
		for l := 0; l < r; l++ {
			if s[l] == s[r] && (r-l <= 2 || dp[l+1][r-1] == true) {
				dp[l][r] = true
			} else {
				dp[l][r] = false
			}
			if dp[l][r] == true {
				res++
			}
		}
	}
	return res
}

# 5
func countSubstrings(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	res := len(s)
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] && judge(s, i, j) == true {
				res++
			}
		}
	}
	return res
}

func judge(s string, i, j int) bool {
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
```

## 剑指OfferII021.删除链表的倒数第n个结点(3)

- 题目

```
给定一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
示例 1：输入：head = [1,2,3,4,5], n = 2 输出：[1,2,3,5]
示例 2：输入：head = [1], n = 1 输出：[]
示例 3：输入：head = [1,2], n = 1 输出：[1]
提示：链表中结点的数目为 sz
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz
进阶：能尝试使用一趟扫描实现吗？
注意：本题与主站 19 题相同
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 遍历     | O(n)       | O(1)       |
| 02   | 快慢指针 | O(n)       | O(1)       |
| 03   | 递归     | O(n)       | O(n)       |

```go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp := &ListNode{Next: head}
	cur := temp
	total := 0
	for cur.Next != nil {
		cur = cur.Next
		total++
	}
	cur = temp
	count := 0
	for cur.Next != nil {
		if total-n == count {
			cur.Next = cur.Next.Next
			break
		}
		cur = cur.Next
		count++
	}
	return temp.Next
}

# 2
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp := &ListNode{Next: head}
	fast, slow := temp, temp
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return temp.Next
}

# 3
var count int

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		count = 0
		return nil
	}
	head.Next = removeNthFromEnd(head.Next, n)
	count = count + 1
	if count == n {
		return head.Next
	}
	return head
}
```

## 剑指OfferII022.链表中环的入口节点(3)

- 题目

```
给定一个链表，返回链表开始入环的第一个节点。 从链表的头节点开始沿着 next 指针进入环的第一个节点为环的入口节点。
如果链表无环，则返回 null。
为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 
如果 pos 是 -1，则在该链表中没有环。注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。
说明：不允许修改给定的链表。
示例 1：输入：head = [3,2,0,-4], pos = 1 输出：返回索引为 1 的链表节点
解释：链表中有一个环，其尾部连接到第二个节点。
示例 2：输入：head = [1,2], pos = 0 输出：返回索引为 0 的链表节点
解释：链表中有一个环，其尾部连接到第一个节点。
示例 3：输入：head = [1], pos = -1 输出：返回 null
解释：链表中没有环。
提示：链表中节点的数目范围在范围 [0, 104] 内
-105 <= Node.val <= 105
pos 的值为 -1 或者链表中的一个有效索引
进阶：是否可以使用 O(1) 空间解决此题？
注意：本题与主站 142 题相同
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

## 剑指OfferII025.链表中的两数相加(3)

- 题目

```
给定两个 非空链表 l1和 l2 来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。
将这两数相加会返回一个新的链表。
可以假设除了数字 0 之外，这两个数字都不会以零开头。
示例1：输入：l1 = [7,2,4,3], l2 = [5,6,4] 输出：[7,8,0,7]
示例2：输入：l1 = [2,4,3], l2 = [5,6,4] 输出：[8,0,7]
示例3：输入：l1 = [0], l2 = [0] 输出：[0]
提示：链表的长度范围为 [1, 100]
0 <= node.val <= 9
输入数据保证链表代表的数字无前导 0
进阶：如果输入链表不能修改该如何处理？换句话说，不能对列表中的节点进行翻转。
注意：本题与主站 445 题相同：
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 反转遍历 | O(n)       | O(n)       |
| 02   | 栈辅助   | O(n)       | O(n)       |
| 03   | 递归     | O(n)       | O(n)       |

```go
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverse(l1)
	l2 = reverse(l2)
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
	return reverse(res.Next)
}

func reverse(head *ListNode) *ListNode {
	var result *ListNode
	var temp *ListNode
	for head != nil {
		temp = head.Next
		head.Next = result
		result = head
		head = temp
	}
	return result
}

# 2
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1 := make([]int, 0)
	stack2 := make([]int, 0)
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}
	var res *ListNode
	carry := 0
	for len(stack1) > 0 || len(stack2) > 0 || carry > 0 {
		if len(stack1) > 0 {
			carry = carry + stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack2) > 0 {
			carry = carry + stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}
		temp := &ListNode{
			Val:  carry % 10,
			Next: res,
		}
		carry = carry / 10
		res = temp
	}
	return res
}

# 3
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	a, b := l1, l2
	length1, length2 := 0, 0
	for a != nil {
		length1++
		a = a.Next
	}
	for b != nil {
		length2++
		b = b.Next
	}
	res, carry := add(l1, l2, length1, length2)
	if carry > 0 {
		return &ListNode{Val: carry, Next: res}
	}
	return res
}

func add(l1, l2 *ListNode, length1, length2 int) (res *ListNode, carry int) {
	if l1 != nil && l2 != nil {
		if l1.Next == nil && l2.Next == nil {
			val := l1.Val + l2.Val
			carry = val / 10
			res = &ListNode{Val: val % 10, Next: nil}
			return
		}
	}
	a := &ListNode{}
	var b, n int
	if length1 > length2 {
		a, b = add(l1.Next, l2, length1-1, length2)
		n = l1.Val + b
	} else if length1 < length2 {
		a, b = add(l1, l2.Next, length1, length2-1)
		n = l2.Val + b
	} else {
		a, b = add(l1.Next, l2.Next, length1-1, length2-1)
		n = l1.Val + l2.Val + b
	}
	res = &ListNode{Val: n % 10, Next: a}
	carry = n / 10
	return
}
```

## 剑指OfferII026.重排链表(3)

- 题目

```
给定一个单链表 L 的头节点 head ，单链表 L 表示为：
L0 → L1 → … → Ln-1 → Ln 
请将其重新排列后变为：
L0 → Ln → L1 → Ln-1 → L2 → Ln-2 → …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
示例 1:输入: head = [1,2,3,4] 输出: [1,4,2,3]
示例 2:输入: head = [1,2,3,4,5] 输出: [1,5,2,4,3]
提示：链表的长度范围为 [1, 5 * 104]
1 <= node.val <= 1000
注意：本题与主站 143 题相同
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 数组辅助 | O(n)       | O(n)       |
| 02   | 反转链表 | O(n)       | O(1)       |
| 03   | 递归     | O(n)       | O(n)       |

```go
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	cur := head
	arr := make([]*ListNode, 0)
	for cur != nil {
		arr = append(arr, cur)
		cur = cur.Next
	}
	res := make([]*ListNode, 0)
	for i := 0; i < len(arr)/2; i++ {
		res = append(res, arr[i], arr[len(arr)-1-i])
	}
	if len(arr)%2 == 1 {
		res = append(res, arr[len(arr)/2])
	}
	cur = head
	for i := 1; i < len(res); i++ {
		cur.Next = res[i]
		cur = cur.Next
	}
	cur.Next = nil
}

# 2
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	second := reverse(slow.Next)
	slow.Next = nil
	cur := head
	count := 0
	for cur != nil && second != nil {
		a := cur.Next
		b := second.Next
		if count%2 == 0 {
			cur.Next = second
			cur = a
		} else {
			second.Next = cur
			second = b
		}
		count++
	}
}

func reverse(head *ListNode) *ListNode {
	var res *ListNode
	for head != nil {
		next := head.Next
		head.Next = res
		res = head
		head = next
	}
	return res
}

# 3
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}
	helper(head, length)
}

func helper(head *ListNode, length int) *ListNode {
	if length == 1 {
		next := head.Next
		head.Next = nil
		return next
	}
	if length == 2 {
		next := head.Next.Next
		head.Next.Next = nil
		return next
	}
	tail := helper(head.Next, length-2)
	next := tail.Next
	temp := head.Next
	head.Next = tail
	tail.Next = temp
	return next
}
```

## 剑指OfferII028.展平多级双向链表(3)

- 题目

```
多级双向链表中，除了指向下一个节点和前一个节点指针之外，它还有一个子链表指针，可能指向单独的双向链表。
这些子列表也可能会有一个或多个自己的子项，依此类推，生成多级数据结构，如下面的示例所示。
给定位于列表第一级的头节点，请扁平化列表，即将这样的多级双向链表展平成普通的双向链表，使所有结点出现在单级双链表中。
示例 1：输入：head = [1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12] 输出：[1,2,3,7,8,11,12,9,10,4,5,6]
解释：输入的多级列表如下图所示：
扁平化后的链表如下图：
示例 2：输入：head = [1,2,null,3] 输出：[1,3,2]
解释：输入的多级列表如下图所示：
  1---2---NULL
  |
  3---NULL
示例 3：输入：head = [] 输出：[]
如何表示测试用例中的多级链表？
以 示例 1 为例：
 1---2---3---4---5---6--NULL
         |
         7---8---9---10--NULL
             |
             11--12--NULL
序列化其中的每一级之后：
[1,2,3,4,5,6,null]
[7,8,9,10,null]
[11,12,null]
为了将每一级都序列化到一起，我们需要每一级中添加值为 null 的元素，以表示没有节点连接到上一级的上级节点。
[1,2,3,4,5,6,null]
[null,null,7,8,9,10,null]
[null,11,12,null]
合并所有序列化结果，并去除末尾的 null 。
[1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]
提示：节点数目不超过 1000
1 <= Node.val <= 10^5
注意：本题与主站 430 题相同
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 递归 | O(n)       | O(n)       |
| 02   | 递归 | O(n)       | O(n)       |
| 03   | 迭代 | O(n)       | O(n)       |

```go
func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	res := &Node{}
	cur := res
	for root != nil {
		cur.Next = root
		root.Prev = cur
		cur = cur.Next
		root = root.Next
		// 处理子节点
		if cur.Child != nil {
			ch := flatten(cur.Child)
			cur.Child = nil
			cur.Next = ch
			ch.Prev = cur
			// 指针移动
			for cur.Next != nil {
				cur = cur.Next
			}
		}
	}
	res.Next.Prev = nil
	return res.Next
}

# 2
var arr []*Node

func flatten(root *Node) *Node {
	arr = make([]*Node, 0)
	dfs(root)
	for i := 0; i < len(arr); i++ {
		if i+1 < len(arr) {
			arr[i].Next = arr[i+1]
		}
		if i > 0 {
			arr[i].Prev = arr[i-1]
		}
		arr[i].Child = nil
	}
	return root
}

func dfs(root *Node) {
	if root == nil {
		return
	}
	arr = append(arr, root)
	dfs(root.Child)
	dfs(root.Next)
}

# 3
func flatten(root *Node) *Node {
	cur := root
	stack := make([]*Node, 0)
	for cur != nil {
		// 处理child
		if cur.Child != nil {
			if cur.Next != nil {
				stack = append(stack, cur.Next)
			}
			cur.Child.Prev = cur
			cur.Next = cur.Child
			cur.Child = nil
			continue
		}
		if cur.Next != nil {
            cur.Child = nil
			cur = cur.Next
			continue
		}
		if len(stack) == 0 {
			break
		}
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur.Next = last
		last.Prev = cur
		cur = last
	}
	return root
}
```

## 剑指OfferII029.排序的循环链表

### 题目

```

```

### 解题思路

```go

```

## 剑指OfferII030.插入、删除和随机访问都是O(1)的容器(2)

- 题目

```
设计一个支持在平均 时间复杂度 O(1) 下，执行以下操作的数据结构：
insert(val)：当元素 val 不存在时返回 true ，并向集合中插入该项，否则返回 false 。
remove(val)：当元素 val 存在时返回 true ，并从集合中移除该项，否则返回 false 。
getRandom：随机返回现有集合中的一项。每个元素应该有 相同的概率 被返回。
示例 :输入: inputs = ["RandomizedSet", "insert", 
"remove", "insert", "getRandom", "remove", "insert", "getRandom"]
[[], [1], [2], [2], [], [1], [2], []]
输出: [null, true, false, true, 2, true, false, 2]
解释:RandomizedSet randomSet = new RandomizedSet();  // 初始化一个空的集合
randomSet.insert(1); // 向集合中插入 1 ， 返回 true 表示 1 被成功地插入
randomSet.remove(2); // 返回 false，表示集合中不存在 2 
randomSet.insert(2); // 向集合中插入 2 返回 true ，集合现在包含 [1,2] 
randomSet.getRandom(); // getRandom 应随机返回 1 或 2 
randomSet.remove(1); // 从集合中移除 1 返回 true 。集合现在包含 [2] 
randomSet.insert(2); // 2 已在集合中，所以返回 false 
randomSet.getRandom(); // 由于 2 是集合中唯一的数字，getRandom 总是返回 2 
提示：-231 <= val <= 231 - 1
最多进行 2 * 105 次 insert ， remove 和 getRandom 方法调用
当调用 getRandom 方法时，集合中至少有一个元素
注意：本题与主站 380 题相同
```

- 解题思路

| No.  | 思路        | 时间复杂度 | 空间复杂度 |
| ---- | ----------- | ---------- | ---------- |
| 01   | 哈希表+数组 | O(1)       | O(n)       |
| 02   | 哈希表      | O(n)       | O(n)       |

```go
type RandomizedSet struct {
	m   map[int]int
	arr []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		m:   make(map[int]int),
		arr: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.arr = append(this.arr, val)
	this.m[val] = len(this.arr) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.m[val]; !ok {
		return false
	}
	index := this.m[val]
	this.arr[index], this.arr[len(this.arr)-1] = this.arr[len(this.arr)-1], this.arr[index]
	this.m[this.arr[index]] = index
	this.arr = this.arr[:len(this.arr)-1]
	delete(this.m, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	if len(this.arr) == 0 {
		return -1
	}
	index := rand.Intn(len(this.arr))
	return this.arr[index]
}

# 2
type RandomizedSet struct {
	m map[int]bool
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		m: make(map[int]bool),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.m[val] = true
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.m[val]; !ok {
		return false
	}
	delete(this.m, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	if len(this.m) == 0 {
		return -1
	}
	index := rand.Intn(len(this.m))
	res := -1
	for res = range this.m {
		index--
		if index == -1 {
			break
		}
	}
	return res
}
```

## 剑指OfferII031.最近最少使用缓存(1)

- 题目

```
运用所掌握的数据结构，设计和实现一个  LRU (Least Recently Used，最近最少使用) 缓存机制 。
实现 LRUCache 类：
LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。
当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
示例：输入 ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出[null, null, null, 1, null, -1, null, -1, 3, 4]
解释 LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4
提示：1 <= capacity <= 3000
0 <= key <= 10000
0 <= value <= 105
最多调用 2 * 105 次 get 和 put
进阶：是否可以在 O(1) 时间复杂度内完成这两种操作？
注意：本题与主站 146 题相同
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 双向链表 | O(1)       | O(n)       |

```go
type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

type LRUCache struct {
	cap    int
	header *Node
	tail   *Node
	m      map[int]*Node
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		cap:    capacity,
		header: &Node{},
		tail:   &Node{},
		m:      make(map[int]*Node, capacity),
	}
	cache.header.next = cache.tail
	cache.tail.prev = cache.header
	return cache
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; ok {
		this.remove(node)
		this.putHead(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.m[key]; ok {
		node.value = value
		this.remove(node)
		this.putHead(node)
		return
	}
	if this.cap <= len(this.m) {
		// 删除尾部
		deleteKey := this.tail.prev.key
		this.remove(this.tail.prev)
		delete(this.m, deleteKey)
	}
	// 插入到头部
	newNode := &Node{key: key, value: value}
	this.putHead(newNode)
	this.m[key] = newNode
}

// 删除尾部节点
func (this *LRUCache) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 插入头部
func (this *LRUCache) putHead(node *Node) {
	next := this.header.next
	this.header.next = node
	node.next = next
	next.prev = node
	node.prev = this.header
}
```

## 剑指OfferII033.变位词组(2)

- 题目

```
给定一个字符串数组 strs ，将 变位词 组合在一起。 可以按任意顺序返回结果列表。
注意：若两个字符串中每个字符出现的次数都相同，则称它们互为变位词。
示例 1:输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"] 
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
示例 2:输入: strs = [""] 输出: [[""]]
示例 3:输入: strs = ["a"] 输出: [["a"]]
提示：1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] 仅包含小写字母
注意：本题与主站 49 题相同：
```

- 解题思路

| No.  | 思路     | 时间复杂度   | 空间复杂度 |
| ---- | -------- | ------------ | ---------- |
| 01   | 哈希辅助 | O(n^2log(n)) | O(n^2)     |
| 02   | 哈希辅助 | O(n^2)       | O(n^2)     |

```go
func groupAnagrams(strs []string) [][]string {
	m := make(map[string]int)
	res := make([][]string, 0)
	for i := 0; i < len(strs); i++ {
		arr := []byte(strs[i])
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
		newStr := string(arr)
		if _, ok := m[newStr]; ok {
			res[m[newStr]] = append(res[m[newStr]], strs[i])
		} else {
			m[newStr] = len(res)
			res = append(res, []string{strs[i]})
		}
	}
	return res
}

#
func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int]int)
	res := make([][]string, 0)
	for i := 0; i < len(strs); i++ {
		arr := [26]int{}
		for j := 0; j < len(strs[i]); j++{
			arr[strs[i][j]-'a']++
		}
		if _, ok := m[arr]; ok {
			res[m[arr]] = append(res[m[arr]], strs[i])
		} else {
			m[arr] = len(res)
			res = append(res, []string{strs[i]})
		}
	}
	return res
}
```

## 剑指OfferII035.最小时间差(2)

- 题目

```
给定一个 24 小时制（小时:分钟 "HH:MM"）的时间列表，找出列表中任意两个时间的最小时间差并以分钟数表示。
示例 1：输入：timePoints = ["23:59","00:00"] 输出：1
示例 2：输入：timePoints = ["00:00","23:59","00:00"] 输出：0
提示：2 <= timePoints <= 2 * 104
timePoints[i] 格式为 "HH:MM"
注意：本题与主站 539 题相同： 
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 排序 | O(nlog(n)) | O(n)       |
| 02   | 排序 | O(nlog(n)) | O(n)       |

```go
func findMinDifference(timePoints []string) int {
	m := make(map[int]bool)
	for i := 0; i < len(timePoints); i++ {
		value := getValue(timePoints[i])
		if _, ok := m[value]; ok {
			return 0
		}
		m[value] = true
	}
	arr := make([]int, 0)
	for k := range m {
		arr = append(arr, k)
	}
	sort.Ints(arr)
	res := math.MaxInt32
	arr = append(arr, arr[0]+1440)
	for i := 1; i < len(arr); i++ {
		if res > arr[i]-arr[i-1] {
			res = arr[i] - arr[i-1]
		}
	}
	return res
}

func getValue(str string) int {
	hour, _ := strconv.Atoi(str[:2])
	minute, _ := strconv.Atoi(str[3:])
	return hour*60 + minute
}

# 2
func findMinDifference(timePoints []string) int {
	arr := make([]int, 0)
	for i := 0; i < len(timePoints); i++ {
		value := getValue(timePoints[i])
		arr = append(arr, value)
	}
	sort.Ints(arr)
	res := math.MaxInt32
	arr = append(arr, arr[0]+1440)
	for i := 1; i < len(arr); i++ {
		if res > arr[i]-arr[i-1] {
			res = arr[i] - arr[i-1]
		}
	}
	return res
}

func getValue(str string) int {
	hour, _ := strconv.Atoi(str[:2])
	minute, _ := strconv.Atoi(str[3:])
	return hour*60 + minute
}
```

## 剑指OfferII036.后缀表达式(1)

- 题目

```
根据 逆波兰表示法，求该后缀表达式的计算结果。
有效的算符包括 +、-、*、/ 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。
说明：整数除法只保留整数部分。
给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。
示例 1：输入：tokens = ["2","1","+","3","*"] 输出：9
解释：该算式转化为常见的中缀算术表达式为：((2 + 1) * 3) = 9
示例 2：输入：tokens = ["4","13","5","/","+"] 输出：6
解释：该算式转化为常见的中缀算术表达式为：(4 + (13 / 5)) = 6
示例 3：输入：tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"] 输出：22
解释：该算式转化为常见的中缀算术表达式为：
  ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22
提示：1 <= tokens.length <= 104
tokens[i] 要么是一个算符（"+"、"-"、"*" 或 "/"），要么是一个在范围 [-200, 200] 内的整数
逆波兰表达式：逆波兰表达式是一种后缀表达式，所谓后缀就是指算符写在后面。
平常使用的算式则是一种中缀表达式，如 ( 1 + 2 ) * ( 3 + 4 ) 。
该算式的逆波兰表达式写法为 ( ( 1 2 + ) ( 3 4 + ) * ) 。
逆波兰表达式主要有以下两个优点：
去掉括号后表达式无歧义，上式即便写成 1 2 + 3 4 + * 也可以依据次序计算出正确结果。
适合用栈操作运算：遇到数字则入栈；遇到算符则取出栈顶两个数字进行计算，并将结果压入栈中。
注意：本题与主站 150 题相同：
```

- 解题思路

| No.  | 思路   | 时间复杂度 | 空间复杂度 |
| ---- | ------ | ---------- | ---------- |
| 01   | 栈辅助 | O(n)       | O(n)       |

```go
func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, v := range tokens {
		length := len(stack)
		if v == "+" || v == "-" || v == "*" || v == "/" {
			a := stack[length-2]
			b := stack[length-1]
			stack = stack[:length-2]
			var value int
			if v == "+" {
				value = a + b
			} else if v == "-" {
				value = a - b
			} else if v == "*" {
				value = a * b
			} else {
				value = a / b
			}
			stack = append(stack, value)
		} else {
			value, _ := strconv.Atoi(v)
			stack = append(stack, value)
		}
	}
	return stack[0]
}
```

## 剑指OfferII037.小行星碰撞(2)

- 题目

```
给定一个整数数组 asteroids，表示在同一行的小行星。
对于数组中的每一个元素，其绝对值表示小行星的大小，正负表示小行星的移动方向（正表示向右移动，负表示向左移动）。
每一颗小行星以相同的速度移动。
找出碰撞后剩下的所有小行星。碰撞规则：两个行星相互碰撞，较小的行星会爆炸。如果两颗行星大小相同，则两颗行星都会爆炸。
两颗移动方向相同的行星，永远不会发生碰撞。
示例 1：输入：asteroids = [5,10,-5] 输出：[5,10]
解释：10 和 -5 碰撞后只剩下 10 。 5 和 10 永远不会发生碰撞。
示例 2：输入：asteroids = [8,-8] 输出：[]
解释：8 和 -8 碰撞后，两者都发生爆炸。
示例 3：输入：asteroids = [10,2,-5] 输出：[10]
解释：2 和 -5 发生碰撞后剩下 -5 。10 和 -5 发生碰撞后剩下 10 。
示例 4：输入：asteroids = [-2,-1,1,2] 输出：[-2,-1,1,2]
解释：-2 和 -1 向左移动，而 1 和 2 向右移动。 由于移动方向相同的行星不会发生碰撞，所以最终没有行星发生碰撞。 
提示：2 <= asteroids.length <= 104
-1000 <= asteroids[i] <= 1000
asteroids[i] != 0
注意：本题与主站 735 题相同：
```

- 解题思路

| No.  | 思路   | 时间复杂度 | 空间复杂度 |
| ---- | ------ | ---------- | ---------- |
| 01   | 栈辅助 | O(n)       | O(n)       |

```go
func asteroidCollision(asteroids []int) []int {
	left := make([]int, 0)
	right := make([]int, 0)
	for i := 0; i < len(asteroids); i++ {
		if asteroids[i] > 0 {
			right = append(right, asteroids[i])
		} else {
			if len(right) > 0 {
				for {
					if len(right) == 0 {
						left = append(left, asteroids[i])
						break
					}
					sum := asteroids[i] + right[len(right)-1]
					if sum == 0 {
						right = right[:len(right)-1]
						break
					} else if sum > 0 {
						break
					} else {
						right = right[:len(right)-1]
					}
				}
			} else {
				left = append(left, asteroids[i])
			}
		}

	}
	return append(left, right...)
}

# 2
func asteroidCollision(asteroids []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(asteroids); i++ {
		value := asteroids[i]
		for value < 0 && len(res) > 0 && res[len(res)-1] > 0 {
			sum := value + res[len(res)-1]
			if sum >= 0 {
				value = 0
			}
			if sum <= 0 {
				res = res[:len(res)-1]
			}
		}
		if value != 0 {
			res = append(res, value)
		}
	}
	return res
}
```

## 剑指OfferII038.每日温度(3)

- 题目

```
请根据每日 气温 列表 temperatures ，重新生成一个列表，要求其对应位置的输出为：要想观测到更高的气温，至少需要等待的天数。
如果气温在这之后都不会升高，请在该位置用 0 来代替。
示例 1:输入: temperatures = [73,74,75,71,69,72,76,73] 输出: [1,1,4,2,1,1,0,0]
示例 2:输入: temperatures = [30,40,50,60] 输出: [1,1,1,0]
示例 3:输入: temperatures = [30,60,90] 输出: [1,1,0]
提示：1 <= temperatures.length <= 105
30 <= temperatures[i] <= 100
注意：本题与主站 739 题相同
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 栈辅助   | O(n)       | O(n)       |
| 02   | 数组辅助 | O(n)       | O(n)       |
| 03   | 暴力法   | O(n^2)     | O(1)       |

```go
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := make([]int, 0) // 栈保存递减数据的下标
	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[last] = i - last
		}
		stack = append(stack, i)
	}
	return res
}

# 2
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	arr := make([]int, 101)
	for i := 0; i < len(arr); i++ {
		arr[i] = math.MaxInt64
	}
	for i := len(temperatures) - 1; i >= 0; i-- {
		temp := math.MaxInt64
		for t := temperatures[i] + 1; t < 101; t++ {
			if arr[t] < temp {
				temp = arr[t]
			}
		}
		if temp < math.MaxInt64 {
			res[i] = temp - i
		}
		arr[temperatures[i]] = i
	}
	return res
}

# 3
func dailyTemperatures(temperatures []int) []int {
	j := 0
	for i := 0; i < len(temperatures); i++ {
		for j = i + 1; j < len(temperatures); j++ {
			if temperatures[j] > temperatures[i] {
				temperatures[i] = j - i
				break
			}
		}
		if j == len(temperatures) {
			temperatures[i] = 0
		}
	}
	return temperatures
}
```

## 剑指OfferII043.往完全二叉树添加节点(1)

- 题目

```
完全二叉树是每一层（除最后一层外）都是完全填充（即，节点数达到最大，第 n 层有 2n-1 个节点）的，
并且所有的节点都尽可能地集中在左侧。
设计一个用完全二叉树初始化的数据结构 CBTInserter，它支持以下几种操作：
CBTInserter(TreeNode root) 使用根节点为 root 的给定树初始化该数据结构；
CBTInserter.insert(int v)  向树中插入一个新节点，节点类型为 TreeNode，值为 v 。
使树保持完全二叉树的状态，并返回插入的新节点的父节点的值；
CBTInserter.get_root() 将返回树的根节点。
示例 1：输入：inputs = ["CBTInserter","insert","get_root"], inputs = [[[1]],[2],[]] 输出：[null,1,[1,2]]
示例 2：输入：inputs = ["CBTInserter","insert","insert","get_root"], inputs = [[[1,2,3,4,5,6]],[7],[8],[]]
输出：[null,3,4,[1,2,3,4,5,6,7,8]]
提示：最初给定的树是完全二叉树，且包含 1 到 1000 个节点。
每个测试用例最多调用 CBTInserter.insert  操作 10000 次。
给定节点或插入节点的每个值都在 0 到 5000 之间。
注意：本题与主站 919 题相同：
```

- 解题思路

| No.  | 思路         | 时间复杂度 | 空间复杂度 |
| ---- | ------------ | ---------- | ---------- |
| 01   | 广度优先搜索 | O(n)       | O(n)       |

```go
type CBTInserter struct {
	root *TreeNode
	arr  []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	arr := make([]*TreeNode, 0)
	queue := make([]*TreeNode, 0)
	arr = append(arr, root)
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
				arr = append(arr, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
				arr = append(arr, queue[i].Right)
			}
		}
		queue = queue[length:]
	}
	return CBTInserter{root: root, arr: arr}
}

func (this *CBTInserter) Insert(v int) int {
	newNode := &TreeNode{Val: v}
	this.arr = append(this.arr, newNode)
	n := len(this.arr)
	target := this.arr[n/2-1]
	if n%2 == 0 {
		target.Left = newNode
	} else {
		target.Right = newNode
	}
	return target.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}
```

## 剑指OfferII044.二叉树每层的最大值(2)

- 题目

```
给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。
示例1：输入: root = [1,3,2,5,3,null,9] 输出: [1,3,9]
解释:
          1
         / \
        3   2
       / \   \  
      5   3   9 
示例2：输入: root = [1,2,3] 输出: [1,3]
解释:
          1
         / \
        2   3
示例3：输入: root = [1] 输出: [1]
示例4：输入: root = [1,null,2] 输出: [1,2]
解释:      
           1 
            \
             2     
示例5：输入: root = [] 输出: []
提示：二叉树的节点个数的范围是 [0,104]
-231 <= Node.val <= 231 - 1
注意：本题与主站 515 题相同：
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 层序遍历 | O(n)       | O(n)       |
| 02   | 递归     | O(n)       | O(n)       |

```go
func largestValues(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		maxValue := math.MinInt32
		for i := 0; i < length; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			if maxValue < queue[i].Val {
				maxValue = queue[i].Val
			}
		}
		res = append(res, maxValue)
		queue = queue[length:]
	}
	return res
}

# 2
var res []int

func largestValues(root *TreeNode) []int {
	res = make([]int, 0)
	if root == nil {
		return res
	}
	dfs(root, 0)
	return res
}

func dfs(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if level >= len(res) {
		res = append(res, math.MinInt32)
	}
	res[level] = max(res[level], root.Val)
	dfs(root.Left, level+1)
	dfs(root.Right, level+1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

## 剑指OfferII045.二叉树最底层最左边的值(2)

- 题目

```
给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
假设二叉树中至少有一个节点。
示例 1:输入: root = [2,1,3] 输出: 1
示例 2:输入: [1,2,3,4,null,5,6,null,null,7] 输出: 7
提示:二叉树的节点个数的范围是 [1,104]
-231 <= Node.val <= 231 - 1 
注意：本题与主站 513 题相同：
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 层序遍历 | O(n)       | O(n)       |
| 02   | 递归     | O(n)       | O(log(n))  |

```go
func findBottomLeftValue(root *TreeNode) int {
	res := 0
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		res = queue[0].Val
		for i := 0; i < length; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[length:]
	}
	return res
}

# 2
var res int
var maxLevel int

func findBottomLeftValue(root *TreeNode) int {
	res = 0
	maxLevel = -1
	if root == nil {
		return res
	}
	dfs(root, 0)
	return res
}

func dfs(root *TreeNode, level int) {
	if root == nil {
		return
	}
	dfs(root.Left, level+1)
	if level > maxLevel {
		maxLevel = level
		res = root.Val
	}
	dfs(root.Right, level+1)
}
```

## 剑指OfferII046.二叉树的右侧视图(2)

- 题目

```
给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
示例 1:输入: [1,2,3,null,5,null,4] 输出: [1,3,4]
示例 2:输入: [1,null,3] 输出: [1,3]
示例 3:输入: [] 输出: []
提示:二叉树的节点个数的范围是 [0,100]
-100 <= Node.val <= 100 
注意：本题与主站 199 题相同：
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 迭代 | O(n)       | O(n)       |
| 02   | 递归 | O(n)       | O(log(n))  |

```go
func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil{
		return res
	}
	list := make([]*TreeNode, 0)
	list = append(list, root)
	for len(list) > 0 {
		length := len(list)
		res = append(res, list[0].Val)
		for i := 0; i < length; i++ {
			node := list[i]
			if node.Right != nil {
				list = append(list, node.Right)
			}
			if node.Left != nil {
				list = append(list, node.Left)
			}
		}
		list = list[length:]
	}
	return res
}

#
var res []int

func rightSideView(root *TreeNode) []int {
	res = make([]int, 0)
	if root == nil {
		return res
	}
	dfs(root, 1)
	return res
}

func dfs(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if level > len(res) {
		res = append(res, root.Val)
	}
	dfs(root.Right, level+1)
	dfs(root.Left, level+1)
}
```

## 剑指OfferII047.二叉树剪枝(1)

- 题目

```
给定一个二叉树 根节点 root ，树的每个节点的值要么是 0，要么是 1。请剪除该二叉树中所有节点的值为 0 的子树。
节点 node 的子树为 node 本身，以及所有 node 的后代。
示例 1:输入: [1,null,0,0,1] 输出: [1,null,0,null,1] 
解释: 只有红色节点满足条件“所有不包含 1 的子树”。
右图为返回的答案。
示例 2:输入: [1,0,1,0,0,0,1] 输出: [1,null,1,null,1]
解释: 
示例 3:输入: [1,1,0,1,1,0,1,0] 输出: [1,1,0,1,1,null,1]
解释: 
提示:二叉树的节点个数的范围是 [1,200]
二叉树节点的值只会是 0 或 1
注意：本题与主站 814 题相同：
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 递归 | O(n)       | O(log(n))  |

```go
func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}
	return root
}
```

## 剑指OfferII049.从根节点到叶节点的路径数字之和(2)

- 题目

```
给定一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
每条从根节点到叶节点的路径都代表一个数字：
例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
计算从根节点到叶节点生成的 所有数字之和 。
叶节点 是指没有子节点的节点。
示例 1：输入：root = [1,2,3] 输出：25
解释：从根到叶子节点路径 1->2 代表数字 12
从根到叶子节点路径 1->3 代表数字 13
因此，数字总和 = 12 + 13 = 25
示例 2：输入：root = [4,9,0,5,1] 输出：1026
解释：从根到叶子节点路径 4->9->5 代表数字 495
从根到叶子节点路径 4->9->1 代表数字 491
从根到叶子节点路径 4->0 代表数字 40
因此，数字总和 = 495 + 491 + 40 = 1026
提示：树中节点的数目在范围 [1, 1000] 内
0 <= Node.val <= 9
树的深度不超过 10
注意：本题与主站 129 题相同： 
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 递归 | O(n)       | O(log(n))  |
| 02   | 迭代 | O(n)       | O(n)       |

```go
var res int

func sumNumbers(root *TreeNode) int {
	res = 0
	dfs(root, 0)
	return res
}

func dfs(root *TreeNode, sum int) {
	if root == nil {
		return
	}
	sum = sum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		res = res + sum
	}
	dfs(root.Left, sum)
	dfs(root.Right, sum)
}

#
func sumNumbers(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}
	list := make([]*TreeNode, 0)
	list = append(list, root)
	for len(list) > 0 {
		length := len(list)
		for i := 0; i < length; i++ {
			node := list[i]
			value := node.Val
			if node.Left == nil && node.Right == nil {
				res = res + value
			}
			if node.Left != nil {
				node.Left.Val = node.Left.Val + value*10
				list = append(list, node.Left)
			}
			if node.Right != nil {
				node.Right.Val = node.Right.Val + value*10
				list = append(list, node.Right)
			}
		}
		list = list[length:]
	}
	return res
}
```

## 剑指OfferII050.向下的路径节点之和(4)

- 题目

```
给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
示例 1：输入：root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8 输出：3
解释：和等于 8 的路径有 3 条，如图所示。
示例 2：输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22 输出：3
提示:二叉树的节点个数的范围是 [0,1000]
-109 <= Node.val <= 109 
-1000 <= targetSum <= 1000 
注意：本题与主站 437 题相同
```

- 解题思路

| No.  | 思路      | 时间复杂度 | 空间复杂度 |
| ---- | --------- | ---------- | ---------- |
| 01   | 递归      | O(n^2)     | O(n)       |
| 02   | 2次递归   | O(n^2)     | O(n)       |
| 03   | 迭代+递归 | O(n^2)     | O(n)       |
| 04   | 保存路径  | O(n^2)     | O(n)       |

```go
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := 0
	var helper func(*TreeNode, int)
	helper = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		sum = sum - node.Val
		// 路径不需要从根节点开始，也不需要在叶子节点结束
		if sum == 0 {
			res++
		}
		helper(node.Left, sum)
		helper(node.Right, sum)
	}
	helper(root, targetSum)
	return res + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

# 2
func helper(node *TreeNode, targetSum int) int {
	if node == nil {
		return 0
	}
	targetSum = targetSum - node.Val
	res := 0
	if targetSum == 0 {
		res = 1
	}
	return res + helper(node.Left, targetSum) + helper(node.Right, targetSum)
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return helper(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

# 3
func helper(node *TreeNode, targetSum int, curSum int) int {
	res := 0
	curSum = curSum + node.Val
	if curSum == targetSum {
		res++
	}
	if node.Left != nil {
		res += helper(node.Left, targetSum, curSum)
	}
	if node.Right != nil {
		res += helper(node.Right, targetSum, curSum)
	}
	return res
}

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
		res += helper(node, sum, tempSum)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return res
}

# 4
func helper(node *TreeNode, sum int, path []int, level int) int {
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
	return res + helper(node.Left, sum, path, level+1) +
		helper(node.Right, sum, path, level+1)
}

func pathSum(root *TreeNode, targetSum int) int {
	return helper(root, targetSum, make([]int, 1001), 0)
}
```

## 剑指OfferII053.二叉搜索树中的中序后继

### 题目

```

```

### 解题思路

```go

```

## 剑指OfferII054.所有大于等于节点的值之和(2)

- 题目

```
给定一个二叉搜索树，请将它的每个节点的值替换成树中大于或者等于该节点值的所有节点值之和。
提醒一下，二叉搜索树满足下列约束条件：
节点的左子树仅包含键 小于 节点键的节点。
节点的右子树仅包含键 大于 节点键的节点。
左右子树也必须是二叉搜索树。
示例 1：输入：root = [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
输出：[30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
示例 2：输入：root = [0,null,1] 输出：[1,null,1]
示例 3：输入：root = [1,0,2] 输出：[3,3,2]
示例 4：输入：root = [3,2,4,1] 输出：[7,9,4,10]
提示：树中的节点数介于 0 和 104 之间。
每个节点的值介于 -104 和 104 之间。
树中的所有值 互不相同 。
给定的树为二叉搜索树。
注意：本题与主站 538 题相同：
本题与主站 1038 题相同：
```

- 解题思路

| No.  | 思路   | 时间复杂度 | 空间复杂度 |
| ---- | ------ | ---------- | ---------- |
| 01   | 递归   | O(n)       | O(log(n))  |
| 02   | 栈辅助 | O(n)       | O(n)       |

```go
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	dfs(root, &sum)
	return root
}

func dfs(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	dfs(root.Right, sum)
	*sum = *sum + root.Val
	root.Val = *sum
	dfs(root.Left, sum)
}

# 2
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	stack := make([]*TreeNode, 0)
	temp := root
	sum := 0
	for {
		if temp != nil {
			stack = append(stack, temp)
			temp = temp.Right
		} else if len(stack) != 0 {
			temp = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			temp.Val = temp.Val + sum
			sum = temp.Val
			temp = temp.Left
		} else {
			break
		}
	}
	return root
}
```

## 剑指OfferII055.二叉搜索树迭代器(2)

- 题目

```
实现一个二叉搜索树迭代器类BSTIterator ，表示一个按中序遍历二叉搜索树（BST）的迭代器：
BSTIterator(TreeNode root) 初始化 BSTIterator 类的一个对象。BST 的根节点 root 会作为构造函数的一部分给出。
指针应初始化为一个不存在于 BST 中的数字，且该数字小于 BST 中的任何元素。
boolean hasNext() 如果向指针右侧遍历存在数字，则返回 true ；否则返回 false 。
int next()将指针向右移动，然后返回指针处的数字。
注意，指针初始化为一个不存在于 BST 中的数字，所以对 next() 的首次调用将返回 BST 中的最小元素。
可以假设 next() 调用总是有效的，也就是说，当调用 next() 时，BST 的中序遍历中至少存在一个下一个数字。
示例：输入inputs = ["BSTIterator", "next", "next", "hasNext", "next", "hasNext", 
"next", "hasNext", "next", "hasNext"]
inputs = [[[7, 3, 15, null, null, 9, 20]], [], [], [], [], [], [], [], [], []]
输出[null, 3, 7, true, 9, true, 15, true, 20, false]
解释 BSTIterator bSTIterator = new BSTIterator([7, 3, 15, null, null, 9, 20]);
bSTIterator.next();    // 返回 3
bSTIterator.next();    // 返回 7
bSTIterator.hasNext(); // 返回 True
bSTIterator.next();    // 返回 9
bSTIterator.hasNext(); // 返回 True
bSTIterator.next();    // 返回 15
bSTIterator.hasNext(); // 返回 True
bSTIterator.next();    // 返回 20
bSTIterator.hasNext(); // 返回 False
提示：树中节点的数目在范围 [1, 105] 内
0 <= Node.val <= 106
最多调用 105 次 hasNext 和 next 操作
进阶：你可以设计一个满足下述条件的解决方案吗？next() 和 hasNext() 操作均摊时间复杂度为 O(1) ，并使用 O(h) 内存。
其中 h 是树的高度。
注意：本题与主站 173 题相同： 
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 数组辅助 | O(1)       | O(n)       |
| 02   | 栈辅助   | O(1)       | O(n)       |

```go
type BSTIterator struct {
	arr  []int
	root *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	arr := make([]int, 0)
	inorder(root, &arr)
	return BSTIterator{
		arr:  arr,
		root: root,
	}
}

func inorder(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, nums)
	*nums = append(*nums, root.Val)
	inorder(root.Right, nums)
}

func (this *BSTIterator) Next() int {
	if len(this.arr) == 0 {
		return -1
	}
	res := this.arr[0]
	this.arr = this.arr[1:]
	return res
}

func (this *BSTIterator) HasNext() bool {
	if len(this.arr) > 0 {
		return true
	}
	return false
}

# 2
type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	res := BSTIterator{}
	res.left(root)
	return res
}

func (this *BSTIterator) left(root *TreeNode) {
	for root != nil {
		this.stack = append(this.stack, root)
		root = root.Left
	}
}

func (this *BSTIterator) Next() int {
	node := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if node.Right != nil {
		this.left(node.Right)
	}
	return node.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}
```

## 剑指OfferII057.值和下标之差都在给定的范围内(2)

- 题目

```
给你一个整数数组 nums 和两个整数 k 和 t 。请你判断是否存在 两个不同下标 i 和 j，
使得 abs(nums[i] - nums[j]) <= t ，同时又满足 abs(i - j) <= k 。
如果存在则返回 true，不存在返回 false。
示例 1：输入：nums = [1,2,3,1], k = 3, t = 0 输出：true
示例 2：输入：nums = [1,0,1,1], k = 1, t = 2 输出：true
示例 3：输入：nums = [1,5,9,1,5,9], k = 2, t = 3 输出：false
提示：0 <= nums.length <= 2 * 104
-231 <= nums[i] <= 231 - 1
0 <= k <= 104
0 <= t <= 231 - 1
注意：本题与主站 220 题相同： 
```

- 解题思路

| No.  | 思路   | 时间复杂度 | 空间复杂度 |
| ---- | ------ | ---------- | ---------- |
| 01   | 暴力法 | O(n^2)     | O(1)       |
| 02   | 桶     | O(n)       | O(n)       |

```go
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if len(nums) <= 1 {
		return false
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums) && j <= i+k; j++ {
			if abs(nums[i], nums[j]) <= t {
				return true
			}
		}
	}
	return false
}

# 2
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if len(nums) <= 1 || t < 0 {
		return false
	}
	m := make(map[int]int)
	width := t + 1
	for i := 0; i < len(nums); i++ {
		key := getKey(nums[i], width)
		if _, ok := m[key]; ok {
			return true
		}
		if value, ok := m[key-1]; ok && abs(nums[i], value) < width {
			return true
		}
		if value, ok := m[key+1]; ok && abs(nums[i], value) < width {
			return true
		}
		m[key] = nums[i]
		if i >= k {
			// 满足i和j的差的绝对值也小于等于ķ
			delete(m, getKey(nums[i-k], width))
		}
	}
	return false
}

func getKey(value, width int) int {
	if value < 0 {
		return (value+1)/width - 1
	}
	return value / width
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
```

## 剑指OfferII058.日程表(2)

- 题目

```
请实现一个 MyCalendar 类来存放你的日程安排。如果要添加的时间内没有其他安排，则可以存储这个新的日程安排。
MyCalendar 有一个 book(int start, int end)方法。它意味着在 start 到 end 时间内增加一个日程安排，
注意，这里的时间是半开区间，即 [start, end), 实数 x 的范围为，  start <= x < end。
当两个日程安排有一些时间上的交叉时（例如两个日程安排都在同一时间内），就会产生重复预订。
每次调用 MyCalendar.book方法时，如果可以将日程安排成功添加到日历中而不会导致重复预订，返回 true。
否则，返回 false 并且不要将该日程安排添加到日历中。
请按照以下步骤调用 MyCalendar 类: MyCalendar cal = new MyCalendar(); MyCalendar.book(start, end)
示例:输入: ["MyCalendar","book","book","book"] [[],[10,20],[15,25],[20,30]]
输出: [null,true,false,true]
解释: MyCalendar myCalendar = new MyCalendar();
MyCalendar.book(10, 20); // returns true 
MyCalendar.book(15, 25); // returns false ，第二个日程安排不能添加到日历中，因为时间 15 已经被第一个日程安排预定了
MyCalendar.book(20, 30); // returns true ，第三个日程安排可以添加到日历中，因为第一个日程安排并不包含时间 20 
提示：每个测试用例，调用 MyCalendar.book 函数最多不超过 1000次。
0 <= start < end <= 109
注意：本题与主站 729 题相同： 
```

- 解题思路

| No.  | 思路   | 时间复杂度 | 空间复杂度 |
| ---- | ------ | ---------- | ---------- |
| 01   | 暴力法 | O(n^2)     | O(n)       |
| 02   | 平衡树 | O(nlog(n)) | O(n)       |

```go
type MyCalendar struct {
	arr [][2]int
}

func Constructor() MyCalendar {
	return MyCalendar{arr: make([][2]int, 0)}
}

func (this *MyCalendar) Book(start int, end int) bool {
	for i := 0; i < len(this.arr); i++ {
		if this.arr[i][0] < end && start < this.arr[i][1] {
			return false
		}
	}
	this.arr = append(this.arr, [2]int{start, end})
	return true
}

# 2
type MyCalendar struct {
	root *Node
}

func Constructor() MyCalendar {
	return MyCalendar{root: nil}
}

func (this *MyCalendar) Book(start int, end int) bool {
	node := &Node{
		start: start,
		end:   end,
		left:  nil,
		right: nil,
	}
	if this.root == nil {
		this.root = node
		return true
	}
	return this.root.Insert(node)
}

type Node struct {
	start int
	end   int
	left  *Node
	right *Node
}

func (this *Node) Insert(node *Node) bool {
	if node.start >= this.end {
		if this.right == nil {
			this.right = node
			return true
		}
		return this.right.Insert(node)
	} else if node.end <= this.start {
		if this.left == nil {
			this.left = node
			return true
		}
		return this.left.Insert(node)
	}
	return false
}
```

## 剑指OfferII060.出现频率最高的k个数字(3)

- 题目

```
给定一个整数数组 nums 和一个整数 k ，请返回其中出现频率前 k 高的元素。可以按 任意顺序 返回答案。
示例 1:输入: nums = [1,1,1,2,2,3], k = 2 输出: [1,2]
示例 2:输入: nums = [1], k = 1 输出: [1]
提示：1 <= nums.length <= 105
k 的取值范围是 [1, 数组中不相同的元素的个数]
题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的
进阶：所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。
注意：本题与主站 347 题相同
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 哈希辅助 | O(nlog(n)) | O(n)       |
| 02   | 堆       | O(nlog(n)) | O(n)       |
| 03   | 桶排序   | O(n)       | O(n)       |

```go
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	arr := make([][2]int, 0)
	for k, v := range m {
		arr = append(arr, [2]int{k, v})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] > arr[j][1]
	})
	res := make([]int, 0)
	for i := 0; i < k; i++ {
		res = append(res, arr[i][0])
	}
	return res
}

# 2
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	var h IntHeap
	heap.Init(&h)
	for k, v := range m {
		heap.Push(&h, [2]int{k, v})
	}
	res := make([]int, 0)
	for h.Len() > 0 && k > 0 {
		k--
		node := heap.Pop(&h).([2]int)
		res = append(res, node[0])
	}
	return res
}

type IntHeap [][2]int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i][1] > h[j][1] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.([2]int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

# 3
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	arr := make([][]int, len(nums)+1)
	temp := make(map[int][]int)
	for key, value := range m {
		temp[value] = append(temp[value], key)
		arr[value] = append(arr[value], key)
	}
	res := make([]int, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		// 避免出现0=>x次的情况
		if _, ok := temp[i]; ok {
			for j := 0; j < len(arr[i]); j++ {
				k--
				if k < 0 {
					break
				}
				res = append(res, arr[i][j])
			}
		}
	}
	return res
}
```

## 剑指OfferII061.和最小的k个数对(2)

- 题目

```
给定两个以升序排列的整数数组 nums1 和 nums2 , 以及一个整数 k 。
定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。
请找到和最小的 k 个数对 (u1,v1),  (u2,v2)  ...  (uk,vk) 。
示例 1:输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3 输出: [1,2],[1,4],[1,6]
解释: 返回序列中的前 3 对数：
    [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
示例 2:输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2 输出: [1,1],[1,1]
解释: 返回序列中的前 2 对数：
     [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
示例 3:输入: nums1 = [1,2], nums2 = [3], k = 3  输出: [1,3],[2,3]
解释: 也可能序列中所有的数对都被返回:[1,3],[2,3]
提示:1 <= nums1.length, nums2.length <= 104
-109 <= nums1[i], nums2[i] <= 109
nums1, nums2 均为升序排列
1 <= k <= 1000
注意：本题与主站 373 题相同：
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 堆   | O(nlog(n)) | O(n)       |
| 02   | 排序 | O(nlog(n)) | O(n^2)     |

```go
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	Heap := &NodeHeap{}
	heap.Init(Heap)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			heap.Push(Heap, Node{
				i: nums1[i],
				j: nums2[j],
			})
			if Heap.Len() > k {
				heap.Pop(Heap)
			}
		}
	}
	res := make([][]int, 0)
	for Heap.Len() > 0 {
		node := heap.Pop(Heap).(Node)
		res = append(res, []int{node.i, node.j})
	}
	return res
}

type Node struct {
	i int
	j int
}

type NodeHeap []Node

func (h NodeHeap) Len() int {
	return len(h)
}

func (h NodeHeap) Less(i, j int) bool {
	return h[i].i+h[i].j > h[j].i+h[j].j
}

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *NodeHeap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}

# 2
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	arr := make([][]int, 0)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			arr = append(arr, []int{nums1[i], nums2[j]})
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0]+arr[i][1] < arr[j][0]+arr[j][1]
	})
	if len(arr) < k {
		return arr
	}
	return arr[:k]
}
```

## 剑指OfferII062.实现前缀树(2)

- 题目

```
Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。
这一数据结构有相当多的应用情景，例如自动补完和拼写检查。
请你实现 Trie 类：
Trie() 初始化前缀树对象。
void insert(String word) 向前缀树中插入字符串 word 。
boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。
示例：输入inputs = ["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
inputs = [[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
输出[null, null, true, false, true, null, true]
解释 Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // 返回 True
trie.search("app");     // 返回 False
trie.startsWith("app"); // 返回 True
trie.insert("app");
trie.search("app");     // 返回 True
提示：1 <= word.length, prefix.length <= 2000
word 和 prefix 仅由小写英文字母组成
insert、search 和 startsWith 调用次数 总计 不超过 3 * 104 次
注意：本题与主站 208 题相同：
```

- 解题思路

| No.  | 思路   | 时间复杂度 | 空间复杂度 |
| ---- | ------ | ---------- | ---------- |
| 01   | trie树 | O(n)       | O(n)       |
| 02   | trie树 | O(n)       | O(n)       |

```go
type Trie struct {
	next   [26]*Trie
	ending int
}

func Constructor() Trie {
	return Trie{
		next:   [26]*Trie{},
		ending: 0,
	}
}

func (this *Trie) Insert(word string) {
	temp := this
	for _, v := range word {
		value := v - 'a'
		if temp.next[value] == nil {
			temp.next[value] = &Trie{
				next:   [26]*Trie{},
				ending: 0,
			}
		}
		temp = temp.next[value]
	}
	temp.ending++
}

func (this *Trie) Search(word string) bool {
	temp := this
	for _, v := range word {
		value := v - 'a'
		if temp = temp.next[value]; temp == nil {
			return false
		}
	}
	if temp.ending > 0 {
		return true
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	temp := this
	for _, v := range prefix {
		value := v - 'a'
		if temp = temp.next[value]; temp == nil {
			return false
		}
	}
	return true
}

# 2
type Trie struct {
	next   map[byte]*Trie
	ending int
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		next:   make(map[byte]*Trie),
		ending: 0,
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	temp := this
	for _, v := range word {
		value := byte(v - 'a')
		if temp.next[value] == nil {
			temp.next[value] = &Trie{
				next:   make(map[byte]*Trie),
				ending: 0,
			}
		}
		temp = temp.next[value]
	}
	temp.ending++
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	temp := this
	for _, v := range word {
		value := byte(v - 'a')
		if temp = temp.next[value]; temp == nil {
			return false
		}
	}
	if temp.ending > 0 {
		return true
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	temp := this
	for _, v := range prefix {
		value := byte(v - 'a')
		if temp = temp.next[value]; temp == nil {
			return false
		}
	}
	return true
}
```

## 剑指OfferII063.替换单词(2)

- 题目

```
在英语中，有一个叫做 词根(root) 的概念，它可以跟着其他一些词组成另一个较长的单词——我们称这个词为 继承词(successor)。
例如，词根an，跟随着单词 other(其他)，可以形成新的单词 another(另一个)。
现在，给定一个由许多词根组成的词典和一个句子，需要将句子中的所有继承词用词根替换掉。
如果继承词有许多可以形成它的词根，则用最短的词根替换它。
需要输出替换之后的句子。
示例 1：输入：dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
输出："the cat was rat by the bat"
示例 2：输入：dictionary = ["a","b","c"], sentence = "aadsfasf absbs bbab cadsfafs" 输出："a a b c"
示例 3：
输入：dictionary = ["a", "aa", "aaa", "aaaa"], sentence = "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa"
输出："a a a a a a a a bbb baba a"
示例 4：输入：dictionary = ["catt","cat","bat","rat"], sentence = "the cattle was rattled by the battery"
输出："the cat was rat by the bat"
示例 5：输入：dictionary = ["ac","ab"], sentence = "it is abnormal that this solution is accepted"
输出："it is ab that this solution is ac"
提示：1 <= dictionary.length <= 1000
1 <= dictionary[i].length <= 100
dictionary[i] 仅由小写字母组成。
1 <= sentence.length <= 10^6
sentence 仅由小写字母和空格组成。
sentence 中单词的总量在范围 [1, 1000] 内。
sentence 中每个单词的长度在范围 [1, 1000] 内。
sentence 中单词之间由一个空格隔开。
sentence 没有前导或尾随空格。
注意：本题与主站 648 题相同：
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 内置函数 | O(n^2)     | O(n)       |
| 02   | 字典树   | O(n)       | O(n)       |

```go
func replaceWords(dictionary []string, sentence string) string {
	sort.Strings(dictionary)
	arr := strings.Split(sentence, " ")
	for i := 0; i < len(arr); i++ {
		for _, v := range dictionary {
			if strings.HasPrefix(arr[i], v) {
				arr[i] = v
				break
			}
		}
	}
	return strings.Join(arr, " ")
}

# 2
func replaceWords(dictionary []string, sentence string) string {
	trie := Constructor()
	for i := 0; i < len(dictionary); i++ {
		trie.Insert(dictionary[i])
	}
	arr := strings.Split(sentence, " ")
	for i := 0; i < len(arr); i++ {
		result := trie.Search(arr[i])
		if result != "" {
			arr[i] = result
		}
	}
	return strings.Join(arr, " ")
}

type Trie struct {
	next   [26]*Trie // 下一级指针，如不限于小写字母，[26]=>[256]
	ending int       // 次数（可以改为bool）
}

func Constructor() Trie {
	return Trie{
		next:   [26]*Trie{},
		ending: 0,
	}
}

// 插入word
func (this *Trie) Insert(word string) {
	temp := this
	for _, v := range word {
		value := v - 'a'
		if temp.next[value] == nil {
			temp.next[value] = &Trie{
				next:   [26]*Trie{},
				ending: 0,
			}
		}
		temp = temp.next[value]
	}
	temp.ending++
}

// 查找
func (this *Trie) Search(word string) string {
	temp := this
	res := ""
	for _, v := range word {
		res = res + string(v)
		value := v - 'a'
		if temp = temp.next[value]; temp == nil {
			return ""
		}
		if temp.ending > 0 {
			return res
		}
	}
	return ""
}
```

## 剑指OfferII064.神奇的字典

### 题目

```

```

### 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 滑动窗口 | O(n^2)     | O(n)       |

```go

```

## 剑指OfferII065.最短的单词编码(3)

- 题目

```
单词数组 words 的 有效编码 由任意助记字符串 s 和下标数组 indices 组成，且满足：
words.length == indices.length
助记字符串 s 以 '#' 字符结尾
对于每个下标 indices[i] ，s 的一个从 indices[i] 开始、
到下一个 '#' 字符结束（但不包括 '#'）的 子字符串 恰好与 words[i] 相等
给定一个单词数组 words ，返回成功对 words 进行编码的最小助记字符串 s 的长度 。
示例 1：输入：words = ["time", "me", "bell"] 输出：10
解释：一组有效编码为 s = "time#bell#" 和 indices = [0, 2, 5] 。
words[0] = "time" ，s 开始于 indices[0] = 0 到下一个 '#' 结束的子字符串，如加粗部分所示 "time#bell#"
words[1] = "me" ，s 开始于 indices[1] = 2 到下一个 '#' 结束的子字符串，如加粗部分所示 "time#bell#"
words[2] = "bell" ，s 开始于 indices[2] = 5 到下一个 '#' 结束的子字符串，如加粗部分所示 "time#bell#"
示例 2：输入：words = ["t"] 输出：2
解释：一组有效编码为 s = "t#" 和 indices = [0] 。
提示：1 <= words.length <= 2000
1 <= words[i].length <= 7
words[i] 仅由小写字母组成
注意：本题与主站 820 题相同： 
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 哈希辅助 | O(n^2)     | O(n)       |
| 02   | 排序遍历 | O(n^2)     | O(n)       |
| 03   | 逆置排序 | O(nlog(n)) | O(n)       |

```go
func minimumLengthEncoding(words []string) int {
	res := 0
	m := make(map[string]bool)
	for i := 0; i < len(words); i++ {
		m[words[i]] = true
	}
	for k := range m {
		for i := 1; i < len(k); i++ {
			delete(m, k[i:])
		}
	}
	for k := range m {
		res = res + len(k) + 1
	}
	return res
}

# 2
func minimumLengthEncoding(words []string) int {
	res := 0
	m := make(map[string]bool)
	for i := 0; i < len(words); i++ {
		m[words[i]] = true
	}
	words = make([]string, 0)
	for k := range m {
		words = append(words, k)
	}
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	for i := len(words) - 1; i >= 0; i-- {
		if m[words[i]] == false {
			continue
		}
		for j := i - 1; j >= 0; j-- {
			if strings.HasSuffix(words[i], words[j]) == true {
				m[words[j]] = false
			}
		}
	}
	for k := range m {
		if m[k] == true {
			res = res + len(k) + 1
		}
	}
	return res
}

# 3
func minimumLengthEncoding(words []string) int {
	res := 0
	arr := make([]string, 0)
	for k := range words {
		arr = append(arr, reverse(words[k]))
	}
	sort.Strings(arr)
	for i := 0; i < len(arr)-1; i++ {
		length := len(arr[i])
		if length <= len(arr[i+1]) && arr[i] == arr[i+1][:length] {
			continue
		}
		res = res + length + 1
	}
	return res + len(arr[len(arr)-1]) + 1
}

func reverse(str string) string {
	res := make([]byte, 0)
	for i := len(str) - 1; i >= 0; i-- {
		res = append(res, str[i])
	}
	return string(res)
}
```

## 剑指OfferII066.单词之和(3)

- 题目

```
实现一个 MapSum 类，支持两个方法，insert 和 sum：
MapSum() 初始化 MapSum 对象
void insert(String key, int val) 插入 key-val 键值对，字符串表示键 key ，整数表示值 val 。
如果键 key 已经存在，那么原来的键值对将被替代成新的键值对。
int sum(string prefix) 返回所有以该前缀 prefix 开头的键 key 的值的总和。
示例：输入：inputs = ["MapSum", "insert", "sum", "insert", "sum"]
inputs = [[], ["apple", 3], ["ap"], ["app", 2], ["ap"]]
输出：[null, null, 3, null, 5]
解释：MapSum mapSum = new MapSum();
mapSum.insert("apple", 3);  
mapSum.sum("ap");           // return 3 (apple = 3)
mapSum.insert("app", 2);    
mapSum.sum("ap");           // return 5 (apple + app = 3 + 2 = 5)
提示：1 <= key.length, prefix.length <= 50
key 和 prefix 仅由小写英文字母组成
1 <= val <= 1000
最多调用 50 次 insert 和 sum
注意：本题与主站 677 题相同
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | trie树   | O(n)       | O(n)       |
| 02   | 哈希辅助 | O(n)       | O(n)       |
| 03   | 哈希辅助 | O(n)       | O(n)       |

```go
type MapSum struct {
	val  int
	next map[int32]*MapSum
}

func Constructor() MapSum {
	return MapSum{
		val:  0,
		next: make(map[int32]*MapSum),
	}
}

func (this *MapSum) Insert(key string, val int) {
	node := this
	for _, v := range key {
		if _, ok := node.next[v]; ok == false {
			temp := Constructor()
			node.next[v] = &temp
		}
		node = node.next[v]
	}
	node.val = val
}

func (this *MapSum) Sum(prefix string) int {
	node := this
	for _, v := range prefix {
		if _, ok := node.next[v]; ok == false {
			return 0
		}
		node = node.next[v]
	}
	res := 0
	queue := make([]*MapSum, 0)
	queue = append(queue, node)
	for len(queue) > 0 {
		temp := queue[0]
		queue = queue[1:]
		res = res + temp.val
		for _, v := range temp.next {
			queue = append(queue, v)
		}
	}
	return res
}

# 2
type MapSum struct {
	m    map[string]int
	data map[string]map[string]bool
}

func Constructor() MapSum {
	return MapSum{
		m:    make(map[string]int),
		data: make(map[string]map[string]bool),
	}
}

func (this *MapSum) Insert(key string, val int) {
	this.m[key] = val
	for i := 1; i <= len(key); i++ {
		str := key[:i]
		if _, ok := this.data[str]; ok == false {
			this.data[str] = make(map[string]bool)
		}
		this.data[str][key] = true
	}
}

func (this *MapSum) Sum(prefix string) int {
	res := 0
	for key := range this.data[prefix] {
		res = res + this.m[key]
	}
	return res
}

# 3
type MapSum struct {
	m map[string]int
}

func Constructor() MapSum {
	return MapSum{
		m: make(map[string]int),
	}
}

func (this *MapSum) Insert(key string, val int) {
	this.m[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	res := 0
	for key, value := range this.m {
		if strings.HasPrefix(key, prefix) {
			res = res + value
		}
	}
	return res
}
```

## 剑指OfferII067.最大的异或(2)

- 题目

```
给定一个整数数组 nums ，返回 nums[i] XOR nums[j] 的最大运算结果，其中 0 ≤ i ≤ j < n 。
示例 1：输入：nums = [3,10,5,25,2,8] 输出：28
解释：最大运算结果是 5 XOR 25 = 28.
示例 2：输入：nums = [0] 输出：0
示例 3：输入：nums = [2,4] 输出：6
示例 4：输入：nums = [8,10,2] 输出：10
示例 5：输入：nums = [14,70,53,83,49,91,36,80,92,51,66,70] 输出：127
提示：1 <= nums.length <= 2 * 104
0 <= nums[i] <= 231 - 1
进阶：你可以在 O(n) 的时间解决这个问题吗？
注意：本题与主站 421 题相同： 
```

- 解题思路

| No.  | 思路        | 时间复杂度 | 空间复杂度 |
| ---- | ----------- | ---------- | ---------- |
| 01   | 哈希+位运算 | O(n)       | O(n)       |
| 02   | trie树      | O(n)       | O(n)       |

```go
func findMaximumXOR(nums []int) int {
	res := 0
	target := 0
	for i := 31; i >= 0; i-- { // 枚举每一位（第i位，从右到左），判断该为能否为1
		m := make(map[int]bool)
		target = target | (1 << i) // target第i位置1
		for j := 0; j < len(nums); j++ {
			m[nums[j]&target] = true // 高位&：取前缀
		}
		temp := res | (1 << i) // 假设结果第i位为1
		// a ^ b = temp
		// temp ^ a = b
		for k := range m {
			if _, ok := m[temp^k]; ok {
				res = temp // 能取到1
				break
			}
		}
	}
	return res
}

# 2
func findMaximumXOR(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	res := 0
	root := Trie{
		next: make([]*Trie, 2), // 0和1
	}
	for i := 0; i < n; i++ {
		temp := &root
		for j := 31; j >= 0; j-- {
			value := (nums[i] >> j) & 1
			if temp.next[value] == nil {
				temp.next[value] = &Trie{
					next: make([]*Trie, 2),
				}
			}
			temp = temp.next[value]
		}
	}
	for i := 0; i < n; i++ {
		temp := &root
		cur := 0
		for j := 31; j >= 0; j-- {
			value := (nums[i] >> j) & 1
			if temp.next[value^1] != nil { // 能取到1
				cur = cur | (1 << j) // 结果在该位可以为1
				temp = temp.next[value^1]
			} else {
				temp = temp.next[value]
			}
		}
		res = max(res, cur)
	}
	return res
}

type Trie struct {
	next []*Trie // 下一级指针，如不限于小写字母，[26]=>[256]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```



# Hard

## 剑指OfferII017.含有所有字符的最短字符串(2)

- 题目

```
给定两个字符串 s 和 t 。返回 s 中包含 t 的所有字符的最短子字符串。如果 s 中不存在符合条件的子字符串，则返回空字符串 "" 。
如果 s 中存在多个符合条件的子字符串，返回任意一个。
注意： 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
示例 1：输入：s = "ADOBECODEBANC", t = "ABC" 输出："BANC" 
解释：最短子字符串 "BANC" 包含了字符串 t 的所有字符 'A'、'B'、'C'
示例 2：输入：s = "a", t = "a" 输出："a"
示例 3：输入：s = "a", t = "aa" 输出：""
解释：t 中两个字符 'a' 均应包含在 s 的子串中，因此没有符合条件的子字符串，返回空字符串。
提示：1 <= s.length, t.length <= 105
s 和 t 由英文字母组成
进阶：你能设计一个在 o(n) 时间内解决此问题的算法吗？
注意：本题与主站 76 题相似（本题答案不唯一）：
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 滑动窗口 | O(n^2)     | O(n)       |
| 02   | 滑动窗口 | O(n)       | O(n)       |

```go
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	window := make(map[byte]int)
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	left, right := -1, -1
	minLength := math.MaxInt32
	for l, r := 0, 0; r < len(s); r++ {
		if r < len(s) && need[s[r]] > 0 {
			window[s[r]]++
		}
		// 找到，然后left往右移
		for check(need, window) == true && l <= r {
			if r-l+1 < minLength {
				minLength = r - l + 1
				left, right = l, r+1
			}
			if _, ok := need[s[l]]; ok {
				window[s[l]]--
			}
			l++
		}
	}
	if left == -1 {
		return ""
	}
	return s[left:right]
}

func check(need, window map[byte]int) bool {
	for k, v := range need {
		if window[k] < v {
			return false
		}
	}
	return true
}

# 2
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	arr := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		arr[t[i]]++
	}
	l, count := 0, 0
	res := ""
	minLength := math.MaxInt32
	for r := 0; r < len(s); r++ {
		arr[s[r]]--
		if arr[s[r]] >= 0 {
			count++
		}
		// left往右边移动
		for count == len(t) {
			if minLength > r-l+1 {
				minLength = r - l + 1
				res = s[l : r+1]
			}
			arr[s[l]]++
			if arr[s[l]] > 0 {
				count--
			}
			l++
		}
	}
	return res
}
```

## 剑指OfferII039.直方图最大矩形面积(3)

- 题目

```
给定非负整数数组 heights ，数组中的数字用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
求在该柱状图中，能够勾勒出来的矩形的最大面积。
示例 1:输入：heights = [2,1,5,6,2,3] 输出：10
解释：最大的矩形为图中红色区域，面积为 10
示例 2：输入： heights = [2,4] 输出： 4
提示：1 <= heights.length <=105
0 <= heights[i] <= 104
注意：本题与主站 84 题相同： 
```

- 解题思路

| No.  | 思路   | 时间复杂度 | 空间复杂度 |
| ---- | ------ | ---------- | ---------- |
| 01   | 单调栈 | O(n)       | O(n)       |
| 02   | 单调栈 | O(n)       | O(n)       |
| 03   | 单调栈 | O(n)       | O(n)       |

```go
func largestRectangleArea(heights []int) int {
	n := len(heights)
	res := 0
	left := make([]int, n)
	right := make([]int, n)
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = n
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	for i := 0; i < n; i++ {
		res = max(res, heights[i]*(right[i]-left[i]-1))
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

# 2
func largestRectangleArea(heights []int) int {
	n := len(heights)
	res := 0
	left := make([]int, n)
	right := make([]int, n)
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		right[i] = n
	}
	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			right[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	for i := 0; i < n; i++ {
		res = max(res, heights[i]*(right[i]-left[i]-1))
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

# 3
func largestRectangleArea(heights []int) int {
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	n := len(heights)
	res := 0
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		// 递增栈
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := i - stack[len(stack)-1] - 1
			res = max(res, height*width)
		}
		stack = append(stack, i)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

## 剑指OfferII040.矩阵中最大的矩形(2)

- 题目

```
给定一个由 0 和 1 组成的矩阵 matrix ，找出只包含 1 的最大矩形，并返回其面积。
注意：此题 matrix 输入格式为一维 01 字符串数组。
示例 1：输入：matrix = ["10100","10111","11111","10010"] 输出：6
解释：最大矩形如上图所示。
示例 2：输入：matrix = [] 输出：0
示例 3：输入：matrix = ["0"] 输出：0
示例 4：输入：matrix = ["1"] 输出：1
示例 5：输入：matrix = ["00"] 输出：0
提示：rows == matrix.length
cols == matrix[0].length
0 <= row, cols <= 200
matrix[i][j] 为 '0' 或 '1'
注意：本题与主站 85 题相同（输入参数格式不同）：
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 单调栈   | O(n^2)     | O(n)       |
| 02   | 动态规划 | O(n^2)     | O(n)       |

```go
func maximalRectangle(matrix []string) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	res := 0
	n, m := len(matrix), len(matrix[0])
	height := make([]int, m) // 高度
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == '0' {
				height[j] = 0
			} else {
				height[j] = height[j] + 1
			}
		}
		res = max(res, getMaxArea(height))
	}
	return res
}

func getMaxArea(heights []int) int {
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	n := len(heights)
	res := 0
	stack := make([]int, 0) // 递增栈
	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := i - stack[len(stack)-1] - 1
			res = max(res, height*width)
		}
		stack = append(stack, i)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

# 2
func maximalRectangle(matrix []string) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	res := 0
	n, m := len(matrix), len(matrix[0])
	left, right, height := make([]int, m), make([]int, m), make([]int, m)
	for i := 0; i < m; i++ {
		right[i] = m
	}
	for i := 0; i < n; i++ {
		curLeft, curRight := 0, m
		// 高度
		for j := 0; j < m; j++ {
			if matrix[i][j] == '1' {
				height[j]++
			} else {
				height[j] = 0
			}
		}
		// 左边
		for j := 0; j < m; j++ {
			if matrix[i][j] == '1' {
				left[j] = max(left[j], curLeft)
			} else {
				left[j] = 0
				curLeft = j + 1
			}
		}
		// 右边
		for j := m - 1; j >= 0; j-- {
			if matrix[i][j] == '1' {
				right[j] = min(right[j], curRight)
			} else {
				right[j] = m
				curRight = j
			}
		}
		for j := 0; j < m; j++ {
			res = max(res, height[j]*(right[j]-left[j]))
		}
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
```

## 剑指OfferII048.序列化与反序列化二叉树(2)

- 题目

```
序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，
同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。
请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，
只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。
示例 1：输入：root = [1,2,3,null,null,4,5] 输出：[1,2,3,null,null,4,5]
示例 2：输入：root = [] 输出：[]
示例 3：输入：root = [1] 输出：[1]
示例 4：输入：root = [1,2] 输出：[1,2]
提示：输入输出格式与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。
你并非必须采取这种方式，也可以采用其他的方法解决这个问题。
树中结点数在范围 [0, 104] 内
-1000 <= Node.val <= 1000
注意：本题与主站 297 题相同：
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 递归 | O(n)       | O(n)       |
| 02   | 迭代 | O(n)       | O(n)       |

```go
type Codec struct {
	res []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	this.res = strings.Split(data, ",")
	return this.dfsDeserialize()
}

func (this *Codec) dfsDeserialize() *TreeNode {
	node := this.res[0]
	this.res = this.res[1:]
	if node == "#" {
		return nil
	}
	value, _ := strconv.Atoi(node)
	return &TreeNode{
		Val:   value,
		Left:  this.dfsDeserialize(),
		Right: this.dfsDeserialize(),
	}
}

# 2
type Codec struct {
	res []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	res := make([]string, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			res = append(res, strconv.Itoa(node.Val))
			queue = append(queue, node.Left, node.Right)
		} else {
			res = append(res, "#")
		}
	}
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 || data == "" {
		return nil
	}
	res := strings.Split(data, ",")
	root := &TreeNode{}
	root.Val, _ = strconv.Atoi(res[0])
	res = res[1:]
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		if res[0] != "#" {
			left, _ := strconv.Atoi(res[0])
			queue[0].Left = &TreeNode{Val: left}
			queue = append(queue, queue[0].Left)
		}
		if res[1] != "#" {
			right, _ := strconv.Atoi(res[1])
			queue[0].Right = &TreeNode{Val: right}
			queue = append(queue, queue[0].Right)
		}
		queue = queue[1:]
		res = res[2:]
	}
	return root
}
```

## 剑指OfferII051.节点之和最大的路径(2)

- 题目

```
路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。
该路径 至少包含一个 节点，且不一定经过根节点。
路径和 是路径中各节点值的总和。
给定一个二叉树的根节点 root ，返回其 最大路径和，即所有路径上节点值之和的最大值。
示例 1：输入：root = [1,2,3] 输出：6
解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6
示例 2：输入：root = [-10,9,20,null,null,15,7] 输出：42
解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42
提示：树中节点数目范围是 [1, 3 * 104]
-1000 <= Node.val <= 1000
注意：本题与主站 124 题相同：
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 递归 | O(n)       | O(log(n))  |
| 02   | 迭代 | O(n)       | O(n)       |

```go
var res int

func maxPathSum(root *TreeNode) int {
	res = math.MinInt32
	dfs(root)
	return res
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := max(dfs(root.Left), 0)
	right := max(dfs(root.Right), 0)
	// 该顶点路径和=root.Val+2边和
	value := left + right + root.Val
	res = max(res, value)
	// 单分支
	return root.Val + max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

# 2
var res int

func maxPathSum(root *TreeNode) int {
	res = math.MinInt32
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	stack := make([]*TreeNode, 0)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		stack = append(stack, node)
	}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = max(res, node.Val)
		var left, right int
		if node.Left == nil {
			left = 0
		} else {
			left = max(node.Left.Val, 0)
		}
		if node.Right == nil {
			right = 0
		} else {
			right = max(node.Right.Val, 0)
		}
		sum := node.Val + left + right
		res = max(res, sum)
		node.Val = node.Val + max(left, right)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

