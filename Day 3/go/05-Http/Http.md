Go 语言的 `net/http` 标准库提供了强大而灵活的 HTTP 客户端和服务器实现，是构建 Web 应用程序、RESTful API 或任何 HTTP 相关服务的基础。下面我会从主要概念、核心组件到实际使用示例，为你全面介绍这个包。

---

## 一、核心概念与类型

### 1. **`http.Handler` 接口**
这是 HTTP 服务端的核心接口，定义了如何处理 HTTP 请求：
```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```
任何实现了 `ServeHTTP` 方法的对象都可以作为 HTTP 处理器（handler）。

### 2. **`http.HandlerFunc` 类型**
它是一个适配器，允许将普通函数作为 `Handler` 使用：
```go
type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}
```
这让我们可以直接用函数来处理请求。

### 3. **`http.Request` 与 `http.ResponseWriter`**
- `Request`：封装了客户端发来的请求信息（URL、头部、表单参数、请求体等）。
- `ResponseWriter`：接口用于构造响应，通过它写入数据（状态码、头部、响应体）给客户端。

### 4. **`http.ServeMux` 多路复用器**
默认的路由分发器，根据请求的 URL 路径匹配到对应的 `Handler`。通常使用 `http.DefaultServeMux` 全局实例，也可以自己创建。

### 5. **`http.Client` 与 `http.Transport`**
- `Client`：发送 HTTP 请求的客户端，可以设置超时、重定向策略等。
- `Transport`：底层连接管理（连接池、代理等），通常使用默认值即可。

---

## 二、快速上手：构建一个简单的 HTTP 服务器

下面是一个极简的 Web 服务器示例，监听 8080 端口，对根路径返回 "Hello, World!"：

```go
package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", helloHandler)      // 注册路由
    http.ListenAndServe(":8080", nil)       // 启动服务器（nil 表示使用默认多路复用器）
}
```

- `http.HandleFunc` 将函数转换为 `HandlerFunc` 并注册到默认的 `ServeMux` 中。
- `http.ListenAndServe` 启动服务器，它会持续监听请求并将请求交给对应的处理器。

---

## 三、深入请求处理

### 1. **获取请求信息**
```go
func handler(w http.ResponseWriter, r *http.Request) {
    // 请求方法
    fmt.Println("Method:", r.Method)

    // URL 路径
    fmt.Println("URL:", r.URL.Path)

    // 查询参数
    query := r.URL.Query()
    name := query.Get("name")

    // POST 表单数据（需要先解析）
    r.ParseForm()
    value := r.Form.Get("key")

    // 获取 Header
    userAgent := r.Header.Get("User-Agent")

    // 读取请求体（例如 JSON）
    body, _ := io.ReadAll(r.Body)
    defer r.Body.Close()
}
```

### 2. **构建响应**
```go
func handler(w http.ResponseWriter, r *http.Request) {
    // 设置响应状态码（默认 200）
    w.WriteHeader(http.StatusCreated)  // 201

    // 设置响应头
    w.Header().Set("Content-Type", "application/json")

    // 写入响应体
    fmt.Fprintf(w, `{"message": "ok"}`)
}
```

### 3. **使用自定义多路复用器**
避免全局默认 mux 的污染，可以自己创建：
```go
mux := http.NewServeMux()
mux.HandleFunc("/hello", helloHandler)
http.ListenAndServe(":8080", mux)
```

---

## 四、路由与路径匹配

`ServeMux` 支持两种路由模式：
- **固定路径**：如 `/hello`，只匹配该路径。
- **子树路径**：以 `/` 结尾，如 `/static/`，匹配所有以该前缀开头的路径。

**注意**：默认 mux 会将请求路径自动进行路径清理（例如将 `//` 转为 `/`），但一般简单的路由完全够用。对于复杂的路由需求（如路径参数、方法匹配），可以引入第三方路由库，如 `gorilla/mux`。

---

## 五、中间件模式

通过包装 `http.Handler` 可以实现中间件，用于日志、鉴权、恢复等：

```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}

func main() {
    handler := http.HandlerFunc(helloHandler)
    http.Handle("/", loggingMiddleware(handler))
    http.ListenAndServe(":8080", nil)
}
```

---

## 六、HTTP 客户端

使用 `http.Client` 发送请求：

```go
client := &http.Client{Timeout: 10 * time.Second}

// GET 请求
resp, err := client.Get("https://api.example.com/data")
if err != nil { ... }
defer resp.Body.Close()
body, _ := io.ReadAll(resp.Body)

// POST JSON
data := `{"name":"test"}`
resp, err = client.Post("https://api.example.com/create", "application/json", strings.NewReader(data))

// 自定义请求
req, _ := http.NewRequest("DELETE", "https://api.example.com/1", nil)
req.Header.Set("Authorization", "Bearer token")
resp, err = client.Do(req)
```

**注意**：务必关闭 `resp.Body` 以防止资源泄漏。

---

## 七、高级功能

### 1. **文件服务器**
```go
http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./assets"))))
```

### 2. **HTTPS 支持**
使用 `http.ListenAndServeTLS` 并提供证书和私钥文件：
```go
http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
```

### 3. **Cookie 处理**
```go
// 设置 Cookie
http.SetCookie(w, &http.Cookie{Name: "session", Value: "123", HttpOnly: true})

// 读取 Cookie
cookie, err := r.Cookie("session")
```

### 4. **模板渲染**
通常结合 `html/template` 包：
```go
tmpl := template.Must(template.ParseFiles("index.html"))
tmpl.Execute(w, data)
```

---

## 八、注意事项

- **并发安全**：`http.Handler` 默认会被并发调用，因此处理器内部要小心共享变量的修改（使用互斥锁或避免共享）。
- **默认客户端**：`http.DefaultClient` 没有超时设置，生产环境建议自定义带超时的 `Client`。
- **请求体关闭**：服务端在处理请求后会自动关闭请求体，但客户端需要手动关闭响应体。
- **路由优先级**：`ServeMux` 中固定路径优先级高于子树路径，相同路径下先注册的 handler 会被覆盖。

---

## 九、总结

`net/http` 包设计简洁，基于接口的组合提供了极高的灵活性。无论是快速搭建原型还是构建生产级 Web 服务，它都能胜任。掌握它的核心接口 `Handler`、请求/响应模型以及客户端的使用，是深入 Go Web 开发的必经之路。

对于更复杂的场景（比如参数校验、依赖注入、更强大的路由），可以借助社区丰富的第三方库，但标准库永远是基础与首选。