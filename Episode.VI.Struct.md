# 结构体

## 结构体定义

在前面我们已经介绍了数组、切片和字典的内建类型，和基本类型不同，这些类型存储的数据都不是一个标量。因此我们也可以将它们称为复合类型。但是前面介绍标量类型和这三种复合类型，都只能涉及同一种类型的数据。如果我们需要在一个数据结构中存储多种不同类型的数据的时候，前面介绍的类型都不太适合。

Go提供了结构体的类型用来存储多种类型数据的集合，结构体`struct`里可以包含除了接口类型之外的其他任何类型，每一个组成部分被称为结构体的一个字段，理论上每个字段都需要一个结构体内部唯一的名称：

```go
// 将一个struct定义为类型myStruct
type myStruct struct {
    fieldA int
    fieldB string
}
// 构建并初始化一个myStruct结构体变量
s1 := myStruct {
    fieldA: 123,
    fieldB: "a string",
}
// 构建并初始化同一个struct结构体的变量
s2 := struct {
    fieldA int
    fieldB string
}{
    fieldA: 123,
    fieldB: "a string",
}
```

从上例中，首先明确的是结构体的类型实际上是由struct关键字及其后花括号括起来的字段定义部分组成的，第一行代码的type只是为这个类型起了一个别名；然后留意初始化结构体的语法，如果你熟悉JSON和Python，这一点难度都没有，Go中的结构体变量就相当于一个JSON对象或者Python中的一个字典，因此初始化的语法也很接近（当然如果你记得Go是一门静态类型的语言的话，很显然你不能为字段设置任意类型的值）；最后我们比较一下变量s1和s2，你应该能看出这两个结构体变量的底层存储结构是一致的，而且实际上我们也为它们的所有字段设置了相同的值。

```go
// 接着上例
// 输出 {123 a string}
fmt.Println(s1)
// 输出 {123 a string}
fmt.Println(s2)
// 输出 true
fmt.Println(s1 == myStruct(s2))
```

