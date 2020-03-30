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
	// 略... 与上例代码一致
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
```

[运行这个例子](https://goplay.space/#KlpBZyaesZI)

观察上面代码，可以看到有两个修改，第一个就是函数的签名，我们在函数名称前面增加了一个参数`(m MyStruct)`，或者更准确来说，我们将原来函数中的第一个参数移到了名称前面。正如前面我们说的方法从属于类一样，这里我们可以认为`String`函数从属于`MyStruct`结构体，不过在Go语言中，一般不描述这种从属关系，我们将函数名称前面的参数叫做接收者（`Receiver`），因此这里我们将`String`函数定义成接收者为`MyStruct`变量的一个方法。第二个修改是输出中，我们不再显式的调用`String`方法了，这里会由Go编译器帮我们进行隐式的调用（参见[EP.X 接口和错误](Episode.X.Interface.md)）。

所以我们可以首先知道的是，Go语言通过带接收者的函数间接实现了OO当中的方法，所以也可以认为Go中的方法是一个特殊的函数。

## 方法定义

Go中方法的定义语法是：

```go
func (r ReceiverType) FuncName([params, ...]) (returns, ...)
```

可以看到，如果将函数名称前面的接收者部分去掉，这个语法与函数定义语法完全一模一样。因此如果将接收者参数放入参数列表称为第一个参数的话，方法定义语法实际上就蜕变成了函数定义语法。

注意上面的`ReceiverType`，这里指的是接收者类型，实际上可以是除了标准类型外的所有合法类型。例如：

```go
type UNIXTimestamp int64
func (u UNIXTimestamp) AsMillis() int64 {
    // ...
}
```

当然还可以是上述所有合法类型的指针类型：

```go
type Coord complex64
func (c *Coord) ResetOrigin() {
    // ...
}
```

然后在方法体内，接收者参数就可以和普通参数一样使用了，除了方法定义上的区别外，方法中的接收者参数与其他普通参数并没有任何地位上的区别。在定义方法的位置上，Go是非常自由的，没有从语法结构中体现出方法与类型的从属关系。具体来讲，就是方法不需要定义在一个固定的类型从属代码结构中，甚至可以不在类型定义的同一个Go源代码文件当中，只需要保证定义在与接收者类型处于同一个包（`package`）当中即可。

> 对比其他语言，我们可以看到Go的方法和其他语言有一定设计概念上的区别：
> - Go当中没有`this`、`self`的关键字，进一步弱化从属关系。接收者就像普通参数那样是具有形参命名的，实际上也更容易保证类型安全。
> - 很显然，Go的方法也没有默认对接收者源数据的修改操作，因为接收者参数就像普通参数一样工作，因此“永远传值”的策略依然有效。你无法通过修改接收者参数的副本来修改原值，**如果你需要修改接收者参数源数据，请传递指针类型**。

## 构造方法

绝大部分面向对象语言都有“构造方法”的概念，用来对一个对象实例进行初始化。例如，Java中的“同名无返回值方法”，Python中的`__new__`和`__init__`，等。当对象实例化时，这些方法会显式或者隐式的进行调用，完成新对象的初始化过程。

然而，Go语言中并没有对象的概念，最接近对象概念的应该就是结构体变量了。结构体可以有多种的初始化方式（参见[EP.VI 结构体](Episode.VI.Struct.md)）。因此在一些情况下，可以认为Go语言中并不需要所谓的“构造方法”。但是也有例外，看下面的结构体定义：

```go
type User struct {
    ID int
    Username string
    signOnTime time.Time
}
```

假设说这个结构体定义在`user`包中，而我们需要在另外一个包（例如`main`包）中使用它。我们会发现这里有个问题，因为`signOnTime`字段在其他包当中是不可见的（参见[EP.XIII 可见性](Episode.XIII.Visibility.md)），因此我们之前的结构体初始化方法就无法设置这个字段的值，假设我们希望`signOnTime`字段初始值为当前系统时间的话，就有些麻烦了：

```go
func main() {
    u := User {
        ID: 9999,
        Username: "John Doe",
        // 当结构体处于不同的包时，下面代码会报错
        signOnTime: time.Now(),
    }
    // 同样的原因，下面代码编译出错
    u.signOnTime = time.Now()
}
```

在这种情况下，有两种方案可以解决，第一种是提供一个包外可见的方法来设置不可见字段的值：

```go
func (u *User) SignOnNow() {
    u.signOnTime = time.Now
}

func main() {
    u := User {
        ID: 9999,
        Username: "John Doe",
    }
    u.SignOnNow()
}
```

另一种方案是为类型`User`提供一个“构造方法”，但是实际上并不是我们通常说的那种构造方法，一般来说是通过在包`package`当中实现一个函数来替代。按照惯例，我们通常将这个函数命名为`New`：

```go
type User struct {
    ID int
    Username string
    signOnTime time.Time
}

