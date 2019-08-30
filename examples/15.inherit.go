package main

import (
	"fmt"
	"time"
)

type UserInfo struct {
	Truename  string
	Nickname  string
	Email     string
	LastLogin time.Time
}

type User struct {
	Username string
	password string
	UserInfo
}

func (ui *UserInfo) UpdateLoginTime() {
	ui.LastLogin = time.Now()
}

func (u *User) CheckUser() bool {
	return u.Username == "admin" && u.password == "123456"
}

func main() {
	user := User{
		"admin",
		"123456",
		UserInfo{
			Truename:  "John Doe",
			Nickname:  "Johnny",
			Email:     "johndoe@hotmail.com",
			LastLogin: time.Now(),
		},
	}
	fmt.Println(user)
	time.Sleep(2 * time.Second)
	fmt.Printf("%s Login state: %v\n", user.Truename, user.CheckUser())
	user.UpdateLoginTime()
	fmt.Println(user)
}
