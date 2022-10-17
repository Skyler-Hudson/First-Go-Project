package main

import "fmt"

func main() {
	userConversation()
	interactiveCalculator()
}

func userConversation() {
	// Prompt the user
	fmt.Println("Hello, how are you today?")
	// Declare Variables
	var response string
	// User Input
	fmt.Scan(&response)

	// Prompt the user
	fmt.Println("What are your 2 favorite colors?")
	// Declare Variables
	var color string
	var color2 string
	// User Input
	fmt.Scan(&color)
	fmt.Scan(&color2)

	// Prompt the user
	fmt.Println("Cool, mine is Black and Blue. What are your 3 favorite hobbies?")
	// Declare Variables
	var hobby1, hobby2, hobby3 string
	fmt.Scan(&hobby1, &hobby2, &hobby3)
}

func interactiveCalculator() {
	// Addition
	// Prompt the User
	fmt.Println("Enter 2 numbers in to be added:")
	// Declare Variables
	var num1, num2 float32
	// Assign the values
	fmt.Scan(&num1, &num2)
	// Doing the math
	var answer = num1 + num2
	// Print the result
	fmt.Println(answer)

	// Subtraction
	// Prompt the user
	fmt.Println("Enter 2 numbers to be subtracted:")
	// Declare Variables
	var num3, num4 float32
	// Assign the Variables
	fmt.Scan(&num3, &num4)
	// Do the math
	var answer2 float32 = num3 - num4
	// Print
	fmt.Println(answer2)
}
