package main

import "fmt"

func main() {

	const lightSpeed = 299792 //km/s
	var distance = 56000000   //km

	fmt.Println(distance/lightSpeed, "seconds")
	distance = 401000000
	fmt.Println(distance/lightSpeed, "seconds")

	// how much time will take to reach mars from earth

	const hoursPerDay = 24
	var speed = 100800
	var dist = 96300000

	fmt.Println(dist/speed/hoursPerDay, "days")
}
