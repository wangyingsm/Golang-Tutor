# 函数

函数是程序设计的最基本组成单元，能提供一个功能的抽象和结构组合。就像数学上函数的定义一样，程序中的函数也可以定义为由输入集合到输出集合的一种映射关系。例如，在人工神经网络中经常用到的激活函数Sigmoid，它的数学定义为：

![sigmoid latex](https://latex.codecogs.com/gif.latex?Sigmoid(t)=\frac{1}{1+e^{-t}}%20(t\in\mathbb{R}))

它表明sigmoid函数在其唯一参数t上定义，参数t的取值范围是所有实数。函数的输出是一个值，取值范围是`(-1, 1)`。这样就建立了一个从输入集合到输出集合的映射。

在计算机程序中，函数具有更广泛的延伸，因为在程序运行过程中，除了输入到输出的映射之后，可能还执行了一些指令改变了计算机（或者称为图灵机）的状态，甚至可能某些函数不返回任何内容，仅为了改变计算机的状态。我们将这样的操作称为函数的副作用（side-effects）。

近年来流行的函数式编程就非常反对这种副作用，认为所有的计算机程序都应该能被分解为有限个lambda操作的组合。Go并非是一门完全的函数式语言，虽然它也支持部分函数式编程的重要概念。

## 函数定义

Go当中函数是一种基本类型，定义一个函数包括三个要素：参数（列表）、返回值（列表）和函数体：

```go
func increment(x int) int {
    return x + 1
}
```

我们分析一下上述的语法，func是定义函数的关键字，类似于Python中的def，JavaScript中的function，Rust中的fn；increment是这个函数的名称，注意函数名称并不是函数必须的要素，参见[匿名函数](#lambda)；该函数接受的参数只有一个，类型是标准整型，参数名称为x；函数返回值也只有一个，类型也是标准整型；函数体非常简单，将x增加1的结果返回。

如果你有C/C++或Java这样的语言知识，你会发现这里的语法都是倒置的，就像Go当中的变量声明一下，返回值写在函数参数的后面，而不是前面。语法倒置首先能提供更加清晰的多返回值声明，其次也能为[方法](Episode.VIII.Method.md)提供独立定义的可能性。

### 参数列表

参数列表写在函数名称后的一对小括号中，和大部分语言一样，两个参数之间使用逗号分隔。多个参数如果具有同一种类型，可以省略前面的类型定义，仅保留最后一个参数的类型即可。例如：

```go
// 多参数列表定义
func add(x int, y int) int {
    return x + y
}
// 可以缩写
func minus(x, y int) int {
    return x - y
}
```

#### 参数传递

Go当中调用函数时参数传递只有一条规则：那就是永远传值。这也是当时C语言设计的方式，因此，无需记忆哪些是基本类型，或者内存空间占用有多大，从而区分参数传递的是值还是引用（reference）。特别是如果你已经熟悉了C++/Python/Java/Rust这类语言之后，这点很容易造成错误。这种错误通常难以察觉和定位。

因为我们还没有介绍[结构体](Episode.VI.Struct.md)，所以用一个数组作为参数来说明这种情况：

```go
// 首先定义一个函数，你预期它会将传递进来的数组元素都增加1
func increaseOne(arr [5]int) {
    // 注意此处range的使用，实际上忽略了第二个返回即元素的值
    for i := range arr {
        arr[i] += 1
    }
}
```

下面我们试着将一个数组代入到参数调用increaseOne函数，完成后将数组打印出来看结果：

```go
// 注意，这里不能用slice，因为数组类型是包含长度的
a := [5]int{0, 10, 20, 30, 40}
// 函数无返回值，我们其实是使用了函数的副作用
increaseOne(a)
// 你可能会预期输出 [1 11 21 31 41]
// 但实际输出 [0 10 20 30 40]
fmt.Println(a)
```

[运行这个例子](https://goplay.space/#k1WnoS00zK1)

造成这个结果的原因就在于，实际上调用函数的时候，Go并不是将a这个数组的引用传递给了参数，而是将整个数组都复制了一份到函数的参数arr处，然后这份a的副本参与了函数体的运算，因此，a的内容并未发生任何改变。

如果你已经编写了一些Go程序了，你可能会反驳我上面的说法，比方说，"用slice作为参数类型，那不就相当于传的引用了吗？"。（如果你不理解这是什么意思，我建议你将上面例子中的代码修改一下，将参数的部分和a的声明部分都改为切片slice的类型，再次运行，你会得到原本预期的结果。）

对于上面这个问题，其实仔细思考还是会得到最后的答案，即Go当中永远传值。回想[EP.II 数组与切片](Episode.II.Slice.md)我们的讨论，切片本身就是一个结构，这个结构当中含有一个指向底层数组第一个元素的指针（我们会在[EP.VII 指针](Episode.VII.Pointer.md)一章详细介绍），因此你可以认为slice本身相当于其他语言中的一个引用。所以当你使用slice进行参数传递时，实际上会复制一份slice的结构（包括指针、长度、容量）到函数的参数处，因此传递的依然是值，而不是引用。这里的关键点在于数组与切片是两个不同的类型，如果参数类型是数组但是实际传递了数组的地址的话，那就是传递的引用，如果参数类型是切片，但是实际传递到函数的是这个切片的副本，那就是传值。

如果上面的讨论已经把你绕晕了的话，那就最后一个方法吧，Go的参数传递遵循最简单的规则，形参实际上是实参的副本，无论这个参数是什么类型。

> 如果你对于堆栈的概念不清楚，或者对于操作系统和编译原理也没有什么兴趣，本节最后这部分内容可以跳过，直接进入[可变长参数](#varargs)。一般我们都知道函数的参数和本地变量都是放在栈区的，而涉及动态内存分配的开销一般都是放在堆区（例如C语言中的malloc，C++和Java中的new）。栈区的读写速度一般都比堆区快，那为什么编程语言不能将所有变量都放置到栈区呢？原因首先是，在操作系统中，栈区的容量是固定的，因此不可能放置很大的内容，这一点就已经决定了没有虚拟机概念的编程语言无法将变量都放置在栈区，例如C/C++以及Rust；其次是栈区内容是会回收的，当线程的函数执行完并返回后，该函数原本占用的栈区空间就能被后续调用的函数使用，这时栈区这段区域的数据是不安全的。这决定了很多语言都需要在堆区中大量使用空间。

> Go是如何解决（或者叫优化）上述问题的？首先，Go并不完全依赖操作系统的线程，而是使用自身的协程，因此Go的栈区大小是可变的，在64位计算机上最小是2MB，最大是1GB，可以容纳大量的数据在栈区；其次，当函数返回时，由于前述的永远传值特性，函数栈区内的的数据也会被拷贝为副本返回到调用者的栈区相关区域，从而不会出现未定义的行为。

> 那么最后的问题就是，如何知道我们声明的变量是在栈区分配还是在堆区分配的内存呢？答案是，我们无法知道，这是由编译器决定的，编译器会根据上下文以及需要分配的内存大小等多种因素决定，该变量需要的内存最终分配在栈区还是堆区，通常写代码的你并不需要知道。如果你很关心这点，你需要更高级的方法如`pprof`，这部分内容超出本书范围，请参考[EP.XX 推荐资源](Episode.XX.References.md)。

#### <a id="varargs">可变长参数</a>

如果一个函数所需的某同种类型的参数个数不定，可能没有，可能多个，这比较适合一种称为可变长参数的场景。可变长参数的语法继承了类似C、C++、Java的语法:

```go
func sum(x ...int) (result int) {
    for _, v := range x {
        result += v
    }
    return result
}
```

在上面的代码段中，sum函数接受一个参数x，类型为`...int`，这个类型就是可变长类型的声明，表示sum能接受0个或多个的标准整数作为输入参数。当进入循环体后，这些整数会组合成一个`[]int`类型的slice，因此我们可以对其进行所有slice的操作，如上例的range。我们来看看这个简单的求和函数是怎么调用的：

```go
// 使用0个参数调用，输出0
fmt.Println(sum())
// 使用1个参数调用，输出数值本身100
fmt.Println(sum(100))
// 使用多个参数调用，输出所有参数的总和6
fmt.Println(sum(1,2,3))
```

[运行这个例子](https://goplay.space/#pI7F7tA5EcN)

如果函数有多个参数，其中一个是可变长参数，那么必须将可变长参数放在参数列表最后：

```go
func appendString(s string, post ...string) string {
    result := s
    for _, p := range post {
        result += p
    }
    return result
}

// ......

// 输出hello Golang wooh!
appendString("hello ", "Go", "lang", " wooh!")
```

假如我们想使用一个切片slice作为参数传递给可变长参数的话，我们可以使用`...`放在slice后面将其打散后传递，例如如果我们想调用上面定义的sum函数对一个slice进行求和：

```go
a := []int{1, 2, 3, 4, 5}
// 输出15
fmt.Println(sum(a...))
```

[运行这个例子](https://goplay.space/#hsdlBIItDpq)

[例子：函数的参数](ep05/function_param.go)

### <a id="returns">返回值列表</a>

Go是真正意义上实现多返回值的语言，对比Python和Rust，实际上这两者的多返回值都是通过tuple类型返回的，而Go并没有tuple类型。当一个函数需要返回多个值时，这些返回值的列表也需要使用小括号括起来：

```go
// 实现一个类似Python的divmod函数
// 返回3个值，分别是商，余数和错误
func divmod(x, y int64) (int64, int64, error) {
    if y == 0 {
        // 被0除，返回错误
        return 0, 0, errors.New("divide by zero")
    }
    // 返回商，余数和无错误
    return x / y, x % y, nil
}
```

类似大多数语言，Go也使用return关键字返回值，多个返回值之间使用逗号分隔。正如上面说到，Go并不是一个完全的函数式语言，因此不能根据最后代码行执行得到的结果作为函数的返回值（就像Rust的函数实现那样），所有的返回都需要显式地使用return关键字。

> 虽然Go可以实现返回任何个数的返回值，但是实际上使用的最多的情况仍然是返回一个值加一个错误，这也是Go规范中提倡的做法。在确实有多个返回值需求的情况下，规范并没有不提倡，就像上例中的divmod函数那样。如果函数的返回值确实太多，这种语法确实难以阅读，考虑使用[结构体](Episode.VI.Struct.md)作为返回值会更加合适。

#### 命名返回值声明

关于返回值，还有一点和其他语言有区别的地方，那就是Go可以给返回值进行声明，就像声明一个变量一样，我们姑且称之为返回值变量，该变量可以在函数体中参与运算。最后在函数返回时，如果并未指定值，则会将返回值变量作为返回值，如下两个函数作用是一样的：

```go
func add1(x, y int) int {
    return x + y
}
func add2(x, y int) (result int) {
    result = x + y
    // 下面的return语句并非可有可无的，不加return的话编译器会
    // 认为这是一个无返回值的情况，拒绝通过
    return
}
```

从上面的例子可以看到，定义一个返回值变量也需要使用小括号括起来，然后这个变量在函数体内就可以直接使用了，不需要再进行声明。根据Effetive Go（参见：[EP.XX 推荐资源](Episode.XX.References.md)）的推荐，这种语法适用于简短的函数使用，如果函数体较长（比如说上百行代码），这种写法反而会影响程序的可读性，在此情况下不推荐使用。

[例子：函数的返回值](ep05/multiple_returns.go)

## 函数的类型和签名

函数的签名是由func关键字后面加上小括号中的参数类型列表以及返回值列表组成的。判断两个函数类型是否相同，需要对比它们的参数类型及顺序是否完全一致，还需要对比它们的返回值类型及顺序是否完全一致。例如：

```go
// f1和f2具有一样的类型
var f1 func(int) bool
var f2 func(int) bool
// f3和f4具有一样的类型
var f3 func(float64, string, bool) (string, error)
var f4 func(float64, string, bool) (string, error)
// f5和f6是不一样的类型
var f5 func(string, int)
var f6 func(string, int) error
// f7和f8是不一样的类型
var f7 func(string, int) (string, bool)
var f8 func(int, string) (bool, string)
```

Go的函数是语言中的一等公民（first class)，函数可以作为变量（如上例），作为参数，作为返回值，甚至结构体的字段。复习一下第一章，函数的零值是nil。

### 函数作为参数

熟悉Python的读者一定很熟悉filter内建函数，它能将一个可迭代集合（Iterable）经过某个条件的函数过滤之后，输出成所有元素都符合该条件的另一个Iterable集合，非常方便。Go当中却没有这个内建实现，我们来试着实现一个简单的filter版本（仅实现标准整型slice的过滤器）：

```go
// 实现一个类型python的filter函数
// 接受的参数第一个是过滤函数，第二个是int的slice
func filter(prediction func(int) bool, iter []int) []int {
    // 返回的slice
    var result []int
    for _, v := range iter {
        // 逐个检查iter中的元素是否符合prediction
        // 这里的prediction是一个函数，接受一个标准整数，返回一个布尔值
        if prediction(v) {
            // 如果符合，将该元素加入到返回值slice中
            result = append(result, v)
        }
    }
    return result
}
// 定义一个判断是否为正数的函数，该函数签名与prediction一致
func isPositive(x int) bool {
    return x > 0
}
// 定义一个判断是否为偶数的函数，该函数签名与prediction一致
func isEven(x int) bool {
    return x % 2 == 0
}
```

上面定义了三个函数，filter是用来过滤slice的，它接受两个参数，第二个参数很清楚明白，就是一个int类型的slice，相当于一个标准整型的集合。我们需要详细观察第一个参数，类型是`func(int) bool`，说明这个参数的类型是一个函数，这个函数的签名是标准整型参数和布尔返回值。任何与这个签名一致的函数都是同一类型，都可以作为参数代入。因此我们在下面定义了两个相同签名的函数，isPositive用来判断输入参数是否是正数，isEven用来判断输入参数是否是偶数。

下面，我们就可以调用这三个函数，来完成对同一个slice的过滤工作：

```go
arr := []int{-50, 17, 3, 28, -9, 0, -33, 26, 15}
// 使用isPositive作为prediction
// 输出 [17 3 28 26 15]
fmt.Println(filter(isPositive, arr))
// 使用isEven作为prediction
// 输出 [-50 28 0 26]
fmt.Println(filter(isEven, arr))
```

[运行这个例子](https://goplay.space/#d8v1RQXWGrs)

从上面的例子可以看到，任何满足`func(int) bool`签名的函数都可以作为prediction参数代入到filter函数进行运算，事实上你可想出无穷多个符合的条件来对slice进行过滤，只要这个条件的函数接受一个标准整数并返回一个布尔值即可。

### 匿名函数

上面的例子中，如果每次写一个prediction函数都需要定义一个函数的话，往往会影响代码的可读性，特别是当这个函数很可能只有一个地方使用到的情况下。为此，Go提供了一种不具名内联函数的语法，其功能和作用类似与JavaScript的匿名function或者Python的lambda，甚至Java8之后也提供了这种匿名函数的语法，以下例子代码接着上一小节的例子：

```go
years := []int{1900, 1996, 2000, 2010, 2012, 2016, 2019}
// 对于上述的年份slice，将里面的闰年过滤出来
// 输出 [1996 2000 2012 2016]
fmt.Println(filter(func(year int) bool {
    return (year % 4 == 0 && year % 100 != 0) || (year % 400 == 0)
}, years))
```

[运行这个例子](https://goplay.space/#HRDZB5UJLir)

上面例子中，我们提供一个新的prediction函数就能将一个年份集合中所有的闰年过滤出来（此处没有考虑输入数据不合法的情况）。但是这个predition函数和前面例子中的isPositive和isEven不相同，它并没有名字，并且这个函数的定义是直接内联写在filter函数的调用语法当中。它等同于像下面一样定义了一个名为isLeapYear的函数，然后将这个函数代入filter的第一个参数一样：

```go
func isLeapYear(year int) bool {
    return (year % 4 == 0 && year % 100 != 0) || (year % 400 == 0)
}
```

匿名函数适用于即来即用的场景，如果你的函数可能被多个不同地方的代码调用的话，那是不合适的。还有一种情况是，如果函数体的代码比较长，也应该避免使用匿名函数，因为这反而会降低代码的可读性。

[例子：使用正则表达式对用户输入进行有效性验证](ep05/function_type.go)

### 闭包

闭包是函数式编程的重要概念之一，也可以称为生成函数的函数。使用闭包可以提供一系列函数的模板，调用者可以从模板生成自己的函数，提供给程序使用。

这里使用一个简单的例子说明。使用过Python的读者一定对里面的range和列表解析有很深刻的印象，我们希望实现一个函数，能适应多种不同的类似列表解析的功能。不熟悉Python也没有关系，不会影响你理解下面的代码片段。

```go
// 指定范围生成slice需要三个参数，范围的开始值start，包括在范围内；
// 范围的结束值end，不包括在范围内；以及步长值step，表示每次增长的量
// 我们的函数将步长值先抽取出来，方便后续我们产生不同步长值的生成函数
// 为简单起见，这里只考虑正数步长以及end大于等于start的情况
func makeRangeSlice(step int) func(int, int) []int {
    // 生成的结果slice
    var slice []int
    // 返回一个函数，接受start和end参数，最终这个匿名函数的返回值
    // 才是生成的slice
    return func(start, end int) []int {
        for i := start; i < end; i += step {
            slice = append(slice, i)
        }
        return slice
    }
}
```

下面我们可以试着使用步长为1，2，3来产生三个函数，让后分别使用它们产生`[0, 10)`区间的slice：

```go
step1 := makeRangeSlice(1)
step2 := makeRangeSlice(2)
step3 := makeRangeSlice(3)
// 输出[0 1 2 3 4 5 6 7 8 9]
fmt.Println(step1(0, 10))
// 输出[0 2 4 6 8]
fmt.Println(step2(0, 10))
// 输出[0 3 6 9]
fmt.Println(step3(0, 10))
```

如果某个步长值我们仅使用一次，也可以用简单的方式进行调用：

```go
// 输出[0 4 8]
fmt.Println(makeRangeSlice(4)(0, 10))
```

[运行上面例子](https://goplay.space/#Xjs2PZSqlRI)

练习：将上述闭包函数修改为支持负数步长，并对参数错误情况进行有效处理。

[例子：闭包](examples/ep05/function_closure.go)

## 特殊的函数：main 和 init

Go中有两个特殊的函数，分别是main和init。main函数是整个程序的入口，这也是大部分类C语言的惯例，相信读者也已经在前面的例子中接触到了main函数了。与C/C++/Rust不同，Go语言中的main函数有着严格的签名：

```go
func main() {
    // 程序从这里开始
}
```

如果你熟悉C语言的main函数签名，Go有点区别，Go的main函数不接受任何参数，也不返回任何值，当你定义程序入口时，必须严格遵循这个函数签名。并且main函数只能处于main package中（详见[EP.XII 包](Episode.XII.Package.md)）。任何在其他package中定义的main函数都不会被编译器认为是程序入口。

> 如果在Go中需要接受命令行参数和返回值给调用者的话，你需要用到os包中的Args变量和Exit()方法（详见[EP.XIX 标准库](Episode.XIX.Stdlib.md)和`go doc`文档）。

另一个比较陌生的特殊函数是init，它负责对每个包`package`进行初始化工作。根据Go语言标准说明，当一个包被导入`import`的时候，以下事件会依次发生：

1. 递归的调用本包导入的其他包`package`的初始化工作；
2. 对本包的所有包定义变量进行初始化工作；
3. 执行本包定义的init函数（如果存在）。

并且，init函数能够保证在一次程序执行过程中只会被执行一次，无论它被导入`import`了多少次（详见[EP.XII 包](Episode.XII.Package.md)）。在main package中，init函数被保证会在main函数之前被执行。init函数也有着与main函数一样的签名：

```go
func init() {
    // 包的其他初始化工作写在这个函数中
    // 例如复杂内存空间的分配，数据库的连接池创建等
}
```

[运行main和init例子](https://goplay.space/#ZC1E9p9-_rE)

[EP.III 字符串](Episode.III.String.md) <|> [EP.VI 结构体](Episode.VI.Struct.md)