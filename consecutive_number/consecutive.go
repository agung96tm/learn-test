package main

import (
	"fmt"
)

func findConsecutiveNumbers(targetSum int) (result [][]int) {
	start := 1
	end := 1
	currentSum := start

	for start <= targetSum/2 {
		if currentSum == targetSum && start != end {
			sequence := make([]int, end-start+1)
			for i := range sequence {
				sequence[i] = start + i
			}
			result = append(result, sequence)
		}
		if currentSum < targetSum {
			end++
			currentSum += end
		} else {
			currentSum -= start
			start++
		}
	}

	return
}

func main() {
	var user_input int
	fmt.Print("input: ")
	_, err := fmt.Scanf("%d", &user_input)
	if err != nil {
		fmt.Println("Invalid input.")
		return
	}

	result := findConsecutiveNumbers(user_input)
	var resultStrings []string

	for _, sequence := range result {
		sequenceString := "["
		for i, num := range sequence {
			sequenceString += fmt.Sprintf("%d", num)
			if i < len(sequence)-1 {
				sequenceString += ", "
			}
		}
		sequenceString += "]"
		resultStrings = append(resultStrings, sequenceString)
	}

	fmt.Print("output: ")
	for i, rString := range resultStrings {
		fmt.Printf("%s", rString)
		if i < len(resultStrings)-1 {
			fmt.Print(", ")
		}
	}
}
