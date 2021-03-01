package main

import (
	"crypto/x509"
	"strconv"
)

type NestedInteger struct{}

func (n NestedInteger) IsInteger() bool {}

func (n NestedInteger) GetInteger() int {}

func (n *NestedInteger) SetInteger(value int) {}

func (n *NestedInteger) Add(elem NestedInteger) {}

func (n NestedInteger) GetList() []*NestedInteger {}

func main() {

}

// leetcode385_迷你语法分析器
func deserialize(s string) *NestedInteger {
	if s[0] != '[' {
		res := &NestedInteger{}
		num, _ := strconv.Atoi(s)
		res.SetInteger(num)
		return res
	}
	stack := make([]NestedInteger, 0)
	str := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			stack = append(stack, NestedInteger{})
		} else if s[i] == ']' {
			if len(str) > 0 {
				node := NestedInteger{}
				num, _ := strconv.Atoi(str)
				node.SetInteger(num)
				stack[len(stack)-1].Add(node)
			}
			str = ""
			if len(stack) > 1 { // 出栈
				last := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack[len(stack)-1].Add(last)
			}
		} else if s[i] == ',' {
			if len(str) > 0 {
				node := NestedInteger{}
				num, _ := strconv.Atoi(str)
				node.SetInteger(num)
				stack[len(stack)-1].Add(node)
			}
			str = ""
		} else {
			str = str + string(s[i])
		}
	}
	return &stack[len(stack)-1]
}
