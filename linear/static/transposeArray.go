package main

import "fmt"

func main() {
	var i, j, rows, columns int

	var mat [10][10]int
	var transMat [10][10]int

	fmt.Print("Enter the Matrix rows and Columns = ")
	fmt.Scan(&rows, &columns)

	fmt.Println("Enter Matrix Items to Transpose = ")
	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			fmt.Scan(&mat[i][j])
		}
	}
	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			transMat[j][i] = mat[i][j]
		}
	}
	fmt.Println(" The Matrix ")
	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			fmt.Print(mat[i][j], "  ")
		}
		fmt.Println()
	}
	fmt.Println(" The Transpose Matrix ")
	for i = 0; i < columns; i++ {
		for j = 0; j < rows; j++ {
			fmt.Print(transMat[i][j], "  ")
		}
		fmt.Println()
	}
}
