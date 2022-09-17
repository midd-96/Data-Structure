package main

import (
	"fmt"
)

func main() {

	var str string
	var key int
	fmt.Println("Enter plain text")
	fmt.Scan(&str)

	fmt.Println("Enter key value")
	fmt.Scan(&key)
	cipherText(str, key)

}
func cipherText(ptext string, k int) {
	runes := []rune(ptext)

	var result []int

	for i := 0; i < len(runes); i++ {

		result = append(result, int(runes[i])+k)
	}

	fmt.Printf("The cipher text : %c", result)
}
