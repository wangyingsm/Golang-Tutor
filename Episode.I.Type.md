# 类型、变量和常量

## Go的类型

Go是一种静态类型(Static Type)的语言，即所有的类型都会在编译期确定，除非覆盖了变量的声明，否则类型不允许在运行期改变。同时Go还是一门类型安全(Type Safe)的语言，意即Go在编译期就会对类型进行检查，任何类型错误都会在编译过程中报告，不会在运行期因为类型错误产生未定义的行为(Undefined Behavior)。

常用计算机语言类型系统比较：

| 语言 | 静态类型 | 类型安全 |
|-|-|-|
| C | 是 | 有限度的安全 |
| C++ | 是 | 有限度的安全 |
| Python | 否 | 是 |
| JavaScript | 否 | 否 |
| Java | 是 | 是 |
| Go | 是 | 是 |
| Rust | 是 | 是 |

### 内建类型

- 整数类型
    - int: 标准整数（通常为32bits，4bytes）[-2147483648, 2147483647]
    - int8: 8位整数（8bits，1byte）[-128, 127]
    - int16: 16位整数（16bits，2bytes）[-32768, 32767]
    - rune, int32: 32位整数（32bits，4bytes）[-2147483648, 2147483647]
    - int64: 64位整数（64bits，8bytes）[-9223372036854775808, 9223372036854775807]
- 无符号整数类型
    - uint: 标准无符号整数（通常为32bits，4bytes）[0, 4294967295]
    - byte, uint8: 8位无符号整数（8bits，1bytes）[0, 255]
    - uint16: 16位无符号整数（16bits，2bytes）[0, 65535]
    - uint32: 32位无符号整数（32bits，4bytes）[0, 4294967295]
    - uint64: 64位无符号整数（64bits，8bytes）[0, 18446744073709551615]
- 浮点数类型
    - float32: 32位浮点数（32bits，4bytes）约[1.4e-45, 3.4e38]
    - float64: 64位浮点数（64bits，8bytes）约[4.9e-324, 1.8e308]
- byte和rune
    - byte代表内存中的一个字节，等同于uint8
    - rune代表UTF8编码的一个字符，等同于int32
- bool布尔类型
    - true
    - false
- string字符串（[详见Ep.III](Episode.III.String.md)）
- 复数类型
    - complex: 标准复数（通常为64bits, 8bytes)
    - complex64: 64位复数（64bits, 8bytes)
    - complex128: 128位复数（128bits, 16bytes)
- func函数类型（[详见Ep.V](Episode.V.Function.md)）
- struct结构体类型（[详见Ep.VI](Episode.VI.Struct.md)）
- `*` 指针类型
- interface接口类型（[详见Ep.X](Episode.X.Interface.md)）
- error错误类型
- chan通道类型（[详见Ep.XV](Episode.XV.Channel.md)）
- uintptr通用指针类型（将不会在本系列中讨论）

> 关于byte等同于uint8类型比较容易理解，一个字节即为一个无符号的8位整数。对于rune定义为int32容易使人困惑，为何一个UTF8字符不定义为一个32位的整数？事实上，int32的范围足以包含所有UTF8字符集的定义了，因此，为了更方便的处理与UTF8字符相关的运算，定义rune为int32比uint32更加灵活。

### 自定义类型

使用关键字type可以将内建类型定义为自定义类型，后续即可使用别名指代原始的类型。

[例子：自定义类型](examples/ep01/custom_type.go)

## 变量

变量指代虚拟内存中的一段空间，该空间用于存储一个程序运行过程中需要使用的值。变量使用名称代表该段空间。

### 变量声明

var关键字用于声明一个变量，如：

```go
var aVariable int
var bVariable bool
```

第一行声明了一个标准整型变量，名称是aVariable，该名称指代一段虚拟内存中的区域，该区域的长度在通常的计算机上为4个字节。第二行声明了一个字符串变量，名称是bVariable，该名称指代一段虚拟内存中的区域，该区域的长度位1个字节。

### 变量初始化

可以使用与变量赋值的方法在变量声明中对变量进行初始化，如：

```go
var aVariable = 123
var bVariable = false
```

编译器会自动判断出初始化变量的类型，并进行变量声明，这种情况下无需对变量类型进行显式声明。相反，下面的语句会让Go编译器提出一个警告：

```go
// compiler warning
var aVariable int = 123
var bVariable bool = false
```

同一组的类型会被推断为其标准类型，如整数类型为标准整数int，如果需要使用非标准类型，显式类型声明则为必须：

```go
var aVariable int16 = 123
var bVariable float32 = 3.14
```

所有声明了的变量都会初始化，如果没有进行显式的初始化，则编译器会自动将该变量声明为其类型的零值。不存在没有初始化的变量。

各类型零值：

