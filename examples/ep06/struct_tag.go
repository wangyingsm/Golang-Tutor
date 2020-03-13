package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
)

// User 例子用结构体
type User struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	RegTime     time.Time `json:"-" format:"yyyy年MM月dd日"`
	RegTimeJSON string    `json:"regTime"`
}

// FormatTime 将User结构体中的时间字段按照自定义格式标记转换并记录在另一个输出的JSON字段中
// 此处使用了方法，接收的结构体参数为指针，原因是需要修改结构体的内容
func (u *User) FormatTime() bool {
	// 使用反射获得结构体类型
	t := reflect.TypeOf(*u)
	// 查找结构体的相关字段定义
	field, ok := t.FieldByName("RegTime")
	if !ok {
		return false
	}
	// 获得format标注字符串
	format, ok := field.Tag.Lookup("format")
	if !ok {
		return false
	}
	// 将yyyy MM dd这样的定义转换成Go的时间格式定义
	format = strings.ReplaceAll(format, "yyyy", "2006")
	format = strings.ReplaceAll(format, "MM", "01")
	format = strings.ReplaceAll(format, "dd", "02")
	// 格式化时间字符串并设置在RegTimeJSON字段中，该字段可以输出成JSON字段内容
	u.RegTimeJSON = u.RegTime.Format(format)
	return true
}

// MarshalJSON 实现JSON Marshaller接口，自定义JSON序列化格式
func (u User) MarshalJSON() ([]byte, error) {
	// 首先格式化时间字段
	u.FormatTime()
	// 使用bytes.Buffer产生最终JSON字符串
	var buf bytes.Buffer
	// 先写一个左大括号
	buf.WriteString("{")
	// 使用反射获取结构体的值
	value := reflect.ValueOf(u)
	// 用一个slice记录每个字段转换结果
	fields := []string{}
	// 遍历结构体每个字段，转换成JSON KV格式字符串
	for i := 0; i < value.NumField(); i++ {
		// 获得该字段的通用值
		val := value.Field(i)
		// 获得该字段的类型
		t := value.Type().Field(i)
		// 获得该字段json标注字符串内容
		key, ok := t.Tag.Lookup("json")
		// json标注为-表示忽略输出，跳过
		if !ok || key == "-" {
			continue
		}
		// 调用marshalField函数获得转换后的JSON KV字符串
		fields = append(fields, marshalField(key, val))
	}
	// 再次遍历slice，将每个字段的KV字符串写入buffer中
	for i, f := range fields {
		buf.WriteString(f)
		// 最后一个JSON字段末尾无逗号
		if i < len(fields)-1 {
			buf.WriteString(",")
		}
	}
	// 最后写入右大括号
	buf.WriteString("}")
	// 返回[]byte，无错误发生
	return buf.Bytes(), nil
}

// marshalField 将每个字段输出成JSON KV格式
func marshalField(key string, val reflect.Value) string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("\"%s\":", key))
	// 反射类型进行对比，此处仅处理了标准整型和字符串类型
	switch val.Kind() {
	case reflect.Int:
		buf.WriteString(fmt.Sprintf("%d", val.Int()))
	case reflect.String:
		buf.WriteString(fmt.Sprintf("\"%s\"", val.String()))
	}
	return buf.String()
}

func main() {
	// 初始化一个User结构体
	u := User{
		ID:      1,
		Name:    "John Doe",
		RegTime: time.Now(),
	}
	// JSON序列化
	j, _ := json.Marshal(&u)
	// 打印JSON字符串
	fmt.Println(string(j))
}
