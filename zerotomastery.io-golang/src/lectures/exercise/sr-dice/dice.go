//--Summary:
//  Create a program that can perform dice rolls. The program should
//  report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must handle any number of dice, rolls, and sides
//
//--Notes:
//* Use packages from the standard library to complete the project

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func roll(sides int) int {
	return rand.Intn(sides) + 1
}

func main() {
	rand.Seed(time.Now().UnixNano())

	dice, sides := 2, 6
	rolls := 2

	for r := 1; r <= rolls; r++ {
		sum := 0

		for d := 1; d <= dice; d++ {
			rolled := roll(sides)
			sum += rolled
			fmt.Println("Roll #", r, "Die #", d, ":", rolled)
		}
		fmt.Println("Total rolled:", sum)

		switch sum := sum; {
		//  - "Snake eyes": when the total roll is 2, and total dice is 2
		case sum == 2 && dice == 2:
			fmt.Println("Snake eyes!")
		//  - "Lucky 7": when the total roll is 7
		case sum == 7:
			fmt.Println("Lucky 7!")
		//  - "Even": when the total roll is even
		case sum%2 == 0:
			fmt.Println("Even!")
		//  - "Odd": when the total roll is odd
		case sum%2 != 0:
			fmt.Println("Odd!")
		}
	}

}
