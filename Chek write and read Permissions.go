package main

import (
	"log"
	"os"
)

func main() {

	//yazma izinleri test
	file, err := os.OpenFile("demo.txt", os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Hata : Yazma izni reddedildi.")
		}
	}
	file.Close()

}
