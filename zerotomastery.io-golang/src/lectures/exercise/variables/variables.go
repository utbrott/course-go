//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//x Store your favorite color in a variable using the `var` keyword
//x Store your birth year and age (in years) in two variables using
//  compound assignment
//x Store your first & last initials in two variables using block assignment
//x Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	var favouriteColour = "blue"
	birthYear, age := 1998, 23
	var (
		firstInitial = "M"
		lastInitial  = "R"
	)
	var ageInDays int
	ageInDays = age * 365

	fmt.Println("Favourite colour:", favouriteColour)
	fmt.Println("Birth year:", birthYear, "Age:", age)
	fmt.Println("Initials", firstInitial, lastInitial)
	fmt.Println("Age in days:", ageInDays)
}
