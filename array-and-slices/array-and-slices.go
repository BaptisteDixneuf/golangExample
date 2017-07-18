package main

import "fmt"

const g_cap int = 5 // Total capacity of our list
var g_groceries [g_cap]string
var g_len int = 0 //total number of grocery items in our current array

func add_groccery(a string) {
	if g_len < g_cap {
		g_groceries[g_len] = a
		g_len++
	} else {
		fmt.Println("Too many items, now we are done  for!!")
	}
}

func list_groceries() {
	fmt.Println("Grocery list is as follow:")
	for i := 0; i < g_len; i++ {
		fmt.Println(g_groceries[i])
	}
}

func main() {
	add_groccery("Bread")
	add_groccery("Cucumbers")
	add_groccery("Salt and Vinegar Potato Chips")
	list_groceries()
}
