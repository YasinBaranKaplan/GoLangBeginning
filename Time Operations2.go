package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2023, time.April, 10, 20, 0, 0, 0, time.UTC)
	fmt.Printf("Çalisiyor: %s\n", t)

	fmt.Println("***********************")

	now := time.Now()
	fmt.Printf("Mevcut zaman: %s\n", now)

	fmt.Println("***********************")

	fmt.Println("Ay: ", now.Month())
	fmt.Println("Gün: ", now.Day())
	fmt.Println("Haftanin günü: ", now.Weekday())

	fmt.Println("***********************")

	tomorrow := now.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v, %v, %v, %v\n", tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year())

	fmt.Println("***********************")

	longFormat := "Monday ,January ,2 ,2023"
	fmt.Println("Tomorrow is ", tomorrow.Format(longFormat))

	shortFormat := "1/2=2023"
	fmt.Println("TMRW is ", tomorrow.Format(shortFormat)) //formatı belirle ve o formata göre yazdırma

}
