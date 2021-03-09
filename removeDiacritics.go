/*
MIT License

Copyright (c) 2021 Elsanussi Mneina

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

// So far the program only removes a single diacritic.
const arabic_fathatan rune = '\u064B'
const arabic_dammatan rune = '\u064C'
const arabic_kasratan rune = '\u064D'
const arabic_fatha rune = '\u064E'
const arabic_damma rune = '\u064F'
const arabic_kasra rune = '\u0650'
const arabic_shadda rune = '\u0651'
const arabic_sukun rune = '\u0652'
const arabic_subscript_alef rune = '\u0656'
const arabic_letter_superscript_alef rune = '\u0670'

// https://golang.org/ref/spec#Rune_literals
// https://golangbyexample.com/iterate-over-a-string-golang/

func TextFilter(text string) string {
	result := ""
	for _, aRune := range text {
		if aRune == arabic_fathatan {

		} else if aRune == arabic_dammatan {
		} else if aRune == arabic_kasratan {

		} else if aRune == arabic_fatha {

		} else if aRune == arabic_damma {

		} else if aRune == arabic_kasra {

		} else if aRune == arabic_shadda {
		} else {
			result += string(aRune)
		}
	}
	return result
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		ucl := TextFilter(scanner.Text())

		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
