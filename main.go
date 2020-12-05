package main

import "fmt"

func main() {
	type tMatrix [5][5]int
	var matrix tMatrix

	// Building the matrix
	for rowIndex, row := range matrix {
		for columnIndex := range row {
			matrix[rowIndex][columnIndex] = calculateFormula(rowIndex+1, columnIndex+1, 5, 5)
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

func calculateFormula(currentRow int, currentColumn int, maxRows int, maxColumns int) int {
	// |---->
	//     /
	//    /
	//   /
	//  ---->|
	// formula := maxColumns*(currentRow-1) + currentColumn

	//  <----|
	//   \
	//    \
	//     \
	// |<----
	// formula := maxColumns*(currentRow-1) + maxColumns - currentColumn + 1

	// |<----
	//     /
	//    /
	//   /
	//  <----|
	// formula := maxColumns*(maxRows-currentRow) + maxColumns - currentColumn + 1

	//  ---->|
	//   \
	//    \
	//     \
	// |---->
	formula := maxColumns*(maxRows-currentRow) + currentColumn

	return formula
}
