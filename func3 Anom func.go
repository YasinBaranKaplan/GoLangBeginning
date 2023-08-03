package main

//Anonym Func

func main() {

	addAnomFunc := func(terms ...int) (numTerms int, sum int) {
		for _, term := range terms {
			sum += term
		}
		numTerms = len(terms)
		return
	}

	numTerms, sum := addAnomFunc(2, 7, 18, 100)
	println("Added : ", numTerms, " terms for a total of", sum)

	//this anom func can use just one time

}
