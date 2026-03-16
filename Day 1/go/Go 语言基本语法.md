## Go 语言基本语法介绍

Go 语言（又称 Golang）是由 Google 开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。它的语法简洁清晰，同时支持面向对象和过程式编程风格。下面我们逐一介绍 Go 的基本语法要素。

---

### 1. 变量

Go 是静态类型语言，变量需要先声明再使用。声明方式有多种：

#### 1.1 使用 `var` 关键字
```go
var name string = "Alice"
var age int
var isStudent bool
```
- 如果声明时不赋初值，则变量会被初始化为该类型的零值（例如 `string` 零值为 `""`，`int` 为 `0`，`bool` 为 `false`）。

#### 1.2 类型推断
```go
var name = "Bob"      // 自动推断为 string
var age = 30           // 自动推断为 int
```

#### 1.3 短变量声明（只能在函数内部使用）
```go
name := "Charlie"      // 等价于 var name string = "Charlie"
age := 25
```

#### 1.4 多变量声明
```go
var x, y int = 1, 2
a, b := "hello", true
```

**注意**：Go 中不允许出现未使用的变量（编译会报错），这有助于保持代码整洁。

---

### 2. 基础数据类型

Go 内置了丰富的基础数据类型：

- **布尔型**：`bool`（`true` / `false`）
- **数字类型**：
    - 有符号整数：`int`（根据平台可能是32或64位）、`int8`、`int16`、`int32`、`int64`
    - 无符号整数：`uint`、`uint8`、`uint16`、`uint32`、`uint64`
    - 浮点数：`float32`、`float64`
    - 复数：`complex64`、`complex128`
- **字符串**：`string`，Go 中字符串是 UTF-8 编码的不可变字节序列。
- **字节和符文**：
    - `byte` 是 `uint8` 的别名，表示一个 ASCII 字符。
    - `rune` 是 `int32` 的别名，表示一个 Unicode 码点。

**示例**：
```go
var flag bool = true
var num int = 42
var pi float64 = 3.14159
var s string = "Hello, 世界"
var b byte = 'A'
var r rune = '中'
```

---

### 3. 复合类型

#### 3.1 数组（固定长度）
```go
var arr [3]int = [3]int{1, 2, 3}
arr2 := [...]int{4, 5, 6} // 自动推断长度
```
- 数组长度是类型的一部分，不可改变。

#### 3.2 切片（动态数组）
切片是对数组的抽象，更灵活常用。
```go
// 创建切片
var s []int                 // nil 切片
s1 := []int{1, 2, 3}        // 字面量创建
s2 := make([]int, 5, 10)    // 长度5，容量10
s3 := arr[1:4]              // 基于数组或切片创建

// 操作
s = append(s, 4)            // 追加元素
length := len(s)             // 长度
capacity := cap(s)           // 容量
```
- 切片是引用类型，指向底层数组。

#### 3.3 映射（Map）
键值对集合，使用 `make` 或字面量创建。
```go
// 声明并初始化
var m map[string]int         // nil map，不能直接赋值
m = make(map[string]int)
m["apple"] = 5

// 字面量
m2 := map[string]int{
    "banana": 3,
    "orange": 7,
}

// 取值与判断存在
value, ok := m2["banana"]    // ok 为 true 表示键存在
delete(m2, "orange")          // 删除键
```
- map 是引用类型，遍历顺序不固定。

#### 3.4 结构体（struct）
将多个字段组合成一种新的类型。
```go
type Person struct {
    Name string
    Age  int
}

// 初始化
p1 := Person{"Alice", 30}               // 顺序初始化（不推荐）
p2 := Person{Name: "Bob", Age: 25}       // 推荐，字段名指定
p3 := new(Person)                         // 返回指针，字段为零值
p3.Name = "Charlie"

// 访问字段
fmt.Println(p1.Name)
```

- 结构体是值类型，赋值或传参会拷贝整个结构体。

#### 3.5 接口（interface）
接口定义了一组方法签名，类型通过实现所有方法来隐式实现接口。
```go
type Animal interface {
    Speak() string
}

type Dog struct{}
func (d Dog) Speak() string {
    return "Woof!"
}

type Cat struct{}
func (c Cat) Speak() string {
    return "Meow"
}

func MakeSound(a Animal) {
    fmt.Println(a.Speak())
}
```
- 接口值可以持有任意实现了该接口的具体类型的值。
- 空接口 `interface{}` 可以表示任何类型（类似 `any`）。

---

### 4. 函数

函数使用 `func` 关键字定义。

#### 4.1 基本定义
```go
func add(x int, y int) int {
    return x + y
}
```
- 参数类型相同时可省略前面的类型：`func add(x, y int) int`

#### 4.2 多返回值
```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```
- 这是 Go 常用的错误处理模式。

