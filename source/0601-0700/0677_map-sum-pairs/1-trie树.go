package main

func main() {

}

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
