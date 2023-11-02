package main

import (
	"fmt"
	"sort"
	"strings"
)

func areAnagramsWithBuiltinFunc(str1, str2 string) bool {
	str1 = strings.ReplaceAll(str1, " ", "")
	str2 = strings.ReplaceAll(str2, " ", "")
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	return sortString(str1) == sortString(str2)
}

func sortString(s string) string {
	sChars := strings.Split(s, "")
	sort.Strings(sChars)
	return strings.Join(sChars, "")
}

func areAnagramsManual(str1, str2 string) bool {
	str1 = strings.ReplaceAll(str1, " ", "")
	str2 = strings.ReplaceAll(str2, " ", "")
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	if len(str1) != len(str2) {
		return false
	}

	charCount1 := make(map[rune]int)
	charCount2 := make(map[rune]int)

	for _, char := range str1 {
		charCount1[char]++
	}

	for _, char := range str2 {
		charCount2[char]++
	}

	for char, count := range charCount1 {
		if count != charCount2[char] {
			return false
		}
	}

	return true
}

func main() {
	var firstStr, lastStr string

	fmt.Print("Input: ")
	fmt.Scan(&firstStr)
	fmt.Print("Input: ")
	fmt.Scan(&lastStr)

	isAnagramManual := areAnagramsManual(firstStr, lastStr)
	fmt.Printf("Is anagram? %v\n", isAnagramManual)

	//isAnagramBuiltin := areAnagramsWithBuiltinFunc(firstStr, lastStr)
	//fmt.Printf("Is anagram (Builtin)? %v\n", isAnagramBuiltin)
}
