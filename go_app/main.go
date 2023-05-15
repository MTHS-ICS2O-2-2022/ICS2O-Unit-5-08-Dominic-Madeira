// Created by: Dominic M.
// Created on: April 2023
//
// This program divides two numbers.
package main

import (
	"fmt"
)

func main() {

	var numberA int
	var numberB int
	var counter int

	// input
	fmt.Println("This program divides two numbers.")
	fmt.Println()
	fmt.Print("Enter your first number: ")
	fmt.Println()
	fmt.Scanln(&numberA)
	fmt.Println()
	fmt.Print("Enter your second number: ")
	fmt.Println()
	fmt.Scanln(&numberB)
	fmt.Println()

	// process
	remainder := numberA
	for remainder >= numberB {
		remainder -= numberB
		counter++
	}

	if remainder > 0 {
		fmt.Print(numberA, " ÷ ", numberB, " = ", counter, " with a remainder of ", remainder, ".")
	} else {
		fmt.Print(numberA, " ÷ ", numberB, " = ", counter)
	}

	fmt.Println()
	fmt.Print("\nDone.")
}
