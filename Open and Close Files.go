package main

import (
	"log"
	"os"
)

func main() {

	//İkinci parametre : dosya açılış amacını ayarlar
	//üçüncü parametre : dosya izinlerini belirler
	file, err := os.OpenFile("demo.txt", os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	/*ikinci parametre tipleri :
	os.O_RDONLY sadece okuma
	os.O_WRONLY sadece yazma
	os.O_RDWR   okuma ve yazma yapılabilir
	os.O_APPEND dosyanın sonuna ekle
	os.O_CREATE dosya yoksa oluştur
	os.O_TRUNC  açılırken dosyayı kes

	bu ayarlar birden fazla olarak da kullanılabilir

	-> os.O_CREATE|os.O_APPEND  gibi
	*/
}