func New(id int, username string) *User {
    return &User {
        ID: id,
        Username: username,
        signOnTime: time.Now(),
    }
}
```

上面代码在`User`所在的包中创建了一个函数`New`（注意不是方法，因此似乎应该在Go中被称为构造函数更为准确），该函数接受`id`和`username`两个参数用来初始化结构体，在函数中将`signOnTime`字段初始化成了当前系统时间，函数返回`User`结构体的指针类型。这里既可以返回原始类型也可以返回指针类型，取决于实际情况，特别是如果结构体数据较大的情况下，指针类型是一个更好的选择。然后我们可以在其他包当中调用这个函数初始化`User`结构体：

```go
import (
    "fmt"
    "xxx/user"
)

func main() {
    u := user.New(9999, "John Doe")
    fmt.Println(u)
}
```

我们看到`user`包中的`New`函数实际上起到了`User`类型的构造方法的作用，在其他包当中都可以使用这个函数来对`User`结构体进行初始化。这其中最常见的应用应该是错误的初始化了：

```go
err := errors.New("未知错误")
```

如果在一个包中有多个类型需要自定义初始化函数（构造方法）的话，可以采用`NewType`作为函数名称，例如上面的`User`类型：

```go
func NewUser(id int, username string) User {
    return User {
        ID: id,
        Username: username,
        signOnTime: time.Now(),
    }
}
```

## 属性的Getter和Setter

OO的概念当中，封装是比较重要的一点。对象的属性应该被封装在对象内部，对象外部的使用者应该通过对应的API对其进行访问或设置，这也是我们经常在其他面向对象语言（例如Java）中看到的对象属性的Getter和Setter。Go在这点上也是惯例性的弱化，意思就是没有特殊情况下，能直接访问的就直接访问，不需要专门编写所谓的Getter和Setter，但是如果确实有这种需要（比方说按照一定规则对字段进行访问），也可以构建相应的Getter和Setter。不过惯例上Getter的命名不要带上Get动词，只需要使用字段名称作为方法名称即可；而Setter的命名就是Set加上字段名称：

```go
// User 用户例子结构体定义
type User struct {
	ID    int
	Name  string
	email string // email字段为包外不可见
}

// Email email字段的Getter
func (u *User) Email() string {
	return u.email
}

// SetEmail email字段的Setter，使用正则表达式验证地址合法性
func (u *User) SetEmail(e string) error {
	r, err := regexp.Compile(`[^@ \t\r\n]+@[^@ \t\r\n]+\.[^@ \t\r\n]+`)
	if err != nil {
		return err
	}
	if !r.Match([]byte(e)) {
		return errors.New("不是合法的电子邮件地址")
	}
	u.email = e
	return nil
}

func main() {
	u := User{
		ID:   9999,
		Name: "John Doe",
	}
	// 设置正确的电子邮件地址
	err := u.SetEmail("johndoe@gmail.com")
	if err != nil {
		panic(err)
	}
	// 输出{9999 John Doe johndoe@gmail.com}
	fmt.Println(u)
	// 设置错误的电子邮件地址
	err = u.SetEmail("non-exist-email")
	if err != nil {
		// panic: 不是合法的电子邮件地址
		panic(err)
	}
}
```

[运行这个例子](https://goplay.space/#WfLJxVgdedR)

上面的例子中我们将`User`结构体的`email`字段设置为包外不可见，因此，如果在其他包当中需要访问或设置`email`字段，就需要用到该字段的Getter和Setter，此处分别为方法`Email`和`SetEmail`。其中Getter只是简单返回了字段的值，而Setter稍微复杂一点，使用了正则表达式来验证电子邮件的正确性，如果不正确会返回一个错误。

## 方法重载

许多OOP支持方法重载（Overload），例如在Java中，经常会写出如下的代码：

```java
public class User {
    // 属性定义，略
    public int getUserId() {
        // 方法体略
    }
    public int getUserId(String prefix) {
        // 方法体略
    }
}
```

但是在类C的语言中，这种做法是错误的，例如Python，如下的定义不会出错，但在运行时会出现诡异的情况：

```python
class User():
    def method(self):
        pass
    def method(self, prefix):
        pass
