package main

func main() {

}

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
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	res := "http://tinyurl.com/"
	this.index++
	count := this.index
	temp := make([]byte, 0)
	for count > 0 {
		temp = append(temp, str[count%62])
		count = count / 62
	}
	res = res + string(temp)
	this.m[res] = longUrl
	return res
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return this.m[shortUrl]
}
