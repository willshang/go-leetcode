package main

func main() {

}

// leetcode1286_字母组合迭代器
type CombinationIterator struct {
	arr   []string
	index int
}

func dfs(str string, k int, index int, cur string, res *[]string) {
	if len(cur) == k {
		*res = append(*res, cur)
		return
	}
	for i := index; i < len(str); i++ {
		dfs(str, k, i+1, cur+string(str[i]), res)
	}
}
func Constructor(characters string, combinationLength int) CombinationIterator {
	res := make([]string, 0)
	dfs(characters, combinationLength, 0, "", &res)
	return CombinationIterator{
		arr:   res,
		index: 0,
	}
}

func (this *CombinationIterator) Next() string {
	if this.index < len(this.arr) {
		this.index++
		return this.arr[this.index-1]
	}
	return ""
}

func (this *CombinationIterator) HasNext() bool {
	return this.index < len(this.arr)
}
