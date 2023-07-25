package main

import (
	"fmt"
	"time"
)

func main() {
	//FOR

	for i := 0; i <= 5; i++ {
		fmt.Println("Deger = ", i)
	}

	for j := 0; j <= 5; {
		fmt.Println("2. Deger = ", j)
		j++

	}

	//FOR-WHILE method

	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
		time.Sleep(300 * time.Millisecond)
	}

}
