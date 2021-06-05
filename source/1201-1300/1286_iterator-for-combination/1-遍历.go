package main

func main() {

}

type CombinationIterator struct {
	arr    []byte
	str    string
	flag   bool
	length int
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	return CombinationIterator{
		arr:    []byte(characters[:combinationLength]),
		str:    characters,
		flag:   false,
		length: combinationLength,
	}
}

func (this *CombinationIterator) Next() string {

}

func (this *CombinationIterator) HasNext() bool {
	return this.flag == false
}
