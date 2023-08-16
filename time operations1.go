package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Unix time at the moment : %v\n", time.Now().Unix()) //unix zamanı gösterir.
	time.Sleep(10 * time.Second)                                    //uygulamayı 10 saniye uyutur.
	fmt.Printf("Unix time at the moment : %v\n", time.Now().Unix()) //10 saniye sonra ki unix zamanı verir.

}
