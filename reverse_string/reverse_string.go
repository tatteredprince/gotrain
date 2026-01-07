package main

// reverseString reverses string
func reverseString(str string) string {
	reversed := ""
	chars := []rune(str)
	for i := len(chars) - 1; i >= 0; i-- {
		reversed += string(chars[i])
	}
	return reversed
}
