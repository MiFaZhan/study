Go语言的 `io` 包是标准库中用于 I/O（输入/输出）操作的核心包。它定义了基本的 I/O 接口和若干实用函数，这些接口被整个标准库广泛实现和使用，比如文件（`os.File`）、内存缓冲（`bytes.Buffer`）、网络连接（`net.Conn`）等。掌握 `io` 包对于编写高效、可复用的 I/O 代码至关重要。

## 核心接口

`io` 包的精髓在于它定义的一组简单、正交的接口。它们抽象了 I/O 操作的共性，让不同的具体类型可以以统一的方式被处理。

- **`io.Reader`**
  ```go
  type Reader interface {
      Read(p []byte) (n int, err error)
  }
  ```
  `Read` 方法将数据读入到字节切片 `p` 中，返回实际读取的字节数 `n` 和可能的错误。当数据流结束时，返回 `io.EOF` 错误（这是一个特殊的错误，表示“文件结束”，通常不是真正的错误）。

- **`io.Writer`**
  ```go
  type Writer interface {
      Write(p []byte) (n int, err error)
  }
  ```
  `Write` 方法将字节切片 `p` 中的数据写入到输出流，返回实际写入的字节数 `n` 和可能的错误。如果 `n < len(p)`，则必须返回一个非 nil 的错误。

- **`io.Closer`**
  ```go
  type Closer interface {
      Close() error
  }
  ```
  `Close` 方法关闭资源（如文件、网络连接），释放相关资源。

- **`io.Seeker`**
  ```go
  type Seeker interface {
      Seek(offset int64, whence int) (int64, error)
  }
  ```
  `Seek` 方法用于设置下一次读/写的偏移量。`whence` 的含义：
    - `io.SeekStart`   = 0：相对于文件起始位置
    - `io.SeekCurrent` = 1：相对于当前偏移量
    - `io.SeekEnd`     = 2：相对于文件结尾

除了这四个基本接口，`io` 包还定义了一些组合接口，以便描述同时具备多种能力的对象：

- **`io.ReadWriter`** = `Reader` + `Writer`
- **`io.ReadCloser`** = `Reader` + `Closer`
- **`io.WriteCloser`** = `Writer` + `Closer`
- **`io.ReadWriteCloser`** = `Reader` + `Writer` + `Closer`
- **`io.ReadSeeker`** = `Reader` + `Seeker`
- **`io.WriteSeeker`** = `Writer` + `Seeker`
- **`io.ReadWriteSeeker`** = `Reader` + `Writer` + `Seeker`

这些组合接口使得函数可以声明更精确的参数需求，只接受那些具备所需全部能力的类型。

## 重要函数

`io` 包提供了一系列实用函数，简化了常见的 I/O 操作。

### 复制数据

- **`io.Copy(dst Writer, src Reader) (written int64, err error)`**  
  将 `src` 中的数据复制到 `dst`，直到 `src` 返回 EOF 或发生错误。它使用一个默认的缓冲区（32KB）进行复制，效率较高。

- **`io.CopyN(dst Writer, src Reader, n int64) (written int64, err error)`**  
  从 `src` 复制恰好 `n` 个字节到 `dst`。如果复制的字节数小于 `n`，会返回错误（除非是 EOF）。

- **`io.CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error)`**  
  与 `Copy` 类似，但允许自定义缓冲区。如果 `buf` 为 `nil`，则会内部分配一个。

### 读取特定数据

- **`io.ReadFull(r Reader, buf []byte) (n int, err error)`**  
  从 `r` 精确地读取 `len(buf)` 个字节填充到 `buf`。如果读取的字节数不足，返回 `io.ErrUnexpectedEOF`。

- **`io.ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)`**  
  至少读取 `min` 个字节，但最多读取 `len(buf)` 个。如果读取的字节数少于 `min`，返回 `io.ErrUnexpectedEOF`（除非是 EOF）。

- **`io.ReadAll(r Reader) ([]byte, error)`**  
  从 `r` 读取所有数据直到 EOF，返回读取的字节切片。此函数原在 `io/ioutil` 包，Go 1.16 后移入 `io` 包。

### 写入字符串

- **`io.WriteString(w Writer, s string) (n int, err error)`**  
  将字符串 `s` 写入 `w`。如果 `w` 实现了 `io.StringWriter` 接口（即拥有 `WriteString` 方法），则直接调用该方法，否则通过转换为 `[]byte` 再调用 `Write`。

### 组合/拆分流

- **`io.MultiReader(readers ...Reader) Reader`**  
  返回一个逻辑上的 Reader，它会按顺序从提供的多个 Reader 中读取数据。

- **`io.MultiWriter(writers ...Writer) Writer`**  
  返回一个 Writer，向其写入的数据会同时写入所有提供的 Writer（类似于 `tee` 命令）。

