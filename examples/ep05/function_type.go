package main

import (
	"fmt"
	reg "regexp"
)

// 定义类型myFunc，接受字符串参数，返回布尔值和错误
// 这是所有下面校验输入的函数的签名，使用自定义类型是为了方便
type myFunc func(string) (bool, error)

// 校验电子邮件正确性，通过返回true，正则表达式格式错误则返回错误
func validEmail(email string) (bool, error) {
	r, err := reg.Compile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if err != nil {
		return false, err
	}
	return r.Match([]byte(email)), nil
}

// 校验手机号码正确性，通过返回true，正则表达式格式错误则返回错误
func validMobileNumber(pn string) (bool, error) {
	r, err := reg.Compile("^[+]*[(]{0,1}[0-9]{1,4}[)]{0,1}[-\\s\\./0-9]*$")
	if err != nil {
		return false, err
	}
	return r.Match([]byte(pn)), nil
}

// 校验用户输入是否满足格式校验，格式校验函数由参数f代入
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
	// 使用匿名函数检测IP地址的有效性
	fmt.Printf("IP address check result: %v", check(ip, func(ip string) (bool, error) {
		r, err := reg.Compile("^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$")
		if err != nil {
			return false, err
		}
		return r.Match([]byte(ip)), nil
	}))
}
