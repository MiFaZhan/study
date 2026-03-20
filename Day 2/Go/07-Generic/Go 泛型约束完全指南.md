
# Go 泛型约束完全指南

## 一、约束类型

### 1. 内置约束

| 约束 | 说明 | 支持的类型 |
|------|------|-----------|
| `any` | 任意类型 | 任何类型都可以 |
| `comparable` | 可比较类型 | 支持 `==` 和 `!=` 的类型 |
| `integer` | 整数类型 | `int`, `int8`, `int16`, `int32`, `int64` |
| `unsigned` | 无符号整数 | `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr` |
| `float` | 浮点数 | `float32`, `float64` |
| `complex` | 复数 | `complex64`, `complex128` |
| `bool` | 布尔值 | `bool` |
| `string` | 字符串 | `string` |

### 2. 自定义约束

```go
// 方式一：使用 interface 定义约束
type Number interface {
    int | int8 | int16 | int32 | int64 | float32 | float64
}

// 方式二：使用 ~ 允许底层类型
type MyInt int
type SignedNumber interface {
    ~int | ~int32 | ~int64   // ~ 表示底层类型
}
```

---

## 二、完整示例代码

```go
package main

import "fmt"

// ============ 约束定义 ============

// 任意类型约束
func PrintAny[T any](v T) {
    fmt.Println(v)
}

// 可比较约束（必须支持 ==）
func Find[T comparable](slice []T, target T) int {
    for i, v := range slice {
        if v == target {
            return i
        }
    }
    return -1
}

// 数字类型约束
type Number interface {
    int | int8 | int16 | int32 | int64 | float32 | float64
}

func Sum[T Number](nums []T) T {
    var total T
    for _, n := range nums {
        total += n
    }
    return total
}

// 浮点数约束
func Average[T float32 | float64](nums []T) T {
    if len(nums) == 0 {
        return 0
    }
    var sum T
    for _, n := range nums {
        sum += n
    }
    return sum / T(len(nums))
}

// 整数约束
func Max[T int | int8 | int16 | int32 | int64](a, b T) T {
    if a > b {
        return a
    }
    return b
}

// 无符号整数约束
func Double[T unsigned](n T) T {
    return n * 2
}

// 布尔和字符串约束
func ToString[T bool | string](v T) string {
    return fmt.Sprintf("%v", v)
}

// ============ 泛型结构体 ============

type Container[T any] struct {
    Value T
}

type Pair[K string, V any] struct {
    Key   K
    Value V
}

// ============ 泛型方法 ============

func (c Container[T]) Get() T {
    return c.Value
}

// ============ 泛型接口 ============

type Adder[T any] interface {
    Add(T) T
}

type IntAdder int

func (i IntAdder) Add(n int) int {
    return int(i) + n
}

type StringAdder string

func (s StringAdder) Add(n string) string {
    return string(s) + n
}

// ============ 主函数 ============

func main() {
    // any 约束
    PrintAny(123)
    PrintAny("hello")
    PrintAny([]int{1, 2, 3})

    // comparable 约束
    nums := []int{1, 2, 3, 4, 5}
    idx := Find(nums, 3)
    fmt.Printf("找到位置: %d\n", idx)

    // Number 约束
    fmt.Printf("整数和: %d\n", Sum([]int{1, 2, 3}))
    fmt.Printf("浮点求和: %.2f\n", Sum([]float64{1.1, 2.2, 3.3}))

    // 浮点数约束
    fmt.Printf("平均值: %.2f\n", Average([]float64{1.0, 2.0, 3.0}))

    // 整数约束
    fmt.Printf("最大值: %d\n", Max(10, 20))

    // 无符号约束
    fmt.Printf("倍增: %d\n", Double(uint(5)))

    // 布尔和字符串约束
    fmt.Printf("转字符串: %s\n", ToString(true))
    fmt.Printf("转字符串: %s\n", ToString("test"))

    // 泛型结构体
    c := Container[int]{Value: 100}
    fmt.Printf("Container: %d\n", c.Get())

    p := Pair[string, int]{Key: "age", Value: 25}
    fmt.Printf("Pair: %s=%d\n", p.Key, p.Value)
}
```

---

## 三、约束组合

```go
// 多个约束用 | 组合
type Numeric interface {
    int | int32 | int64 | float32 | float64
}

// 添加方法约束
type Addable[T any] interface {
    Add(T) T
}
```

---

## 四、注意事项

1. **Go 版本**：需要 Go 1.18+
2. **comparable**：切片、映射、函数、通道不支持比较
3. **性能**：泛型零运行时开销，编译时类型擦除
        