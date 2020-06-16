package main

func main() {

}

// 剑指offer_面试题09_用两个栈实现队列
type stack []int

func (s *stack) Push(value int) {
	*s = append(*s, value)
}
func (s *stack) Pop() int {
	value := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return value
}

type CQueue struct {
	tail stack
	head stack
}

func Constructor() CQueue {
	return CQueue{}
}

// 1.入队，tail栈保存
// 2.出队, head不为空，出head；head为空，tail出到head里，最后出head
func (this *CQueue) AppendTail(value int) {
	this.tail.Push(value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.head) != 0 {
		return this.head.Pop()
	} else if len(this.tail) != 0 {
		for len(this.tail) > 0 {
			this.head.Push(this.tail.Pop())
		}
		return this.head.Pop()
	}
	return -1
}
