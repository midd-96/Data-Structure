package main

import (
	"fmt"
	"strings"
)

func main() {
	var str, subStr string

	fmt.Println("Enter the string")
	fmt.Scan(&str)
	fmt.Println("Enter the sub-string")
	fmt.Scan(&subStr)

	res := checkContain(str, subStr)

	if res == true {
		fmt.Printf("String '%s' contains sub-string '%s'", str, subStr)
	} else {
		fmt.Printf("String '%s' does not contains substring '%s'", str, subStr)
	}
}
func checkContain(a, b string) bool {

	if strings.Contains(a, b) == true {
		return true
	} else {
		return false
	}
}
