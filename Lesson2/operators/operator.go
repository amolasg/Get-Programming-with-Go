package main

import "fmt"

func main() {
	// assignment operator
	var weight = 149.0
	weight = weight * 0.3783
	weight *= 0.3783 //shortcut decleration of above declaration

	//increment operator
	var age = 41
	fmt.Println("1", age) // 41
	age = age + 1
	fmt.Println("2", age) // 42
	age += 1
	fmt.Println("3", age) // 43
	age++
	fmt.Println("4", age) //44

}
