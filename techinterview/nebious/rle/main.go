package main

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrWrongSymbol = errors.New("Wrong symbol")
)

func rle(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}
	if len(s) == 1 {
		return s, nil
	}
	runes := []rune(s)
	rc := runes[0]
	c := 1
	sr := strings.Builder{}
	add := func() {
		sr.WriteRune(rc)
		if c > 1 {
			sr.WriteString(fmt.Sprintf("%d", c))
		}
	}
	for _, r := range runes[1:] {
		if r < rune('A') || rune('Z') < r {
			return "", ErrWrongSymbol
		}
		if rc == r {
			c++
		} else {
			add()
			c = 1
			rc = r
		}
	}
	add()
	return sr.String(), nil
}

func rle2(s string) (string, error) {
	rc := rune(-1)
	c := 0
	sr := strings.Builder{}
	add := func() {
		sr.WriteRune(rc)
		if c > 1 {
			sr.WriteString(fmt.Sprintf("%d", c))
		}
	}
	for _, r := range s {
		if r < rune('A') || rune('Z') < r {
			return "", ErrWrongSymbol
		}
		if rc == r {
			c++
		} else if rc == rune(-1) {
			c++
			rc = r
		} else {
			add()
			c = 1
			rc = r
		}
	}
	add()
	return sr.String(), nil
}

func writeNumber(sr *strings.Builder, c int) {
	rs := make([]byte, 0, 20)
	for c > 0 {
		rs = append(rs, byte('0'+c%10))
		c = c / 10
	}
	for i := len(rs) - 1; i >= 0; i-- {
		sr.WriteByte(rs[i])
	}
}

func rle3(s string) (string, error) {
	rc := []rune(s[:1])[0]
	c := 0
	sr := strings.Builder{}
	add := func() {
		sr.WriteRune(rc)
		if c > 1 {
			writeNumber(&sr, c)
		}
	}
	ra := rune('A')
	rz := rune('Z')
	for _, r := range s {
		if r < ra || rz < r {
			return "", ErrWrongSymbol
		}
		if rc == r {
			c++
		} else {
			add()
			c = 1
			rc = r
		}
	}
	add()
	return sr.String(), nil
}
