package main

import (
	"fmt"
	"strings"
)

func mod(dividend, divisor int) int {
	return ((dividend % divisor) + divisor) % divisor
}

func createPyramid(height, rotation int) {
	rotation = mod(rotation, 4)
	var pyramid []string

	if rotation == 0 {
		for i := 1; i <= height; i++ {
			spaces := strings.Repeat(" ", height-i)
			stars := strings.Repeat("*", 2*i-1)
			pyramid = append(pyramid, spaces+stars)
		}
	} else if rotation == 1 {
		for i := 1; i <= height; i++ {
			pyramid = append(pyramid, strings.Repeat("*", i))
		}
		for i := height - 1; i > 0; i-- {
			pyramid = append(pyramid, strings.Repeat("*", i))
		}
	} else if rotation == 2 {
		for i := height; i > 0; i-- {
			spaces := strings.Repeat(" ", height-i)
			stars := strings.Repeat("*", 2*i-1)
			pyramid = append(pyramid, spaces+stars)
		}
	} else if rotation == 3 {
		for i := 1; i <= height; i++ {
			spaces := strings.Repeat(" ", height-i)
			stars := strings.Repeat("*", i)
			pyramid = append(pyramid, spaces+stars)
		}
		for i := height - 1; i > 0; i-- {
			spaces := strings.Repeat(" ", height-i)
			stars := strings.Repeat("*", i)
			pyramid = append(pyramid, spaces+stars)
		}
	}

	if len(pyramid) == 0 {
		fmt.Println("Invalid input")
	}

	for _, p := range pyramid {
		fmt.Println(p)
	}
}

func main() {
	var height, rotation int

	fmt.Print("Height: ")
	fmt.Scan(&height)
	fmt.Print("Rotation: ")
	fmt.Scan(&rotation)

	createPyramid(height, rotation)
}
