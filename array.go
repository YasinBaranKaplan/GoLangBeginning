package main

import "fmt"

func main() {
	//array of number

	myArray := [3]int{}

	myArray[0] = 55
	myArray[1] = 78
	myArray[2] = 100

	fmt.Println(myArray)

	//array of colour
	var colors = [4]string{"blue", "red", "green", "brown"}
	fmt.Println(colors)
	fmt.Println(colors[1])
	fmt.Println(len(colors))

	//---------------------------------//
	var cars = [3]string{"BMW", "Tesla", "TOGG"}

	i := 0

	for i <= len(cars)-1 {
		fmt.Println(cars[i])
		i++
	}

}
