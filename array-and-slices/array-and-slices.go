package main

import "fmt"

var g_groceries []string

func add_groccery(a string) {
	fmt.Println("Capacity", cap(g_groceries))
	g_groceries = append(g_groceries, a)
}

func list_groceries() {
	fmt.Println("Grocery list is as follow:")
	/*for i := 0; i < len(g_groceries); i++ {
		fmt.Println(g_groceries[i])
	}*/
	for _, data := range g_groceries {
		fmt.Println(data)
	}
}

func main() {
	add_groccery("Bread")
	add_groccery("Cucumbers")
	add_groccery("Salt and Vinegar Potato Chips")
	add_groccery("Salt and Vinegar Potato Chips")
	add_groccery("Salt and Vinegar Potato Chips")
	add_groccery("Salt and Vinegar Potato Chips")
	add_groccery("Salt and Vinegar Potato Chips")
	list_groceries()
}