u = User()
# 调用时出错，TypeError: method() missing 1 required positional argument: 'prefix'
u.method()
```

这表明python只是简单的将后面定义的方法覆盖了前面的方法定义。事实上，OOP的这种重载特性也是被吐槽的重点，因为这是代码爆炸的主要原因之一，能大幅降低代码的可读性。通常如果一个类有很多的方法重载的话，需要考虑两点可能性：一就是这个对象的抽象本身有问题，应该拆分，或者更学术化的说是违反了抽象对象的正交性设计；二就是过度设计，为了程序员的自我感觉而编写的代码，哪怕可能永远都用不上。这样的代码会导致后续的维护非常困难，一定要记住：**代码不是写给机器看的，代码是写给未来的人看的，甚至包括你自己**。

Go从编译器层面上杜绝了这种情况，也就是说，Go不允许出现方法重载，如果一个类型中有两个同名方法的话，Go编译器会拒绝通过。

那么假设我们有一个需求，需要根据类型的实际情况调用不同的方法来实现不同的逻辑该怎么办。Go有一种更直接更优雅的方式来实现这个需求。假设我们的`User`用户结构体可以接受用户名或电子邮件注册，然后我们根据用户注册的不同方式，要对相应的字段进行检查。与其写两个检查的方法，然后在使用到的地方进行硬编码调用，不如将检查的方法放置在结构体中，任何需要检查的地方统一进行调用即可：

```go
// User 用户例子结构体定义
type User struct {
	ID            int
	Username      string
	Email         string
	validateField func(u *User) bool // 这个字段保存着该用户字段验证的相应函数
}

// SetValidateField 根据User结构体实际情况确定其使用的字段验证函数
// 注意此处User结构体更加适合使用构造函数而不是Setter
func (u *User) SetValidateField() {
	if u.Username != "" {
		u.validateField = validateUsername
		return
	}
	if u.Email != "" {
		u.validateField = validateEmail
		return
	}
	u.validateField = func(u *User) bool {
		return false
	}
}

// ValidateField 将字段验证函数封装成结构体的方法
func (u *User) ValidateField() bool {
	return u.validateField(u)
}

// 验证用户名正确性的帮助函数
func validateUsername(u *User) bool {
	r, err := regexp.Compile(`^[a-z0-9_-]{3,15}$`)
	if err != nil {
		return false
	}
	return r.Match([]byte(u.Username))
}

// 验证电子邮件地址正确性的帮助函数
func validateEmail(u *User) bool {
	r, err := regexp.Compile(`[^@ \t\r\n]+@[^@ \t\r\n]+\.[^@ \t\r\n]+`)
	if err != nil {
		return false
	}
	return r.Match([]byte(u.Email))
}

func main() {
	u := User{
		ID:       9999,
		Username: "john doe",
	}
	// 设置相应的用户名验证函数
	u.SetValidateField()
	if !u.ValidateField() {
		fmt.Printf("用户名设置错误: %s\n", u.Username)
	} else {
		fmt.Println("用户名验证通过")
		fmt.Println(u)
	}
	u = User{
		ID:    99999,
		Email: "john.doe@gmail.com",
	}
	// 设置相应的电子邮件地址验证函数
	u.SetValidateField()
	if !u.ValidateField() {
		fmt.Printf("电子邮件设置错误: %s\n", u.Email)
	} else {
		fmt.Println("电子邮件验证通过")
		fmt.Println(u)
	}
}
```

[运行这个例子](https://goplay.space/#pQ5E6FZDYIg)

上面的例子有点长，但是建议读者将它仔细读完。这里的技巧就是将验证结构体字段正确性的函数作为结构体的字段保存下来，回忆之前我们说过，结构体字段可以是任何类型，当然也可以是函数类型。然后在验证时只需要调用这个函数（代码中我们使用方法对其进行了封装）即可以得到验证的结果。

你可以会说，我可以写一个方法就能完成的事情，为什么要做的这么复杂呢？这里的好处有两点：

- 验证相应字段的逻辑只需要执行一次，后续无论验证多少次，都不会重新再次进行这部分的逻辑判断，因为相应的判断逻辑实际上已经记录在这个结构体中了。这在事实上实现了更好的封装。
- 当对整体结构体或判断逻辑进行扩展时，只需要修改验证字段的Setter和新增或修改相应的帮助函数，只要这个函数满足签名`func (u *User) bool`即可。这种修改不会对任何外部调用者产生影响，甚至调用者不修改原来的代码也可以应用到新的逻辑之上，这些调用包括初始化结构体，设置验证字段以及调用方法进行字段验证，所有的对外接口都是保持良好的。大大增强了代码的可拓展性和健壮性。因此这种写法是很好的一种代码实践。

本章的最后，我们用一个详细的双链表例子来说明方法在Go语言中的应用。

[例子：双链表的方法实现](examples/ep08/method_linklist.go)

> 注意：例子中的双链表代码实现不要直接用在生产中，该实现中没有考虑通用性和性能优化问题。如果在生产中需要使用双链表请使用标准库`container/list`（参见[EP.XIX 标准库](Episode.XIX.Stdlib.md)）。

[EP.VII 指针](Episode.VII.Pointer.md) <|> [EP.IX 结构体嵌套](Episode.IX.Nested.md)