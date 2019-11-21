# 字典类型

Go中的字典`map`是一种内建类型，其实现和使用方式与C++中的Map，Python中的dict，以及Java和Rust中的HashMap基本相同。都是提供一个在内存中基于键值对（`key-value`）存储数据的结构。

字典变量在Go当中也是一种引用，类似于切片slice，slice的底层存储是数组，而map的底层存储是哈希表（HashTable）。其存储结构如下图：

![map底层结构](imgs/map_underly.png)

> 初学者可以跳过这部分内容直接进入[创建和初始化字典](#creating)一节。Go的map实现是：首先使用一个底层数组存储桶`bucket`，也就是所有字典中的值都会分布在桶当中，键的hash值的低二进制位决定着其存储所在的桶的序号，即图中的`LOB Hash`；然后每个桶都维护着另一个数组的指针，每个桶最多能有8个数据结点，如果发生了hash碰撞，导致结点数超过了8个，会使用额外的数组链接到该桶当中，数据结点的序号是通过键hash值的高二进制数位决定的，即图中的`HOB Hash`；结点中维护着指向最终存储键值对的指针，这些键值对被对齐打包放置在内存的一段区域内，即图中的K和V，其中的E代表为了内存对齐填充的空字节。

> 当在字典中查找某个key时，只需要将其进行Hash计算，即可立刻定位到桶和数据结点，然后在与存储的K进行对比，即可判断出是否定位到了这个key。因此查找的复杂度近似为`O(1)`。

## <a id="creating">创建和初始化字典</a>

在使用字典之前，首先需要确认key的类型和value的类型，map对于value的类型没有任何约束，但是对于key，要求为可以进行相等比较的类型，例如数值类型、string、指针和结构体（参见[EP.VI 结构体](Episode.VI.Struct.md)和[EP.VII 指针](Episode.VII.Pointer.md)），数组、切片就不能作为key的类型。更进一步，为保持hash计算的稳定性，key的类型应该尽量保持不变，否则很可能产生奇怪的错误，因此应该尽量避免使用结构体、指针这样的key类型。

完整的字典类型应该包括key类型和value类型，用`map[keyType]valueType`表示。例如，`map[string]int`表示一个key是字符串value是标准整型的字典类型。

最常用的创建一个空字典的方式是使用内建函数`make`，就像创建一个空的切片一样：

```go
// 创建一个字符串key和标准整型value的map，无初始容量
m := make(map[string]int)
// 创建一个标准整型key和布尔型value的map，初始容量为10
n := make(map[int]bool, 10)
```

### 未初始化的map

如果你以下述方式声明了一个map变量：

```go
var m map[string]int
```

这表示，m的类型是`map[string]int`，且m本身并未初始化，go当中所有未初始化的变量都会自动初始化为该类型的零值（参见[EP.I 类型、变量和常量](Episode.I.Type.md)），而map的零值是nil。此时对m变量进行的字典操作都会得到一个`nil dereference`的错误（实际上是一个panic），意思是该map还未初始化，无法使用map的相关操作。你仍然需要使用make对其进行初始化。

### 有初始化值的map

如果我们在创建map和初始化时，已经有一些`key-value`值需要初始化到字典里面，我们可以这样写：

```go
m := map[string]int {
    "a": 1,
    "Alice": 10,
    "Bob": -7,
}
```

留意花括号中的内容，即为初始化值的部分，熟悉JSON语法和Python语法的话，你一定不会感到陌生，这和JSON的对象以及Python的dict对象语法是基本一致的。这种语法可以代替make对map进行初始化：

```go
m := map[int]bool{}
```

上述写法与`make(map[int]bool)`的作用相同，都是创建并初始化了一个空的字典。

> 如果你仔细观察就会注意到：上述的花括号写法与JSON语法有一点点差别，就是最后一项之后还可以允许带有一个额外的逗号，当采用缩进方式书写代码时，能够保证字典中的每行格式是一致的，特别是采用了go fmt对代码进行标准格式化后，代码能够更加美观和易读。Python和Rust也支持这种额外的逗号语法，但是在JSON当中，这种写法会导致一个错误。

## 插入或更新字典值

类似Python，Go采用了简单的赋值语句来插入或者更新字段值。中括号中的key值如果在字典中不存在，则插入新的值，如果存在，就会将原来的value更新成新的值。例如：

```go
m := make(map[string]int)
// "a"不存在，插入该值
m["a"] = 1
// 输出为map[a:1]
fmt.Println(m)
// "Alice"不存在，插入该值
m["Alice"] = 10
// 输出为map[Alice:10 a:1]
fmt.Println(m)
// "a"已经存在，将原来的值修改为0
m["a"] = 0
// 输出为map[Alice:10 a:0]
fmt.Println(m)
```

[运行这个例子](https://goplay.space/#nkxNPB-R20F)

## 获取某个值

使用中括号中的key值可以插入或修改字典值，它的逆操作——获取字典中某个特定key的值，也是使用中括号，例如：

```go
m := map[string]int {
    "Alice": 0,
    "Bob": 1,
}
// 输出1
fmt.Println(m["Bob"])
```

### 获取字典中不存在的键值

这是Go初学者很容易犯的错误之一，Go的map本身不提供检测key是否存在的机制，如果一个键在字典中不存在，按照上面的方法去获取字典值然后直接使用会导致难以定位的错误。

当一个键在字典中不存在，使用中括号语法去获取值，不会产生任何错误，语句会直接返回值类型所对应的零值，例如：

```go
m := map[string]int {
    "Alice": 0,
    "Bob": 1,
}
// 输出0
fmt.Println(m["John"])
```

这往往不是你希望的结果，实际上中括号取值语句包含两个输出，第一个就是前面看到的字典值，第二个是一个成功标志，类型是bool，如果标志为true，表示获取成功，如果标志为false，则表示获取失败，该key在字典中不存在。因此，在map中获取字典值最保险的方式应该是：

```go
m := map[string]int {
    "Alice": 0,
    "Bob": 1,
}
// 由于John在字典中不存在，因此下面语句没有输出
if e, ok := m["John"]; ok {
    fmt.Println(e)
}
```

> 当你使用map的时候，请牢记这一点，因为如果你没有判断获取状态，或者说你没有去取状态标志，那么程序不会发生运行期错误，但是可能会产生莫名其妙的逻辑错误。Go编译器甚至都不会给出一个警告（这也是Rust爱好者喜欢吐槽Go编译器的一点，太过于仁慈）。

> 上面的if语法可能会让其他类C语言的开发人员觉得奇怪，因为其他语言当中都没有这种语法。这是Go特有的赋值加分支语法，可以在分支判断条件之前加一个赋值语句，然后使用分号和判断条件分隔开。这种语法的好处是在不降低代码可读性的前提下减少语句行数，并且能够将变量作用范围控制在最小的语句块当中。

## 删除字典中的键值项

使用内建的delete函数可以从字典中删除一个key所对应的项目，例如：

```go
m := map[int]bool {
    1: true,
    2: true,
}
delete(m, 1)
// 输出map[2:true]
fmt.Println(m)
```

如果第二个参数key在字典中不存在，delete函数就相当于没有进行任何操作一样，这点要注意，delete函数不会返回操作是否成功的标志。

## 字典的长度和容量

使用内建的len函数即可求出map的长度，使用方法与使用len函数求slice的长度是一样的。但是，你不能对map使用cap函数，意思即map并没有当前容量的概念，因为map的容量是可动态伸缩的，所以如果你试图使用cap函数求map容量是，编译器会拒绝通过。

```go
m := map[int]bool {
    1: true,
    2: true,
}
m[10] = false
// 输出3
fmt.Println(len(m))
```

## 字典循环迭代

range关键字同样作为用来对map进行循环迭代的方式，与slice的range稍微有点区别的是，切片的range返回的两个值分别是元素序号和元素值，而字典的range返回的两个值分别是字典的key和字典的值。例如：

```go
m := map[int]bool {
    1: true,
    2: true,
}
// 逐行输出字典中的键值对，直到迭代完成整个字典，但输出顺序不确定
for k, v := range m {
    fmt.Printf("%d -> %v\n", k, v)
}
```

值得强调的是，循环迭代字典的时候，元素的顺序是不确定的，这个不确定不是某种确定的乱序，而是每次迭代产生的顺序都可能会不一样，这是Go的运行时`runtime`故意为之的。如果需要对字典的键或值进行排序的话，需要将它们保存到一个切片中，再按照[EP.X 接口](Episode.X.Interface.md)中介绍的方法进行排序。

## Set集合

Go没有内置实现的set集合类型，但可以使用map间接实现set，具体可以有两种实现方法。

- 使用一个值为bool类型的map实现set集合。这也是官方推荐的方法，将一个键key加入set时，可以将字典中该key的值设置为true。当判断一个键key是否已经在该set中存在时，只需要读取该key对应的字典值，然后判断是否为true即可。因为如果该key在map中不存在，取值语句将会返回bool类型的零值，即false。这种方法简单易用，但会产生额外的内存开销。

```go
s := map[string]bool {
    "hello": true,
    "world": true,
}
// 输出键不存在的信息
if !s["not-exists"] {
    fmt.Println("key does not exist in set")
}
```

- 第二种方法是选择struct{}空结构体作为值的类型，因为空结构体是不占任何内存的，因此这个方案能节省内存开销。但是，判断键key是否在set里面存在就需要使用到取字典值时返回的第二个参数获取状态标志。程序逻辑有点别扭。（详情参见[EP.VI 结构体](Episode.VI.Struct.md)）

```go
s := map[string]struct{} {
    // 这里的struct{}是类型，后面的{}是初始化结构，表示空的结构体
    "hello": struct{}{},
    "world": struct{}{},
}
// 输出键不存在的信息
if _, ok := s["not-exists"]; !ok {
    fmt.Println("key does not exist in set")
}
```

> 在Go中，单一的下划线作为变量名称有特殊的函数，相当于明确地告诉Go编译器该变量的值将被丢弃，因为Go代码中不允许含有声明了却未使用的变量（参见[EP.I 类型、变量和常量](Episode.I.Type.md)）。因此如果该变量确实后续无须用到，则可以使用下划线变量来承载，变量值将直接被丢弃。如果你熟悉*nix操作系统的话，你可以认为Go中的下划线变量就像系统的`/dev/null`设备一样。

内存空间不是特别紧张的情况下，尽量采用第一种方案来实现。

[例子 字典](examples/ep04/map_type.go)

[EP.III 字符串](Episode.III.String.md) <|> [EP.V 函数](Episode.V.Function.md)
