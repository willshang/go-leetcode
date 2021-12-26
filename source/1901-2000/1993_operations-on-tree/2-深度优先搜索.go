package main

func main() {

}

// leetcode1993_树上的操作
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
	v := num
	for {
		if v >= 0 {
			if _, ok := this.m[v]; ok {
				return false
			}
			v = this.parent[v]
		} else {
			break
		}
	}
	if this.hasLock(num) == false {
		return false
	}
	this.m[num] = user
	this.dfsUnLock(num)
	return true
}

func (this *LockingTree) hasLock(num int) bool {
	for i := 0; i < len(this.next[num]); i++ {
		v := this.next[num][i]
		if _, ok := this.m[v]; ok {
			return true
		}
		if this.hasLock(v) {
			return true
		}
	}
	return false
}

func (this *LockingTree) dfsUnLock(num int) {
	for i := 0; i < len(this.next[num]); i++ {
		v := this.next[num][i]
		delete(this.m, v)
		this.dfsUnLock(v)
	}
}