- **`io.TeeReader(r Reader, w Writer) Reader`**  
  返回一个 Reader，从它读取数据时会同时将数据写入 `w`（类似于 `tee` 命令）。常用于调试或数据校验。

### 限制读取

- **`io.LimitReader(r Reader, n int64) Reader`**  
  返回一个 Reader，它从 `r` 读取，但只能读取最多 `n` 个字节。之后返回 EOF。

## 特殊接口

- **`io.ReaderFrom`**
  ```go
  type ReaderFrom interface {
      ReadFrom(r Reader) (n int64, err error)
  }
  ```
  如果某个类型（如 `bytes.Buffer`）实现了 `ReadFrom`，那么 `io.Copy` 会优先调用它，以可能实现更高效的直接复制（例如直接写入底层数组）。

- **`io.WriterTo`**
  ```go
  type WriterTo interface {
      WriteTo(w Writer) (n int64, err error)
  }
  ```
  类似地，如果 `src` 实现了 `WriteTo`，`io.Copy` 会调用它来将数据直接写入 `dst`，避免中间缓冲区。

- **`io.ByteReader`** 和 **`io.ByteWriter`**  
  分别定义 `ReadByte() (byte, error)` 和 `WriteByte(c byte) error` 方法，用于逐字节操作。

- **`io.RuneReader`**  
  定义 `ReadRune() (r rune, size int, err error)`，用于读取 Unicode 字符。

## 错误处理

`io` 包定义了一些标准错误：

- **`io.EOF`**：当 `Read` 无法再返回数据时返回，表示输入流已结束。这通常不是错误，而是正常终止条件。
- **`io.ErrUnexpectedEOF`**：在 `ReadFull` 或 `ReadAtLeast` 中，如果读取的字节数少于期望值但又不是因为 EOF 导致的，则返回此错误。
- **`io.ErrShortWrite`**：当 `Write` 写入的字节数少于提供的长度且没有返回显式错误时，`io` 包中的某些函数可能会返回此错误。
- **`io.ErrClosedPipe`**：对已关闭的管道进行读写时返回。
- **`io.ErrNoProgress`**：某些循环复制中，如果连续多次没有读写任何字节，可能返回此错误以避免死循环。

## 与其他包的协同

`io` 包的接口被广泛实现，因此你可以用统一的方式处理各种 I/O 资源：

- **文件**：`os.File` 实现了 `Reader`、`Writer`、`Closer`、`Seeker` 等。
- **内存缓冲区**：`bytes.Buffer` 实现了 `Reader`、`Writer`、`ReaderFrom`、`WriterTo` 等；`strings.Reader` 实现了 `Reader`、`Seeker` 等。
- **网络连接**：`net.Conn` 实现了 `Reader`、`Writer`、`Closer`。
- **压缩/加密流**：`compress/gzip.Reader`/`Writer`、`crypto/cipher.StreamReader`/`StreamWriter` 都基于 `io.Reader`/`Writer` 构建。
- **HTTP 响应体**：`http.Response.Body` 是一个 `io.ReadCloser`。

这种基于接口的设计使得 Go 的 I/O 操作非常灵活，你可以轻松地将数据从文件复制到网络连接，或从 HTTP 响应读取并写入压缩流，而无需关心底层具体类型。

## 示例

```go
package main

import (
    "bytes"
    "fmt"
    "io"
    "os"
    "strings"
)

func main() {
    // 1. 使用 Reader 和 Writer
    r := strings.NewReader("Hello, Go!")
    var buf bytes.Buffer
    n, err := io.Copy(&buf, r)
    if err != nil {
        fmt.Println("copy error:", err)
        return
    }
    fmt.Printf("copied %d bytes: %s\n", n, buf.String())

    // 2. 使用 MultiWriter 同时写入文件和控制台
    file, _ := os.Create("test.txt")
    defer file.Close()
    mw := io.MultiWriter(os.Stdout, file)
    io.WriteString(mw, "This goes to both stdout and file.\n")

    // 3. 使用 LimitReader 限制读取长度
    limited := io.LimitReader(strings.NewReader("1234567890"), 5)
    data, _ := io.ReadAll(limited)
    fmt.Printf("limited read: %s\n", data) // 输出 "12345"

    // 4. 使用 TeeReader 同时读取并写入
    var teeBuf bytes.Buffer
    tee := io.TeeReader(strings.NewReader("tee example"), &teeBuf)
    io.Copy(os.Stdout, tee) // 输出到 stdout，同时 teeBuf 也会得到数据
    fmt.Printf("\ncaptured in buffer: %s\n", teeBuf.String())
}
```

## 总结

`io` 包通过简单而强大的接口抽象了所有 I/O 操作，是 Go 语言中组合和泛化能力的典范。理解并熟练使用这些接口和函数，能够帮助你写出更简洁、更通用、更易于测试的 I/O 相关代码。