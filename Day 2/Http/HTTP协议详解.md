# HTTP 协议详解

## 1. HTTP 协议工作原理

HTTP (HyperText Transfer Protocol) 是一种应用层协议，基于 TCP/IP 通信协议来传递数据。

### 工作流程

1. **客户端发起请求**：浏览器或客户端向服务器发送 HTTP 请求
2. **建立 TCP 连接**：通过三次握手建立连接（HTTP/1.1 默认端口 80，HTTPS 默认端口 443）
3. **发送请求报文**：客户端发送包含请求方法、URL、协议版本、请求头和请求体的报文
4. **服务器处理请求**：服务器接收并解析请求，执行相应的业务逻辑
5. **返回响应报文**：服务器返回包含状态码、响应头和响应体的报文
6. **关闭连接**：完成数据传输后关闭 TCP 连接（HTTP/1.1 支持持久连接）

### 特点

- **无状态**：每次请求都是独立的，服务器不保存客户端状态（可通过 Cookie/Session 解决）
- **无连接**：每次连接只处理一个请求（HTTP/1.1 支持持久连接）
- **简单快速**：客户端只需传送请求方法和路径
- **灵活**：允许传输任意类型的数据对象

## 2. HTTP 请求方法

### 常用方法

| 方法 | 说明 | 特点 |
|------|------|------|
| **GET** | 获取资源 | 幂等、可缓存、参数在 URL 中 |
| **POST** | 提交数据 | 非幂等、不可缓存、参数在请求体中 |
| **PUT** | 更新资源（完整替换） | 幂等 |
| **DELETE** | 删除资源 | 幂等 |
| **PATCH** | 部分更新资源 | 非幂等 |
| **HEAD** | 获取响应头（不返回响应体） | 用于检查资源是否存在 |
| **OPTIONS** | 获取服务器支持的方法 | 用于 CORS 预检请求 |

### 幂等性说明

- **幂等**：多次执行相同操作，结果相同（GET、PUT、DELETE）
- **非幂等**：多次执行可能产生不同结果（POST、PATCH）

## 3. HTTP 状态码

状态码由三位数字组成，第一位定义响应类别。

### 1xx - 信息性状态码

- **100 Continue**：客户端应继续请求
- **101 Switching Protocols**：服务器正在切换协议（如升级到 WebSocket）

### 2xx - 成功状态码

- **200 OK**：请求成功
- **201 Created**：资源创建成功（常用于 POST）
- **204 No Content**：请求成功但无返回内容（常用于 DELETE）
- **206 Partial Content**：部分内容（用于断点续传）

### 3xx - 重定向状态码

- **301 Moved Permanently**：永久重定向
- **302 Found**：临时重定向
- **304 Not Modified**：资源未修改，可使用缓存

### 4xx - 客户端错误

- **400 Bad Request**：请求语法错误
- **401 Unauthorized**：未授权，需要身份验证
- **403 Forbidden**：服务器拒绝请求
- **404 Not Found**：资源不存在
- **405 Method Not Allowed**：请求方法不被允许
- **409 Conflict**：请求冲突（如资源已存在）
- **429 Too Many Requests**：请求过多，限流

### 5xx - 服务器错误

- **500 Internal Server Error**：服务器内部错误
- **502 Bad Gateway**：网关错误
- **503 Service Unavailable**：服务不可用（维护或过载）
- **504 Gateway Timeout**：网关超时

## 4. HTTP 报文组成

### 请求报文结构

```
GET /api/users HTTP/1.1              ← 请求行
Host: example.com                     ← 请求头
User-Agent: Mozilla/5.0
Content-Type: application/json
Authorization: Bearer token123
                                      ← 空行
{"name": "张三", "age": 25}          ← 请求体（可选）
```

#### 请求行

- **请求方法**：GET、POST 等
- **请求 URL**：资源路径
- **协议版本**：HTTP/1.1、HTTP/2 等

#### 请求头（Headers）

常见请求头：

- `Host`：目标服务器域名
- `User-Agent`：客户端信息
- `Accept`：可接受的响应内容类型
- `Content-Type`：请求体的数据类型
- `Authorization`：身份验证信息
- `Cookie`：客户端 Cookie
- `Referer`：来源页面
- `Connection`：连接管理（keep-alive）

#### 请求体（Body）

包含要发送的数据，常见格式：
- `application/json`：JSON 数据
- `application/x-www-form-urlencoded`：表单数据
- `multipart/form-data`：文件上传
- `text/plain`：纯文本

### 响应报文结构

```
HTTP/1.1 200 OK                       ← 状态行
Content-Type: application/json        ← 响应头
Content-Length: 45
Set-Cookie: session=abc123
Cache-Control: no-cache
                                      ← 空行
{"id": 1, "name": "张三", "age": 25} ← 响应体
```

