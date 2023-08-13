package main

import "fmt"

func main() {

	fmt.Println("User oluşturma v1")

	user1 := &User{
		ID:       1,
		Fname:    "Yasin",
		Lname:    "Kaplan",
		UserName: "YasinKaplan",
		Age:      22,
		Pay: &Payment{
			Salary: 30000,
			Bonus:  7800,
		},
	}

	fmt.Println(user1)
	fmt.Println(user1.GetFullName(), user1.UserName, user1.GetPayment())

	fmt.Println("User oluşturma v2")
	user2 := NewUser()
	user2.Fname = "Mehmet"
	user2.Lname = "Delali"
	user2.UserName = "MUCİTMehmet"
	user2.Age = 41
	user2.Pay = &Payment{Salary: 64000, Bonus: 6500}

	fmt.Println(user2.GetFullName(), user2.GetUserName(), user2.GetPayment())

}

// kullanıcı yapısı
type User struct {
	ID       int
	Fname    string
	Lname    string
	UserName string
	Age      int
	Pay      *Payment
}

// Kullanıcının yapıcı metodu
func NewUser() *User {
	u := new(User)
	u.Pay = NewPayment()
	return u
}

// ödeme yapısı
type Payment struct {
	Salary float64
	Bonus  float64
}

// ödemenin yapıcı metodu
func NewPayment() *Payment {
	p := new(Payment)
	return p
}

// metotlar
func (u User) GetFullName() string {
	return u.Fname + " " + u.Lname
}

func (u *User) GetUserName() string {
	return u.UserName
}

func (u *User) GetPayment() float64 {
	pay := u.Pay.Salary + u.Pay.Bonus
	return pay
}
