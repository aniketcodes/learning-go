package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}

func (u User) GetStatus() bool {
	return u.Status
}

func (u *User) UpdateStatus() bool {
	u.Status = !u.Status
	return u.Status
}

func main() {
	aniket := User{"Aniket", 20, "aniket@gmail.com", true}
	fmt.Println(aniket)
	fmt.Println(aniket.GetStatus())
	fmt.Println(aniket.UpdateStatus())
	fmt.Println(aniket)
}
