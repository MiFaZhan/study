
在 **GoFrame** 框架中，Web 服务的开发非常高效。下面我将为你详细介绍你提到的四个核心知识点：路由注册、路由参数获取、控制器开发以及数据库操作。

---

## 1. 路由注册

GoFrame 提供了多种路由注册方式，其中最常用的是 **`BindHandler`** 和 **控制器注解（`@handler`）**。

### 1.1 使用 `g.Server().BindHandler()`

通过 `BindHandler` 可以将一个函数直接绑定到指定的路由路径上，支持灵活的路由规则。

```go
package main

import (
    "github.com/gogf/gf/v2/frame/g"
    "github.com/gogf/gf/v2/net/ghttp"
)

func main() {
    s := g.Server()
    
    // 绑定 GET 请求
    s.BindHandler("/hello", func(r *ghttp.Request) {
        r.Response.WriteJson(g.Map{
            "message": "Hello World",
        })
    })
    
    // 绑定 POST 请求，支持 RESTful 风格
    s.BindHandler("/user/:id", func(r *ghttp.Request) {
        id := r.Get("id")
        r.Response.WriteJson(g.Map{
            "user_id": id,
        })
    })
    
    s.Run()
}
```

### 1.2 使用控制器注解（`@handler`）

在 GoFrame 2.x 中，引入了基于注解的路由注册方式，通过在控制器结构体或方法上添加 `@handler` 注解，可以自动注册路由，使代码更加简洁。

```go
package controller

import (
    "github.com/gogf/gf/v2/net/ghttp"
)

type UserController struct{}

// @handler GET /user/:id
func (c *UserController) GetUser(r *ghttp.Request) {
    id := r.Get("id")
    r.Response.WriteJson(g.Map{
        "id": id,
        "name": "John",
    })
}

// @handler POST /user
func (c *UserController) CreateUser(r *ghttp.Request) {
    // 处理创建用户
}
```

需要在启动时扫描注解：

```go
s := g.Server()
s.Group("/", func(group *ghttp.RouterGroup) {
    group.Bind([]interface{}{
        &controller.UserController{},
    })
})
```

---

## 2. 路由参数获取

在 GoFrame 的 `ghttp.Request` 对象中，提供了丰富的方法来获取不同类型的参数。

### 2.1 URL 路径参数

路径参数通常定义在路由规则中，如 `/user/:id`，可以通过 `r.Get("id")` 获取。

```go
s.BindHandler("/user/:id", func(r *ghttp.Request) {
    id := r.Get("id")          // 返回 interface{}，一般用字符串
    idInt := r.GetInt("id")    // 直接转换为 int
    r.Response.WriteJson(g.Map{"id": idInt})
})
```

### 2.2 查询参数（Query String）

查询参数即 URL `?key=value` 部分，同样使用 `r.Get()` 系列方法。

```go
s.BindHandler("/search", func(r *ghttp.Request) {
    keyword := r.GetString("keyword")
    page := r.GetInt("page", 1)  // 第二个参数为默认值
    r.Response.WriteJson(g.Map{
        "keyword": keyword,
        "page":    page,
    })
})
```

### 2.3 表单数据

POST 表单数据（`application/x-www-form-urlencoded` 或 `multipart/form-data`）也通过相同方式获取：

```go
s.BindHandler("/form", func(r *ghttp.Request) {
    name := r.GetString("name")
    age := r.GetInt("age")
    // 处理表单数据
})
```

### 2.4 JSON 数据

当请求的 `Content-Type` 为 `application/json` 时，可以使用 `r.GetJson()` 或直接解析到结构体。

```go
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

s.BindHandler("/json", func(r *ghttp.Request) {
    var user User
    if err := r.Parse(&user); err != nil {
        r.Response.WriteJsonExit(g.Map{"error": err.Error()})
    }
    // 使用 user 对象
})
```

或者直接获取 JSON 数据：

```go
jsonData := r.GetJson()
name := jsonData.Get("name").String()
```

---

## 3. 控制器开发（RESTful 风格）

在 GoFrame 中，控制器通常继承自 `ghttp.Controller` 或直接使用普通结构体，并按方法名（如 `Get`、`Post`、`Delete`）自动映射 HTTP 方法。

