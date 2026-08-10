package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Print("Enter a positive integer n: ")
	_, err := fmt.Scan(&n)
	if err != nil || n < 1 {
		fmt.Println("Invalid input. Please enter a positive integer.")
		return
	}

	var choice string
	fmt.Print("Choose operation ('sum' or 'product'): ")
	fmt.Scan(&choice)
	choice = strings.ToLower(choice)

	if choice == "sum" {
		sum := 0
		for i := 1; i <= n; i++ {
			sum += i
		}
		fmt.Printf("The sum of numbers from 1 to %d is: %d\n", n, sum)

	} else if choice == "product" {
		// Using int64 to accommodate larger results, as factorials grow extremely fast
		var product int64 = 1
		for i := 1; i <= n; i++ {
			product *= int64(i)
		}
		fmt.Printf("The product of numbers from 1 to %d is: %d\n", n, product)

	} else {
		fmt.Println("Invalid choice. Please run the program again and enter 'sum' or 'product'.")
	}
}

