package counter

import "strings"

func Increment(s string) string {
	if len(s) > 0 && s[0] == '-' {
		abs := s[1:]
		if abs == "" || abs == "0" {
			return "1"
		}
		dec := decrement(abs)
		if dec == "0" {
			return "0"
		}
		return "-" + dec
	}
	return incrementPositive(s)
}

func incrementPositive(s string) string {
	b := []byte(s)
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] != '9' {
			b[i]++
			return string(b)
		}
		b[i] = '0'
	}
	return "1" + strings.Repeat("0", len(b))
}

func decrement(s string) string {
	b := []byte(s)
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] != '0' {
			b[i]--
			break
		}
		b[i] = '9'
	}
	i := 0
	for i < len(b)-1 && b[i] == '0' {
		i++
	}
	return string(b[i:])
}
