package main

func main() {

}

type MapSum struct {
	m    map[string]int
	data map[string]map[string]bool
}

func Constructor() MapSum {
	return MapSum{
		m:    make(map[string]int),
		data: make(map[string]map[string]bool),
	}
}

func (this *MapSum) Insert(key string, val int) {
	this.m[key] = val
	for i := 1; i <= len(key); i++ {
		str := key[:i]
		if _, ok := this.data[str]; ok == false {
			this.data[str] = make(map[string]bool)
		}
		this.data[str][key] = true
	}
}

func (this *MapSum) Sum(prefix string) int {
	res := 0
	for key := range this.data[prefix] {
		res = res + this.m[key]
	}
	return res
}
