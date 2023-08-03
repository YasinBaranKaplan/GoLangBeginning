package main

//Make simple func

func main() {
	sayHello()
	sayMessage("Hi")

	Pmessage := "Hi there "
	sayPMessage(&Pmessage)

	println(sum(12, 24))

}
func sayHello() {
	println("Hello")
}

func sayMessage(message string) {
	println(message)
}

func sayPMessage(Pmessage *string) {
	*Pmessage = "Hi Pointer message"
	println(*Pmessage)

}

func sum(x, y int) int {
	return x + y
}
