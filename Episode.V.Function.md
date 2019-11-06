# 函数

函数是程序设计的最基本组成单元，能提供一个功能的抽象和结构组合。就像数学上函数的定义一样，程序中的函数也可以定义为由输入集合到输出集合的一种映射关系。例如，在人工神经网络中经常用到的激活函数Sigmoid，它的数学定义为：

![sigmoid latex](https://latex.codecogs.com/gif.latex?Sigmoid(t)=\frac{1}{1+e^{-t}}%20(t\in\mathbb{R}))

它表明sigmoid函数在其唯一参数t上定义，参数t的取值范围是所有实数。函数的输出是一个值，取值范围是$(-1, 1)$

Go当中函数是一种基本类型，定义一个函数包括三个要素：参数（列表）、返回值（列表）和函数体：

```go
func increment(x int) int {
    return x + 1
}
```

我们分析一下上述的语法，func是定义函数的关键字，类似于Python中的def，JavaScript中的function，Rust中的fn；increment是这个函数的名称，注意函数名称并不是函数必须的要素，参见[匿名函数](#lambda)；该函数接受的参数只有一个，类型是标准整型，参数名称为x；函数返回值也只有一个，类型也是标准整型；函数体非常简单，将x增加1的结果返回。

如果你有C/C++或Java这样的语言知识，你会发现这里的语法都是倒置的，就像Go当中的变量声明一下，返回值写在函数参数的后面，而不是前面。语法倒置首先能提供更加清晰的多返回值声明，其次也能为[方法](Episode.VIII.Method.md)提供独立定义的可能性。

函数的签名是由func关键字后面加上小括号中的参数类型列表以及返回值列表组成的。判断两个函数类型是否相同，需要对比它们的参数类型及顺序是否完全一致，还需要对比它们的返回值类型及顺序是否完全一致。例如：

```go
// f1和f2具有一样的类型
var f1 func(int) bool
var f2 func(int) bool
// f3和f4具有一样的类型
var f3 func(float64, string, bool) (string, error)
var f4 func(float64, string, bool) (string, error)

Go的函数是语言中的一等公民（first class)，函数可以作为变量，作为参数，作为返回值，甚至结构体的字段。