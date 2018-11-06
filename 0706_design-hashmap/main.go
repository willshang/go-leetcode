package main

func main() {
	
}
type MyHashMap struct {
	table []int
}


/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{
		table:make([]int,1000001),
	}
}


/** value will always be positive. */
func (this *MyHashMap) Put(key int, value int)  {
	this.table[key] = value+1
}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	return this.table[key] - 1
}


/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
	this.table[key] = 0
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */