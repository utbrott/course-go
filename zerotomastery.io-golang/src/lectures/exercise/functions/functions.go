//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func greetUser(user string) {
	fmt.Println("Hello", user, "!")
}

func returnMessage(message string) string {
	return message
}

func addThree(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

func returnNumber() int {
	return 1
}

func returnTwo() (int, int) {
	return 2, 3
}

func main() {
	greetUser("Mark")
	fmt.Println(returnMessage("A second hello"))

	ret1, ret2 := returnTwo()
	sumOfThree := addThree(returnNumber(), ret1, ret2)
	fmt.Println("The sum of", returnNumber(), ret1, ret2, "equals", sumOfThree)
}
