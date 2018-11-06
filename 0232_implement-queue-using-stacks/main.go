package main

func main() {
	
}

type MyQueue struct {
	a []int
}

func Constructor() MyQueue  {
	return MyQueue{}
}

func (this *MyQueue) Push(x int)  {
	this.a = append(this.a,x)
}

func (this *MyQueue) Pop() int  {
	if len(this.a) == 0{
		return 0
	}
	first := this.a[0]
	this.a = this.a[1:]
	return first
}

func (this *MyQueue) Peek() int  {
	if len(this.a) == 0{
		return 0
	}
	return this.a[0]
}
func (this *MyQueue) Empty() bool  {
	if len(this.a) == 0{
		return true
	}
	return false
}

/*type MyQueue struct {
	a, b *Stack
}


func Constructor() MyQueue {
	return MyQueue{
		a:NewStack(),
		b:NewStack(),
	}
}


func (this *MyQueue) Push(x int)  {
	this.a.Push(x)
}


func (this *MyQueue) Pop() int {
	if this.b.Len() == 0{
		for this.a.Len() > 0{
			this.b.Push(this.a.Pop())
		}
	}
	return this.b.Pop()
}


func (this *MyQueue) Peek() int {
	res := this.Pop()
	this.b.Push(res)
	return res
}


func (this *MyQueue) Empty() bool {
	return this.a.Len() == 0 && this.b.Len() == 0
}

type Stack struct {
	nums []int
}

func NewStack() *Stack  {
	return &Stack{nums:[]int{}}
}

func (s *Stack) Push(n int)  {
	s.nums = append(s.nums,n)
}

func (s *Stack) Pop() int  {
	res := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return res
}

func (s *Stack) Len() int  {
	return len(s.nums)
}

func (s *Stack) IsEmpty() bool  {
	return s.Len() == 0
}*/
/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

