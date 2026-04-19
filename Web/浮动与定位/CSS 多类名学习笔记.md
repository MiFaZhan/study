# CSS 多类名学习笔记

## 什么是多类名？

一个 HTML 元素可以同时应用**多个 class（类）**，多个类名之间用**空格**隔开。

```html
<div class="类名1 类名2 类名3">内容</div>
```

## 为什么使用多类名？

**好处：样式可以复用，避免重复代码**

比如有一个基础的盒子样式：

```css
.box {
    width: 100px;
    height: 100px;
    margin: 10px;
}
```

然后可以为不同定位方式创建额外的类：

```css
.relative {
    position: relative;
}

.absolute {
    position: absolute;
}

.fixed {
    position: fixed;
}
```

在 HTML 中组合使用：

```html
<div class="box relative">相对定位</div>
<div class="box absolute">绝对定位</div>
<div class="box fixed">固定定位</div>
```

这样：
- `box` 负责**尺寸和间距**
- `relative/absolute/fixed` 负责**定位方式**
- 各司其职，代码更简洁

## 生活中的类比

就像一个人穿衣服：

```html
<div class="帽子 上衣 裤子 鞋子">一个人</div>
```

| 类名 | 作用 |
|------|------|
| 帽子 | 戴在头上 |
| 上衣 | 穿在上身 |
| 裤子 | 穿在下身 |
| 鞋子 | 穿在脚上 |

每个类负责一部分样式，组合起来就是完整的"人"。

## 注意事项

1. **类名顺序不重要**：`class="box relative"` 和 `class="relative box"` 效果一样
2. **相同属性会覆盖**：如果多个类有相同属性，后面的会覆盖前面的
3. **类名要语义化**：尽量使用有意义的名称，如 `header`、`footer`、`active`