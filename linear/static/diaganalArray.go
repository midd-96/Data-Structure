package main

import "fmt"

func main() {
	var i, j, rows, columns int

	var mat [10][10]int

	fmt.Print("Enter the Matrix Rows and Columns ")
	fmt.Scan(&rows, &columns)

	fmt.Println("Enter Matrix")
	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			fmt.Scan(&mat[i][j])
		}
	}
	fmt.Println("The Matrix")
	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			fmt.Print(mat[i][j], " ")
		}
		fmt.Println()
	}
	if rows == columns {
		for i = 0; i < rows; i++ {

			mat[i][i], mat[i][rows-i-1] = mat[i][rows-i-1], mat[i][i]

		}
		fmt.Println("The Matrix after Interchanging Diagonals ")
		for i = 0; i < rows; i++ {
			for j = 0; j < columns; j++ {
				fmt.Print(mat[i][j], "  ")
			}
			fmt.Println()
		}
	} else {
		fmt.Println("The Given Matrix is Not a Square Matrix")
	}
}
