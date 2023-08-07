package main

import "fmt"

type Human struct {
	FName string
	LName string
	Age   int
}

//Varsayilan bos yapici metod
func NewHuman() *Human {
	h := new(Human)
	return h
}

func NewHuman2(firstname, lastname string, age int) *Human {
	h := new(Human)
	h.FName = firstname
	h.LName = lastname
	h.Age = age
	return h

}

func main() {

	//v1 : struct yapili birimi cagirma
	x := Human{FName: "Yasin Baran", LName: "Kaplan", Age: 22}
	fmt.Println(x.FName, x.LName, x.Age)

	//v2
	y := new(Human)
	y.Age = 23
	y.FName = "Bursa"
	y.LName = "Str"

	fmt.Println(y.FName, y.LName, y.Age)

	//v3 construct ile
	q := NewHuman()
	q.Age = 28
	q.FName = "Mehmet"
	q.LName = "Huso"

	fmt.Println(q.FName, q.LName, q.Age)

	w := NewHuman2("Demyr", "Dasdelen", 26)

	fmt.Println(w.FName, w.LName, w.Age)

}
