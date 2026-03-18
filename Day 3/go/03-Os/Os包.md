Go语言的 `os` 包提供了与操作系统进行交互的平台无关接口。它封装了文件系统操作、环境变量、进程控制、信号处理等功能，是标准库中最常用的包之一。下面从几个主要方面介绍 `os` 包。

---

### 1. 文件与目录操作

`os` 包提供了丰富的函数和类型来操作文件和目录。

- **打开/创建文件**
    - `os.Open(name string) (*File, error)`：以只读方式打开文件。
    - `os.Create(name string) (*File, error)`：创建或截断一个文件（如果已存在则清空），以读写方式打开。
    - `os.OpenFile(name string, flag int, perm FileMode) (*File, error)`：以更灵活的方式打开文件，可指定标志（如 `O_RDONLY`、`O_WRONLY`、`O_RDWR`、`O_APPEND`、`O_CREATE` 等）和权限。

- **读取/写入文件**  
  返回的 `*File` 类型实现了 `io.Reader` 和 `io.Writer` 接口，因此可以使用 `Read`、`Write` 方法或 `io` 包的工具函数（如 `io.Copy`）进行操作。  
  示例：
  ```go
  file, err := os.Open("test.txt")
  if err != nil {
      log.Fatal(err)
  }
  defer file.Close()
  data := make([]byte, 100)
  n, err := file.Read(data)
  ```

- **文件信息与属性**
    - `os.Stat(name string) (FileInfo, error)`：获取文件信息，返回的 `FileInfo` 接口包含 `Name()`、`Size()`、`Mode()`、`ModTime()`、`IsDir()` 等方法。
    - `os.Lstat` 类似，但不会跟随符号链接。
    - `os.Chmod(name string, mode FileMode) error`：修改文件权限。
    - `os.Chown`、`os.Chtimes` 等用于修改所有者与时间戳。

- **目录操作**
    - `os.Mkdir(name string, perm FileMode) error`：创建单个目录。
    - `os.MkdirAll(path string, perm FileMode) error`：递归创建多级目录。
    - `os.Remove(name string) error`：删除文件或空目录。
    - `os.RemoveAll(path string) error`：递归删除所有内容。
    - `os.ReadDir(name string) ([]DirEntry, error)`：读取目录返回条目列表（Go 1.16+），替代了旧的 `Readdir`。

- **重命名/移动**
    - `os.Rename(oldpath, newpath string) error`：重命名或移动文件/目录。

- **临时文件与目录**
    - `os.CreateTemp(dir, pattern string) (*File, error)`：在指定目录（或系统默认临时目录）创建临时文件。
    - `os.MkdirTemp(dir, pattern string) (string, error)`：创建临时目录。

---

### 2. 环境变量

- `os.Getenv(key string) string`：获取环境变量的值，不存在则返回空字符串。
- `os.Setenv(key, value string) error`：设置环境变量（仅影响当前进程及其子进程）。
- `os.Unsetenv(key string) error`：删除环境变量。
- `os.Environ() []string`：返回所有环境变量的副本，格式为 `"KEY=value"`。
- `os.ExpandEnv(s string) string`：将字符串中的 `${var}` 或 `$var` 替换为对应的环境变量值。

---

### 3. 进程与用户

- **进程 ID 与程序参数**
    - `os.Getpid() int`：获取当前进程 ID。
    - `os.Getppid() int`：获取父进程 ID。
    - `os.Args []string`：命令行参数，`os.Args[0]` 为程序名。

- **退出与状态**
    - `os.Exit(code int)`：以指定状态码退出程序，不会执行 defer。
    - `os.ProcAttr` 和 `os.StartProcess`：启动外部进程，通常推荐使用 `os/exec` 包简化操作。

- **用户信息**
    - `os.UserHomeDir() (string, error)`：获取当前用户的家目录（Go 1.12+）。
    - `os.UserCacheDir()`、`os.UserConfigDir()` 等获取特定配置目录。

---

### 4. 信号处理

`os` 包定义了 `os.Signal` 接口，常与 `os/signal` 包配合使用。
- 常见的信号如 `os.Interrupt`（Ctrl+C）、`os.Kill`（强制终止）在 `os` 包中作为 `os.Signal` 类型的变量存在。
- 使用 `signal.Notify` 可以捕获并处理这些信号。

---

### 5. 标准输入/输出/错误

- `os.Stdin`、`os.Stdout`、`os.Stderr` 是 `*os.File` 类型的变量，分别对应标准输入、输出和错误。
- 可以通过它们进行重定向或与其他 `io` 操作组合使用。

---

### 6. 文件路径与系统信息

- `os.PathSeparator` 和 `os.PathListSeparator`：路径分隔符（`/` 或 `\`）和路径列表分隔符（`:` 或 `;`）。
- `os.Getwd() (string, error)`：获取当前工作目录。
- `os.Chdir(dir string) error`：改变工作目录。
- `os.Hostname() (string, error)`：获取主机名。
- `os.TempDir() string`：返回用于临时文件的默认目录。

---

### 7. 错误处理

`os` 包定义了一些错误变量，便于判断操作失败的原因：
- `os.ErrExist`：文件已存在。
- `os.ErrNotExist`：文件不存在。
- `os.ErrPermission`：权限不足。
- `os.ErrInvalid`：无效参数。
- `os.IsExist(err) bool`、`os.IsNotExist(err) bool`、`os.IsPermission(err) bool` 等辅助函数。

---

### 8. 示例代码

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // 获取环境变量
    home := os.Getenv("HOME")
    fmt.Println("HOME:", home)

    // 创建目录
    err := os.MkdirAll("test/subdir", 0755)
    if err != nil {
        fmt.Fprintf(os.Stderr, "mkdir error: %v\n", err)
        os.Exit(1)
    }

    // 创建文件并写入
    file, err := os.Create("test/subdir/hello.txt")
    if err != nil {
        fmt.Fprintf(os.Stderr, "create file error: %v\n", err)
        os.Exit(1)
    }
    defer file.Close()
    file.WriteString("Hello, os package!\n")

    // 重命名文件
    os.Rename("test/subdir/hello.txt", "test/hello.txt")

    // 读取文件信息
    info, err := os.Stat("test/hello.txt")
    if err != nil {
        fmt.Fprintf(os.Stderr, "stat error: %v\n", err)
        os.Exit(1)
    }
    fmt.Printf("File size: %d bytes\n", info.Size())

    // 删除目录（递归）
    os.RemoveAll("test")
}
```

---

### 总结

`os` 包是 Go 语言与底层操作系统交互的核心工具。它通过统一的接口封装了不同操作系统之间的差异，使得开发者可以编写跨平台的文件、进程和环境操作代码。掌握 `os` 包对于系统编程、工具开发以及日常文件处理都至关重要。