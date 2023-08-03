package main

import "fmt"

//func : return multiple results
//variadic func

func main() {

	q, z := swap("Yasin", "Baran")
	fmt.Println(q, z)

	sayHello("Hello ", "new ", "go ", "developer ")

	fmt.Println(add(1, 2, 3, 4, 5, 6, 7, 8, 9))

	numLen, multiple := multip(7, 8, 9)
	fmt.Println("Added : ", numLen, "result of multiplication = ", multiple)

}

func swap(x, y string) (string, string) {
	return y, x

	//"xyz","abc" ---> "abc","xyz"
}

func sayHello(messages ...string) { //... is how to use
	for _, message := range messages {
		fmt.Printf(message)
	}
}

func add(terms ...int) int {
	result := 0
	for _, term := range terms {
		result += term
	}
	return result
}

func multip(numbers ...int) (int, int) {
	result := 1
	for _, num := range numbers {
		result *= num
	}
	return len(numbers), result
}
