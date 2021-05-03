package main

import "strconv"

func main() {

}

// leetcode535_TinyURL的加密与解密
type Codec struct {
	m     map[string]string
	index int
}

func Constructor() Codec {
	return Codec{
		m:     make(map[string]string),
		index: 1,
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	res := "http://tinyurl.com/" + strconv.Itoa(this.index)
	this.m[res] = longUrl
	this.index++
	return res
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return this.m[shortUrl]
}
