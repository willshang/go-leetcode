package main

func main() {

}

type MyStack struct {
	a []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.a = append(this.a, x)
}
func (this *MyStack) Pop() int {
	if len(this.a) == 0 {
		return 0
	}
	last := this.a[len(this.a)-1]
	this.a = this.a[0 : len(this.a)-1]
	return last
}

func (this *MyStack) Top() int {
	if len(this.a) == 0 {
		return 0
	}
	return this.a[len(this.a)-1]
}

func (this *MyStack) Empty() bool {
	if len(this.a) == 0 {
		return true
	}
	return false
}

/*type MyStack struct {
	a, b *Queue
}


func Constructor() MyStack {
	return MyStack{
		a:NewQueue(),
		b:NewQueue(),
	}
}


func (this *MyStack) Push(x int)  {
	if this.a.Len() == 0{
		this.a, this.b = this.b,this.a
	}
	this.a.Push(x)
}


func (this *MyStack) Pop() int {
	if this.a.Len() == 0{
		this.a, this.b = this.b , this.a
	}

	for this.a.Len() > 1{
		this.b.Push(this.a.Pop())
	}
	return this.a.Pop()
}


func (this *MyStack) Top() int {
	res := this.Pop()
	this.Push(res)
	return res
}


func (this *MyStack) Empty() bool {
	return (this.a.Len() + this.b.Len()) == 0
}

type Queue struct {
	nums []int
}

func NewQueue() *Queue  {
	return &Queue{nums:[]int{}}
}

func (q *Queue) Push(n int)  {
	q.nums = append(q.nums,n)
}

func (q *Queue) Pop() int  {
	res := q.nums[0]
	q.nums = q.nums[1:]
	return res
}

func (q *Queue) Len() int  {
	return len(q.nums)
}
func (q *Queue) IsEmpty() bool  {
	return q.Len() == 0
}*/
/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
