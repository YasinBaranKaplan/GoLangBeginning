package main

import "fmt"

func main() {

	//If

	x := 10
	y := 11

	if x < y {
		fmt.Println("Buyuktur")
	} else if x == y {
		fmt.Println("Esittir")
	} else {
		fmt.Println("Kucuktur")
	}

	//------------------------------------------------//

	foo := 1
	if foo == 1 {
		println("bar")
	}

	//------------------------------------------------//

	//if we use variable in  if case what happen ?

	if a := 10; a == 1 {
		println("true")
	} else {
		println("false")
	}
	//fmt.Println(a)

	// -->fmt.Println(a) we cant use like this out the if case, cause assigned in if case

}
