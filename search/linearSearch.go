package main

import "fmt"

func linearsearch(datalist []int, key int) bool {
	for _, item := range datalist {
		if item == key {
			return true
		}
	}
	return false
}

func main() {
	var num int
	items := []int{95, 78, 46, 58, 45, 86, 99, 251, 320}
	fmt.Println("Numbers are: ", items)
	fmt.Println("Enter number to search")
	fmt.Scan(&num)
	status := (linearsearch(items, num))
	if status == true {
		fmt.Println("The number exist")
	} else {
		fmt.Println("Number not exist")
	}
}
