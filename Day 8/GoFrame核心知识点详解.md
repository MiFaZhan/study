# GoFrame 核心知识点详解

## 目录
1. [框架基础](#框架基础)
2. [路由与控制器](#路由与控制器)
3. [数据库操作](#数据库操作)
4. [中间件](#中间件)
5. [配置管理](#配置管理)
6. [日志处理](#日志处理)
7. [数据验证](#数据验证)
8. [响应处理](#响应处理)

---

## 框架基础

### 1.1 项目结构

GoFrame 采用标准的分层架构设计：

```
项目根目录/
├── api/                    # API 接口定义层（对外暴露）
│   └── module/
│       ├── module.go       # 自动生成的接口定义
│       └── v1/
│           └── module.go   # 手动定义的请求/响应结构
│
├── internal/               # 内部实现层（不对外暴露）
│   ├── cmd/               # 命令行入口
│   ├── controller/        # 控制器层
│   ├── logic/             # 业务逻辑层
│   ├── dao/               # 数据访问层
│   ├── model/             # 数据模型层
│   │   ├── entity/        # 数据库表实体
│   │   └── do/            # 数据操作对象
│   ├── service/           # 服务接口层
│   └── consts/            # 常量定义
│
├── manifest/              # 配置清单
│   ├── config/           # 应用配置
│   ├── i18n/             # 国际化
│   └── docker/           # Docker 配置
│
├── resource/             # 静态资源
└── main.go              # 程序入口
```

### 1.2 各层职责

| 层级 | 职责 | 特点 |
|------|------|------|
| API | 定义接口契约 | 使用 g.Meta 标签定义路由和验证规则 |
| Controller | 处理 HTTP 请求 | 参数绑定、调用 Logic |
| Logic | 业务逻辑实现 | 可被多个 Controller 复用 |
| Service | 业务服务接口 | 面向接口编程，便于测试 |
| DAO | 数据库操作 | 自动生成，类型安全 |
| Model | 数据模型 | Entity（查询结果）、DO（操作对象）|


### 1.3 程序启动流程

```go
// main.go
package main

import (
    _ "project/internal/packed"  // 资源打包
    "github.com/gogf/gf/v2/frame/g"
    "github.com/gogf/gf/v2/os/gctx"
    "project/internal/cmd"
    _ "github.com/gogf/gf/contrib/drivers/mysql/v2"  // MySQL 驱动
)

func main() {
    // 1. 设置国际化语言
    g.I18n().SetLanguage("zh-CN")
    
    // 2. 测试数据库连接
    err := g.DB().PingMaster()
    if err != nil {
        panic(err)
    }
    
    // 3. 启动 HTTP 服务
    cmd.Main.Run(gctx.GetInitCtx())
}
```

**启动流程说明：**
1. 导入 `packed` 包，将配置文件打包到二进制
2. 设置全局语言（用于验证错误信息国际化）
3. 测试数据库连接是否正常
4. 运行 cmd 命令启动 HTTP 服务器

---

## 路由与控制器

### 2.1 路由注册方式

#### 方式一：自动路由注册（推荐）

```go
// internal/cmd/cmd.go
var Main = gcmd.Command{
    Name: "main",
    Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
        s := g.Server()
        s.Group("/", func(group *ghttp.RouterGroup) {
            // 注册全局中间件
            group.Middleware(ghttp.MiddlewareHandlerResponse)
            
            // 注册 v1 版本路由
            group.Group("/v1", func(group *ghttp.RouterGroup) {
                // 绑定控制器（自动注册路由）
                group.Bind(
                    users.NewV1(),
                    words.NewV1(),
                )
            })
        })
        s.Run()
        return nil
    },
}
```

#### 方式二：手动路由注册

```go
s := g.Server()
s.BindHandler("/user/login", func(r *ghttp.Request) {
    r.Response.Write("Login")
})
```

### 2.2 API 层定义

```go
// api/users/v1/users.go
package v1

import "github.com/gogf/gf/v2/frame/g"

// 登录请求
type LoginReq struct {
    g.Meta   `path:"/users/login" method:"post" summary:"用户登录" tags:"用户"`
    Username string `v:"required|length:3,12" json:"username" dc:"用户名"`
    Password string `v:"required|length:6,18" json:"password" dc:"密码"`
}

// 登录响应
type LoginRes struct {
    Token string `json:"token" dc:"JWT Token"`
}
```

**g.Meta 标签说明：**
- `path`: 路由路径
- `method`: HTTP 方法（get/post/put/delete）
- `summary`: API 摘要（用于文档）
- `tags`: API 分组标签

**v 标签验证规则：**
- `required`: 必填
- `length:min,max`: 长度范围
- `email`: 邮箱格式
- `between:min,max`: 数值范围
- `regex`: 正则表达式


### 2.3 控制器实现

```go
// internal/controller/users/users_v1_login.go
package users

import (
    "context"
    v1 "project/api/users/v1"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
    // 调用业务逻辑层
    token, err := c.users.Login(ctx, req.Username, req.Password)
    if err != nil {
        return nil, err
    }
    
    return &v1.LoginRes{Token: token}, nil
}
```

**控制器职责：**
- 接收 HTTP 请求参数（自动绑定到结构体）
- 调用 Logic 层处理业务逻辑
- 返回响应数据（自动序列化为 JSON）

### 2.4 路由参数获取

#### 方式一：结构体自动绑定（推荐）

```go
// API 定义
type DetailReq struct {
    g.Meta `path:"/words/{id}" method:"get"`
    Id     uint `v:"required|min:1" json:"id" dc:"单词ID"`
}

// 控制器自动接收
func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {
    // req.Id 已自动绑定
}
```

#### 方式二：手动获取参数

```go
func Handler(r *ghttp.Request) {
    // 获取路径参数
    id := r.Get("id").Int()
    
    // 获取查询参数
    page := r.GetQuery("page").Int()
    size := r.GetQuery("size", 10).Int()  // 默认值 10
    
    // 获取表单数据
    username := r.GetForm("username").String()
    
    // 获取 JSON 数据
    var req LoginReq
    if err := r.Parse(&req); err != nil {
        r.Response.WriteJson(g.Map{"error": err.Error()})
        return
    }
    
    // 获取 Header
    token := r.Header.Get("Authorization")
}
```

### 2.5 RESTful 路由设计

| 操作 | Method | URI | 说明 |
|------|--------|-----|------|
| 列表 | GET | /words | 获取单词列表 |
| 详情 | GET | /words/{id} | 获取单个单词 |
| 创建 | POST | /words | 创建新单词 |
| 更新 | PUT | /words/{id} | 更新单词 |
| 删除 | DELETE | /words/{id} | 删除单词 |

---

## 数据库操作

### 3.1 数据库配置

```yaml
# manifest/config/config.yaml
database:
  default:
    link: "mysql:root:password@tcp(127.0.0.1:3306)/dbname"
    debug: true              # 开启 SQL 调试日志
    charset: "utf8mb4"       # 字符集
    dryRun: false            # 空跑模式（不执行 SQL）
    maxIdle: 10              # 最大空闲连接数
    maxOpen: 100             # 最大打开连接数
    maxLifetime: "30s"       # 连接最大存活时间
```

### 3.2 DAO 层自动生成

```bash
# 生成 DAO、Entity、DO
make dao
# 或
gf gen dao
```

**生成的文件：**
- `internal/dao/users.go` - DAO 对外接口
- `internal/dao/internal/users.go` - DAO 内部实现
- `internal/model/entity/users.go` - 表实体
- `internal/model/do/users.go` - 操作对象

### 3.3 基础 CRUD 操作

#### 插入数据

```go
// 插入单条
result, err := dao.Users.Ctx(ctx).Data(do.Users{
    Username: "oldme",
    Password: "123456",
    Email:    "oldme@example.com",
}).Insert()

// 获取插入 ID
id, _ := result.LastInsertId()

// 批量插入
_, err = dao.Users.Ctx(ctx).Data([]do.Users{
    {Username: "user1", Password: "pass1"},
    {Username: "user2", Password: "pass2"},
}).Insert()
```

#### 查询数据

```go
// 查询单条
var user entity.Users
err := dao.Users.Ctx(ctx).Where("id", 1).Scan(&user)

// 查询多条
var users []entity.Users
err := dao.Users.Ctx(ctx).Where("status", 1).Scan(&users)

// 查询所有
list, err := dao.Users.Ctx(ctx).All()

// 查询单个字段
username, err := dao.Users.Ctx(ctx).Where("id", 1).Value("username")

// 统计数量
count, err := dao.Users.Ctx(ctx).Where("status", 1).Count()
```

#### 更新数据

```go
// 更新指定字段
_, err := dao.Users.Ctx(ctx).
    Where("id", 1).
    Data(do.Users{
        Username: "newname",
        Email:    "new@example.com",
    }).Update()

// 更新单个字段
_, err = dao.Users.Ctx(ctx).Where("id", 1).Update(g.Map{
    "status": 1,
})

// 自增/自减
_, err = dao.Users.Ctx(ctx).Where("id", 1).Increment("login_count", 1)
_, err = dao.Users.Ctx(ctx).Where("id", 1).Decrement("balance", 100)
```

#### 删除数据

```go
// 删除
_, err := dao.Users.Ctx(ctx).Where("id", 1).Delete()

// 批量删除
_, err = dao.Users.Ctx(ctx).Where("status", 0).Delete()
```


### 3.4 条件查询

#### Where 条件

```go
// 单个条件
dao.Users.Ctx(ctx).Where("id", 1)
dao.Users.Ctx(ctx).Where("status", 1)

// 多个条件（AND）
dao.Users.Ctx(ctx).
    Where("status", 1).
    Where("age >", 18)

// 使用 And 方法
dao.Users.Ctx(ctx).
    Where("status", 1).
    And("age >", 18)

// 使用 Or 方法
dao.Users.Ctx(ctx).
    Where("status", 1).
    Or("vip_level >", 3)

// IN 查询
dao.Users.Ctx(ctx).WhereIn("id", []int{1, 2, 3})

// BETWEEN 查询
dao.Users.Ctx(ctx).WhereBetween("age", 18, 30)

// NULL 查询
dao.Users.Ctx(ctx).WhereNull("deleted_at")
dao.Users.Ctx(ctx).WhereNotNull("email")

// 模糊查询
dao.Users.Ctx(ctx).WhereLike("username", "%admin%")
```

#### 复杂条件

```go
// 使用 Map
dao.Users.Ctx(ctx).Where(g.Map{
    "status":     1,
    "age >":      18,
    "vip_level": []int{1, 2, 3},  // IN 查询
})

// 使用原生 SQL
dao.Users.Ctx(ctx).Where("age > ? AND status = ?", 18, 1)

// 子查询
dao.Users.Ctx(ctx).Where("id IN(?)", 
    dao.Orders.Ctx(ctx).Fields("user_id").Where("amount >", 1000))
```

### 3.5 排序和分页

```go
// 排序
dao.Users.Ctx(ctx).Order("id DESC")
dao.Users.Ctx(ctx).OrderAsc("created_at")
dao.Users.Ctx(ctx).OrderDesc("id")

// 多字段排序
dao.Users.Ctx(ctx).Order("status DESC, id ASC")

// 分页方式一：Page
dao.Users.Ctx(ctx).Page(1, 10)  // 第 1 页，每页 10 条

// 分页方式二：Limit + Offset
dao.Users.Ctx(ctx).Limit(10).Offset(0)

// 分页查询并统计总数
var list []entity.Users
var total int
err := dao.Users.Ctx(ctx).
    Where("status", 1).
    OrderDesc("id").
    Page(page, size).
    ScanAndCount(&list, &total, true)
```

### 3.6 字段操作

```go
// 指定查询字段
dao.Users.Ctx(ctx).Fields("id, username, email")

// 排除字段
dao.Users.Ctx(ctx).FieldsEx("password, salt")

// 使用列常量（推荐）
var cls = dao.Users.Columns()
dao.Users.Ctx(ctx).Fields(cls.Id, cls.Username, cls.Email)

// 聚合函数
count, _ := dao.Users.Ctx(ctx).Count()
sum, _ := dao.Users.Ctx(ctx).Sum("amount")
avg, _ := dao.Users.Ctx(ctx).Avg("age")
max, _ := dao.Users.Ctx(ctx).Max("score")
min, _ := dao.Users.Ctx(ctx).Min("price")
```

### 3.7 分组和关联

```go
// 分组查询
dao.Orders.Ctx(ctx).
    Fields("user_id, COUNT(*) as order_count, SUM(amount) as total_amount").
    Group("user_id").
    Having("order_count >", 5).
    Scan(&result)

// 左连接
dao.Users.Ctx(ctx).
    LeftJoin("orders", "users.id = orders.user_id").
    Fields("users.*, COUNT(orders.id) as order_count").
    Group("users.id").
    Scan(&result)

// 内连接
dao.Users.Ctx(ctx).
    InnerJoin("orders", "users.id = orders.user_id").
    Where("orders.status", 1).
    Scan(&result)
```

### 3.8 事务处理

#### 方式一：手动事务

```go
// 开启事务
tx, err := g.DB().Begin(ctx)
if err != nil {
    return err
}

// 使用 defer 确保事务处理
defer func() {
    if err != nil {
        tx.Rollback()
    } else {
        tx.Commit()
    }
}()

// 在事务中执行操作
_, err = tx.Model("users").Data(g.Map{
    "username": "oldme",
    "balance":  1000,
}).Insert()
if err != nil {
    return err
}

_, err = tx.Model("orders").Data(g.Map{
    "user_id": 1,
    "amount":  100,
}).Insert()
if err != nil {
    return err
}

// 事务提交（在 defer 中处理）
```

#### 方式二：闭包事务（推荐）

```go
err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
    // 扣减用户余额
    _, err := tx.Model("users").
        Where("id", userId).
        Decrement("balance", amount)
    if err != nil {
        return err
    }
    
    // 创建订单
    _, err = tx.Model("orders").Data(g.Map{
        "user_id": userId,
        "amount":  amount,
    }).Insert()
    if err != nil {
        return err
    }
    
    // 返回 nil 自动提交，返回 error 自动回滚
    return nil
})
```

### 3.9 原生 SQL

```go
// 执行查询
result, err := g.DB().Query(ctx, "SELECT * FROM users WHERE id = ?", 1)

// 执行更新
result, err := g.DB().Exec(ctx, "UPDATE users SET status = ? WHERE id = ?", 1, 1)

// 获取单个值
value, err := g.DB().GetValue(ctx, "SELECT username FROM users WHERE id = ?", 1)

// 获取单条记录
record, err := g.DB().GetOne(ctx, "SELECT * FROM users WHERE id = ?", 1)

// 获取多条记录
records, err := g.DB().GetAll(ctx, "SELECT * FROM users WHERE status = ?", 1)
```


---

## 中间件

### 4.1 中间件概念

中间件是在请求到达控制器之前或响应返回客户端之前执行的函数，用于处理通用逻辑。

**常见用途：**
- 身份认证（JWT 验证）
- 权限检查
- 日志记录
- 跨域处理（CORS）
- 请求限流
- 响应格式统一

### 4.2 中间件注册

```go
// internal/cmd/cmd.go
s := g.Server()
s.Group("/", func(group *ghttp.RouterGroup) {
    // 全局中间件（所有路由生效）
    group.Middleware(ghttp.MiddlewareHandlerResponse)
    group.Middleware(middleware.CORS)
    
    // 分组中间件（仅该组路由生效）
    group.Group("/v1", func(group *ghttp.RouterGroup) {
        // 公开接口（无需认证）
        group.Bind(users.NewV1())
        
        // 需要认证的接口
        group.Group("/", func(group *ghttp.RouterGroup) {
            group.Middleware(middleware.Auth)  // JWT 认证
            group.Bind(
                account.NewV1(),
                words.NewV1(),
            )
        })
    })
})
```

### 4.3 自定义中间件

#### JWT 认证中间件

```go
// internal/logic/middleware/auth.go
package middleware

import (
    "net/http"
    "project/internal/consts"
    "github.com/gogf/gf/v2/net/ghttp"
    "github.com/golang-jwt/jwt/v5"
)

func Auth(r *ghttp.Request) {
    // 获取 Token
    tokenString := r.Header.Get("Authorization")
    if tokenString == "" {
        r.Response.WriteStatus(http.StatusUnauthorized)
        r.Response.WriteJson(g.Map{
            "code":    401,
            "message": "未提供认证令牌",
        })
        r.Exit()
        return
    }
    
    // 解析 Token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(consts.JwtKey), nil
    })
    
    if err != nil || !token.Valid {
        r.Response.WriteStatus(http.StatusForbidden)
        r.Response.WriteJson(g.Map{
            "code":    403,
            "message": "认证令牌无效",
        })
        r.Exit()
        return
    }
    
    // 提取用户信息
    if claims, ok := token.Claims.(jwt.MapClaims); ok {
        r.SetCtxVar("userId", claims["id"])
        r.SetCtxVar("username", claims["username"])
    }
    
    // 继续执行后续中间件和控制器
    r.Middleware.Next()
}
```

#### CORS 跨域中间件

```go
func CORS(r *ghttp.Request) {
    r.Response.CORSDefault()
    r.Middleware.Next()
}

// 或自定义 CORS 配置
func CORS(r *ghttp.Request) {
    r.Response.Header().Set("Access-Control-Allow-Origin", "*")
    r.Response.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
    r.Response.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
    
    // OPTIONS 请求直接返回
    if r.Method == "OPTIONS" {
        r.Response.WriteStatus(http.StatusOK)
        r.Exit()
        return
    }
    
    r.Middleware.Next()
}
```

#### 日志记录中间件

```go
func Logger(r *ghttp.Request) {
    // 记录请求开始时间
    startTime := time.Now()
    
    // 记录请求信息
    g.Log().Infof(r.Context(), "请求开始: %s %s", r.Method, r.URL.Path)
    
    // 执行后续处理
    r.Middleware.Next()
    
    // 记录响应信息
    duration := time.Since(startTime)
    g.Log().Infof(r.Context(), "请求完成: %s %s [%d] %v", 
        r.Method, r.URL.Path, r.Response.Status, duration)
}
```

#### 请求限流中间件

```go
import "golang.org/x/time/rate"

var limiter = rate.NewLimiter(10, 20)  // 每秒 10 个请求，桶容量 20

func RateLimit(r *ghttp.Request) {
    if !limiter.Allow() {
        r.Response.WriteStatus(http.StatusTooManyRequests)
        r.Response.WriteJson(g.Map{
            "code":    429,
            "message": "请求过于频繁，请稍后再试",
        })
        r.Exit()
        return
    }
    r.Middleware.Next()
}
```

### 4.4 中间件执行顺序

```
请求 → 中间件1 → 中间件2 → 中间件3 → 控制器
                                        ↓
响应 ← 中间件1 ← 中间件2 ← 中间件3 ← 控制器
```

**示例：**
```go
func Middleware1(r *ghttp.Request) {
    g.Log().Info(r.Context(), "中间件1 - 前置处理")
    r.Middleware.Next()
    g.Log().Info(r.Context(), "中间件1 - 后置处理")
}

func Middleware2(r *ghttp.Request) {
    g.Log().Info(r.Context(), "中间件2 - 前置处理")
    r.Middleware.Next()
    g.Log().Info(r.Context(), "中间件2 - 后置处理")
}

// 输出顺序：
// 中间件1 - 前置处理
// 中间件2 - 前置处理
// [控制器执行]
// 中间件2 - 后置处理
// 中间件1 - 后置处理
```

---

## 配置管理

### 5.1 配置文件

```yaml
# manifest/config/config.yaml
server:
  address: ":8000"
  serverRoot: "resource/public"
  openapiPath: "/api.json"
  swaggerPath: "/swagger"

logger:
  level: "all"
  stdout: true
  path: "logs"

database:
  default:
    link: "mysql:root:password@tcp(127.0.0.1:3306)/dbname"
    debug: true

redis:
  default:
    address: "127.0.0.1:6379"
    db: 0
    pass: ""

# 自定义配置
app:
  name: "MyApp"
  version: "1.0.0"
  jwtKey: "your-secret-key"
  uploadPath: "uploads"
```

### 5.2 读取配置

```go
// 读取字符串
name := g.Cfg().MustGet(ctx, "app.name").String()

// 读取整数
port := g.Cfg().MustGet(ctx, "server.port", 8000).Int()

// 读取布尔值
debug := g.Cfg().MustGet(ctx, "database.default.debug").Bool()

// 读取到结构体
type AppConfig struct {
    Name    string
    Version string
    JwtKey  string
}

var appConfig AppConfig
err := g.Cfg().MustGet(ctx, "app").Scan(&appConfig)
```

### 5.3 多环境配置

```
manifest/config/
├── config.yaml          # 默认配置
├── config.dev.yaml      # 开发环境
├── config.test.yaml     # 测试环境
└── config.prod.yaml     # 生产环境
```

**启动时指定环境：**
```bash
# 开发环境
gf.gfcli.env=dev go run main.go

# 生产环境
gf.gfcli.env=prod go run main.go
```

**或通过环境变量：**
```bash
export GF_GCFG_FILE=config.prod.yaml
go run main.go
```


---

## 日志处理

### 6.1 日志配置

```yaml
# manifest/config/config.yaml
logger:
  level: "all"                    # 日志级别：all, debug, info, warn, error
  stdout: true                    # 是否输出到控制台
  path: "logs"                    # 日志文件路径
  file: "{Y-m-d}.log"            # 日志文件名格式
  rotateSize: "100M"             # 按大小切分
  rotateExpire: "7d"             # 日志保留时间
  rotateBackupLimit: 10          # 备份文件数量
  rotateBackupExpire: "30d"      # 备份保留时间
  rotateBackupCompress: 9        # 压缩级别（0-9）
  rotateCheckInterval: "1h"      # 检查间隔
  stdoutColorDisabled: false     # 是否禁用颜色
  writerColorEnable: false       # 文件日志是否启用颜色
```

### 6.2 日志使用

#### 基础日志

```go
import "github.com/gogf/gf/v2/frame/g"

// 不同级别的日志
g.Log().Debug(ctx, "调试信息")
g.Log().Info(ctx, "普通信息")
g.Log().Notice(ctx, "通知信息")
g.Log().Warning(ctx, "警告信息")
g.Log().Error(ctx, "错误信息")
g.Log().Critical(ctx, "严重错误")

// 格式化日志
g.Log().Infof(ctx, "用户 %s 登录成功", username)
g.Log().Errorf(ctx, "数据库连接失败: %v", err)

// 打印日志（不带级别）
g.Log().Print(ctx, "打印信息")
g.Log().Printf(ctx, "格式化打印: %s", message)
```

#### 结构化日志

```go
// 使用 Map
g.Log().Info(ctx, g.Map{
    "action":   "login",
    "username": "oldme",
    "ip":       "192.168.1.1",
    "status":   "success",
})

// 使用链式调用
g.Log().
    Cat("user").              // 分类
    File("user.log").         // 指定文件
    Line().                   // 显示行号
    Stack().                  // 显示堆栈
    Info(ctx, "用户操作日志")
```

#### 自定义日志实例

```go
// 创建日志实例
logger := g.Log("custom")

// 配置日志实例
logger.SetLevel("info")
logger.SetPath("logs/custom")
logger.SetFile("custom-{Y-m-d}.log")

// 使用日志实例
logger.Info(ctx, "自定义日志")
```

### 6.3 日志最佳实践

```go
// 业务逻辑中记录日志
func (u *Users) Login(ctx context.Context, username, password string) (token string, err error) {
    // 记录操作开始
    g.Log().Infof(ctx, "用户登录尝试: username=%s", username)
    
    // 查询用户
    var user entity.Users
    err = dao.Users.Ctx(ctx).Where("username", username).Scan(&user)
    if err != nil {
        g.Log().Errorf(ctx, "查询用户失败: username=%s, error=%v", username, err)
        return "", err
    }
    
    if user.Id == 0 {
        g.Log().Warningf(ctx, "用户不存在: username=%s", username)
        return "", gerror.New("用户名或密码错误")
    }
    
    // 验证密码
    if user.Password != u.encryptPassword(password) {
        g.Log().Warningf(ctx, "密码错误: username=%s", username)
        return "", gerror.New("用户名或密码错误")
    }
    
    // 生成 Token
    token, err = u.generateToken(user)
    if err != nil {
        g.Log().Errorf(ctx, "生成Token失败: userId=%d, error=%v", user.Id, err)
        return "", err
    }
    
    // 记录成功
    g.Log().Infof(ctx, "用户登录成功: userId=%d, username=%s", user.Id, username)
    return token, nil
}
```

---

## 数据验证

### 7.1 验证规则

#### 常用规则

```go
type CreateReq struct {
    // 必填
    Username string `v:"required" dc:"用户名"`
    
    // 长度范围
    Password string `v:"required|length:6,18" dc:"密码"`
    
    // 邮箱格式
    Email string `v:"required|email" dc:"邮箱"`
    
    // 数值范围
    Age int `v:"required|between:1,150" dc:"年龄"`
    
    // 最小值/最大值
    Score int `v:"min:0|max:100" dc:"分数"`
    
    // 正则表达式
    Phone string `v:"required|regex:^1[3-9]\\d{9}$" dc:"手机号"`
    
    // 枚举值
    Status int `v:"required|in:0,1,2" dc:"状态"`
    
    // 日期格式
    Birthday string `v:"date" dc:"生日"`
    
    // URL 格式
    Website string `v:"url" dc:"网站"`
    
    // IP 地址
    IP string `v:"ip" dc:"IP地址"`
    
    // 相同值（确认密码）
    ConfirmPassword string `v:"required|same:Password" dc:"确认密码"`
    
    // 不同值
    NewPassword string `v:"required|different:OldPassword" dc:"新密码"`
}
```

#### 组合规则

```go
type RegisterReq struct {
    // 多个规则用 | 分隔
    Username string `v:"required|length:3,12|regex:^[a-zA-Z0-9_]+$" dc:"用户名"`
    
    // 可选字段（不填时不验证）
    Nickname string `v:"length:2,20" dc:"昵称"`
    
    // 条件验证（当 Type=1 时必填）
    Address string `v:"required-if:Type,1" dc:"地址"`
    
    // 条件验证（当 Type 不为空时必填）
    City string `v:"required-with:Type" dc:"城市"`
}
```

### 7.2 自定义验证规则

```go
// 注册自定义规则
func init() {
    rule := "unique-username"
    gvalid.RegisterRule(rule, func(ctx context.Context, in gvalid.RuleFuncInput) error {
        username := in.Value.String()
        
        // 检查用户名是否已存在
        count, err := dao.Users.Ctx(ctx).Where("username", username).Count()
        if err != nil {
            return err
        }
        
        if count > 0 {
            return gerror.New("用户名已存在")
        }
        
        return nil
    })
}

// 使用自定义规则
type RegisterReq struct {
    Username string `v:"required|unique-username" dc:"用户名"`
}
```

### 7.3 自定义错误信息

```go
type CreateReq struct {
    Username string `v:"required|length:3,12#请输入用户名|用户名长度应为3-12个字符" dc:"用户名"`
    Email    string `v:"required|email#请输入邮箱|邮箱格式不正确" dc:"邮箱"`
}
```

### 7.4 国际化验证消息

```toml
# manifest/i18n/zh-CN/validation.toml
"gf.gvalid.rule.required"             = "{field}字段不能为空"
"gf.gvalid.rule.required-if"          = "{field}字段不能为空"
"gf.gvalid.rule.required-unless"      = "{field}字段不能为空"
"gf.gvalid.rule.required-with"        = "{field}字段不能为空"
"gf.gvalid.rule.required-with-all"    = "{field}字段不能为空"
"gf.gvalid.rule.required-without"     = "{field}字段不能为空"
"gf.gvalid.rule.required-without-all" = "{field}字段不能为空"
"gf.gvalid.rule.date"                 = "{field}字段值`{value}`日期格式不正确"
"gf.gvalid.rule.datetime"             = "{field}字段值`{value}`日期时间格式不正确"
"gf.gvalid.rule.date-format"          = "{field}字段值`{value}`日期格式不满足: {pattern}"
"gf.gvalid.rule.email"                = "{field}字段值`{value}`邮箱地址格式不正确"
"gf.gvalid.rule.phone"                = "{field}字段值`{value}`手机号码格式不正确"
"gf.gvalid.rule.telephone"            = "{field}字段值`{value}`电话号码格式不正确"
"gf.gvalid.rule.passport"             = "{field}字段值`{value}`账号格式不合法，必需以字母开头，只能包含字母、数字和下划线，长度在6~18之间"
"gf.gvalid.rule.password"             = "{field}字段值`{value}`密码格式不合法，密码格式为任意6-18位的可见字符"
"gf.gvalid.rule.password2"            = "{field}字段值`{value}`密码格式不合法，密码格式为任意6-18位的可见字符，必须包含大小写字母和数字"
"gf.gvalid.rule.password3"            = "{field}字段值`{value}`密码格式不合法，密码格式为任意6-18位的可见字符，必须包含大小写字母、数字和特殊字符"
"gf.gvalid.rule.postcode"             = "{field}字段值`{value}`邮政编码格式不正确"
"gf.gvalid.rule.resident-id"          = "{field}字段值`{value}`身份证号码格式不正确"
"gf.gvalid.rule.bank-card"            = "{field}字段值`{value}`银行卡号格式不正确"
"gf.gvalid.rule.qq"                   = "{field}字段值`{value}`QQ号码格式不正确"
"gf.gvalid.rule.ip"                   = "{field}字段值`{value}`IP地址格式不正确"
"gf.gvalid.rule.ipv4"                 = "{field}字段值`{value}`IPv4地址格式不正确"
"gf.gvalid.rule.ipv6"                 = "{field}字段值`{value}`IPv6地址格式不正确"
"gf.gvalid.rule.mac"                  = "{field}字段值`{value}`MAC地址格式不正确"
"gf.gvalid.rule.url"                  = "{field}字段值`{value}`URL地址格式不正确"
"gf.gvalid.rule.domain"               = "{field}字段值`{value}`域名格式不正确"
"gf.gvalid.rule.length"               = "{field}字段值`{value}`字段长度应当为{min}到{max}个字符"
"gf.gvalid.rule.min-length"           = "{field}字段值`{value}`字段长度应当大于或等于{min}个字符"
"gf.gvalid.rule.max-length"           = "{field}字段值`{value}`字段长度应当小于或等于{max}个字符"
"gf.gvalid.rule.size"                 = "{field}字段值`{value}`字段大小应当为{size}"
"gf.gvalid.rule.between"              = "{field}字段值`{value}`字段大小应当在{min}到{max}之间"
"gf.gvalid.rule.min"                  = "{field}字段值`{value}`字段大小应当大于或等于{min}"
"gf.gvalid.rule.max"                  = "{field}字段值`{value}`字段大小应当小于或等于{max}"
"gf.gvalid.rule.json"                 = "{field}字段值`{value}`JSON格式不正确"
"gf.gvalid.rule.xml"                  = "{field}字段值`{value}`XML格式不正确"
"gf.gvalid.rule.array"                = "{field}字段值`{value}`应当为数组"
"gf.gvalid.rule.integer"              = "{field}字段值`{value}`应当为整数"
"gf.gvalid.rule.float"                = "{field}字段值`{value}`应当为浮点数"
"gf.gvalid.rule.boolean"              = "{field}字段值`{value}`应当为布尔值"
"gf.gvalid.rule.same"                 = "{field}字段值`{value}`应当与{field1}字段值相同"
"gf.gvalid.rule.different"            = "{field}字段值`{value}`应当与{field1}字段值不同"
"gf.gvalid.rule.in"                   = "{field}字段值`{value}`应当在{pattern}范围内"
"gf.gvalid.rule.not-in"               = "{field}字段值`{value}`不应当在{pattern}范围内"
"gf.gvalid.rule.regex"                = "{field}字段值`{value}`字段值不合法"
```


---

## 响应处理

### 8.1 统一响应格式

```go
// 标准响应结构
type Response struct {
    Code    int         `json:"code"`    // 业务状态码
    Message string      `json:"message"` // 提示信息
    Data    interface{} `json:"data"`    // 响应数据
}
```

### 8.2 响应中间件

```go
// 使用 GoFrame 内置的响应中间件
s.Group("/", func(group *ghttp.RouterGroup) {
    group.Middleware(ghttp.MiddlewareHandlerResponse)
    // ... 路由注册
})
```

**中间件处理逻辑：**
- 控制器返回 `(data, error)`
- 如果 `error != nil`，返回错误响应
- 如果 `error == nil`，返回成功响应

### 8.3 响应方法

```go
// JSON 响应
r.Response.WriteJson(g.Map{
    "code":    0,
    "message": "success",
    "data":    data,
})

// JSONP 响应
r.Response.WriteJsonP(data)

// XML 响应
r.Response.WriteXml(data)

// 字符串响应
r.Response.Write("Hello World")
r.Response.Writef("Hello %s", name)

// 文件下载
r.Response.ServeFile("path/to/file.pdf")
r.Response.ServeFileDownload("path/to/file.pdf", "download.pdf")

// 重定向
r.Response.RedirectTo("/login")
r.Response.RedirectBack()

// 设置状态码
r.Response.WriteStatus(http.StatusNotFound)
r.Response.WriteStatusExit(http.StatusForbidden)
```

### 8.4 错误处理

```go
import "github.com/gogf/gf/v2/errors/gerror"

// 创建错误
err := gerror.New("用户不存在")
err := gerror.Newf("用户 %s 不存在", username)

// 带错误码的错误
err := gerror.NewCode(gcode.New(1001, "用户不存在", nil))

// 包装错误
err := gerror.Wrap(err, "登录失败")
err := gerror.Wrapf(err, "用户 %s 登录失败", username)

// 在控制器中返回错误
func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
    token, err := c.users.Login(ctx, req.Username, req.Password)
    if err != nil {
        // 直接返回错误，中间件会自动处理
        return nil, err
    }
    return &v1.LoginRes{Token: token}, nil
}
```

### 8.5 自定义响应中间件

```go
func CustomResponse(r *ghttp.Request) {
    r.Middleware.Next()
    
    // 如果已经有响应内容，不再处理
    if r.Response.BufferLength() > 0 {
        return
    }
    
    var (
        msg  string
        err  = r.GetError()
        res  = r.GetHandlerResponse()
        code = gerror.Code(err)
    )
    
    if err != nil {
        // 错误响应
        if code == gcode.CodeNil {
            code = gcode.CodeInternalError
        }
        msg = err.Error()
    } else {
        // 成功响应
        code = gcode.CodeOK
        msg = "success"
    }
    
    r.Response.WriteJson(g.Map{
        "code":    code.Code(),
        "message": msg,
        "data":    res,
    })
}
```

---

## 总结

### 开发流程

1. **定义 API 接口**（api 层）
   - 定义请求/响应结构
   - 使用 g.Meta 标签定义路由
   - 使用 v 标签定义验证规则

2. **生成代码**
   ```bash
   make ctrl  # 生成 Controller
   make dao   # 生成 DAO/Entity/DO
   ```

3. **实现业务逻辑**（logic 层）
   - 调用 DAO 层操作数据库
   - 处理业务规则
   - 返回结果或错误

4. **实现控制器**（controller 层）
   - 调用 Logic 层
   - 返回响应数据

5. **注册路由**（cmd 层）
   - 使用 group.Bind() 注册控制器
   - 配置中间件

6. **测试接口**
   - 访问 Swagger 文档
   - 使用 Postman/Apifox 测试

### 最佳实践

1. **分层清晰**：严格遵循分层架构，不跨层调用
2. **错误处理**：使用 gerror 包装错误，提供清晰的错误信息
3. **日志记录**：关键操作记录日志，便于排查问题
4. **参数验证**：使用验证规则，确保数据合法性
5. **事务处理**：涉及多表操作使用事务，保证数据一致性
6. **配置管理**：敏感信息放配置文件，不硬编码
7. **代码生成**：充分利用代码生成工具，提高开发效率
8. **接口文档**：使用 OpenAPI 自动生成文档，保持文档同步

### 常用命令

```bash
# 启动项目
go run main.go

# 生成 Controller
make ctrl
gf gen ctrl

# 生成 DAO
make dao
gf gen dao

# 生成 Service
make service
gf gen service

# 构建项目
make build
gf build

# 更新 GoFrame
make up
gf up
```

### 学习资源

- 官方文档：https://goframe.org
- GitHub：https://github.com/gogf/gf
- 社区论坛：https://goframe.org/pages/viewpage.action?pageId=1114119
