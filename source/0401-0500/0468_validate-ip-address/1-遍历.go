package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(validIPAddress("172.16.254.1"))
}

// leetcode468_验证IP地址
var defaultRes string = "Neither"

func validIPAddress(IP string) string {
	if strings.Contains(IP, ":") {
		return checkIPv6(IP)
	}
	return checkIPv4(IP)
}

func checkIPv4(ip string) string {
	arr := strings.Split(ip, ".")
	if len(arr) != 4 {
		return defaultRes
	}
	for i := 0; i < len(arr); i++ {
		num, err := strconv.Atoi(arr[i])
		if err != nil || num > 255 || num < 0 {
			return defaultRes
		}
		if len(arr[i]) > 1 && arr[i][0] == '0' || (i == 0 && num == 0) {
			return defaultRes
		}
	}
	return "IPv4"
}

func checkIPv6(ip string) string {
	arr := strings.Split(ip, ":")
	if len(arr) != 8 {
		return defaultRes
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == "" || len(arr[i]) > 4 {
			return defaultRes
		}
		for j := 0; j < len(arr[i]); j++ {
			if (arr[i][j] > 'F' && arr[i][j] < 'a') ||
				(arr[i][j] > 'f' && arr[i][j] <= 'z') {
				return defaultRes
			}
		}
	}
	return "IPv6"
}