[运行上面例子](https://goplay.space/#U1eEa63nthL)

可以看到s1和s2具有的类型是不同的，但是可以使用类型转换语法互相转换（此处将s2转换为myStruct类型），并且可以使用相等条件语法进行比较。熟悉C++或Java的读者可能会奇怪为何这个相等比较会返回true，回想[上一章](Episode.V.Function.md)我们说过，Go采用的是最简单的永远传值模式，因此这个相等比较的结构体的内容，而不是所谓的*引用*，后面[EP.VII 指针](Episode.VII.Pointer.md)一章中我们会更加深入的讨论这点。此处你只需要记住，只要没有明确使用指针或取址操作的情况下，Go变量永远代表它的值本身。

## 结构体的字段

结构体中可以含有任意个字段，每个字段都由可选的字段名称以及必须的字段类型组成。通常来说，我们都是使用有名称的字段，因为这样能让我们的代码有着明确意义，同时也更加容易维护。每个字段的定义就是字段名称加字段类型，任何Go中的合法类型包括自定义类型都可以作为结构体字段的类型。结构体每个字段惯例会写在不同的行中，每两行之间无需任何分隔符号（事实上是编译器自动在每行航模加上了分号）；如果多个字段需要定义在同一行，两个字段之间需要使用分号分隔，这种影响阅读的书写方式是不推荐的：

```go
// 尽量避免以下写法
type MyStruct struct {fieldA int; fieldB string}
```

在不考虑内存优化的情况下，字段定义的顺序是不重要的，也就是说上例中AB字段的顺序也可以变为BA，而不影响程序功能。

结构体字段可以没有名称，但是这个做法的前提是结构体中不能包含重复的字段类型，例如：

```go
type MyStruct struct {
    int
    string
}
```

上例中的语法是正确的，结构体有两个字段，分别是标准整型和字符串类型。由于字段名称缺失，因此在初始化结构体时，无法按照名称来设置相应的值，提供的初始化值需要按照字段定义顺序进行，此时的字段定义顺序是重要的。我们可以如下初始化上面的结构体：

```go
m := MyStruct {
    0,
    "hello",
}
```

详见[结构体初始化](#结构体初始化)。

同样，由于字段名称缺失，当引用字段时，也需要使用字段类型来进行引用：

```go
fmt.Println(m.int)    // output 0
fmt.Println(m.string) // output "hello"
```

从上面的代码也可以看到这种忽略字段名称的方式，无法支持多个字段有着同样的类型。例如下面的结构体定义就是**错误的**:

```go
// 错误的结构体定义
type MyStruct struct {
    int
    string
    int
}
```

> 熟悉Python或者Rust语言的读者知道它们有元组Tuple类型，例如Python的`rgb = (0, 0, 0)`或者Rust的`struct RGB (u8, u8, u8)`。这些元组中的字段可以使用序号进行访问，如Python的`rgb[0]`或者Rust的`rgb.1`。然而Go并不支持元组，因此没有序号访问字段的概念。因此上面的元组只能实现成结构体：
> 
> ```go
> type RGB struct {
>     r uint8
>     g uint8
>     b uint8
> }
> ```
> 
> 好处在于代码更加明确易读。不过元组有其优势，特别类似RGB这种结构体，每个字段的定义基本都是明确的。笔者个人希望后续Go能考虑加入这个特性。

### 空结构体

一个不含任何字段定义的结构体称为空结构体，其类型为`struct {}`。因为空结构体不含有任何字段，因此它占用的内存空间为0。这也是Go语言中占用内存最小的类型，按照信息理论，它仅能表示两种状态：存在或不存在，因此经常被用在一些二态标记的存储和传输上：

```go
func main() {
    empty := struct{}{}
    // 输出为0
    fmt.Println(unsafe.Sizeof(empty))
}
```

[运行这个例子](https://goplay.space/#RjKkxrU5Mty)

### 结构体初始化

前面我们已经接触了一些初始化结构体的方法，其标准语法与数组或切片的初始化语法类似，结构体类型后面的一对大括号中书写字段的初始化值。例如：

```go
s := struct {
    id   int64
    name string
}{
    id:   1,
    name: "John Doe",
}
```

后一对大括号当中就是初始化的语法，同样这也是你很熟悉的类JSON语法，与map类型初始化一样，每个字段初始化后面都带有逗号。

当然结构体初始化并不需要对每个字段都进行初始化赋值，可以仅对结构体中的部分字段进行初始化，此时忽略的字段可以不在初始化语法中出现。例如：

```go
s := struct {
    id   int64
    name string
}{
    name: "John Doe",
}
fmt.Printf("%+v\n", s) // 输出{id:0 name:John Doe}
```

[运行这个例子](https://goplay.space/#fzJqGAO5iBR)

上例中我们可以看到，初始化时我们忽略了字段`id`的值。前面我们说过，Go当中所有变量都会初始化，不存在着未初始化的变量，这条规则对于结构体的字段也是通用的，未初始化的字段会被自动设置成该字段类型的零值。因此输出中的id值为0。

字段初始化的顺序也可以不按照结构体字段定义的顺序，例如：

```go
s := struct {
    id   int64
    name string
}{
    name: "John Doe", 
    id:   1,
}
```

这样的初始化方式的效果与前一个例子中的情况是一致的。在**逻辑**上，你甚至可以认为命名字段的结构体类似于字典，字段名称就是字典的key，而字典是无所谓顺序的。但是在忽略字段名称的结构体初始化时，初始化字段的顺序是重要的，不能乱。

### *字段顺序优化*

如果你对于Go语言程序性能优化暂时还不太感兴趣，你可以直接掉过本小节，进入[字段标注字符串](#字段标注字符串)。

Go当中有着它从C语言继承过来的内存对齐机制，因此结构体中的字段并不是一定占用了其原始数据类型的内存长度，而是在很多情况下，都会占用更多的内存空间。我们来看一个最简单的例子：

```go
type User struct {
    flag bool  // 布尔类型仅占用1个字节
    id   int64 // int64类型占用8个字节
}

func main() {
    u := User{
        flag: true,
        id:   0x100000000000000,
    }
    // 使用unsafe.Sizeof获得变量u的长度，此处输出了16
    fmt.Println(unsafe.Sizeof(u))
}
```

[运行这个例子](https://goplay.space/#YDuGdHKfmIy)

如果按照字段真实占用的内存计算，你会预期输出的空间字节大小应该是9，但是实际上输出的是16。这是因为Go编译器在编译时，会采用内存对齐的方式对User结构体进行封装，因此flag字段也占用了8个字节的长度，而实际上该字段仅仅使用了第一个字节的内容，后面的7个字节是为了保持对齐而产生的占位字节。通常来说，这不会是什么大问题，毕竟现在的计算机内存通常会以GB作为单位进行计算。但是如果你将Go作为一门系统语言而不是仅仅用来做应用开发的话，这样的内存占用区别可能会对系统性能造成较大的影响。特别当结构体的字段很多或者结构体嵌套内容较长，且该结构体被系统程序大量使用时，这个差别可能会最终到达GB的量级，这种代码的优化变成了需要。**注意：再次提醒，一般情况下这个优化的工作都是不需要的，而且永远不应该成为你编写代码时的第一考虑，请记住“过早地优化是犯罪”**

在刚才的例子中，无论我们如何调整字段的顺序，我们都没有办法节省内存空间了，假设我们观察下面的例子，在原来的基础上增加了一个字段：

```go
type User struct {
    flag       bool  // 布尔类型仅占用1个字节
    id         int64 // int64类型占用8个字节
    loginTimes int   // int类型占用4个字节
}

func main() {
    u := User{
        flag: true,
        id:   0x100000000000000,
    }
    // 使用unsafe.Sizeof获得变量u的长度，此处输出了24
    fmt.Println(unsafe.Sizeof(u))
}
```

[运行这个例子](https://goplay.space/#VO4Tt2Hss4Y)

上例中我们在结构体中增加了一个字段`loginTimes`，其类型是标准整型，正常应该占用4个字节，由于我们已经知道了对齐的原理，因此很合理的可以猜测到输出的结果会是24。我们用下面的表格来展示u变量的内存占用和对齐情况：

| 字节0        | 字节1        | 字节2        | 字节3        | 字节4 | 字节5 | 字节6 | 字节7 |
|:----------:|:----------:|:----------:|:----------:|:---:|:---:|:---:|:---:|
| flag       | E          | E          | E          | E   | E   | E   | E   |
| id         | id         | id         | id         | id  | id  | id  | id  |
| loginTimes | loginTimes | loginTimes | loginTimes | E   | E   | E   | E   |

表格中的`E`代表着用来对齐的空字节。在这个例子的情况下，字段定义的顺序可以进行调整并优化程序的内存占用情况，我们仅将`id`和`loginTimes`字段的位置进行互换，其他代码保持不变：

```go
type User struct {
    flag       bool  // 布尔类型仅占用1个字节
    loginTimes int   // int类型占用4个字节
    id         int64 // int64类型占用8个字节
}
```

[运行这个例子](https://goplay.space/#raxonmMG7P9)

结果让人困惑，输出的值是16，而不是24，一个字段定义顺序的改变直接节省了8个字节的内存占用，原因在于Go编译器的对齐改成了下面的方式：

| 字节0  | 字节1 | 字节2 | 字节3 | 字节4        | 字节5        | 字节6        | 字节7        |
|:----:|:---:|:---:|:---:|:----------:|:----------:|:----------:|:----------:|
| flag | E   | E   | E   | loginTimes | loginTimes | loginTimes | loginTimes |
| id   | id  | id  | id  | id         | id         | id         | id         |

假设User结构体在程序中维护了一百万个用户的信息，那么这个简单的字段顺序修改将带来将近8MB的内存空间节省。如果这个结构体字段数量很多并且含有嵌套结构体的话，那么有可能会节省上百MB的内存空间。最后用一个例子结束这一小节，当我们往`User`结构体继续添加一个2字节长度的字段时，最优化的内存占用情况可以如下：

```go
type User struct {
    id         int64 // int64类型占用8个字节
    flag       bool  // 布尔类型仅占用1个字节
    friends    int16 //int16类型占用2个字节
    loginTimes int   // int类型占用4个字节
}
```

[运行这个例子](https://goplay.space/#iMCbUrylzLo)

### 字段标注字符串

这个小节会讲述Go语言结构体字段定义非常有趣和创新的一点，可以为每个定义的字段添加标注的字符串。这个标注字符串不但能被Go编译器识别，还能在运行时得到解析，为结构体字段添加额外的功能。

下面以我们最常用的JSON数据序列化和反序列化为例讨论标注字符串的基本使用：

```go
type Item struct {
    ID          uint64 `json:"id"`          // json中的key为id
    ItemName    string `json:"itemName"`    // json中的key为itemName
    Description string `json:"description"` // json中的key为description
    Price       uint   `json:"price"`       // json中的key为price
    onSale      bool   // 不输出到json当中
}

func main() {
    // 初始化一个item结构体
    item := Item{
        ID:       10000,
        ItemName: "《Go编程语言》",
        Price:    5890,
        onSale:   true,
    }
    // 将该结构体序列化成一个json的字节数组，发生错误则退出程序
    j, err := json.Marshal(item)
    if err != nil {
        panic(err)
    }
    // 将序列化后的json字符串打印出来
    // 此处输出{"id":10000,"itemName":"《Go编程语言》","description":"","price":5890}
    fmt.Println(string(j))
    // 将JSON字节数组反序列化成一个新的Item结构体，发生错误则退出程序
    var i Item
    err = json.Unmarshal(j, &i)
    if err != nil {
        panic(err)
    }
    // 将反序列化后的结构体与原结构体比较，此处输出false
    fmt.Println(item == i)
    // 设置onSale字段为true后再次进行比较，此处输出true
    i.onSale = true
    fmt.Println(item == i)
}
```

[运行这个例子](https://goplay.space/#MZf_MCJd_43)

这是目前为止我们看到的最长的一段代码，里面还有一些内容我们还没有介绍到，因此我们需要比较详细的过一遍。

首先看到的是Item结构体的定义，其中一共有5个字段，每个字段都有名称和类型，但是前四个字段的后面还有一个额外的字符串标注内容，例如````json:"id"````，这个标注的含义其实就是表示该字段应该能被序列化到JSON数据结构（或从JSON反序列化回来）当中，并且在JSON中使用的键名称是`id`，另外三个字段`ItemName`、`Description`和`Price`也是如此。而最后一个字段`onSale`因为没有定义这样的标注，所以将不会被序列化和反序列化。这里有三点必须注意的：

- 一是标注字符串使用Go中的原始字符串进行编写，也就是代码中看到的成对出现的````````符号，其中的任何符号都会当成原始字符存在，不需要额外转义。但是这种写法不是必须的，例如`ID`字段的标注也可以写成`"json:\"id\""`，两者结果是一致的。但是显然后面的写法很复杂且难以阅读。

- 二是标注中的语法是固定的，不要为了美观增加任何的空白字符，这会导致运行时一些莫名其妙的错误，后面我们很快就会看到。

- 三是所有需要序列化或反序列化的字段都必须使用大写字母作为命名的首字母，这是由于标准库中json模块需要能够访问这些字段才能对其进行操作，详见[EP.XIII 可见性](Episode.XIII.Visibility.md)。这是Go新手非常容易踩的坑之一，因此需要特别注意。

然后初始化一个Item结构体，初始化中字段`Description`留空，意味着该字段使用了字符串的零值`""`。之后我们使用`"encoding/json"`包对`item`结构体进行了序列化，这是通过Marshal函数完成的，返回值包括一个字节数组和一个错误，字节数组是结构体序列化后JSON的字节内容，而本章后半段会介绍到的[错误](#错误)表示序列化过程中是否有错误发生。打印得到的JSON字符串中共有四个键，对应着结构体定义中标注了的四个字段，`Description`字段具有零值，因此JSON中也是一个空白字符串。最后我们将JSON重新反序列化成`Item`结构体，与原结构体内容进行比较，第一次比较结果是不相同的，原因在于`onSale`字段没有反序列化，因此具有原始的零值`false`，将反序列化后的结构体`onSale`字段赋值为`true`后，两者比较的结果是相同的，与我们期待的结果一致。

从上面的例子中我们还注意到`item`结构体的`Description`字段是字符串零值，序列化成为JSON后，这个字段也是具有空字符串值，如果我们希望这种零值字段不包含在序列化的JSON中，我们可以修改一下结构体的定义，程序其他部分内容不变：

```go
// 本程序将输出{"id":10000,"itemName":"《Go编程语言》","price":5890}
type Item struct {
    ID          uint64 `json:"id"`                    // json中的key为id
    ItemName    string `json:"itemName"`              // json中的key为itemName
    Description string `json:"description,omitempty"` // json中的key为description
    Price       uint   `json:"price"`                 // json中的key为price
    onSale      bool   // 不输出到json当中
}
```

[运行这个例子](https://goplay.space/#SFysbw2A-tE)

上面的结构体定义中唯一的修改是`Description`字段的标注字符串，增加了`omitempty`标注，这个标注表示如果字段值为零值，那么它将不会在序列化的JSON当中出现，因此JSON字符串的输出中就没有`"Description"`字段了。**注意，这里标注的格式非常严格，添加空白字符将会失效，例如写成`json:"description, omitempty"`，不会产生任何编译期和运行期错误，很容易中招**。

上面讨论了标注字符串最常见的使用场景，JSON转换。实际上标注字符串还能用在更多的标注定义上，例如`YAML`格式的序列化和反序列化，`YAML`格式的转换不在Go标准库中提供，需要用到第三方库`"gopkg.in/yaml.v2"`，有关Go语言包结构和管理的内容请参见[EP.XII 包](Episode.XII.Package.md)。有关`YAML`标注的使用请参考下面的例子，从代码中也可以看出字段有两个或以上独立标注时的语法，两个标注之间使用空格分隔。

[结构体-字段例子](examples/ep06/struct_type.go)

#### *自定义标注*

本小节属于Go当中比较高级一些的技巧，如果你认为你暂时没有使用的需求，可以跳过直接阅读[嵌套结构体](#嵌套结构体)。

上面讨论的都是标准库或者第三方库帮我们实现了标注字符串的例子，如果我们需要编写自己的工具包对结构体进行处理的话，我们需要实现自定义的标注字符串，然后使用反射机制按照标注字符串的定义对结构体进行相应的处理。

下面用一个简单的例子，来说明自定义标记的使用。我们假设有这样一个需求，我们的结构体中有时间类型的字段，对应这Go中time.Time结构体数据。如果采用json标准序列化，会得到类似`2020-03-13T17:02:26.857285407+08:00`这样的序列化结果。这通常不是我们希望得到的时间格式。我们希望能够在结构体中自定义时间格式的输出方式，并且和大部分语言（例如Java，Python等）兼容，这些格式将使用诸如`yyyy`这样的格式化字符串来定义（而不是Go中的`2006-01-02`的方法）。然后我们实现自定义的方式将结构体进行JSON序列化。为了简单起见，下面的例子仅支持如下时间格式标记：

| 格式字符串 | 格式说明 |
|-|-|
| `yyyy` | 四位的年份 |
| `MM` | 两位的月份 |
| `dd` | 两位的日期 |

> 特别注意：千万不要直接在生产系统中使用下面的例子，该例子大量忽略了异常分支处理，也忽略了大部分的数据类型转换，并且不支持嵌套的结构体转换。直接应用在生产系统使用下面的例子可能会导致panic。读者可以在这个例子的基础上进行功能补全和异常分支处理后，作为生产系统来使用。

[完整的代码例子在这里](examples/ep06/struct_tag.go)

```go
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
```

例子当中使用了很多我们仍未接触到的内容，包括[EP.VII 指针](Episode.VII.Pointer.md)、[EP.VIII 方法](Episode.VIII.Method.md)、[EP.X 接口](Episode.X.Interface.md)和反射（参见[EP.XIX 标准库](Episode.XIX.Stdlib.md)）。我们会在后面的章节一一介绍。

此处最重要的是表明我们能使用反射机制对自定义的结构体字段进行解析和处理，我们还能覆盖标准的JSON序列化方法。程序中的注释较为详细，因此不再做详细代码说明。

## 读取和设置结构体字段

当一个结构体初始化后，它的字段就能够被读取或者赋值了。要引用结构体的字段使用简单的`结构体变量.字段名称`即可，如果该字段是匿名字段，字段名称就是该字段的类型。例如：

```go
type MyStruct struct {
	ID   int
	Name string
	int64
}

func main() {
	m := MyStruct{
		ID:    1,
		Name:  "John Doe",
		int64: 0x1000000000000,
	}
	fmt.Println(m.ID)
	fmt.Println(m.Name)
	fmt.Println(m.int64)
}
```

[运行这个例子](https://goplay.space/#KYMPOfij-hs)

我们还可以通过引用字段来修改字段的值，例如同样的结构体定义：

```go
func main() {
	m := MyStruct{
		ID:    1,
		Name:  "John Doe",
		int64: 0x1000000000000,
	}
	m.ID++
	m.Name = "Thanos"
	m.int64 -= 0x1000000000000
	fmt.Println(m.ID)
	fmt.Println(m.Name)
	fmt.Println(m.int64)
}
```

[运行这个例子](https://goplay.space/#y0nc2f0DHY5,11)

[EP.V 函数](Episode.V.Function.md) <|> [EP.VII 指针](Episode.VII.Pointer.md)