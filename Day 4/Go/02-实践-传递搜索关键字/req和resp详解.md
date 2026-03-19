
## req 和 resp 的含义

### req (Request)
**req** 是 **Request** 的简写，表示 **HTTP请求对象**。

```go
req, err := http.NewRequest("GET", fullUrl, nil)
```
- `req` 是一个 `*http.Request` 类型的指针
- 包含了请求的所有信息：URL、方法、Header、Body等
- 可以自定义请求的各种参数

### resp (Response)
**resp** 是 **Response** 的简写，表示 **HTTP响应对象**。

```go
resp, err := client.Do(req)
```
- `resp` 是一个 `*http.Response` 类型的指针
- 包含了服务器返回的所有信息：状态码、Header、Body等
- 通过它可以获取服务器的响应数据

---

## http.Request 结构体详解

`http.Request` 是 Go 标准库中定义的结构体，包含了 HTTP 请求的所有信息。

### 常用字段

| 字段 | 类型 | 说明 |
|------|------|------|
| `Method` | string | HTTP 方法，如 "GET"、"POST"、"PUT"、"DELETE" |
| `URL` | *url.URL | 请求的完整 URL 对象 |
| `Proto` | string | HTTP 协议版本，如 "HTTP/1.1" |
| `Header` | http.Header | 请求头，本质是 `map[string][]string` |
| `Body` | io.ReadCloser | 请求体，用于 POST/PUT 等携带数据 |
| `ContentLength` | int64 | 请求体的长度 |
| `Host` | string | 目标主机名 |
| `RemoteAddr` | string | 客户端地址（服务端使用） |
| `Cookies()` | []*Cookie | 获取请求中的所有 Cookie |

### 创建请求的几种方式

```go
// 方式1：GET 请求，无请求体
req, err := http.NewRequest("GET", "https://example.com", nil)

// 方式2：POST 请求，带 JSON 数据
body := strings.NewReader(`{"name":"golang"}`)
req, err := http.NewRequest("POST", "https://example.com/api", body)

// 方式3：POST 请求，带表单数据
form := url.Values{}
form.Add("username", "admin")
form.Add("password", "123456")
req, err := http.NewRequest("POST", "https://example.com/login", strings.NewReader(form.Encode()))
```

### 设置请求头

```go
req.Header.Set("Content-Type", "application/json")
req.Header.Set("User-Agent", "MyApp/1.0")
req.Header.Set("Authorization", "Bearer token123")

// 添加同名字段（追加而非覆盖）
req.Header.Add("Accept", "application/json")
req.Header.Add("Accept", "text/html")

// 获取请求头
ua := req.Header.Get("User-Agent")
```

---

## http.Response 结构体详解

`http.Response` 是 Go 标准库中定义的结构体，包含了 HTTP 响应的所有信息。

### 常用字段

| 字段 | 类型 | 说明 |
|------|------|------|
| `StatusCode` | int | HTTP 状态码，如 200、404、500 |
| `Status` | string | 状态文本，如 "200 OK" |
| `Proto` | string | HTTP 协议版本 |
| `Header` | http.Header | 响应头 |
| `Body` | io.ReadCloser | 响应体，需要手动关闭 |
| `ContentLength` | int64 | 响应体的长度 |
| `Cookies()` | []*Cookie | 获取响应中的所有 Cookie |

### 常见状态码

| 状态码 | 常量 | 含义 |
|--------|------|------|
| 200 | http.StatusOK | 成功 |
| 301 | http.StatusMovedPermanently | 永久重定向 |
| 302 | http.StatusFound | 临时重定向 |
| 400 | http.StatusBadRequest | 请求错误 |
| 401 | http.StatusUnauthorized | 未授权 |
| 403 | http.StatusForbidden | 禁止访问 |
| 404 | http.StatusNotFound | 未找到 |
| 500 | http.StatusInternalServerError | 服务器内部错误 |

### 读取响应

```go
resp, err := client.Do(req)
if err != nil {
    panic(err)
}
defer resp.Body.Close()  // 重要：必须关闭，防止资源泄漏

// 检查状态码
if resp.StatusCode == http.StatusOK {
    fmt.Println("请求成功")
}

// 读取响应体
body, err := io.ReadAll(resp.Body)
if err != nil {
    panic(err)
}
fmt.Println(string(body))

// 获取响应头
contentType := resp.Header.Get("Content-Type")
server := resp.Header.Get("Server")
```

---

## 完整流程示意

```
┌──────────────────────────────────────────────────────────────┐
│                         客户端                               │
│                                                              │
│  1. 创建 Request                                             │
│     req, _ := http.NewRequest("GET", url, nil)              │
│                                                              │
│  2. 设置请求头 (可选)                                        │
│     req.Header.Set("User-Agent", "...")                     │
│                                                              │
│  3. 发送请求                                                 │
│     resp, _ := client.Do(req)                               │
│                                                              │
│  4. 处理 Response                                            │
│     - 检查 resp.StatusCode                                   │
│     - 读取 resp.Body                                         │
│     - defer resp.Body.Close()                               │
│                                                              │
└──────────────────────────────────────────────────────────────┘
                               ↕
                         HTTP 协议通信
                               ↕
┌──────────────────────────────────────────────────────────────┐
│                         服务器                                │
└──────────────────────────────────────────────────────────────┘
```

---

## 重要注意事项

1. **必须关闭 resp.Body**
   ```go
   defer resp.Body.Close()  // 否则会导致连接泄漏
   ```

2. **GET 请求通常没有 Body**，传 nil 即可

3. **POST 请求需要设置 Content-Type**
   ```go
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   // 或
   req.Header.Set("Content-Type", "application/json")
   ```

4. **Header 是 map[string][]string**，同一个 key 可以有多个值

---

## 类比理解

可以把它们想象成：
- **req** = 你写给服务器的信（请求）
- **resp** = 服务器回给你的信（响应）

这两个变量名是 Go HTTP 编程中**约定俗成**的简写，几乎所有 Go 开发者都能立即理解它们的含义。
        