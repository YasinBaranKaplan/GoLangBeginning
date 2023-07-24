package main

import "fmt"

func main() {
	foo := 3
	switch {
	case foo == 1:
		println("one")
	case foo == 2:
		println("two")
	case foo > 2:
		println("greater")
	}

	//----------------------------------------//

	var point float64
	fmt.Println("Enter number between one and five ")
	fmt.Scanf("%v", &point)
	//"%v" and %variable for scan

	switch {
	case point == 1:
		fmt.Println("Love")
	case point == 2:
		fmt.Println("Money")
	case point == 3:
		fmt.Println("Gamble")
	case point == 4:
		fmt.Println("Gun")
	case point == 5:
		fmt.Println("Death")
	}

}
