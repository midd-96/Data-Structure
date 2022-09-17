package main

import "fmt"

var arr1 [100][100]int
var arr2 [100][100]int
var limit, i, j int

func getArray() {
	fmt.Println("Enter first array")
	for i = 0; i < limit; i++ {
		for j = 0; j < limit; j++ {
			fmt.Scan(&arr1[i][j])
		}
	}
	fmt.Println("Enter second array")
	for i = 0; i < limit; i++ {
		for j = 0; j < limit; j++ {
			fmt.Scan(&arr2[i][j])
		}
	}
}
func swapArray() {
	for i = 0; i < limit; i++ {
		for j = 0; j < limit; j++ {
			arr1[i][j], arr2[i][j] = arr2[i][j], arr1[i][j]
		}
		fmt.Println()
	}
}

func dispArray() {
	fmt.Println("First array: ")
	for i = 0; i < limit; i++ {
		for j = 0; j < limit; j++ {
			fmt.Print(arr1[i][j], "  ")
		}
		fmt.Println()
	}
	fmt.Println("Second array: ")
	for i = 0; i < limit; i++ {
		for j = 0; j < limit; j++ {
			fmt.Print(arr2[i][j], "  ")
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("Enter limit of the array")
	fmt.Scan(&limit)
	getArray()
	swapArray()
	dispArray()
}
