package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func Unpacked(str string) (string, error) {

	var prev_rune, curr_rune, next_rune rune
	var is_slash bool
	var n int

	runes := []rune(str)
	res := make([]rune, 0)

	prev_rune = runes[0]
	is_slash = false
	n = len(runes)

	for i := 1; i < n; i++ {
		curr_rune = runes[i]
		if i < n-1 {
			next_rune = runes[i+1]
		}

		fmt.Printf("[i = %d] [prev = %s] [curr = %s] [next = %s] [is_slash = %t]\n",
			i, string(prev_rune), string(curr_rune), string(next_rune), is_slash)

		if curr_rune == rune('\\') {

			if unicode.IsDigit(next_rune) || next_rune == rune('\\') {
				is_slash = true
				continue
			} else {
				panic(errors.New("Некорректная строка"))
			}
		}

		if is_slash {
			//mt.Println("Add 1")
			res = append(res, prev_rune)

			if i == n-1 {
				//fmt.Println("Add 2")
				res = append(res, curr_rune)
			}

			prev_rune = curr_rune
			is_slash = false

			continue
		}

		if unicode.IsDigit(curr_rune) && prev_rune != rune('\\') {
			r := []rune(strings.Repeat(string(prev_rune), int(curr_rune-'0')))
			//fmt.Println("Add 3")
			//fmt.Println(string(r))

			res = append(res, r...)
		}

		if !unicode.IsDigit(prev_rune) {
			//fmt.Println("Add 4")
			res = append(res, prev_rune)
		}

		if i == n-1 {
			//fmt.Println("Add 5")
			res = append(res, curr_rune)
		}

		prev_rune = curr_rune
	}

	return string(res), nil
}

func main() {

	str := `q1w2e3\45`
	res, _ := Unpacked(str)

	fmt.Println(str, " -> ", string(res))
}
