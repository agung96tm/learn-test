package main

import (
	"fmt"
	"strconv"
	"strings"
)

//func removeItemsInList(originItems []int, removeItems []int) []int {
//	// Create a map of items to remove for fast lookups
//	removeMap := make(map[int]bool)
//	for _, item := range removeItems {
//		removeMap[item] = true
//	}
//
//	// Create a result slice for items that should not be removed
//	var result []int
//	for _, item := range originItems {
//		if !removeMap[item] {
//			result = append(result, item)
//		}
//	}
//
//	return result
//}

func removeItemsInList(originItems []int, removeItems []int) []int {
	var result []int

	for _, item := range originItems {
		shouldRemove := false
		for _, rmItem := range removeItems {
			if item == rmItem {
				shouldRemove = true
				break
			}
		}
		if !shouldRemove {
			result = append(result, item)
		}
	}

	return result
}

func main() {
	var oriItems, rmItems []int

	fmt.Print("Items: ")
	var itemsInput string
	fmt.Scanln(&itemsInput)
	oriItemsStr := strings.Fields(itemsInput)
	for _, itemStr := range oriItemsStr {
		item, err := strconv.Atoi(itemStr)
		if err == nil {
			oriItems = append(oriItems, item)
		}
	}

	fmt.Print("Remove items: ")
	var removeItemsInput string
	fmt.Scanln(&removeItemsInput)
	rmItemsStr := strings.Fields(removeItemsInput)
	for _, rmItemStr := range rmItemsStr {
		rmItem, err := strconv.Atoi(rmItemStr)
		if err == nil {
			rmItems = append(rmItems, rmItem)
		}
	}

	newList := removeItemsInList(oriItems, rmItems)
	fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(newList)), " "), "[]"))
}
