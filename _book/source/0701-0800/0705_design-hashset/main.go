package main

func main() {

}

type MyHashSet struct {
	table []bool
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{
		table: make([]bool, 1000001),
	}
}

func (this *MyHashSet) Add(key int) {
	this.table[key] = true
}

func (this *MyHashSet) Remove(key int) {
	this.table[key] = false
}

/** Returns true if this set did not already contain the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return this.table[key]
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
