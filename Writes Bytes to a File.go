package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"WRİTEONLY.txt",
		os.O_WRONLY|os.O_CREATE,
		0666)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := []byte("Bu dosyaya bu yaziyi yazdik\n")
	byteWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes\n", byteWritten)

}
