package main

import "fmt"

func main() {
	//Slice

	array := [...]int{200, 300, 400}
	slice1 := array[:]
	fmt.Println(slice1)

	slice1[0] = 250
	fmt.Println(slice1)
	fmt.Println(array)

	colors := []string{"blue", "white", "green"}
	colors = append(colors, "black")
	fmt.Println(colors)

	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	fmt.Println(cap(colors))
	fmt.Println(len(colors))

}
