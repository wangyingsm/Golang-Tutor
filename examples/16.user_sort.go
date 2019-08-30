package main

import (
	"fmt"
	"sort"
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

type Users []*User

func (u Users) Len() int {
	return len(u)
}

func (u Users) Less(i, j int) bool {
	return u[i].LastLogin.Before(u[j].LastLogin)
}

func (u Users) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func (u Users) String() string {
	result := ""
	for _, v := range u {
		result = fmt.Sprintf("%s \n %v", result, v)
	}
	return result
}

func main() {
	users := Users{
		&User{
			"admin",
			"123456",
			UserInfo{
				Truename:  "John Doe",
				LastLogin: time.Now(),
			},
		},
		&User{
			"user",
			"654321",
			UserInfo{
				Truename:  "Alice Bob",
				LastLogin: time.Now().Add(-2 * time.Hour),
			},
		},
	}
	fmt.Println(users)
	sort.Sort(users)
	fmt.Println(users)
}
