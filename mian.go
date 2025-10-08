package main

import (
	"fmt"
)

func main() {
	var num1, num2 float 
	var operator string
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)		
	fmt.Print("Enter second number: ")		
	fmt.Scanln(&num2)
	fmt.Print("Enter operator: ")
	fmt.Scanln(&operator)
	switch operator {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		fmt.Println(num1 / num2)
	default:
		fmt.Println("Invalid operator")
	}
}

// Simple calculator program in Go