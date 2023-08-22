package main

import (
	"io"
	"log"
	"os"
)

func main() {
	originalFile, err := os.Open("demo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()

	//yeni dosya oluşturma
	newFile, err := os.Create("demo_copy.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer newFile.Close()

	//byte'ları kaynaktan hedefe kopyala
	bytesWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes", bytesWritten)

	//dosya içeriğini işle
	//belleği diske boşalt

	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}

}
