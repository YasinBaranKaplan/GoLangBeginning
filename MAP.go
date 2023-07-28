package main

import "fmt"

func main() {

	product := make(map[int]string)
	product[110] = "Cherry"
	product[111] = "Watermelon"
	product[112] = "Cucumber"
	product[113] = "Patato"

	fmt.Println(product)

	//product listesinde 111 anahtarli veriyi değiskene atama
	karpuz := product[111]
	fmt.Println("your product is", karpuz)

	//product listesinde 113 anahtarli veriyi silme
	delete(product, 111)
	//fmt.Println(product[111])//
	fmt.Println(product)

	product[111] = "Tomato"

	for k, v := range product {
		fmt.Printf("%v = %v\n ", k, v)
	}

	//product map nesnesinin elemanı adedince kapasiteli bir key listesi olusturma
	keys := make([]int, len(product))

	i := 0
	for k := range product {
		keys[i] = k
		i++
	}
	fmt.Println(keys)

	//key listesinde indeks degerlerine gore product nesnesinde bulunan urunleri yazdır
	for i := range keys {
		fmt.Println(product[keys[i]])
	}

}
