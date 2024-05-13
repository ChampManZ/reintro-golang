package main

import "fmt"

type User struct {
	Username string
	Fullname string
	Email    string
	UID      int
}

func main() {
	username := "johndoe"
	fullname := "John Doe"
	email := "johndoe@test.com"
	uid := 1

	u1 := User{
		Username: username,
		Fullname: fullname,
		Email:    email,
		UID:      uid,
	}

	fmt.Println(u1)
	fmt.Println(u1.Username)
	fmt.Println(u1.Fullname)
	fmt.Println(u1.Email)
	fmt.Println(u1.UID)
}
