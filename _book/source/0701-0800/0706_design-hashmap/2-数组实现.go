package main

func main() {

}

type MyHashMap struct {
	keys  [10000][]int
	value [10000][]int
}

func Constructor() MyHashMap {
	return MyHashMap{
		keys:  [10000][]int{},
		value: [10000][]int{},
	}
}

func (m *MyHashMap) Put(key int, value int) {
	for k, v := range m.keys[key%10000] {
		if v == key {
			m.value[key%10000][k] = value
			return
		}
	}
	m.keys[key%10000] = append(m.keys[key%10000], key)
	m.value[key%10000] = append(m.value[key%10000], value)
}

func (m *MyHashMap) Get(key int) int {
	for k, v := range m.keys[key%10000] {
		if v == key {
			return m.value[key%10000][k]
		}
	}
	return -1
}

func (m *MyHashMap) Remove(key int) {
	for k, v := range m.keys[key%10000] {
		if v == key {
			m.keys[key%10000] = append(m.keys[key%10000][:k], m.keys[key%10000][k+1:]...)
			m.value[key%10000] = append(m.value[key%10000][:k], m.value[key%10000][k+1:]...)
		}
	}
}
