package main

import (
	"fmt"
)

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
	// ↦ → → →
	//      ↙
	//    ↙
	//  ↙
	// → → → ⇥
	formula := maxColumns*(currentRow-1) + currentColumn

	//  ← ← ← ↤
	//   ↘
	//     ↘
	//       ↘
	//  ⇤ ← ← ←
	// formula := maxColumns*(currentRow-1) + maxColumns - currentColumn + 1

	//  ⇤ ← ← ←
	//       ↗
	//     ↗
	//   ↗
	//  ← ← ← ↤
	// formula := maxColumns*(maxRows-currentRow) + maxColumns - currentColumn + 1

	//  → → → ⇥
	//   ↖
	//     ↖
	//       ↖
	//  ↦ → → →
	// formula := maxColumns*(maxRows-currentRow) + currentColumn

	// ↧      ↗↓
	// ↓    ↗  ↓
	// ↓  ↗    ↓
	// ↓↗      ⤓
	// formula := maxRows*(currentColumn-1) + currentRow

	// ↑↘      ⤒
	// ↑  ↘    ↑
	// ↑    ↘  ↑
	// ↥      ↘↑
	// formula := maxRows*(currentColumn-1) + maxRows - currentRow + 1

	// ↓↖      ↧
	// ↓  ↖    ↓
	// ↓    ↖  ↓
	// ⤓      ↖↓
	// formula := maxRows*(maxColumns-currentColumn) + currentRow

	// ⤒      ↙↑
	// ↑    ↙  ↑
	// ↑  ↙    ↑
	// ↑↙      ↥
	// formula := maxRows*(maxColumns-currentColumn) + maxRows - currentRow + 1

	//  ↦ → → →
	//         ↓
	//  ← ← ← ←
	// ↓
	//  → → → ⇥
	// formula := (int(math.Max(0, math.Pow(-1, float64(currentRow+1)))))*(maxColumns*(currentRow-1)+currentColumn) + ((int(math.Max(0, math.Pow(-1, float64(currentRow))))) * (maxColumns*(currentRow-1) + maxColumns - currentColumn + 1))

	//  ← ← ← ↤
	// ↓
	//  → → → →
	//         ↓
	//  ⇤ ← ← ←
	// formula := ((int(math.Max(0, math.Pow(-1, float64(currentRow+1))))) * (maxColumns*(currentRow-1) + maxColumns - currentColumn + 1)) + (int(math.Max(0, math.Pow(-1, float64(currentRow)))) * (maxColumns*(currentRow-1) + currentColumn))

	//  ⇤ ← ← ←
	//         ↑
	//  → → → →
	// ↑
	//  ← ← ← ↤
	// formula := ((int(math.Max(0, math.Pow(-1, float64(currentRow+1))))) * (maxColumns*(maxRows-currentRow) + maxColumns - currentColumn + 1)) + (int(math.Max(0, math.Pow(-1, float64(currentRow)))) * (maxColumns*(maxRows-currentRow) + currentColumn))

	//  → → → ⇥
	// ↑
	//  ← ← ← ←
	//         ↑
	//  ↦ → → →
	// formula := ((int(math.Max(0, math.Pow(-1, float64(currentRow+1))))) * (maxColumns*(maxRows-currentRow) + currentColumn)) + (int(math.Max(0, math.Pow(-1, float64(currentRow)))) * (maxColumns*(maxRows-currentRow) + maxColumns - currentColumn + 1))

	// ↧   ↑ → ↓
	// ↓   ↑   ↓
	// ↓   ↑   ↓
	// ↓   ↑   ↓
	// ↓ → ↑   ⤓
	// formula := ((int(math.Max(0, math.Pow(-1, float64(currentColumn+1))))) * (maxRows*(currentColumn-1) + currentRow)) + (int(math.Max(0, math.Pow(-1, float64(currentColumn)))) * (maxRows*(currentColumn-1) + maxRows - currentRow + 1))

	// ↑ → ↓   ⤒
	// ↑   ↓   ↑
	// ↑   ↓   ↑
	// ↑   ↓   ↑
	// ↥   ↓ → ↑
	// formula := ((int(math.Max(0, math.Pow(-1, float64(currentColumn+1))))) * (maxRows*(currentColumn-1) + maxRows - currentRow + 1)) + (int(math.Max(0, math.Pow(-1, float64(currentColumn)))) * (maxRows*(currentColumn-1) + currentRow))

	// ⤒   ↓ ← ↑
	// ↑   ↓   ↑
	// ↑   ↓   ↑
	// ↑   ↓   ↑
	// ↑ ← ↓   ↥
	// formula := ((int(math.Max(0, math.Pow(-1, float64(currentColumn+1))))) * (maxRows*(maxColumns-currentColumn) + maxRows - currentRow + 1)) + (int(math.Max(0, math.Pow(-1, float64(currentColumn)))) * (maxRows*(maxColumns-currentColumn) + currentRow))

	// ↓ ← ↑   ↧
	// ↓   ↑   ↓
	// ↓   ↑   ↓
	// ↓   ↑   ↓
	// ⤓   ↑ ← ↓
	// formula := ((int(math.Max(0, math.Pow(-1, float64(currentColumn+1))))) * (maxRows*(maxColumns-currentColumn) + currentRow)) + (int(math.Max(0, math.Pow(-1, float64(currentColumn)))) * (maxRows*(maxColumns-currentColumn) + maxRows - currentRow + 1))

	return formula
}
