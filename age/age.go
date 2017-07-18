package main

import "fmt"

func main() {

	age := 20

	if age < 13 {
		fmt.Println("Wow, you're young!")
	} else if age < 20 {
		fmt.Println("You're teenager")
	} else if age < 30 {
		fmt.Println("You'are in your tweenties")
	} else if age < 40 {
		fmt.Println("You'are in your thirties")
	} else if age < 50 {
		fmt.Println("You'are getting there! ")
	} else {
		fmt.Println("Over the hill !")
	}
}
