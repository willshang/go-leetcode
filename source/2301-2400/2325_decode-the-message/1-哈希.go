package main

func main() {

}

func decodeMessage(key string, message string) string {
	m := make(map[byte]byte)
	cur := byte('a')
	for i := 0; i < len(key); i++ {
		if key[i] != ' ' && m[key[i]] == 0 {
			m[key[i]] = cur
			cur++
		}
	}
	arr := []byte(message)
	for i := 0; i < len(message); i++ {
		if message[i] != ' ' {
			arr[i] = m[message[i]]
		}
	}
	return string(arr)
}
