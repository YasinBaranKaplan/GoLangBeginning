package main

import (
	"fmt"
	"log"
	"os"
)

var (
	fileInfo *os.FileInfo
	err      error
)

func main() {
	//dosya bilgisini döndürür
	//eger dosya yoksa hata döner.

	fileInfo, err := os.Stat("demo.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File name : ", fileInfo.Name())            //dosya adı
	fmt.Println("Size in bytes : ", fileInfo.Size())        //dosya boyutu
	fmt.Println("Permissions : ", fileInfo.Mode())          //dosya kullanabilirliği
	fmt.Println("Last Modified : ", fileInfo.ModTime())     //dosya değiştirilme tarihi
	fmt.Println("Is Dictionary : ", fileInfo.IsDir())       //dosya klasör mü file mi ?
	fmt.Println("System Interface Type : ", fileInfo.Sys()) //dosya ile ilgili sistem bilgisi
	fmt.Println("System Info : %+v\n\n", fileInfo.Sys())    //dosya ile ilgili sistem bilgisi 2

}
