# CSS 最佳实践

---

## 1. box-sizing 重置（必用）

```css
*, *::before, *::after {
    box-sizing: border-box;
}
```
防止 padding/border 撑大元素宽度。

---

## 2. CSS Reset / Normalize

```css
/* 简单版 */
* {
    margin: 0;
    padding: 0;
}
```
推荐使用 [normalize.css](https://necolas.github.io/normalize.css/)，保留有用默认值。

---

## 3. 选择器规范

```css
/* ✅ ID 少用，class 为主 */
.page-container { }
.card-title { }

/* ✅ BEM 命名法 */
.block__element--modifier { }
.modal__header--active { }

/* ❌ 避免嵌套过深 */
.nav .list .item .link .text { }

/* ✅ 嵌套一般不超过 3 层 */
.nav { }
.nav .item { }
.nav .item .link { }
```

---

## 4. 单位使用

```css
/* 字体/间距用 rem，便于无障碍和响应式 */
font-size: 1rem;
padding: 1rem;

/* 固定边框用 px */
border: 1px solid #000;

/* 动画用无单位数字 */
transition: transform 0.3s, opacity 0.3s;

/* 媒体查询用 em（部分浏览器不支持 rem） */
@media (min-width: 48em) { }
```

---

## 5. 简写属性

```css
/* ✅ */
margin: 0 auto;
padding: 10px 20px;
background: #fff url(bg.png) no-repeat center;
transition: all 0.3s;

/* 明确要覆盖的才拆开 */
margin-top: 10px; /* 只需改顶部时单独写 */
```

---

## 6. 颜色规范

```css
/* ✅ 推荐简写 */
color: #fff;
background: #000;
border-color: #333;

/* ❌ 不必要 */
color: #ffffff;
background-color: #000000;
```

---

## 7. Flex 布局

### 垂直居中

```css
.parent {
    display: flex;
    justify-content: center;
    align-items: center;
}
```

### 消除块级空格

```css
/* 用 gap 代替子元素 margin */
.flex-container {
    display: flex;
    gap: 1rem;
}
```

---

## 8. Grid 布局

```css
.grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
    gap: 1rem;
}
```

---

## 9. 响应式设计

```css
/* ✅ min-width 从小到大写 */
@media (min-width: 768px) { ... }
@media (min-width: 1024px) { ... }

/* ❌ 不用 max-width 从大到小 */
```

---

## 10. 动画性能

```css
/* ✅ 只用 transform/opacity（GPU 加速） */
transition: transform 0.3s, opacity 0.3s;

/* ❌ 避免动画布局属性 */
transition: width 0.3s, height 0.3s, margin 0.3s;
```

---

## 11. 图片自适应

```css
img {
    max-width: 100%;
    height: auto;
}
```

---

## 12. 移动端优化

```css
/* 去除点击高亮 */
a, button {
    -webkit-tap-highlight-color: transparent;
}

/* 圆角兼容性 */
border-radius: 8px;
```

---

## 13. 字体抗锯齿

```css
body {
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
}
```

---

## 14. CSS 变量

```css
:root {
    --primary: #3498db;
    --spacing: 1rem;
    --border-radius: 8px;
}

.button {
    background: var(--primary);
    padding: var(--spacing);
    border-radius: var(--border-radius);
}
```

---

## 15. 禁止文本选中

```css
.no-select {
    user-select: none;
}
```

---

## 16. 媒体查询位置

```css
/* ✅ 放文件底部或单独文件 */
.card { ... }

/* 响应式覆盖放后面 */
@media (min-width: 768px) {
    .card { ... }
}
```

---

## 速查表

| 规范 | 推荐程度 |
|------|---------|
| `box-sizing: border-box` | ⭐⭐⭐ 必用 |
| CSS Reset | ⭐⭐⭐ 必用 |
| BEM 命名 | ⭐⭐ 推荐 |
| `img { max-width: 100% }` | ⭐⭐⭐ 必用 |
| 动画用 transform/opacity | ⭐⭐⭐ 必用 |
| CSS 变量 | ⭐⭐⭐ 推荐 |
| 颜色/简写属性 | ⭐ 可选 |
| `!important` 避免使用 | ⭐⭐⭐ 必免 |
