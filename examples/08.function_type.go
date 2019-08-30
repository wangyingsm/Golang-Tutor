package main

import (
	"fmt"
	reg "regexp"
)

type myFunc func(string) (bool, error)

func validEmail(email string) (bool, error) {
	r, err := reg.Compile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if err != nil {
		return false, err
	}
	return r.Match([]byte(email)), nil
}

func validMobileNumber(pn string) (bool, error) {
	r, err := reg.Compile("^[+]*[(]{0,1}[0-9]{1,4}[)]{0,1}[-\\s\\./0-9]*$")
	if err != nil {
		return false, err
	}
	return r.Match([]byte(pn)), nil
}

func check(s string, f myFunc) bool {
	valid, err := f(s)
	if err != nil {
		fmt.Println(err)
	}
	return valid
}

func main() {
	email := "myemail@example.com"
	fmt.Printf("email check result: %v\n", check(email, validEmail))
	pn := "13800138000"
	fmt.Printf("phone number check result: %v\n", check(pn, validMobileNumber))
	ip := "127.0.0.1"
	fmt.Printf("IP address check result: %v", check(ip, func(ip string) (bool, error) {
		r, err := reg.Compile("^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$")
		if err != nil {
			return false, err
		}
		return r.Match([]byte(ip)), nil
	}))
}
