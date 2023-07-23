package main

import (
	"fmt"
	"strconv"
)

func main() {
	var myString = "1"
	var x = 10
	var f float32 = 2.0

	fmt.Println(myString, x, f)

	//CONVERT (STRİNG TO İNT)
	number, _ := strconv.Atoi(myString)
	// ( _ ) is empty identifer.
	fmt.Println(number)

	result := number + 2

	fmt.Println(result)

	fmt.Println("Result" + strconv.Itoa(x))

	//Atoi string to int ------ Itoa int to string

}
