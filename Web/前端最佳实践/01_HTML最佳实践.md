# HTML 最佳实践

---

## 1. 语义化标签

使用语义化标签代替 div：

```html
<!-- ✅ 好 -->
<header></header>
<nav></nav>
<main>
    <article>
        <section></section>
    </article>
</main>
<aside></aside>
<footer></footer>

<!-- ❌ 差 -->
<div class="header"><div class="nav">...
```

| 标签 | 用途 |
|------|------|
| `<header>` | 页面/区域头部 |
| `<nav>` | 导航栏 |
| `<main>` | 主内容区 |
| `<article>` | 文章/帖子 |
| `<section>` | 章节/区块 |
| `<aside>` | 侧边栏 |
| `<footer>` | 页面/区域底部 |

---

## 2. 表单最佳实践

### label 与 input 关联

```html
<!-- ✅ 好 - 点击 label 聚焦 input -->
<label for="username">用户名</label>
<input type="text" id="username" name="username" />

<!-- ❌ 差 -->
<span>用户名</span>
<input type="text" />
```

### 输入类型

```html
<!-- ✅ 使用正确类型，浏览器会提供对应键盘/验证 -->
<input type="email" />
<input type="tel" />
<input type="number" />
<input type="date" />
<input type="url" />

<!-- ❌ 全部用 text -->
<input type="text" />
```

### 占位符不能替代 label

```html
<!-- ❌ 差 - 占位符消失后用户不知道要填什么 -->
<input type="text" placeholder="请输入用户名" />

<!-- ✅ 好 -->
<label for="username" class="sr-only">用户名</label>
<input type="text" id="username" placeholder="请输入用户名" />
```

---

## 3. 属性顺序

```html
<input 
    type="text" 
    name="username" 
    id="username" 
    class="input-field"
    placeholder="用户名"
    required
    autocomplete="username"
/>
<!-- type → name → id → class → placeholder → 其他 -->
```

---

## 4. 嵌套规则

```html
<!-- 内联元素不能包含块级元素 -->
<!-- ❌ 错误 -->
<a><div>错误</div></a>
<span><p>错误</p></span>

<!-- ✅ 正确 -->
<div><a>正确</a></div>
<p><span>正确</span></p>
```

---

## 5. 资源路径

```html
<!-- 相对路径用于本地资源 -->
<img src="./assets/logo.png" alt="logo" />
<link href="./styles/main.css" rel="stylesheet" />

<!-- 协议相对 URL 用于外部资源 -->
<a href="//external.com">外部链接</a>

<!-- 绝对 URL 带协议 -->
<a href="https://example.com">完整URL</a>
```

---

## 6. 图片与多媒体

```html
<!-- 必须加 alt -->
<img src="photo.jpg" alt="产品图片描述" />

<!-- alt 为空用于装饰性图片 -->
<img src="decorative.svg" alt="" role="presentation" />

<!-- 指定尺寸避免布局抖动 -->
<img src="photo.jpg" alt="" width="200" height="150" />
```

---

## 7. SEO 基础

```html
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>页面标题 - 网站名</title>
    <meta name="description" content="页面描述，不超过160字符">
    <link rel="canonical" href="https://example.com/page">
</head>
```

---

## 8. 可访问性（a11y）

```html
<!-- 链接文本要有意义 -->
<!-- ❌ -->
<a href="#">点击这里</a>
<a href="javascript:void(0)">更多</a>

<!-- ✅ -->
<a href="/products">查看全部产品</a>

<!-- 按钮用 <button>，链接用 <a> -->
<!-- ❌ -->
<div onclick="submit()">提交</div>
<a href="#" onclick="submit()">提交</a>

<!-- ✅ -->
<button type="submit">提交</button>
```

---

## 9. HTML5 模板

```html
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="">
    <title>页面标题</title>
    <link rel="stylesheet" href="styles.css">
</head>
<body>
    <header>
        <nav></nav>
    </header>
    
    <main></main>
    
    <footer></footer>
    
    <script src="main.js"></script>
</body>
</html>
```

---

## 速查表

| 规范 | 说明 |
|------|------|
| 语义化标签 | 用 `<nav>` 代替 `<div class="nav">` |
| label 关联 | `for` 与 `id` 对应 |
| input type | email/tel/number 等代替 text |
| alt 属性 | 图片必须加描述 |
| 嵌套规则 | 内联不包块级 |
| 链接文本 | 要有意义，不可用"点击这里" |
