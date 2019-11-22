package main

import "fmt"

func main() {
	//Testing fmt package
	age := 26
	name := "amol"
	add := "Pune"

	fmt.Printf("My age is %v and name is %v city is %v\n", age, name, add)
	fmt.Println("My age is ", age+1, "In March 2020")
	fmt.Println("My Current age", age)

	// Printf can help you align text,specify a width as part of verb

	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "ABC", 100)
}
