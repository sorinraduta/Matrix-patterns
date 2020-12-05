package main

import "fmt"

type tMatrix [5][5]int

func main() {
	var matrix tMatrix

	// Building the matrix
	c := 0
	for rowIndex, row := range matrix {
		for columnIndex := range row {
			c++
			matrix[rowIndex][columnIndex] = c
		}
	}

	// Printing the matrix
	for rowIndex, row := range matrix {
		fmt.Printf("\n ____    ____    ____    ____    ____  \n")
		for columnIndex := range row {
			fmt.Printf("| %0.2d |\t", matrix[rowIndex][columnIndex])
		}
		fmt.Printf("\n ‾‾‾‾    ‾‾‾‾    ‾‾‾‾    ‾‾‾‾    ‾‾‾‾  ")
	}
}