#### 4.3 命名返回值
可以在函数签名中给返回值命名，函数体内直接使用这些变量，最后通过 `return` 返回。
```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return  // 裸返回，返回 x 和 y
}
```

#### 4.4 函数作为值
函数是一等公民，可以赋值给变量，作为参数或返回值。
```go
func main() {
    add := func(a, b int) int {
        return a + b
    }
    result := add(3, 4)
}
```

#### 4.5 可变参数
```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
// 调用：sum(1,2,3)
```

---

### 5. 指针

Go 支持指针，但限制了指针运算（不能进行算术操作）。

#### 5.1 声明和使用
```go
var x int = 10
var p *int = &x   // p 是指向 x 的指针
fmt.Println(*p)   // 解引用，输出 10
*p = 20           // 通过指针修改 x
```

#### 5.2 函数传递指针
```go
func zeroVal(val int) {
    val = 0
}
func zeroPtr(ptr *int) {
    *ptr = 0
}
func main() {
    x := 5
    zeroVal(x)     // x 仍然是 5
    zeroPtr(&x)    // x 变为 0
}
```
- 使用指针可以在函数内部修改外部变量的值。

#### 5.3 new 函数
`new(T)` 创建类型 T 的零值，并返回其指针。
```go
p := new(int)   // *int 类型的指针，指向零值 0
```

---

### 6. 流程控制

#### 6.1 if 语句
Go 中的 `if` 可以包含一个初始化语句（作用域仅限于该 `if` 块）。
```go
if score := 85; score >= 60 {
    fmt.Println("及格")
} else {
    fmt.Println("不及格")
}
// 这里不能访问 score
```
- 条件表达式不需要括号，但执行体必须有大括号。

#### 6.2 for 循环
Go 只有 `for` 这一种循环结构，但能实现多种形式。

**标准三部分循环**：
```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

**类似 while 的循环**：
```go
sum := 1
for sum < 1000 {
    sum += sum
}
```

**无限循环**：
```go
for {
    // 无限循环，通常配合 break 使用
}
```

**range 子句**（遍历数组、切片、map、字符串等）：
```go
nums := []int{2, 4, 6}
for index, value := range nums {
    fmt.Println(index, value)
}
for key, value := range map[string]int{"a": 1, "b": 2} {
    fmt.Println(key, value)
}
for i, ch := range "hello" {   // ch 是 rune 类型
    fmt.Println(i, ch)
}
```
- 如果不需要索引或键，可以使用下划线 `_` 忽略。

#### 6.3 switch 语句
Go 的 `switch` 非常灵活，每个 `case` 默认带有 `break`，不会自动向下穿透，除非显式使用 `fallthrough`。

**基本用法**：
```go
switch day {
case "Monday":
    fmt.Println("开始工作")
case "Friday":
    fmt.Println("准备周末")
default:
    fmt.Println("普通日子")
}
```

**多条件匹配**：
```go
switch day {
case "Saturday", "Sunday":
    fmt.Println("周末")
default:
    fmt.Println("工作日")
}
```

**无表达式的 switch**（相当于 if-else 链）：
```go
score := 85
switch {
case score >= 90:
    fmt.Println("优秀")
case score >= 60:
    fmt.Println("及格")
default:
    fmt.Println("不及格")
}
```

**带初始化语句**：
```go
switch num := getNumber(); num {
case 1:
    fmt.Println("one")
case 2:
    fmt.Println("two")
}
```

#### 6.4 defer 语句
虽然不是流程控制，但常用于资源释放。`defer` 会在函数返回前执行（后进先出）。
```go
func main() {
    defer fmt.Println("world")
    fmt.Println("hello")
}
// 输出 hello world
```

---

### 7. 类型别名与自定义类型

- **类型别名**：`type MyInt = int`，只是为类型起个新名字，编译后仍是原类型。
- **自定义类型**：`type MyInt int`，创建了一个新类型，拥有不同的方法集，需要进行类型转换才能与 int 交互。

---

### 8. 方法

Go 不是纯粹的面向对象语言，但可以为类型定义方法（接收者可以是值或指针）。
```go
type Rectangle struct {
    Width, Height float64
}

// 值接收者
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// 指针接收者（可以修改结构体）
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}
```

- 调用方法时，Go 会自动处理指针和值的转换（如果接收者是指针，用值调用时也会自动取地址；反之亦然）。

---

### 总结

Go 的语法设计追求简洁和高效，没有类和继承，但通过结构体和接口实现了强大的组合能力。其内置的并发原语（goroutine 和 channel）更是语言的一大亮点。掌握上述基础语法后，就可以开始编写简单的 Go 程序了。

建议多动手练习，熟悉 Go 的工具链（如 `go fmt` 自动格式化代码、`go build` 编译等），并阅读官方文档《Effective Go》来深入理解最佳实践。