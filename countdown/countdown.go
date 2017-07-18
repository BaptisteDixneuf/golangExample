package main

import "fmt"
import "time"

func main() {

	for i:= 10; i> 0; i--{
		fmt.Println(i)
		time.Sleep(time.Second)
		i=i-1
	}
	fmt.Println("Happy New Year !")
}