### 3.1 简单 RESTful 控制器

```go
type UserController struct{}

func (c *UserController) Get(r *ghttp.Request) {
    r.Response.WriteJson(g.Map{"method": "GET", "action": "list"})
}

func (c *UserController) Post(r *ghttp.Request) {
    r.Response.WriteJson(g.Map{"method": "POST", "action": "create"})
}

func (c *UserController) Put(r *ghttp.Request) {
    r.Response.WriteJson(g.Map{"method": "PUT", "action": "update"})
}

func (c *UserController) Delete(r *ghttp.Request) {
    r.Response.WriteJson(g.Map{"method": "DELETE", "action": "delete"})
}

// 注册
s.Group("/user", func(group *ghttp.RouterGroup) {
    group.Bind(&UserController{})
})
```

这样，访问 `/user` 时，GET 请求会调用 `Get` 方法，POST 请求调用 `Post` 方法，以此类推。

### 3.2 带路径参数的 RESTful

可以在路由中定义资源 ID：

```go
s.Group("/user", func(group *ghttp.RouterGroup) {
    group.GET("/:id", func(r *ghttp.Request) {
        id := r.Get("id")
        r.Response.WriteJson(g.Map{"user_id": id})
    })
})
```

也可以使用控制器方法直接接收参数：

```go
// @handler GET /user/:id
func (c *UserController) Detail(r *ghttp.Request) {
    id := r.Get("id")
    // ...
}
```

---

## 4. 数据库操作

GoFrame 内置了强大的 ORM 组件 `gdb`，支持链式操作，非常易用。

### 4.1 数据库配置

在 `config/config.yaml` 中配置数据库连接：

```yaml
database:
  default:
    type: "mysql"
    link: "root:123456@tcp(127.0.0.1:3306)/test"
    debug: true
```

### 4.2 基本增删改查

```go
// 插入
result, err := g.DB().Insert("user", g.Map{
    "name": "John",
    "age":  30,
})
id, _ := result.LastInsertId()

// 查询单条
user, err := g.DB().Model("user").Where("id", 1).One()
if err == nil {
    fmt.Println(user["name"].String())
}

// 查询多条
users, err := g.DB().Model("user").Where("age >= ?", 18).All()
for _, user := range users {
    fmt.Println(user["name"])
}

// 更新
_, err = g.DB().Model("user").Where("id", 1).Data(g.Map{"age": 31}).Update()

// 删除
_, err = g.DB().Model("user").Where("id", 1).Delete()
```

### 4.3 在控制器中使用数据库

通常在控制器中通过 `g.DB()` 获取数据库对象，也可以将数据库服务注入到控制器中（推荐使用依赖注入）。

```go
func (c *UserController) Get(r *ghttp.Request) {
    id := r.GetInt("id")
    user, err := g.DB().Model("user").Where("id", id).One()
    if err != nil {
        r.Response.WriteJsonExit(g.Map{"error": err.Error()})
    }
    r.Response.WriteJson(user)
}
```

### 4.4 事务处理

使用 `g.DB().Transaction()` 进行事务操作：

```go
err := g.DB().Transaction(r.Context(), func(ctx context.Context, tx *gdb.TX) error {
    _, err := tx.Model("user").Data(g.Map{"name": "Tom"}).Insert()
    if err != nil {
        return err
    }
    _, err = tx.Model("account").Data(g.Map{"user_id": 1, "balance": 100}).Insert()
    return err
})
if err != nil {
    // 事务回滚
}
```

---

## 总结

- **路由注册**：`BindHandler` 适合快速绑定匿名函数；控制器注解和 `Group().Bind()` 更适合结构化、RESTful 风格的 API。
- **参数获取**：通过 `ghttp.Request` 的 `Get`、`GetInt`、`Parse` 等方法可以轻松获取各种格式的请求参数。
- **控制器开发**：通过方法名映射 HTTP 方法，实现 RESTful 接口。
- **数据库操作**：使用 `gdb` ORM 进行链式查询和事务处理，简化数据库交互。

掌握这些核心概念，你就能高效地使用 GoFrame 构建 Web 应用了。如果想深入了解某个方面，可以查阅官方文档或继续提问。