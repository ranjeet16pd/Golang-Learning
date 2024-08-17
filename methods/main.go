package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status string
	Age    string
}

func (u User) GetStatus() {
	fmt.Println("is user active :", u.Status)
	fmt.Println("email id is  :", u.Email)
}

func (u User) NewEmail() {
	u.Email = "rj@gmail.com"
	println(u.Email)

}

func main() {

	ranjeet := User{
		Name:   "Ranjeet",
		Email:  "ranjeet16pd@gmail.com",
		Status: "active",
		Age:    "22",
	}

	ranjeet.GetStatus()

	ranjeet.NewEmail()

	println(ranjeet.Email)

}
