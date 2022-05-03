package main

func main() {

}

// leetcode2227_加密解密字符串
type Encrypter struct {
	arr [26]string
	m   map[string]int
}

func Constructor(keys []byte, values []string, dictionary []string) Encrypter {
	arr := [26]string{}
	for i := 0; i < len(keys); i++ {
		arr[int(keys[i]-'a')] = values[i]
	}
	e := Encrypter{
		arr: arr,
		m:   map[string]int{},
	}
	// 加密所有值
	for i := 0; i < len(dictionary); i++ {
		e.m[e.Encrypt(dictionary[i])]++
	}
	return e
}

func (this *Encrypter) Encrypt(word1 string) string {
	res := make([]byte, 0)
	for i := 0; i < len(word1); i++ {
		v := this.arr[int(word1[i]-'a')]
		if v == "" {
			return ""
		}
		res = append(res, v...)
	}
	return string(res)
}

func (this *Encrypter) Decrypt(word2 string) int {
	return this.m[word2]
}
