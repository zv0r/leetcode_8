package main

import "fmt"

func main() {
	fmt.Println(myAtoi("42") == 42)
	fmt.Println(myAtoi("-042") == -42)
	fmt.Println(myAtoi("-42") == -42)
	fmt.Println(myAtoi("4193 with words") == 4193)
	fmt.Println(myAtoi("words and 987") == 0)
	fmt.Println(myAtoi("words and -987") == 0)
	fmt.Println(myAtoi("1337c0d3") == 1337)
	fmt.Println(myAtoi("-91283472332") == -2147483648)
	fmt.Println(myAtoi("+1") == 1)
	fmt.Println(myAtoi("+-12") == 0)
	fmt.Println(myAtoi("9223372036854775808") == 2147483647)
}

func myAtoi(s string) int {
	res := 0
	multiplier := 1
	parsed := false
	for _, c := range s {
		if !parsed && c == ' ' {
			continue
		}
		if !parsed && (c == '-' || c == '+') {
			if c == '-' {
				multiplier = -1
			}
			parsed = true
			continue
		}
		if c >= '0' && c <= '9' {
			parsed = true
			res = res*10 + int(c-'0')
			if res > 2147483647 && multiplier == 1 {
				return 2147483647
			} else if res > 2147483648 && multiplier == -1 {
				return -2147483648
			}
			continue
		}
		break
	}
	res *= multiplier
	return res
}