| 类型 | 对应零值 |
|-|-|
| 整数类型 | 0 |
| 无符号整数类型 | 0 |
| 浮点数类型 | 0.0 |
| 布尔类型 | false |
| 字符串 | "" |
| 复数类型 | complex(0.0, 0.0) |
| 结构体类型 | 所有字段都为零值的结构体 |
| 指针类型 | nil |
| 函数类型 | nil |
| 接口类型 | 空接口 |
| 错误类型 | nil |
| 通道类型 | nil |

> Go语言中的空值nil与其他语言中的空值有比较大的区别，最明显的差别就是应用范围，Go明显刻意限制了空值引用的使用范围。关于空值引用带来的问题，可参考[Tony Horae: Null References: The Billion Dollar Mistake](https://www.infoq.com/presentations/Null-References-The-Billion-Dollar-Mistake-Tony-Hoare/)。尤其需要注意的是，很多Go的组合类型如字符串，切片，字典，结构体等，其零值均不是空值nil。

[例子：变量声明和初始化](examples/ep01/var_definition.go)

### 短变量声明

当定义一个非包所属变量时，可以采用一种简洁的语法声明变量。

```go
aVariable := 123
bVariable := false
```

注意不要混淆了短变量声明语法和变量赋值语法，:=符号代表的含义包括两重，首先声明一个变量，然后将符号后面的值初始化到变量中。而赋值语法=仅仅是将等号右边的值赋予一个已经声明过的变量中。因此 `a := 0`等同于 `var a = 0`，但是不等同于 `a = 0`。

> 注意在var声明加初始化语法中，类型是可省的，如果显式指定编译器会给出一个警告。但是在短变量声明语法中，类型是不允许存在的，因此 `a int := 0` 会产生一个编译错误。

[例子：短变量声明](examples/ep01/short_var_definition.go)

## 类型转换

Go语言中无法对类型自动进行转换，哪怕是源类型和目标类型是兼容的也不允许，因此类型转换必须显式的进行。例如下面的语句是错误的:

```go
a := 0  // a is a type int variable
var b int64 = a // compile error
```

兼容的类型之间可以直接进行转换，语法是使用目标类型名称将需要转换的值小括号括起来，如果你熟悉python语法，这对你来说将会一点都不陌生。例如上述语句应该写为：

```go
a := 0
var b int64 = int64(a)  // ok
```

如果目标类型可容纳的范围比原值要小，那么这种类型转换将会产生溢出，其处理方式与其他溢出一致，原值将会被截断。例如：

```go
a := 100000 // a is larger than int16, which is 11000101010001101 in binary
var b int16 = int16(b)  // b is 1000101010001101 in binary, which is -30067 in decimal
```

> 任何不兼容的类型之间都无法进行直接转换，通常需要自己定义函数进行处理，Go标准库中也提供了对于整数和字符串转换的基本方法。需要注意的一点是，对于C语言、python、nodejs的开发人员来说，整数和布尔类型是一致或可以自动转换的，但是在Go当中，这是不被允许的。

## 变量作用域

Go中的变量作用范围被限定在最小代码块中，代码块有两个层级，代码包（package）和代码块（使用花括号括起来的代码部分）。无论是函数，方法，条件分支if，循环for等结构，都是使用花括号将代码部分封装的。因此任何声明在代码块中的变量仅在该代码块中有效。定义在代码块之外的变量是对于整个包（详见[EP.XII 包](Episode.XII.Package.md)）有效的，例如：

```go
package example

import "fmt"

var packageVar = "package string"   // a package var is available inside package example

// function f is also available inside package example
func f() {
    funcVar := 123  // funcVar is available inside function f
    for i:=0; i<100; i++ {
        // loop variable i is available inside for structure
        fmt.Println(i)
    }   // i is no longer available here
    if err := callFunc(funcVar); err != nil {
        // err and funcVar are available inside if structure
        fmt.Errorf("invalid value %d: %v\n", funcVar, err)
    }   // err is no longer available here
}   // funcVar is no longer available here
```

> 值得注意的是，如果声明了一个变量但是没有使用（通常情况下这意味着程序的逻辑有错误，哪怕没有错误，也会多占用内存），其他语言的做法通常是编译时给出一个警告，Go的方式不一样，编译时会直接给出编译错误。这意味着你构建一个可能有错误的程序。

### 变量遮蔽

在Go中，内部结构块声明一个外部结构块中的同名变量是允许的，不会产生编译错误，例如：

```go
... {
    a := 0
    ... {
        a := "a string" // ok
        // variable a is a string inside inner block
        ...
    }   // string variable a is no longer valid
    fmt.Println(a)  // output 0
}
```

上例中，我们说里层的a变量遮蔽了外层的a变量。因此，在里层的代码块中，a是一个字符串变量，离开了里层代码块将会失效，在外层的代码块中，a是一个整型变量。

> Go的初学者很容易犯这样的错误，事实上，这个不属于Go的设计不良，到了后面（详见[EP.V 函数](Episode.V.Function.md)），我们会发现这个设计其实是有意义的。

[例子：变量遮蔽](examples/ep01/short_var_definition.go)

[下一篇 EP.II 数组与切片](Episode.II.Slice.md)
