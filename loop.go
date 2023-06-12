package main

import "fmt"

func main() {

	//fmt.Println("hello world !")
	//fmt.Println("hello world 123!")

	for i := 1; i < 15; i++ {

		if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}

		if i == 14 {
			fmt.Println("fizzbuzz")
		}
	}
}
