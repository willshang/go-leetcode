package main

func main() {

}

type LockingTree struct {
	m      map[int]int   // 判断该节点是否上锁
	parent map[int]int   // 父节点
	next   map[int][]int // 保存子节点数组
}

func Constructor(parent []int) LockingTree {
	temp := make(map[int]int)
	next := make(map[int][]int)
	for i := 0; i < len(parent); i++ {
		temp[i] = parent[i]
		next[parent[i]] = append(next[parent[i]], i)
	}
	return LockingTree{m: make(map[int]int), parent: temp, next: next}
}

func (this *LockingTree) Lock(num int, user int) bool {
	if _, ok := this.m[num]; ok {
		return false
	}
	this.m[num] = user
	return true
}

func (this *LockingTree) Unlock(num int, user int) bool {
	if v, ok := this.m[num]; ok == false || v != user {
		return false
	}
	delete(this.m, num)
	return true
}

func (this *LockingTree) Upgrade(num int, user int) bool {
	if _, ok := this.m[num]; ok == true {
		return false // 条件1
	}
	queue := make([]int, 0)
	queue = append(queue, this.next[num]...)
	flag := false
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node < len(this.parent) {
			queue = append(queue, this.next[node]...)
			if _, ok := this.m[node]; ok {
				flag = true
				break
			}
		}
	}
	if flag == false {
		return false // 条件2：至少有1个上锁的子孙节点
	}
	flag = false
	p := this.parent[num]
	for p != -1 {
		if _, ok := this.m[p]; ok {
			flag = true
		}
		p = this.parent[p]
	}
	if flag == true { // 条件3：指定节点没有任何上锁的祖先节点
		return false
	}
	this.m[num] = user // 下面是子孙节点解锁
	queue = make([]int, 0)
	queue = append(queue, this.next[num]...)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node < len(this.parent) {
			queue = append(queue, this.next[node]...)
			delete(this.m, node)
		}
	}
	return true
}