#### 状态行

- **协议版本**：HTTP/1.1
- **状态码**：200
- **状态描述**：OK

#### 响应头（Headers）

常见响应头：

- `Content-Type`：响应体的数据类型
- `Content-Length`：响应体长度
- `Set-Cookie`：设置 Cookie
- `Cache-Control`：缓存控制
- `ETag`：资源版本标识
- `Location`：重定向地址
- `Access-Control-Allow-Origin`：CORS 跨域设置

#### 响应体（Body）

服务器返回的实际数据内容。

## 5. RESTful API

REST (Representational State Transfer) 是一种软件架构风格，用于设计网络应用程序的 API。

### RESTful 设计原则

#### 1. 资源导向（Resource-Oriented）

- 一切皆资源，每个资源有唯一的 URI
- 使用名词而非动词表示资源

```
✅ 好的设计
GET /users          # 获取用户列表
GET /users/123      # 获取特定用户
POST /users         # 创建用户
PUT /users/123      # 更新用户
DELETE /users/123   # 删除用户

❌ 不好的设计
GET /getUsers
POST /createUser
POST /deleteUser
```

#### 2. 使用标准 HTTP 方法

| 操作 | HTTP 方法 | 示例 | 说明 |
|------|-----------|------|------|
| 查询列表 | GET | `GET /users` | 获取所有用户 |
| 查询单个 | GET | `GET /users/123` | 获取 ID 为 123 的用户 |
| 创建 | POST | `POST /users` | 创建新用户 |
| 完整更新 | PUT | `PUT /users/123` | 完整替换用户信息 |
| 部分更新 | PATCH | `PATCH /users/123` | 部分更新用户信息 |
| 删除 | DELETE | `DELETE /users/123` | 删除用户 |

#### 3. 无状态（Stateless）

- 每个请求包含所有必要信息
- 服务器不保存客户端状态
- 通过 Token（如 JWT）进行身份验证

#### 4. 统一接口

使用标准的 HTTP 状态码和方法，保持接口一致性。

#### 5. 资源的表现形式

通过 `Content-Type` 指定数据格式（JSON、XML 等）：

```http
Accept: application/json
Content-Type: application/json
```

### RESTful API 最佳实践

#### 1. 版本控制

```
https://api.example.com/v1/users
https://api.example.com/v2/users
```

#### 2. 过滤、排序、分页

```
GET /users?status=active&sort=created_at&page=1&limit=20
```

#### 3. 嵌套资源

```
GET /users/123/orders          # 获取用户的订单
POST /users/123/orders         # 为用户创建订单
GET /users/123/orders/456      # 获取用户的特定订单
```

#### 4. 错误处理

返回清晰的错误信息：

```json
{
  "error": {
    "code": "INVALID_INPUT",
    "message": "用户名不能为空",
    "field": "username"
  }
}
```

#### 5. HATEOAS（可选）

在响应中包含相关资源的链接：

```json
{
  "id": 123,
  "name": "张三",
  "links": {
    "self": "/users/123",
    "orders": "/users/123/orders",
    "profile": "/users/123/profile"
  }
}
```

### RESTful vs 传统 API

| 特性 | RESTful | 传统 API |
|------|---------|----------|
| URL 设计 | 资源导向 `/users/123` | 动作导向 `/getUser?id=123` |
| HTTP 方法 | 充分利用 GET/POST/PUT/DELETE | 主要使用 GET/POST |
| 状态码 | 语义化状态码 | 通常只用 200 |
| 可读性 | 高 | 较低 |
| 缓存 | 易于实现 | 较难 |

### 示例：用户管理 API

```go
// 获取用户列表
GET /api/v1/users?page=1&limit=10
Response: 200 OK
[
  {"id": 1, "name": "张三", "email": "zhangsan@example.com"},
  {"id": 2, "name": "李四", "email": "lisi@example.com"}
]

// 获取单个用户
GET /api/v1/users/1
Response: 200 OK
{"id": 1, "name": "张三", "email": "zhangsan@example.com"}

// 创建用户
POST /api/v1/users
Body: {"name": "王五", "email": "wangwu@example.com"}
Response: 201 Created
{"id": 3, "name": "王五", "email": "wangwu@example.com"}

// 更新用户
PUT /api/v1/users/1
Body: {"name": "张三三", "email": "zhangsan@example.com"}
Response: 200 OK
{"id": 1, "name": "张三三", "email": "zhangsan@example.com"}

// 删除用户
DELETE /api/v1/users/1
Response: 204 No Content
```

---

## 总结

- HTTP 是无状态的请求-响应协议
- 使用合适的 HTTP 方法和状态码
- RESTful API 以资源为中心，遵循统一接口原则
- 良好的 API 设计提高可维护性和可扩展性
