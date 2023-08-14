package main

//BU kod Udemy Cihan Özhan Golang kursundan alntılanmıştır.
import (
	"fmt"
	"strconv"
)

func CarExecute(c Carface) {
	fmt.Print("\n")
	fmt.Println("Araç Bilgisi : \n" + c.Information())
	fmt.Print("\n")

	msg := ""

	isRun := c.Run()
	if isRun {
		msg = "Calismiyor"
	} else {
		msg = "Calisiyor"
	}

	fmt.Println("Arac " + msg + ".")

	isStop := c.Stop()
	if isStop {
		msg = "Durdu"
	} else {
		msg = "Duramadi fren tutmuyor"
	}
	fmt.Println("Arac " + msg + ".")

}

func main() {
	ferr := new(Ferrari)
	ferr.Brand = "Ferrari"
	ferr.Model = "F50"
	ferr.Color = "Red"
	ferr.Speed = 340
	ferr.Price = 1
	ferr.Special = true

	fmt.Println(ferr.Information())

	merc := new(Mercedes)
	merc.Brand = "Mercedes"
	merc.Model = "CLX"
	merc.Color = "Black"
	merc.Speed = 290
	merc.Price = 2
	fmt.Println(merc.Information())

	//CarExecute
	CarExecute(ferr)
	CarExecute(merc)

}

// base strcuts
type Car struct {
	Brand string
	Model string
	Color string
	Speed int
	Price float64
}

//InterFace

type Carface interface {
	Run() bool
	Stop() bool
	Information() string
}

type SpecialProduction struct {
	Special bool
}

type Ferrari struct {
	Car //composition
	SpecialProduction
}

func (_ Ferrari) Run() bool {
	return true
}

func (_ Ferrari) Stop() bool {
	return true
}

func (x *Ferrari) Information() string {
	ret := "\t" + x.Brand + " " + x.Model + "\n" + "\t" + "Color: " + x.Color + "\t" + "Speed: " + strconv.Itoa(x.Speed) + "\n" + "\t" + "Price: " + strconv.Itoa(int(x.Price)) + "Million"
	add := "Yes"
	if x.Special {
		ret += "\n" + "\t" + "Special : " + add
	}
	return ret
}

// Lambo
type Lambo struct {
	Car //composition
	SpecialProduction
}

func (_ Lambo) Run() bool {
	return true
}

func (_ Lambo) Stop() bool {
	return true
}

func (x *Lambo) Information() string {
	ret := "\t" + x.Brand + " " + x.Model + "\n" + "\t" + "Color: " + x.Color + "\t" + "Speed: " + strconv.Itoa(x.Speed) + "\n" + "\t" + "Price: " + strconv.Itoa(int(x.Price)) + "Million"
	add := "Yes"
	if x.Special {
		ret += "\n" + "\t" + "Special : " + add
	}
	return ret
}

// Mercedes
type Mercedes struct {
	Car //composition
	SpecialProduction
}

func (_ Mercedes) Run() bool {
	return true
}

func (_ Mercedes) Stop() bool {
	return true
}

func (x *Mercedes) Information() string {
	ret := "\t" + x.Brand + " " + x.Model + "\n" + "\t" + "Color: " + x.Color + "\t" + "Speed: " + strconv.Itoa(x.Speed) + "\n" + "\t" + "Price: " + strconv.Itoa(int(x.Price)) + "Million"

	return ret
}
