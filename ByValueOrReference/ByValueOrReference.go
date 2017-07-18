package main

import "fmt"

func reset_value(i int) {
	fmt.Println("In reset_value - i is:", i, "Located at : ", &i)
	i = 0
	fmt.Println("In reset_value - i is:", i, "Located at : ", &i)
}

func main() {
	j := 15
	fmt.Println("In main - j is:", j, "Located at : ", &j)
	reset_value(j)
	fmt.Println("In main - j is:", j, "Located at : ", &j)
}
