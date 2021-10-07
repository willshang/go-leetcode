# Trie树(前缀树/字典树)

- 定义：**Trie树**，又经常叫**前缀树**，**字典树**等等
- 百度百科：又称**单词查找树**，Trie树，是一种树形结构，是一种哈希树的变种。
- 典型应用是用于统计，排序和保存大量的字符串（但不仅限于字符串），所以经常被搜索引擎系统用于文本词频统计。
- 它的优点是：利用字符串的公共前缀来减少查询时间，最大限度地减少无谓的字符串比较，查询效率比哈希树高。

## 0、定义

### 1、操作

## 2、Go实现

- 实现参考：leetcode208.实现 Trie (前缀树)

- Insert操作
- Search操作

```go
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

// 查找前缀
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
```

## 3、Leetcode

| Title                                                        | Tag                                                          | 难度   | 完成情况 |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------ | -------- |
| [14.最长公共前缀](https://leetcode-cn.com/problems/longest-common-prefix) | 字符串                                                       | Easy   | 完成     |
| [208.实现 Trie (前缀树)](https://leetcode-cn.com/problems/implement-trie-prefix-tree/) | 设计、字典树                                                 | Medium | 完成     |
| [211.添加与搜索单词-数据结构设计](https://leetcode-cn.com/problems/add-and-search-word-data-structure-design/) | 设计、字典树、回溯算法                                       | Medium | 完成     |
| [648.单词替换](https://leetcode-cn.com/problems/replace-words/) | 字典树、哈希表                                               | Medium | 完成     |
| [676.实现一个魔法字典](https://leetcode-cn.com/problems/implement-magic-dictionary/) | 设计、字典树、哈希表、字符串                                 | Medium | 完成     |
| [720.词典中最长的单词](https://leetcode-cn.com/problems/longest-word-in-dictionary/) | 字典树、哈希表                                               | Easy   | 完成     |
| [1268.搜索推荐系统](https://leetcode-cn.com/problems/search-suggestions-system/) | 字典树、数组、字符串                                         | Medium | 完成     |
| [面试题17.13.恢复空格](https://leetcode-cn.com/problems/re-space-lcci/) | 字典树、数组、哈希表、字符串、<br />动态规划、哈希函数、滚动哈希 | Medium | 完成     |
| [面试题17.17.多次搜索](https://leetcode-cn.com/problems/multi-search-lcci/) | 字典树、数组、哈希表、字符串、<br />字符串匹配、滑动窗口     | Medium | 完成     |
| [剑指OfferII064.神奇的字典](https://leetcode-cn.com/problems/US1pGT/) | 设计、字典树、哈希表、字符串                                 | Medium | 完成     |

