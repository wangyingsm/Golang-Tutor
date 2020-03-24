# 方法

在[EP.VI 结构体](Episode.VI.Struct.md)一章中，我们已经初步了解了Go在面向对象设计方面提供的最基础特性，后面的几章内容中，我们会更加深入的了解面向对象在Go当中的实现。

前面我们知道了Go中最基本的代码结构是函数，函数从属于包（详见[EP.XII 包](Episode.XII.Package.md)），任何满足可见性的地方都可以调用函数。除此之外，我们知道面向对象程序设计中有一个很重要的概念叫做方法，在大多数面向对象语言中（如Java、Python），方法都是从属于类的。例如下面简单的Python代码片段：

```python
class Clazz(object):
    def __init__(self):
        pass
    def method(self, **kwargs):
        print(kwargs)
```

`Clazz`类有一个初始化方法`__init__`和一个成员方法`method`，当我们初始化一个类实例时`c = Clazz()`，`__init__`方法会隐含调用，完成实例初始化工作。然后就可以使用`c.method(a=0, b=False)`来调用类的成员变量了。

在Go当中，刻意回避了类的概念，设计的初衷是为了避免面向对象一些固有的缺点，但是Go依然为OO提供了相应的解决方案，方法也是其中之一。与上面的语言不同，Go当中的方法是从属于类型的。其语法与函数语法定义保持一致，并且不要求写在从属类型的结构当中，可以书写在包中任何位置上。可以认为Go语言的方法是一种特殊的函数。我们先来看如何使用我们已经了解的函数模拟出方法的效果：

```go
type MyStruct struct {
	ID   int
	Name string
}

func MyStructToString(m MyStruct) string {
	var buf bytes.Buffer
	if m.ID >= 0 {
		buf.WriteString(fmt.Sprintf("ID -> %d\n", m.ID))
	}
	if m.Name == "" {
		buf.WriteString("Name -> noname")
	} else {
		buf.WriteString(fmt.Sprintf("Name -> %s", m.Name))
	}
	return buf.String()
}

func main() {
	m := MyStruct{
		ID:   1,
		Name: "John Doe",
	}
	fmt.Println(MyStructToString(m))
	m.ID = -1
	m.Name = ""
	fmt.Println(MyStructToString(m))
}
```

[运行这个例子](https://goplay.space/#d8IzlCMp2sy)

上例中`bytes.Buffer`是标准库提供的内存缓冲区结构体，`fmt.Sprintf`是格式化文本函数，这些部分重要内容可以在[EP.XIX 标准库](Episode.XIX.Stdlib.md)中找到。其他部分代码我们应该不需要再做解释了。这里的函数作用是接收一个`MyStruct`结构体参数，然后按照要求转换成字符串。但是如果我们深入思考一下就会发现，实际上这个函数从逻辑上就应该从属于`MyStruct`结构体，因为功能上它就是对这个结构体进行操作的，因此如果将这个函数实现成结构体的方法显然更加合适：

```go
func (m MyStruct) String() string {
	//
}

func main() {
	m := MyStruct{
		ID:   1,
		Name: "John Doe",
	}
	fmt.Println(m)
	m.ID = -1
	m.Name = ""
	fmt.Println(m)
}