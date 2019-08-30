package main

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Param struct {
	Name   string `json:"name" yaml:"name"`
	Mobile string `json:"mobile,omitempty" yaml:"mobile"`
	State  bool   `json:"state" yaml:"state"`
	Age    int    `json:"age" yaml:"age"`
}

func main() {
	p := Param{
		Name:   "John Doe",
		Mobile: "13800138000",
		State:  true,
		Age:    30,
	}
	fmt.Printf("%v\n", p)
	getOld(p)
	j, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(j))
	p.Mobile = ""
	j, _ = json.Marshal(p)
	fmt.Println(string(j))
	y, _ := yaml.Marshal(p)
	fmt.Println(string(y))
}

func getOld(p Param) {
	p.Age++
	fmt.Printf("Happy birthday %s for his/her %d years old.\n", p.Name, p.Age)
}
