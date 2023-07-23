package main

import "fmt"

func main() {

	//--------------------Variables-------------------//
	fmt.Println("Hello Go")

	var name string = "Yasin"
	var name2 = "Yasin"
	name3 := "Yasin"

	fmt.Println(name, name2, name3)

	var number0 int = 0
	var number1 = 1
	number2 := 2

	fmt.Println(number0 + number1 + number2)

	var myFloat float32 = 32.2
	var myFloat2 = 32.3
	myFloat3 := 32.5

	var myTotalFloat float32 = myFloat + float32(myFloat2) + float32(myFloat3)

	fmt.Println(myTotalFloat)

	//--------------------Operators--------------------------------//
	/*

		A=10  B=20

		A+B -----> + =	 30
		A-B -----> - =	-10
		A*B -----> * =  200
		B/A -----> / =    2
		B%A (mod) ----> % = 0
		A++ ----> A = 11
		A-- ----> A =  9


		(A == B) is not true.
		(A != B) is true.
		(A > B) is not true.
		(A < B) is true.
		(A >= B) is not true.
		(A <= B) is true.



	*/
}
