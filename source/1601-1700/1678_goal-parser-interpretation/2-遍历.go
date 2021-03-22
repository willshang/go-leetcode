package main

func main() {

}

func interpret(command string) string {
	res := ""
	for i := 0; i < len(command); {
		if command[i] == 'G' {
			res = res + "G"
			i = i + 1
		} else if command[i] == '(' {
			if command[i+1] == ')' {
				res = res + "o"
				i = i + 2
			} else {
				res = res + "al"
				i = i + 4
			}
		}
	}
	return res
}
