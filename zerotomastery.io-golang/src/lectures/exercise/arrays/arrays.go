//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

//  - Products must include the price and the name
type ListItem struct {
	name  string
	price int
}

//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items

func printInfo(list [4]ListItem) {
	var cost, totalItems int

	for i := 0; i < len(list); i++ {
		item := list[i]
		cost += item.price

		if item.name != "" {
			totalItems++
		}
	}

	fmt.Println("Last item:", list[totalItems-1])
	fmt.Println("Total items:", totalItems)
	fmt.Println("Total cost:", cost)
}

func main() {
	//* Using an array, create a shopping list with enough room
	//  for 4 products
	//  - Products must include the price and the name
	shoppingList := [4]ListItem{
		{"Bread", 2},
		{"Milk", 4},
		{"Banana", 1},
	}

	printInfo(shoppingList)

	//* Add a fourth product to the list and print out the
	//  information again
	shoppingList[3] = ListItem{"Salad", 2}
	printInfo(shoppingList)
}
