package main

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v2" // 进行YAML格式转换的第三方包
)

// Param 测试用结构体
type Param struct {
	// 下面四个字段都定义了JSON格式和YAML格式的标注，注意名称首字母大写
	Name string `json:"name" yaml:"name"`
	// Mobile字段为零值时，在JSON中忽略
	Mobile string `json:"mobile,omitempty" yaml:"mobile"`
	State  bool   `json:"state" yaml:"state"`
	Age    int    `json:"age" yaml:"age"`
}

func main() {
	// 初始化Param结构体
	p := Param{
		Name:   "John Doe",
		Mobile: "13800138000",
		State:  true,
		Age:    30,
	}
	fmt.Printf("%v\n", p)
	// 调用getOld函数
	getOld(p)
	// 序列化成JSON字节数组
	j, err := json.Marshal(p)
	// 发生错误，打印错误内容，非0退出程序
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// 控制台打印JSON字符串
	fmt.Println(string(j))
	// 设置Mobile字段为零值，注意这里不能写成 p.Mobile = nil，字符串没有nil值
	p.Mobile = ""
	// 再次序列化JSON字节数组，使用下划线变量忽略错误
	j, _ = json.Marshal(p)
	// 控制台打印JSON字符串，观察与前面的输出区别
	fmt.Println(string(j))
	// 序列化YAML字节数组，使用下划线变量忽略错误
	y, _ := yaml.Marshal(p)
	// 控制台打印YAML字符串
	fmt.Println(string(y))
}

// getOld 将结构体中的Age字段增加1，打印生日祝福语
// 由于传递的参数是结构体本身，根据永远传值原则，代入的原始结构体不变
func getOld(p Param) {
	p.Age++
	fmt.Printf("Happy birthday %s for his/her %d years old.\n", p.Name, p.Age)
}
