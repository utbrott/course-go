//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - slice the assembly line so it contains only the two new parts
//  - print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

//* Create a function to print out the contents of the assembly line
func printSlice(line []Part) {
	fmt.Println("-- Current line --")
	for i := 0; i < len(line); i++ {
		part := line[i]
		fmt.Println(part)
	}
}

func main() {
	//* Using a slice, create an assembly line that contains type Part
	parts := make([]Part, 3)

	//  - Create an assembly line having any three parts
	parts[0] = "part 1"
	parts[1] = "part 2"
	parts[2] = "part 3"
	printSlice(parts)

	//  - Add two new parts to the line
	parts = append(parts, "part 4", "part 5")
	printSlice(parts)

	fmt.Println("processing parts...")

	//  - slice the assembly line so it contains only the two new parts
	parts = parts[3:]
	printSlice(parts)
}
