package main

import "fmt"

func main() {
	s := "Привет!"
	var runes []rune

	i := 0
loop:
	for i < len(s) {
		var r rune
		b := s[i]
		switch {
		case b&0b10000000 == 0b00000000: // 1-byte character: 0xxxxxxx
			fmt.Println("One byte")
			r = rune(b)
			i += 1
		case b&0b11100000 == 0b11000000: // 2-byte character: 110xxxxx 10xxxxxx
			fmt.Println("Two byte")
			if i+1 >= len(s) {
				break loop // Handle incomplete byte sequence
			}
			r = (rune(b&0b00011111) << 6) | (rune(s[i+1] & 0b00111111))
			i += 2
		case b&0b11110000 == 0b11100000: // 3-byte character: 1110xxxx 10xxxxxx 10xxxxxx
			fmt.Println("Three byte")
			if i+2 >= len(s) {
				break loop // Handle incomplete byte sequence
			}
			r = (rune(b&0b00001111) << 12) | (rune(s[i+1]&0b00111111) << 6) | (rune(s[i+2] & 0b00111111))
			i += 3
		case b&0b11111000 == 0b11110000: // 4-byte character: 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
			fmt.Println("Four byte")
			if i+3 >= len(s) {
				break loop // Handle incomplete byte sequence
			}
			r = (rune(b&0b00000111) << 18) | (rune(s[i+1]&0b00111111) << 12) | (rune(s[i+2]&0b00111111) << 6) | (rune(s[i+3] & 0b00111111))
			i += 4
		default:
			fmt.Println("Invalid UTF-8 sequence")
			break loop
		}

		runes = append(runes, r)
	}

	// Print the result
	fmt.Printf("Decoded runes: %U\n", runes)
	fmt.Println(string(runes))
}
