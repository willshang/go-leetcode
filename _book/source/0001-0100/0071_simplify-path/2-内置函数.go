package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(simplifyPath("/home/"))
}

func simplifyPath(path string) string {
	return filepath.Clean(path)
}
