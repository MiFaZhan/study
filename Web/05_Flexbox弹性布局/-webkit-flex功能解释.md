
## `display: -webkit-flex;` 功能解释

这是一个 **CSS Flexbox 布局** 属性，用于将元素设置为弹性容器。

### 详细说明：

| 属性 | 说明 |
|------|------|
| `display: -webkit-flex;` | WebKit 内核浏览器（Safari、Chrome 早期版本）的私有前缀写法 |
| `display: flex;` | 标准 Flexbox 属性，现代浏览器通用 |

### 作用：

将 `.flex-container` 元素变成一个 **弹性容器（Flex Container）**，其子元素（`.flex-item`）会自动变为 **弹性项目（Flex Items）**，可以：

- **水平或垂直排列** - 子元素会自动在一行或一列中排列
- **自动换行** - 当空间不足时，子元素可以自动换行
- **对齐控制** - 方便控制子元素的对齐方式（居中、两端对齐等）
- **等高布局** - 子元素自动等高

### 在你的代码中：

```css
.flex-container {
    display: -webkit-flex;  /* 旧版 WebKit 浏览器 */
    display: flex;          /* 现代浏览器标准写法 */
    width: 400px;
    height: 250px;
    background-color: lightgrey;
}
```

这样设置后，`flex-container` 内的三个 `div` 盒子会水平排列成一个弹性布局，宽度和高度会自动调整。

> 💡 **最佳实践**：通常同时写两个属性（带前缀和不带前缀），以确保兼容新旧浏览器。现代项目可能只需要 `display: flex;`。
        