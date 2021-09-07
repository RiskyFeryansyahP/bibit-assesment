package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Please refactor the code below to make it more concise, efficient and readable with good logic flow.")

	fmt.Printf("%s\n", findFirstStringInBracket("(Test)"))
}

// here i am use the standart of happy path Go, which
// Align the happy path to the left;
// you should quickly be able to scan down one column to see the expected execution flow
func findFirstStringInBracket(str string) string {
	if len(str) < 0 {
		return ""
	}

	indexFirstBracketFound := strings.Index(str, "(")
	if indexFirstBracketFound < 0 {
		return ""
	}

	runes := []rune(str)

	wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
	indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
	if indexClosingBracketFound < 0 {
		return ""
	}

	runes = []rune(wordsAfterFirstBracket)

	return string(runes[1 : indexClosingBracketFound-1])
}
