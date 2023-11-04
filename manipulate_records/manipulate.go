package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	EDIT = "edit"
	EXIT = "exit"
)

func printRecordsByAction(records map[string]string, action string, printableActions []string, printReversed bool) {
	if contains(printableActions, action) {
		keys := make([]string, 0, len(records))
		for key := range records {
			keys = append(keys, key)
		}

		if printReversed {
			for i := len(keys) - 1; i >= 0; i-- {
				key := keys[i]
				value := records[key]
				fmt.Printf("%s %s\n", key, value)
			}
		} else {
			for _, key := range keys {
				value := records[key]
				fmt.Printf("%s %s\n", key, value)
			}
		}
	}
}

func contains(actions []string, action string) bool {
	for _, a := range actions {
		if a == action {
			return true
		}
	}
	return false
}

func inputAction(records map[string]string) (map[string]string, string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	parts := strings.Split(input, " ")
	action := parts[0]

	if action == EDIT && len(parts) == 3 {
		key := parts[1]
		value := parts[2]
		records[key] = value
	} else if action == EXIT {
		return records, action
	} else {
		fmt.Println("Invalid action or values")
	}

	return records, action
}

func inputLengthRecords() int {
	var length int
	fmt.Scanln(&length)
	return length
}

func inputRecordsByLength(length int) map[string]string {
	records := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < length; i++ {
		scanner.Scan()
		input := scanner.Text()
		parts := strings.Split(input, " ")
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			records[key] = value
		} else {
			fmt.Println("Invalid input, skipping this record.")
		}
	}

	return records
}

func main() {
	lengthRecords := inputLengthRecords()
	records := inputRecordsByLength(lengthRecords)

	for {
		records, action := inputAction(records)
		printRecordsByAction(records, action, []string{EDIT}, true)

		if action == EXIT {
			break
		}
	}
}
