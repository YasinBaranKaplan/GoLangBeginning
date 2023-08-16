package main

//concurrency pararel çalışma değildir. Eş zamanlılıktır

import "time"

func main() {
	//basit bir gorouting kullanımı
	go xFunc() //buradaki işlem bitmeden alt satıra geçer ve bu işlemi çalıştırmaya devam eder (THREAD) --İşletim sistemi seviyesinde bir thread değildir goroutin
	time.Sleep(100 * time.Millisecond)
}

func xFunc() {
	for l := byte('a'); l <= byte('z'); l++ {
		println(string(l))
	}
}

//pararel programlama sıralama gözetmez çünkü farklı çekirdeklerde işlemler yapılır.
