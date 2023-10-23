package main

func main() {

}

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
