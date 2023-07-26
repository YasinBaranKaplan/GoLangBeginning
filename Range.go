package main

import "fmt"

func main() {

	a := [...]string{"x", "y", "z", "q"}

	for i := range a {
		fmt.Println("array item ", i, "is ", a[i])
	}
	// range herhangi bir "map","slice" veya "dizi" gibi formların değerlerinde dolaşır.

	//Map

	capital := map[string]string{"Turkey": "Ankara", "Japan": "Tokyo", "Canada": "Ottava", "France": "Paris"}
	for key := range capital {
		fmt.Println("Map item: Capital of ", key, "is ", capital[key])
	}

	//alternatif kullanım
	/*
		capital := map[string]string{"Turkey": "Ankara", "Japan": "Tokyo", "Canada": "Ottava", "France": "Paris"}
		for key, val := range capital {
			fmt.Println("Map item: Capital of ", key, "is ", val)
		}
	*/
}
