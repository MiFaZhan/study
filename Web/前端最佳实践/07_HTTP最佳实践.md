# HTTP 最佳实践

---

## 1. HTTP 方法

| 方法 | 用途 | 幂等 | 安全 |
|------|------|------|------|
| GET | 获取资源 | ✅ | ✅ |
| POST | 创建资源 | ❌ | ❌ |
| PUT | 更新资源（整体） | ✅ | ❌ |
| PATCH | 更新资源（部分） | ❌ | ❌ |
| DELETE | 删除资源 | ✅ | ❌ |
| HEAD | 获取元信息 | ✅ | ✅ |
| OPTIONS | 预检请求 | ✅ | ✅ |

---

## 2. RESTful API 设计

```bash
# ✅ 资源用名词复数
GET    /api/users          # 获取用户列表
GET    /api/users/123      # 获取单个用户
POST   /api/users           # 创建用户
PUT    /api/users/123       # 更新用户（整体）
PATCH  /api/users/123       # 更新用户（部分）
DELETE /api/users/123       # 删除用户

# ❌ 避免动词在 URL 中
POST /api/getUser
POST /api/createUser
POST /api/deleteOrder
```

### 嵌套资源

```bash
GET    /api/users/123/orders          # 获取用户123的订单
GET    /api/users/123/orders/456     # 获取用户123的订单456
POST   /api/users/123/orders          # 为用户123创建订单
```

---

## 3. HTTP 状态码

### 2xx 成功
| 状态码 | 说明 |
|--------|------|
| 200 | OK（标准成功） |
| 201 | Created（创建成功） |
| 204 | No Content（成功无返回） |

### 3xx 重定向
| 状态码 | 说明 |
|--------|------|
| 301 | Moved Permanently（永久重定向） |
| 302 | Found（临时重定向） |
| 304 | Not Modified（缓存） |

### 4xx 客户端错误
| 状态码 | 说明 |
|--------|------|
| 400 | Bad Request（参数错误） |
| 401 | Unauthorized（未登录） |
| 403 | Forbidden（无权限） |
| 404 | Not Found（资源不存在） |
| 409 | Conflict（冲突） |
| 422 | Unprocessable Entity（验证失败） |
| 429 | Too Many Requests（限流） |

### 5xx 服务端错误
| 状态码 | 说明 |
|--------|------|
| 500 | Internal Server Error |
| 502 | Bad Gateway |
| 503 | Service Unavailable |
| 504 | Gateway Timeout |

---

## 4. 请求头规范

```bash
# 基础
Content-Type: application/json
Accept: application/json

# 认证
Authorization: Bearer <token>
Authorization: Basic <base64>

# 缓存
If-None-Match: "etag-value"
If-Modified-Since: Wed, 21 Oct 2015 07:28:00 GMT

# 其他
User-Agent: Mozilla/5.0
Accept-Language: zh-CN,zh;q=0.9
```

---

## 5. 响应格式

### 成功响应
```json
{
    "code": 200,
    "message": "success",
    "data": {
        "id": 1,
        "name": "张三",
        "email": "zhang@example.com"
    }
}
```

### 列表响应
```json
{
    "code": 200,
    "message": "success",
    "data": [
        { "id": 1, "name": "张三" },
        { "id": 2, "name": "李四" }
    ],
    "pagination": {
        "page": 1,
        "pageSize": 10,
        "total": 100,
        "totalPages": 10
    }
}
```

### 错误响应
```json
{
    "code": 400,
    "message": "参数错误",
    "errors": [
        { "field": "email", "message": "邮箱格式不正确" },
        { "field": "password", "message": "密码至少8位" }
    ]
}
```

---

## 6. 缓存策略

```bash
# 强缓存（不过期不请求）
Cache-Control: max-age=3600        # 1小时
Expires: Wed, 21 Oct 2026 07:28:00 GMT

# 协商缓存（需验证）
Cache-Control: no-cache            # 每次都验证
Cache-Control: private              # 仅浏览器缓存
Cache-Control: public               # 可被 CDN 缓存

# ETag 验证
ETag: "33a64df551425fcc55e4d42a148795d9"
# 请求时带 If-None-Match: "33a64df5..."
# 若匹配返回 304 Not Modified
```

### 常用 Cache-Control 值

| 值 | 说明 |
|----|------|
| max-age=秒 | 缓存有效期 |
| no-cache | 每次验证 |
| no-store | 不缓存 |
| private | 仅浏览器缓存 |
| public | 可被代理缓存 |

---

## 7. CORS 跨域

```bash
# 允许所有来源（开发用）
Access-Control-Allow-Origin: *

# 允许指定来源
Access-Control-Allow-Origin: https://example.com

# 允许的方法
Access-Control-Allow-Methods: GET, POST, PUT, DELETE

# 允许的头
Access-Control-Allow-Headers: Content-Type, Authorization

# 预检缓存时间
Access-Control-Max-Age: 86400  # 24小时
```

---

## 8. HTTP2 / HTTP3 优势

```
HTTP/1.1: 串行请求，多个请求需排队
HTTP/2:   多路复用，并行请求
HTTP/3:   基于 QUIC，进一步减少延迟
```

---

## 9. HTTPS

```
HTTP  →  明文传输
HTTPS →  TLS 加密传输

# TLS 1.2  vs  1.3
1.3 更安全，握手更快（0-RTT/1-RTT）
```

---

## 10. 常用工具

### 浏览器 DevTools
- Network 面板：查看请求/响应详情
- 禁用缓存：模拟首次请求

### 命令行工具
```bash
# curl 示例
curl -X GET "https://api.example.com/users" \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json"

curl -X POST "https://api.example.com/users" \
  -d '{"name":"张三"}' \
  -H "Content-Type: application/json"
```

### 接口测试
- Postman
- Insomnia
- Apifox

---

## 速查表

| 分类 | 要点 |
|------|------|
| URL 命名 | 名词复数，无动词 |
| GET | 查询，不修改 |
| POST | 创建 |
| PUT | 整体更新 |
| PATCH | 部分更新 |
| DELETE | 删除 |
| 状态码 | 2xx成功/4xx客户端/5xx服务端 |
| 认证 | Authorization Bearer token |
| 缓存 | Cache-Control / ETag |
| 错误 | 4xx参数/401未登录/403无权限 |
